package status

type Repository interface {
	Create(model Status) error
	CreateFromRdServicesAndRdEndpointsSlice(rdServices, rdEndpoints []string, status string) ([]string, []string, error)
	GetByRdServiceAndRdEndpoint(rdService, rdEndpoint string) ([]*Status, error)
	DeleteWhereOneRdServiceAndManyRdEndpointsEqual(rdServices, rdEndpoints []string) ([]string, []string, error)
}
