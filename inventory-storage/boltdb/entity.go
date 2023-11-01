package inventorystorage_boltdb

import (
	"encoding/json"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/boltdb/bolt"
)

func (i *Service) ReadAllEntities() ([]interfaces.Entity, error) {
	res := []interfaces.Entity{}

	err := i.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(entityBucketName)).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			entity := interfaces.Entity{}
			err := json.Unmarshal(v, &entity)
			if err != nil {
				i.logger.Errorf("entity unmarshal", "err", err.Error())
				continue
			}
			res = append(res, entity)
		}

		return nil
	})

	return res, err
}

func (i *Service) CreateEntity(entity interfaces.Entity) error {
	return i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(entityBucketName))

		data, err := json.Marshal(entity)
		if err != nil {
			return err
		}
		return b.Put([]byte(entity.ID), data)
	})
}

func (i *Service) ReadEntity(ID string) (interfaces.Entity, error) {
	entity := interfaces.Entity{}

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(entityBucketName))
		data := b.Get([]byte(ID))
		return json.Unmarshal(data, &entity)
	})

	return entity, err
}

func (i *Service) UpdateEntity(entity interfaces.Entity) error {
	return i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(entityBucketName))

		data, err := json.Marshal(entity)
		if err != nil {
			return err
		}
		return b.Put([]byte(entity.ID), data)
	})
}

func (i *Service) DeleteEntity(ID string) error {
	return i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(entityBucketName))
		return b.Delete([]byte(ID))
	})
}
