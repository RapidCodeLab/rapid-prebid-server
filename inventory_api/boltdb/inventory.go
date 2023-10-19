package inventoryapiboltdb

import (
	"encoding/json"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/boltdb/bolt"
)

func (i *Service) ReadAllInventories() ([]interfaces.Inventory, error) {
	res := []interfaces.Inventory{}

	err := i.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(inventoryBucketName)).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			inventory := interfaces.Inventory{}
			err := json.Unmarshal(v, &inventory)
			if err != nil {
				i.logger.Errorf("inventory unmarshal", "err", err.Error())
				continue
			}
			res = append(res, inventory)
		}

		return nil
	})

	return res, err
}

func (i *Service) CreateInventory(inv interfaces.Inventory) error {
	return i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(inventoryBucketName))

		data, err := json.Marshal(inv)
		if err != nil {
			return err
		}
		return b.Put([]byte(inv.ID), data)
	})
}

func (i *Service) ReadInventory(ID string) (interfaces.Inventory, error) {
	inv := interfaces.Inventory{}

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(inventoryBucketName))
		data := b.Get([]byte(ID))
		return json.Unmarshal(data, &inv)
	})

	return inv, err
}

func (i *Service) UpdateInventory(inv interfaces.Inventory) error {
	return i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(inventoryBucketName))

		data, err := json.Marshal(inv)
		if err != nil {
			return err
		}
		return b.Put([]byte(inv.ID), data)
	})
}

func (i *Service) DeleteInventory(inv interfaces.Inventory) error {
	return i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(inventoryBucketName))
		return b.Delete([]byte(inv.ID))
	})
}
