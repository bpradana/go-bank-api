package userBalances

type SchemaDepositWithdraw struct {
	Amount int `json:"amount" validate:"required,min=1"`
}

type SchemaTransfer struct {
	Amount   int    `json:"amount" validate:"required,min=1"`
	Username string `json:"username" validate:"required"`
}
