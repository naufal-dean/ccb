package status

type Repository interface {
	Create(model Status)
	CreateFromOneRdServiceAndManyRdEndpoints(rdService string, rdEndpoints []string, status string)
	GetByRdServiceAndRdEndpoint(rdService, rdEndpoint string) []*Status
	DeleteWhereOneRdServiceAndManyRdEndpointsEqual(rdService string, rdEndpoints []string)
}
