package inventoryapiboltdb

import (
	"context"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/boltdb/bolt"
)

const (
	inventoryBucketName = "inventories"
	entityBucketName    = "entities"
)

type Service struct {
	db     *bolt.DB
	logger interfaces.Logger
}

func New(
	ctx context.Context,
	path string,
	l interfaces.Logger,
) (*Service, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		err := db.Close()
		if err != nil {
			l.Errorf("db close", "err", err.Error())
		}
	}()

	s := &Service{
		db:     db,
		logger: l,
	}

	return s, s.initBuckets()
}

func (i *Service) initBuckets() error {
	err := i.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(inventoryBucketName))
		return err
	})
	if err != nil {
		return err
	}

	return i.db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(entityBucketName))
		return err
	})
}
