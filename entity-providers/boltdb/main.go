package boltdb_entity_provider

import "github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"

type provider struct {
	storage interfaces.EntityReadStorager
}

func New(
	s interfaces.EntityReadStorager,
) interfaces.EntityProvider {
	return &provider{
		storage: s,
	}
}

func (i *provider) Provide(
	entityID string,
) (interfaces.Entity, error) {
	return i.storage.ReadEntity(entityID)
}
