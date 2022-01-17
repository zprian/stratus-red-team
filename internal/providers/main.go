package providers

type StratusProvider interface {
	GetAwsProvider() AwsProvider
}

type stratusProviderImpl struct {
	awsProvider AwsProvider
}

func (m *stratusProviderImpl) GetAwsProvider() AwsProvider {
	if m.awsProvider == nil {
		m.awsProvider = NewAwsProvider()
	}
	return m.awsProvider
}

func NewStratusProvider() StratusProvider {
	return &stratusProviderImpl{}
}
