package status

type Repository interface {
	Create(model Status)
	CreateFromOneServiceAndManyEndpoints(service string, endpoints []string, status string)
	GetByServiceAndEndpoint(service, endpoint string) []*Status
	DeleteWhereOneServiceAndManyEndpointsEqual(service string, endpoints []string)
}
