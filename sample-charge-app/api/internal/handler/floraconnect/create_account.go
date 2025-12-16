package floraconnecthandler

import (
	"context"
	"log/slog"

	accountv1 "github.com/flora/pkg/pb/account/v1"

	"github.com/flora/api/internal/usecase"
)

// CreateAccountUsecase defines the dependency we need from the usecase layer.
type CreateAccountUsecase interface {
	Execute(req *usecase.CreateAccountRequestDTO) (*usecase.CreateAccountResponseDTO, error)
}

type AccountHandler struct {
	uc CreateAccountUsecase
}

func NewAccountHandler(uc CreateAccountUsecase) *AccountHandler {
	return &AccountHandler{uc: uc}
}

// CreateAccount is the gRPC handler for AccountService.CreateAccount.
func (h *AccountHandler) CreateAccount(ctx context.Context, req *accountv1.CreateAccountRequest) (*accountv1.CreateAccountResponse, error) {
	if h == nil || h.uc == nil {
		return nil, context.Canceled
	}

	dtoReq := &usecase.CreateAccountRequestDTO{
		ExternalAccountID: req.GetExternalAccountId(),
		// TODO: Set TenantID from context or header
		TenantID: "",
	}

	res, err := h.uc.Execute(dtoReq)
	if err != nil {
		slog.Error("Failed to create account", "error", err)
		return nil, handleError(err)
	}

	return &accountv1.CreateAccountResponse{
		Account: &accountv1.Account{
			Id:                res.ID,
			ExternalAccountId: req.GetExternalAccountId(),
			TenantId:          dtoReq.TenantID,
		},
	}, nil
}
