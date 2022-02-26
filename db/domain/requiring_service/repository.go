package requiring_service

type Repository interface {
	Create(model RequiringService) error
	GetRgServicesByEndpoint(endpoint string) ([]string, error)
	GetDependencyMapByEndpoints(endpoints []string) (map[string][]string, map[string][]string, error)
}
