package auth

import (
	"context"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (b *base) Register(ctx context.Context) error {
	now := time.Now()

	for i := 0; i <= 100; i++ {
		password, err := bcrypt.GenerateFromPassword([]byte("user123"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		_, err = b.repository.CreateUser(ctx, &domain.User{
			CreatedAt:  now,
			UpdatedAt:  now,
			Username:   fmt.Sprintf("user%d", i),
			FullName:   fmt.Sprintf("User %d", i),
			Password:   string(password),
			Role:       enums.USERRole,
			BaseSalary: 5000000,
		})
		if err != nil {
			return err
		}
	}

	password, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = b.repository.CreateUser(ctx, &domain.User{
		CreatedAt:  now,
		UpdatedAt:  now,
		Username:   "admin",
		FullName:   "Administrator",
		Password:   string(password),
		Role:       enums.ADMINRole,
		BaseSalary: 7500000,
	})
	if err != nil {
		return err
	}

	return nil
}
