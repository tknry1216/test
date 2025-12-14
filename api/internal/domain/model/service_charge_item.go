package model

type ServiceChargeItem struct {
	id              string
	name            string
	amount          int64
	serviceChargeID string
	oneShotUsageID  *string // nullable
	subscriptionID  *string // nullable
}

func NewServiceChargeItem(
	name string,
	amount int64,
	serviceChargeID string,
	oneShotUsageID *string,
	subscriptionID *string,
) *ServiceChargeItem {
	return &ServiceChargeItem{
		name:            name,
		amount:          amount,
		serviceChargeID: serviceChargeID,
		oneShotUsageID:  oneShotUsageID,
		subscriptionID:  subscriptionID,
	}
}

func (sci *ServiceChargeItem) ID() string              { return sci.id }
func (sci *ServiceChargeItem) Name() string            { return sci.name }
func (sci *ServiceChargeItem) Amount() int64           { return sci.amount }
func (sci *ServiceChargeItem) ServiceChargeID() string { return sci.serviceChargeID }
func (sci *ServiceChargeItem) OneShotUsageID() *string { return sci.oneShotUsageID }
func (sci *ServiceChargeItem) SubscriptionID() *string { return sci.subscriptionID }

func ReconstructServiceChargeItem(
	id string,
	name string,
	amount int64,
	serviceChargeID string,
	oneShotUsageID *string,
	subscriptionID *string,
) *ServiceChargeItem {
	return &ServiceChargeItem{
		id:              id,
		name:            name,
		amount:          amount,
		serviceChargeID: serviceChargeID,
		oneShotUsageID:  oneShotUsageID,
		subscriptionID:  subscriptionID,
	}
}
