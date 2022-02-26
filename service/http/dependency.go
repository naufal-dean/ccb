package http

import (
	"github.com/naufal-dean/ccb/db/domain/required_service"
	"github.com/naufal-dean/ccb/db/domain/requiring_service"
)

func buildDependency(requiredService, requiredEndpoint string, initialHeader map[string]string) (*requiring_service.RequiringService, *required_service.RequiredService, bool) {
	// Parse requiring service data
	// If any of metadata header not exists, then no need to store dependency tree as it will not be used
	requiringService, ok := initialHeader[requiringServiceHeader]
	if !ok {
		return nil, nil, false
	}
	currentService, ok := initialHeader[currentServiceHeader]
	if !ok {
		return nil, nil, false
	}
	currentEndpoint, ok := initialHeader[currentEndpointHeader]
	if !ok {
		return nil, nil, false
	}

	// Create models
	requiringServiceModel := &requiring_service.RequiringService{
		RgService: requiringService,
		Service:   currentService,
		Endpoint:  currentEndpoint,
	}
	requiredServiceModel := &required_service.RequiredService{
		Endpoint:   currentEndpoint,
		RdService:  requiredService,
		RdEndpoint: requiredEndpoint,
	}
	return requiringServiceModel, requiredServiceModel, true
}

// NOTE:
// requiring service
// source service - endpoint
//
// required service
// endpoint - target service + endpoint
//
// service down -> detect curr endpoint that use the service -> list all service that requiring the endpoint
// broadcast to requiring service: current service, list opened endpoint
// if there are requiring service, then there are X-Endpoint and X-Method header, then current endpoint and method will be stored in db
// else (no requiring service, or direct from client) current method and method is unused, just use random placeholder (chosen: not storing the dependency at all)
