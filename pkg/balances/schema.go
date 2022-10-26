package balances

type SchemaDepositWithdraw struct {
	Amount int `json:"amount" validate:"required"`
}
