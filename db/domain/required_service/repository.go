package required_service

type Repository interface {
	Create(model RequiredService)
	GetEndpointsByService(service string) (endpoints []string)
}
