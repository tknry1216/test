package model

import "time"

type ServiceCharge struct {
	id             string
	accountID      string
	startDate      time.Time
	endDate        time.Time
	amount         int64
	latestStatusID *string // nullable
}

func NewServiceCharge(
	accountID string,
	startDate time.Time,
	endDate time.Time,
	amount int64,
) *ServiceCharge {
	return &ServiceCharge{
		accountID: accountID,
		startDate: startDate,
		endDate:   endDate,
		amount:    amount,
	}
}

func (sc *ServiceCharge) ID() string              { return sc.id }
func (sc *ServiceCharge) AccountID() string       { return sc.accountID }
func (sc *ServiceCharge) StartDate() time.Time    { return sc.startDate }
func (sc *ServiceCharge) EndDate() time.Time      { return sc.endDate }
func (sc *ServiceCharge) Amount() int64           { return sc.amount }
func (sc *ServiceCharge) LatestStatusID() *string { return sc.latestStatusID }

func (sc *ServiceCharge) UpdateLatestStatusID(statusID string) {
	sc.latestStatusID = &statusID
}

func ReconstructServiceCharge(
	id string,
	accountID string,
	startDate time.Time,
	endDate time.Time,
	amount int64,
	latestStatusID *string,
) *ServiceCharge {
	return &ServiceCharge{
		id:             id,
		accountID:      accountID,
		startDate:      startDate,
		endDate:        endDate,
		amount:         amount,
		latestStatusID: latestStatusID,
	}
}
