package interfaces

type InventoryProvider interface {
	Provide(uuid string) ([]byte, error)
}
