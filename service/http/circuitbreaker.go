package http

import (
	"log"
	"time"

	"github.com/naufal-dean/ccb/lib/circuitbreaker"
	"github.com/naufal-dean/ccb/service/internal"
)

var cbs = make(map[string]*circuitbreaker.CircuitBreaker)

func (s HttpServer) getCircuitBreaker(name string) *circuitbreaker.CircuitBreaker {
	if cb, ok := cbs[name]; ok {
		return cb
	}

	st := circuitbreaker.Settings{
		Name: name,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.75
		},
		Timeout: time.Duration(60) * time.Second,
		OnStateChange: func(name string, from circuitbreaker.State, to circuitbreaker.State) {
			log.Printf("OnChangeState: %s, %v, %v\n", name, from, to)
			// Get affected endpoint (only broadcast new created rd endpoints)
			endpoints, err := s.app.Repositories.RequiredService.GetEndpointsByRdService(name)
			if err != nil {
				return
			}
			// Get affected requiring service
			serviceDepMap, endpointDepMap, err := s.app.Repositories.RequiringService.GetDependencyMapByEndpoints(endpoints)
			if err != nil {
				return
			}
			// Broadcast status
			switch to {
			case circuitbreaker.StateOpen:
				for serviceAddr := range serviceDepMap {
					err := internal.BroadcastOpenCircuits(serviceAddr, serviceDepMap[serviceAddr], endpointDepMap[serviceAddr])
					if err != nil {
						log.Printf("OpenCircuits: error on BroadcastOpenCircuits to %s: %v\n", serviceAddr, err)
					}
				}
			case circuitbreaker.StateClosed:
				for serviceAddr := range serviceDepMap {
					err := internal.BroadcastCloseCircuits(serviceAddr, serviceDepMap[serviceAddr], endpointDepMap[serviceAddr])
					if err != nil {
						log.Printf("CloseCircuits: error on BroadcastCloseCircuits to %s: %v\n", serviceAddr, err)
					}
				}
			}
		},
	}

	cb := circuitbreaker.NewCircuitBreaker(st)
	cbs[name] = cb
	return cb
}
