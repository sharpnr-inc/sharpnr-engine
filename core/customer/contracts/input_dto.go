package org_customer_contracts

type AddCustomerRequestDto struct {
	Name   *string `json:"name"`                      // Pointer allows nil/null
	Email  *string `json:"email"`                     // Pointer allows nil/null
	Mobile string  `json:"mobile" binding:"required"` // Non-pointer + required
}
type UpdateCustomerRequestDto struct {
	Name   *string `json:"name"`
	Email  *string `json:"email"`
	Mobile *string `json:"mobile"`
}
