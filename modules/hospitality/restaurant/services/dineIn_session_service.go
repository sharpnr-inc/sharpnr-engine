package restaurantservices

import (
	"context"
	"encoding/json"
	"fmt"
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	restaurant_events "sharpnr-engine/modules/hospitality/restaurant/events"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	payment_constants "sharpnr-engine/payment_engine/constant"
	payment_engine_contracts "sharpnr-engine/payment_engine/contracts"
	payment_engine_models "sharpnr-engine/payment_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	invoice_engine_contracts "sharpnr-engine/service_engines/invoicing_engine/contracts"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *restaurantService) StartDineSession(ctx context.Context,
	restaurantID uuid.UUID,
	orgID uuid.UUID,
	employeeID uuid.UUID,
	tableID uuid.UUID,
) (uuid.UUID, error) { // Removed the named 'diningSessionID' for cleaner return

	var sessionID uuid.UUID

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.restaurantRepo.UpdateTableStatus(txCtx, tableID, restaurant_constants.TableStatusOccupied); err != nil {
			return err
		}

		// 1. Generate the ID first so we can capture it
		sessionID = uuid.New()

		session := &restaurant_models.RestaurantDiningSession{
			ID:           sessionID,
			RestaurantID: restaurantID,
			TableID:      tableID,
			Status:       restaurant_constants.DineInSessionOpened,
			StartedBy:    employeeID,
			StartedAt:    time.Now(),
			CreatedAt:    time.Now(),
		}

		if err := s.restaurantRepo.CreateRestaurantDiningSession(txCtx, session); err != nil {
			return err
		}

		return nil
	})

	// 2. Check if the transaction failed
	if err != nil {
		return uuid.Nil, err
	}

	// 3. Return the captured ID
	return sessionID, nil
}

/* ----------------------------------- <*> ---------------------------------- */

func (s *restaurantService) AddOrdersInSession(
	ctx context.Context,
	dineInSessionID uuid.UUID,
	req *restaurant_contracts.DiningSessionOrder_req_dto,
	orgID uuid.UUID,
	restaurantID uuid.UUID,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Initialize the Event Payload
		eventPayload := restaurant_events.OrderCreatedPayload{
			OrderID:        uuid.New(),
			OrganizationID: orgID,
			RestaurantID:   restaurantID,
			SessionID:      dineInSessionID,
			Items:          make([]restaurant_events.OrderItemEventDTO, 0),
		}

		orderSequence, errr := s.restaurantRepo.GetOrderSequenceNumber(txCtx, restaurantID, orgID, time.Now())
		if errr != nil {
			return errr
		}

		// 2. Create Order Header
		ord := &restaurant_models.RestaurantOrders{
			ID:                        eventPayload.OrderID,
			OrganizationID:            orgID,
			RestaurantID:              restaurantID,
			RestaurantDiningSessionID: &dineInSessionID,
			OrderNumber:               orderSequence,
			Status:                    "PENDING",
			CreatedAt:                 time.Now(),
			UpdatedAt:                 time.Now(),
		}

		if err := s.restaurantRepo.CreateOrder(txCtx, ord); err != nil {
			return err
		}

		for _, itm := range req.Items {

			// Build the Item Record for the DB
			itemRecord := &restaurant_models.RestaurantOrderItems{
				ID:                uuid.New(),
				RestaurantOrderID: ord.ID,
				MenuItemID:        itm.MenuItemID,
				MenuItemName:      itm.ItemName,
				Quantity:          itm.Quantity,
				UnitPrice:         itm.BasePriceAtOrder,
				Status:            "PENDING", // unesseory
				Notes:             itm.SpecialInstructions,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			}

			if err := s.restaurantRepo.CreateOrderItem(txCtx, itemRecord); err != nil {
				return err
			}

			// Prepare Item for the Event
			itemEvent := restaurant_events.OrderItemEventDTO{
				MenuItemID:   itm.MenuItemID,
				MenuItemName: itm.ItemName,
				Quantity:     itm.Quantity,
				Notes:        itm.SpecialInstructions,
				Variants:     make([]restaurant_events.VariantEventDTO, 0),
				Addons:       make([]restaurant_events.AddonEventDTO, 0),
			}

			// 3. Handle Variants
			for _, vr := range itm.SelectedVariants {
				variantRecord := &restaurant_models.RestaurantOrderItemsVariants{
					ID:                    uuid.New(),
					RestaurantOrderItemID: itemRecord.ID,
					VariantID:             vr.VariantID,
					VariantGroupID:        vr.VariantGroupID,
					Name:                  vr.VariantName,
					PriceAtOrder:          vr.PriceAtOrder,
					CreatedAt:             time.Now(),
				}
				if err := s.restaurantRepo.CreateOrderItemVariant(txCtx, variantRecord); err != nil {
					return err
				}

				// Add to Event DTO
				itemEvent.Variants = append(itemEvent.Variants, restaurant_events.VariantEventDTO{
					VariantID:    vr.VariantID,
					VariantName:  vr.VariantName,
					PriceAtOrder: vr.PriceAtOrder,
				})
			}

			// 4. Handle Add-ons
			for _, adn := range itm.SelectedAddOns {
				addonRecord := &restaurant_models.RestaurantOrderItemsAddOns{
					ID:                    uuid.New(),
					RestaurantOrderItemID: itemRecord.ID,
					AddOnID:               adn.ID,
					Name:                  adn.Name,
					Quantity:              adn.Qty,
					PriceAtOrder:          adn.PriceAtOrder,
					CreatedAt:             time.Now(),
				}
				if err := s.restaurantRepo.CreateOrderItemAddon(txCtx, addonRecord); err != nil {
					return err
				}

				// Add to Event DTO
				itemEvent.Addons = append(itemEvent.Addons, restaurant_events.AddonEventDTO{
					AddOnID:      adn.ID,
					Name:         adn.Name,
					Qty:          adn.Qty,
					PriceAtOrder: adn.PriceAtOrder,
				})
			}

			// Append the fully built item DTO to the payload
			eventPayload.Items = append(eventPayload.Items, itemEvent)
		}

		// 5. Publish the clean DTO
		err := s.seroEngine.Publish(txCtx, restaurant_events.OrderCreatedInRestaurant, eventPayload)
		if err != nil {
			return err
		}

		return nil
	})
}

/* ----------------------------------- <*> ---------------------------------- */

func (s *restaurantService) GetDineSessionOrders(ctx context.Context, dineInSessionID uuid.UUID) (*restaurant_contracts.DiningSessionOrders_resp_dto, error) {

	// 1. Fetch orders from Repo
	orders, err := s.restaurantRepo.GetOrdersByDiningSessionID(ctx, dineInSessionID)
	if err != nil {
		return nil, err
	}

	resp := &restaurant_contracts.DiningSessionOrders_resp_dto{
		DiningSessionID: dineInSessionID,
		Orders:          make([]restaurant_contracts.Order_resp_dto, 0),
	}

	var runningSessionTotal float64

	for _, ord := range orders {
		orderDTO := restaurant_contracts.Order_resp_dto{
			ID:          ord.ID,
			OrderNumber: ord.OrderNumber,
			Status:      ord.Status,
			CreatedAt:   ord.CreatedAt,
			Items:       make([]restaurant_contracts.OrderItem_resp_dto, 0),
		}

		// 2. Fetch Items for this specific order
		items, _ := s.restaurantRepo.GetOrderItemsByOrderID(ctx, ord.ID)

		for _, itm := range items {
			itemDTO := restaurant_contracts.OrderItem_resp_dto{
				ID:         itm.ID,
				MenuItemID: itm.MenuItemID,
				Name:       itm.MenuItemName, // Assumes your model/join has this
				Quantity:   itm.Quantity,
				UnitPrice:  itm.UnitPrice,
				Status:     itm.Status,
				Notes:      itm.Notes,
			}

			// 3. Fetch Variants & Addons (Optimization: Use a Join in Repo instead of loops)
			variants, _ := s.restaurantRepo.GetOrderItemVariantByOrderID(ctx, itm.ID)
			for _, v := range variants {
				itemDTO.Variants = append(itemDTO.Variants, restaurant_contracts.Variant_resp_dto{
					ID:    v.VariantID,
					Name:  v.Name,
					Price: v.PriceAtOrder,
				})
			}

			addons, _ := s.restaurantRepo.GetOrderItemAddOnByOrderID(ctx, itm.ID)
			for _, a := range addons {
				itemDTO.AddOns = append(itemDTO.AddOns, restaurant_contracts.Addon_resp_dto{
					ID:       a.AddOnID,
					Name:     a.Name,
					Quantity: a.Quantity,
					Price:    a.PriceAtOrder,
				})
			}

			orderDTO.Items = append(orderDTO.Items, itemDTO)
		}

		resp.Orders = append(resp.Orders, orderDTO)
	}

	resp.SessionTotal = runningSessionTotal
	return resp, nil
}

/* ---------------------------------- <***> --------------------------------- */

/* ----------------------- Closing the session Method ----------------------- */

func (s *restaurantService) CloseDineInSession(
	ctx context.Context,
	orgID, emplID, restaurantID, dineInSessionID uuid.UUID,
	request restaurant_contracts.CloseSessionRequest,
) (*restaurant_contracts.CloseSessionResponse, error) {

	session, err := s.restaurantRepo.GetDiningSessionByID(ctx, dineInSessionID)
	if err != nil {
		return nil, err
	}

	sessionOrders, err := s.restaurantRepo.GetOrdersByDiningSessionID(ctx, dineInSessionID)
	if err != nil {
		return nil, err
	}

	// Step 1
	var items []invoice_engine_contracts.InvoiceItemRequest
	for _, so := range sessionOrders {

		sessionOrderItems, err := s.restaurantRepo.GetOrderItemsByOrderID(ctx, so.ID)
		if err != nil {
			return nil, err
		}

		for _, item := range sessionOrderItems {

			menuItemTaxCategoryID, _, err := s.menuService.GetTaxCategoryOfMenuItem(ctx, orgID, emplID, item.MenuItemID)
			if err != nil {
				return nil, err
			}
			fmt.Println("the taxCategory of :-->", menuItemTaxCategoryID)
			items = append(items, invoice_engine_contracts.InvoiceItemRequest{
				Description:   item.MenuItemName,
				Quantity:      float64(item.Quantity),
				Rate:          item.UnitPrice,
				TaxCategoryID: &menuItemTaxCategoryID,
				HSNCode:       nil,
			})
		}
	}

	rest, err := s.restaurantRepo.GetRestaurantByID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}

	invoice, err := s.invoiceEngine.GenerateInvoice(ctx, invoice_engine_contracts.CreateInvoiceRequest{
		OrganizationID: orgID,
		LegalEntityID:  rest.LegalEntityID,
		CustomerID:     request.CustomerID,
		CustomerName:   request.CustomerName,
		CustomerGST:    request.CustomerGST,
		Source:         "RESTAURANT_SESSION",
		SourceRefID:    &dineInSessionID,
		Items:          items,
	})
	if err != nil {
		return nil, err
	}

	invoice, err = s.invoiceEngine.IssueInvoice(ctx, invoice.ID)
	if err != nil {
		return nil, err
	}

	invoiceIDStr := invoice.ID.String()
	sessionIDStr := dineInSessionID.String()

	// SourceMeta carries context that payment engine echoes back in webhook
	// Your webhook handler uses this to know which session to close
	sourceMeta, _ := json.Marshal(map[string]string{
		"session_id": sessionIDStr,
		"invoice_id": invoiceIDStr,
	})
	sourceMetaStr := string(sourceMeta)

	paymentOrder, err := s.paymentsEngine.CreateOrder(ctx, payment_engine_contracts.CreateOrderRequest{
		TenantID:         orgID,
		LegalEntityID:    *rest.LegalEntityID,
		PaymentAccountID: request.PaymentAccountID,
		PayerType:        "CUSTOMER",
		PayerName:        request.CustomerName,
		AmountPaise:      int64(invoice.Total * 100),
		CurrencyCode:     "INR",
		Source:           "INVOICE",
		SourceRefID:      &invoiceIDStr,
		SourceMeta:       &sourceMetaStr,
		CreatedBy:        emplID.String(),
	})

	if err != nil {
		return nil, fmt.Errorf("create payment order: %w", err)
	}

	payment, err := s.paymentsEngine.Pay(ctx, payment_engine_contracts.PayRequest{
		PaymentOrderID: paymentOrder.ID,
		AmountPaise:    int64(invoice.Total * 100),
		Method:         request.PaymentMethod,
		CreatedBy:      emplID.String(),
	})

	if err != nil {
		return nil, fmt.Errorf("process payment: %w", err)
	}

	if payment.Status == payment_constants.StatusProcessing {
		if err := s.restaurantRepo.UpdateDiningSessionStatus(ctx, dineInSessionID, "AWAITING_PAYMENT"); err != nil {
			return nil, err
		}
		return &restaurant_contracts.CloseSessionResponse{
			SessionID:     dineInSessionID,
			InvoiceID:     invoice.ID,
			InvoiceNumber: invoice.InvoiceNumber,
			Total:         invoice.Total,
			PaymentStatus: "PENDING",
		}, nil
	}

	if err := s.settleInvoice(ctx, payment, invoice.ID, emplID.String()); err != nil {
		return nil, fmt.Errorf("settle invoice: %w", err)
	}

	if err := s.restaurantRepo.UpdateTableStatus(ctx, session.TableID, restaurant_constants.TableStatusAvailable); err != nil {
		return nil, err
	}

	if err := s.restaurantRepo.UpdateDiningSessionStatus(ctx, dineInSessionID, "CLOSED"); err != nil {
		return nil, err
	}

	return &restaurant_contracts.CloseSessionResponse{
		SessionID:     dineInSessionID,
		InvoiceID:     invoice.ID,
		InvoiceNumber: invoice.InvoiceNumber,
		Total:         invoice.Total,
		PaymentStatus: "PAID",
	}, nil

}

/* ---------------------------------- <***> --------------------------------- */

func (s *restaurantService) settleInvoice(
	ctx context.Context,
	payment *payment_engine_models.Payment,
	invoiceID uuid.UUID,
	by string,
) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Tell payment engine this payment covers the invoice
		_, err := s.paymentsEngine.AllocatePayment(txCtx, payment_engine_contracts.AllocateRequest{
			PaymentID:      payment.ID,
			TargetType:     payment_constants.AllocationTargetInvoice,
			TargetID:       invoiceID.String(),
			AllocatedPaise: payment.AmountPaise,
			CreatedBy:      by,
		})
		if err != nil {
			return fmt.Errorf("allocate payment: %w", err)
		}

		// 2. Tell invoice engine payment was received
		// PaymentRef is opaque — invoice engine stores it as a string, never queries payment tables
		if err := s.invoiceEngine.RecordPayment(txCtx, invoice_engine_contracts.RecordPaymentRequest{
			InvoiceID:      invoiceID,
			PaymentRef:     payment.ID.String(),
			PaymentMethod:  payment.Method,
			AllocatedPaise: payment.AmountPaise,
			PaidAt:         *payment.PaidAt,
		}); err != nil {
			return fmt.Errorf("record payment on invoice: %w", err)
		}

		// here we have to

		return nil
	})
}
