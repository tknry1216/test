package model

type Account struct {
	id                string
	externalAccountID string
	tenantID          string
}

func NewAccount(externalAccountID, tenantID string) *Account {
	return &Account{
		externalAccountID: externalAccountID,
		tenantID:          tenantID,
	}
}

func (a *Account) ID() string                { return a.id }
func (a *Account) ExternalAccountID() string { return a.externalAccountID }
func (a *Account) TenantID() string          { return a.tenantID }

func ReconstructAccount(id, externalAccountID, tenantID string) *Account {
	return &Account{
		id:                id,
		externalAccountID: externalAccountID,
		tenantID:          tenantID,
	}
}
