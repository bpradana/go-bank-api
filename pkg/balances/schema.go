package balances

type SchemaDeposit struct {
	Amount int `json:"amount" validate:"required,min=1"`
}
