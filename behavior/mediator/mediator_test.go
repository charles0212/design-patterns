package mediator

import "testing"

func TestMediator(t *testing.T) {
	t.Run("mediator", func(t *testing.T) {
		stationManager := newStationManger()

		passengerTrain := &passengerTrain{
			mediator: stationManager,
		}
		freightTrain := &freightTrain{
			mediator: stationManager,
		}

		passengerTrain.arrive()
		freightTrain.arrive()
		passengerTrain.depart()
		freightTrain.depart()
	})
}
