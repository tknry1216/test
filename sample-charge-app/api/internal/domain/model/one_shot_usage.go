package model

import "time"

type OneShotUsage struct {
	id             string
	accountID      string
	catalogID      string
	idempotencyKey string
	usedAt         time.Time
}

func NewOneShotUsage(
	accountID string,
	catalogID string,
	idempotencyKey string,
	usedAt time.Time,
) *OneShotUsage {
	return &OneShotUsage{
		accountID:      accountID,
		catalogID:      catalogID,
		idempotencyKey: idempotencyKey,
		usedAt:         usedAt,
	}
}

func (o *OneShotUsage) ID() string             { return o.id }
func (o *OneShotUsage) AccountID() string      { return o.accountID }
func (o *OneShotUsage) CatalogID() string      { return o.catalogID }
func (o *OneShotUsage) IdempotencyKey() string { return o.idempotencyKey }
func (o *OneShotUsage) UsedAt() time.Time      { return o.usedAt }

func ReconstructOneShotUsage(
	id string,
	accountID string,
	catalogID string,
	idempotencyKey string,
	usedAt time.Time,
) *OneShotUsage {
	return &OneShotUsage{
		id:             id,
		accountID:      accountID,
		catalogID:      catalogID,
		idempotencyKey: idempotencyKey,
		usedAt:         usedAt,
	}
}
