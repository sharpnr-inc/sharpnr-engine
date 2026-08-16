package stockwaveEngineDtos

type NewInventoryAccountRequestDto struct {
	Name     string `json:"name"`
	Type     string `json:"type"` // Likely "KITCHEN", "BAR", etc.
	IsShared bool   `json:"is_shared"`
}
