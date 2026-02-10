package accountingdtos

type GetAccountBooksOfOrg_Response_dto struct {
	AccountBooks []AccountingBook_Response_dto `json:"account_books"`
}

type AccountingBook_Response_dto struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Module *string `json:"module,omitempty"`

	SubModule *string `json:"sub_module,omitempty"`

	Status string `json:"status"`
}
