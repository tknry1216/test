package usecase

import (
	"time"

	"github.com/flora/api/internal/domain/model"
)

type reserveServiceChargeUsecase struct{}

type ReserveServiceChargeRequestDTO struct {
	AccountID string
	StartDate time.Time
	EndDate   time.Time
}

type ReserveServiceChargeResponseDTO struct {
	ServiceChargeID string
}

func NewReserveServiceChargeUsecase() *reserveServiceChargeUsecase {
	return &reserveServiceChargeUsecase{}
}

func (uc *reserveServiceChargeUsecase) Execute(request *ReserveServiceChargeRequestDTO) (*ReserveServiceChargeResponseDTO, error) {
	usages, err := listTargetUsages()
	if err != nil {
		return nil, err
	}

	errCh := make(chan error, len(usages))
	doneCh := make(chan struct{})
	for _, usage := range usages {
		go func(u *model.OneShotUsage) {
			serviceCharge := convertOneShotUsageToServiceCharge(u)
			if err := createServiceCharge(serviceCharge); err != nil {
				errCh <- err
				return
			}
			errCh <- nil
		}(usage)
	}

	go func() {
		for i := 0; i < len(usages); i++ {
			if err := <-errCh; err != nil {
				break
			}
		}
		close(doneCh)
	}()

	select {
	case <-doneCh:
		// All finished
	case err := <-errCh:
		if err != nil {
			return nil, err
		}
	}

	return &ReserveServiceChargeResponseDTO{ServiceChargeID: "123"}, nil
}

func listTargetUsages() ([]*model.OneShotUsage, error) {
	return nil, nil
}

// convOneShotUsageToServiceCharge
func convertOneShotUsageToServiceCharge(usage *model.OneShotUsage) *model.ServiceCharge {
	return &model.ServiceCharge{
		// TODO
	}
}

func createServiceCharge(serviceCharge *model.ServiceCharge) error {
	return nil
}
