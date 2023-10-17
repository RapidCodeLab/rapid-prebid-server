package interfaces

type InventoryEntityProvider interface {
	ProvideEntity(entityID string) (InventoryEntity, error)
}
