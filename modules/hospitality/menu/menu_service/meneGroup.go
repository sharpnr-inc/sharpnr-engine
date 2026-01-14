package menuservice

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

// CreateMenuGroup implements MenuService.
func (m *menuService) CreateMenuGroup(ctx context.Context, group *hospitalityMenu_models.MenuGroup) error {
	panic("unimplemented")
}

// GetAllMenuGroups implements MenuService.
func (m *menuService) GetAllMenuGroups(ctx context.Context, orgID uuid.UUID) ([]hospitalitydtos.Menu_GetAllMenuGroup_RESPONSE, error) {
	// Get models from repo
	groups, err := m.menuRepo.GetAllMenuGroups(ctx, orgID)
	if err != nil {
		return nil, err
	}

	// Map model → DTO
	result := make([]hospitalitydtos.Menu_GetAllMenuGroup_RESPONSE, 0, len(groups))

	for _, g := range groups {
		result = append(result, hospitalitydtos.Menu_GetAllMenuGroup_RESPONSE{
			Id:          g.ID,
			Name:        g.Name,
			Type:        g.Type,
			Description: g.Description,
		})
	}

	return result, nil
}
