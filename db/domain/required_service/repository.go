package required_service

type Repository interface {
	Create(model RequiredService) error
	GetEndpointsByRdService(rdService string) ([]string, error)
	GetEndpointsByRdServicesAndRdEndpoints(rdServices, rdEndpoints []string) ([]string, error)
}
