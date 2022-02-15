package status

type Repository interface {
	Create(model Status) error
	CreateFromOneRdServiceAndManyRdEndpoints(rdService string, rdEndpoints []string, status string) ([]string, error)
	GetByRdServiceAndRdEndpoint(rdService, rdEndpoint string) ([]*Status, error)
	DeleteWhereOneRdServiceAndManyRdEndpointsEqual(rdService string, rdEndpoints []string) ([]string, error)
}
