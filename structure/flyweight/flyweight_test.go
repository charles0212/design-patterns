package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	t.Run("flyweight", func(t *testing.T) {
		game := newGame()

		//Add Terrorist
		game.addTerrorist(TerroristDressType)
		game.addTerrorist(TerroristDressType)
		game.addTerrorist(TerroristDressType)
		game.addTerrorist(TerroristDressType)

		//Add CounterTerrorist
		game.addCounterTerrorist(CounterTerroristDressType)
		game.addCounterTerrorist(CounterTerroristDressType)
		game.addCounterTerrorist(CounterTerroristDressType)

		dressFactoryInstance := getDressFactorySingleInstance()

		for dressType, dress := range dressFactoryInstance.dressMap {
			fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
		}
	})
}
