package model

import "time"

type BillingType string

const (
	BillingTypeOneShot   BillingType = "ONE_SHOT"
	BillingTypeRecurring BillingType = "RECURRING"
)

type BillingCycle string

const (
	BillingCycleMonthly   BillingCycle = "MONTHLY"
	BillingCycleYearly    BillingCycle = "YEARLY"
	BillingCycleQuarterly BillingCycle = "QUARTERLY"
)

type CatalogPrice struct {
	id           string
	catalogID    string
	amount       int64
	billingType  BillingType
	billingCycle *BillingCycle // nullable
	startDate    time.Time
	endDate      *time.Time // nullable
}

func NewCatalogPrice(
	catalogID string,
	amount int64,
	billingType BillingType,
	billingCycle *BillingCycle,
	startDate time.Time,
	endDate *time.Time,
) *CatalogPrice {
	return &CatalogPrice{
		catalogID:    catalogID,
		amount:       amount,
		billingType:  billingType,
		billingCycle: billingCycle,
		startDate:    startDate,
		endDate:      endDate,
	}
}

func (cp *CatalogPrice) ID() string                  { return cp.id }
func (cp *CatalogPrice) CatalogID() string           { return cp.catalogID }
func (cp *CatalogPrice) Amount() int64               { return cp.amount }
func (cp *CatalogPrice) BillingType() BillingType    { return cp.billingType }
func (cp *CatalogPrice) BillingCycle() *BillingCycle { return cp.billingCycle }
func (cp *CatalogPrice) StartDate() time.Time        { return cp.startDate }
func (cp *CatalogPrice) EndDate() *time.Time         { return cp.endDate }

func ReconstructCatalogPrice(
	id string,
	catalogID string,
	amount int64,
	billingType BillingType,
	billingCycle *BillingCycle,
	startDate time.Time,
	endDate *time.Time,
) *CatalogPrice {
	return &CatalogPrice{
		id:           id,
		catalogID:    catalogID,
		amount:       amount,
		billingType:  billingType,
		billingCycle: billingCycle,
		startDate:    startDate,
		endDate:      endDate,
	}
}
