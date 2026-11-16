package idea

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type StateFn func(t *TrafficLight) StateFn
type Event int

const (
	PedestrianPushButton Event = iota
)

type TrafficLight struct {
	Ev               chan Event
	RDur, YDur, GDur time.Duration
}

func (t *TrafficLight) Run() error {
	state := redState
	for state != nil {
		ptr := reflect.ValueOf(state).Pointer()
		statefn := runtime.FuncForPC(ptr)
		if statefn == nil {
			panic("unreachable")
		}
		fmt.Println("Light:", statefn.Name())
		state = state(t)
	}
	return nil
}

func redState(t *TrafficLight) StateFn {
	<-time.After(t.RDur)
	return greenState
}
func yellowState(t *TrafficLight) StateFn {
	<-time.After(t.YDur)
	return redState
}
func greenState(t *TrafficLight) StateFn {
	dur := time.After(t.GDur)
	for {
		select {
		case ev, ok := <-t.Ev:
			if !ok {
				return nil
			}
			switch ev {
			case PedestrianPushButton:
				return yellowState
			}
		case <-dur:
			return yellowState
		}
	}
}
