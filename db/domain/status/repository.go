package status

type Repository interface {
	Create(model Status) error
	CreateFromRdServicesAndRdEndpointsSlice(rdServices, rdEndpoints []string, status string, expiry int64) ([]string, []string, error)
	GetByRdServiceAndRdEndpoint(rdService, rdEndpoint string) ([]*Status, error)
	DeleteWhereRdServicesAndRdEndpointsInSlice(rdServices, rdEndpoints []string) ([]string, []string, error)
}
