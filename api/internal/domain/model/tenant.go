package model

type Tenant struct {
	id   string
	name string
}

func NewTenant(name string) *Tenant {
	return &Tenant{
		name: name,
	}
}

func (t *Tenant) ID() string   { return t.id }
func (t *Tenant) Name() string { return t.name }

func ReconstructTenant(id, name string) *Tenant {
	return &Tenant{
		id:   id,
		name: name,
	}
}

