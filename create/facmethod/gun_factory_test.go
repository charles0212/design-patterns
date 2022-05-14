package facmethod

import (
	"fmt"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	t.Run("FactoryMethod", func(t *testing.T) {
		nak47, _ := getGun("ak47")
		musket, _ := getGun("musket")

		printDetails(nak47)
		printDetails(musket)
	})
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
