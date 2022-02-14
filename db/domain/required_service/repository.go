package required_service

type Repository interface {
	Create(model RequiredService)
	GetEndpointsByRdService(rdService string) (endpoints []string)
	GetEndpointsByRdServiceAndRdEndpoints(rdService string, rdEndpoints []string) (endpoints []string)
}
