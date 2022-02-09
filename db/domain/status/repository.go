package status

type Repository interface {
	Create(model Status)
	CreateFromOneServiceAndManyEndpoints(service string, endpoints []string, status string)
	GetById(id int) *Status
	GetByServiceAndEndpoint(service, endpoint string) []*Status
}
