package http

import (
	"log"
	"time"

	"github.com/naufal-dean/ccb/lib/circuitbreaker"
)

var cbs = make(map[string]*circuitbreaker.CircuitBreaker)

func getCircuitBreaker(name string) *circuitbreaker.CircuitBreaker {
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
		},
	}

	cb := circuitbreaker.NewCircuitBreaker(st)
	cbs[name] = cb
	return cb
}
