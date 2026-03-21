package payment_engine_service

import (
	"context"
	"fmt"

	payment_constants "sharpnr-engine/payment_engine/constant"
	contracts "sharpnr-engine/payment_engine/contracts"
	models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

func (s *paymentEngine) AllocatePayment(
	ctx context.Context,
	req contracts.AllocateRequest,
) (*models.PaymentAllocation, error) {

	payment, err := s.paymentRepo.GetPaymentByID(ctx, req.PaymentID)
	if err != nil {
		return nil, fmt.Errorf("payment not found: %w", err)
	}
	if payment.Status != payment_constants.StatusSuccess {
		return nil, fmt.Errorf("can only allocate a successful payment, status: %s", payment.Status)
	}

	alreadyAllocated, err := s.paymentRepo.GetTotalAllocatedByPayment(ctx, req.PaymentID)
	if err != nil {
		return nil, fmt.Errorf("fetch existing allocations: %w", err)
	}

	available := payment.AmountPaise - alreadyAllocated
	if req.AllocatedPaise > available {
		return nil, fmt.Errorf(
			"allocation of %d paise exceeds available balance of %d paise",
			req.AllocatedPaise, available,
		)
	}

	if !isValidAllocationTarget(req.TargetType) {
		return nil, fmt.Errorf("invalid allocation target type: '%s'", req.TargetType)
	}

	allocation := &models.PaymentAllocation{
		ID:             uuid.New(),
		TenantID:       payment.TenantID,
		PaymentID:      req.PaymentID,
		TargetType:     req.TargetType,
		TargetID:       req.TargetID,
		AllocatedPaise: req.AllocatedPaise,
		Status:         "ACTIVE",
		CreatedBy:      req.CreatedBy,
	}

	if err := s.paymentRepo.CreateAllocation(ctx, allocation); err != nil {
		return nil, fmt.Errorf("create allocation: %w", err)
	}
	return allocation, nil
}

func (s *paymentEngine) ReverseAllocation(
	ctx context.Context,
	req contracts.ReverseAllocationRequest,
) error {

	allocation, err := s.paymentRepo.GetAllocationByID(ctx, req.AllocationID)
	if err != nil {
		return fmt.Errorf("allocation not found: %w", err)
	}
	if allocation.Status == "REVERSED" {
		return fmt.Errorf("allocation is already reversed")
	}
	return s.paymentRepo.ReverseAllocation(ctx, req.AllocationID, req.By, req.Note)
}

func (s *paymentEngine) GetAllocationsByTarget(
	ctx context.Context,
	targetType string,
	targetID string,
) ([]*models.PaymentAllocation, error) {

	allocations, err := s.paymentRepo.GetAllocationsByTarget(ctx, targetType, targetID)
	if err != nil {
		return nil, fmt.Errorf("get allocations by target: %w", err)
	}
	return allocations, nil
}

func (s *paymentEngine) GetOutstandingAmount(
	ctx context.Context,
	targetType string,
	targetID string,
) (int64, error) {

	total, err := s.paymentRepo.GetTotalAllocatedByTarget(ctx, targetType, targetID)
	if err != nil {
		return 0, fmt.Errorf("get outstanding amount: %w", err)
	}
	return total, nil
}

func isValidAllocationTarget(targetType string) bool {
	switch targetType {
	case payment_constants.AllocationTargetInvoice,
		payment_constants.AllocationTargetAdvance,
		payment_constants.AllocationTargetCreditNote:
		return true
	}
	return false
}
