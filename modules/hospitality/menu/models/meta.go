package hospitalityMenu_models

import "github.com/google/uuid"

type MenuItemKitchenRegister struct {
	ID         uuid.UUID //
	MenuItemID uuid.UUID // unique and index
	KitchenID  uuid.UUID
	OwnersID   uuid.UUID
}
