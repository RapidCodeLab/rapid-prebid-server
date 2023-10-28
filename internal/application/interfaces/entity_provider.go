package interfaces

type EntityProvider interface {
	Provide(entityID string) (Entity, error)
}
