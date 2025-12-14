package repositoryif

import (
	"context"

	"github.com/flora/api/internal/domain/model"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error)
}
