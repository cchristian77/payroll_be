package response

import (
	"github.com/cchristian77/payroll_be/domain"
	"time"
)

type Reimbursement struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Description string `json:"description"`
	Amount      uint64 `json:"amount"`
	Status      string `json:"status"`
}

func NewReimbursementFromDomain(r *domain.Reimbursement) *Reimbursement {
	if r == nil {
		return nil
	}

	return &Reimbursement{
		ID:          r.ID,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		Description: r.Description,
		Amount:      r.Amount,
		Status:      r.Status,
	}
}
