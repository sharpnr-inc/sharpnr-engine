package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sharpnr-inc/sharpnr-engine/core/interfaces"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
)

type terminalService struct {
	repo interfaces.TerminalRepository
}

func NewTerminalService(repo interfaces.TerminalRepository) *terminalService {
	return &terminalService{repo: repo}
}

func (s *terminalService) InitializeDefaultAccountingTerminals(orgUUID uuid.UUID, deptID uuid.UUID) error {
	// Create a default terminal for the Accounting department
	terminal := &models.Terminal{
		ID:             uuid.New(), // Corrected field name
		OrganizationID: orgUUID,
		Name:           "Accounting Terminal 01",
		Description:    "Default terminal for accounting operations",
		// Add other necessary fields for a terminal
	}

	if err := s.repo.Create(terminal); err != nil {
		return fmt.Errorf("failed to create default accounting terminal in repo: %w", err)
	}

	fmt.Println("Default Accounting Terminal initialized.")
	return nil
}
