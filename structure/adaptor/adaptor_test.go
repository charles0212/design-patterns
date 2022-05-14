package adaptor

import "testing"

func TestAdaptor(t *testing.T) {
	t.Run("Adapter", func(t *testing.T) {
		client := &client{}
		mac := &mac{}

		client.insertLightningConnectorIntoComputer(mac)

		windowsMachine := &windows{}
		windowsMachineAdapter := &windowsAdapter{
			windowMachine: windowsMachine,
		}
		client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
	})
}
