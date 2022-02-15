package required_service

type Repository interface {
	Create(model RequiredService) error
	GetEndpointsByRdService(rdService string) ([]string, error)
	GetEndpointsByRdServiceAndRdEndpoints(rdService string, rdEndpoints []string) ([]string, error)
}
