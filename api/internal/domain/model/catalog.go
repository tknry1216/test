package model

type Catalog struct {
	id       string
	name     string
	tenantID string
}

func NewCatalog(name, tenantID string) *Catalog {
	return &Catalog{
		name:     name,
		tenantID: tenantID,
	}
}

func (c *Catalog) ID() string       { return c.id }
func (c *Catalog) Name() string     { return c.name }
func (c *Catalog) TenantID() string { return c.tenantID }

func ReconstructCatalog(id, name, tenantID string) *Catalog {
	return &Catalog{
		id:       id,
		name:     name,
		tenantID: tenantID,
	}
}
