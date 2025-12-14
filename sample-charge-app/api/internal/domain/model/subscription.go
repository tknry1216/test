package model

import "time"

type Subscription struct {
	id              string
	accountID       string
	catalogID       string
	startDate       time.Time
	endDate         *time.Time // nullable
	idempotencyKey  string
	nextBillingDate time.Time
}

func NewSubscription(
	accountID string,
	catalogID string,
	startDate time.Time,
	endDate *time.Time,
	idempotencyKey string,
	nextBillingDate time.Time,
) *Subscription {
	return &Subscription{
		accountID:       accountID,
		catalogID:       catalogID,
		startDate:       startDate,
		endDate:         endDate,
		idempotencyKey:  idempotencyKey,
		nextBillingDate: nextBillingDate,
	}
}

func (s *Subscription) ID() string              { return s.id }
func (s *Subscription) AccountID() string       { return s.accountID }
func (s *Subscription) CatalogID() string       { return s.catalogID }
func (s *Subscription) StartDate() time.Time    { return s.startDate }
func (s *Subscription) EndDate() *time.Time     { return s.endDate }
func (s *Subscription) IdempotencyKey() string  { return s.idempotencyKey }
func (s *Subscription) NextBillingDate() time.Time { return s.nextBillingDate }

func ReconstructSubscription(
	id string,
	accountID string,
	catalogID string,
	startDate time.Time,
	endDate *time.Time,
	idempotencyKey string,
	nextBillingDate time.Time,
) *Subscription {
	return &Subscription{
		id:              id,
		accountID:       accountID,
		catalogID:       catalogID,
		startDate:       startDate,
		endDate:         endDate,
		idempotencyKey:  idempotencyKey,
		nextBillingDate: nextBillingDate,
	}
}

