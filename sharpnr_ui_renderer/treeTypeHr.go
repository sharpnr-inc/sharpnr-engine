package sharpnruirenderer

type SharpHRTree struct {
	HasAccess              bool                   `json:"has_access"`
	SharpHospitalityModule SharpHospitalityModule `json:"sub_module_section"`
}
