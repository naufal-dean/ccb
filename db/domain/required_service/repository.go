package required_service

type Repository interface {
	Create(model RequiredService) error
	GetEndpointsByRdService(rdService string) ([]string, error)
	GetEndpointsByRdServiceAndRdEndpoints(rdServices, rdEndpoints []string) ([]string, error)
}
