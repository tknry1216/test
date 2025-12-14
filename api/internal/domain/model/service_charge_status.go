package model

type ServiceChargeStatusType string

const (
	ServiceChargeStatusReserved  ServiceChargeStatusType = "RESERVED"
	ServiceChargeStatusCompleted ServiceChargeStatusType = "COMPLETED"
	ServiceChargeStatusFailed    ServiceChargeStatusType = "FAILED"
	ServiceChargeStatusCancelled ServiceChargeStatusType = "CANCELLED"
)

type ServiceChargeStatus struct {
	id              string
	name            string
	status          ServiceChargeStatusType
	serviceChargeID string
}

func NewServiceChargeStatus(
	name string,
	status ServiceChargeStatusType,
	serviceChargeID string,
) *ServiceChargeStatus {
	return &ServiceChargeStatus{
		name:            name,
		status:          status,
		serviceChargeID: serviceChargeID,
	}
}

func (scs *ServiceChargeStatus) ID() string                      { return scs.id }
func (scs *ServiceChargeStatus) Name() string                    { return scs.name }
func (scs *ServiceChargeStatus) Status() ServiceChargeStatusType { return scs.status }
func (scs *ServiceChargeStatus) ServiceChargeID() string         { return scs.serviceChargeID }

func ReconstructServiceChargeStatus(
	id string,
	name string,
	status ServiceChargeStatusType,
	serviceChargeID string,
) *ServiceChargeStatus {
	return &ServiceChargeStatus{
		id:              id,
		name:            name,
		status:          status,
		serviceChargeID: serviceChargeID,
	}
}
