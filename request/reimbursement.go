package request

type UpsertReimbursement struct {
	ID uint64 `json:"id"`

	Description string `json:"description" validate:"required"`
	Amount      uint64 `json:"amount" validate:"required"`
}
