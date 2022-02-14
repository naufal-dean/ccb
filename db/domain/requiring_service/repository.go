package requiring_service

type Repository interface {
	Create(model RequiringService)
	GetRgServicesByEndpoint(endpoint string) (rgServices []string)
}
