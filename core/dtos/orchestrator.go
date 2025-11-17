package dto

/*
	TODO So we need a Generic and Org Response based on type
*/

type OrchestratorDetails[T any] struct {
	Data T `json:"data"`
}

type OrchestratorOrg_Hospitality_Details_Response struct {
}
