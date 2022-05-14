package facade

import (
	"fmt"
	"log"
	"testing"
)

func TestFacade(t *testing.T) {
	t.Run("facade", func(t *testing.T) {
		fmt.Println()
		walletFacade := newWalletFacade("abc", 1234)
		fmt.Println()

		err := walletFacade.addMoneyToWallet("abc", 1234, 10)
		if err != nil {
			log.Fatalf("Error: %s\n", err.Error())
		}

		fmt.Println()
		err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
		if err != nil {
			log.Fatalf("Error: %s\n", err.Error())
		}
	})
}
