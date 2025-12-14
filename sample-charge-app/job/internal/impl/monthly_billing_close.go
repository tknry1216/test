package impl

type monthlyBillingCloseImpl struct {
}

func NewMonthlyBillingCloseImpl() *monthlyBillingCloseImpl {
	return &monthlyBillingCloseImpl{}
}

func (impl *monthlyBillingCloseImpl) Execute() error {
	return nil
}
