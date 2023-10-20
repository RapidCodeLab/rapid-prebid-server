package interfaces

type InventoryEntityProvider interface {
	ProvideEntity(entityID string) (Entity, error)
}
