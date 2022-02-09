package requiring_service

type Repository interface {
	Create(model RequiringService)
	GetServicesByEndpoint(endpoint string) (services []string)
}
