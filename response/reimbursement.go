package response

import "time"

type Reimbursement struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Description string `json:"description"`
	Amount      uint64 `json:"amount"`
	Status      string `json:"status"`
}
