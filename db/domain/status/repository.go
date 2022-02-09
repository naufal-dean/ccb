package status

type Repository interface {
	Create(model Status)
	GetById(id int) *Status
	GetByServiceAndEndpoint(service, endpoint string) []*Status
}
