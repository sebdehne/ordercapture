package persistence

import (
	"encoding/json"
	"github.com/sebdehne/ordercapture/domain/v1"
)

type KeyValueStore interface {
	// If the expectedVersion matches (0 for non-existing), then this function stores the value in the map
	//
	// "ok" indicates whether the expectedVersion matched and the value was updated or not
	// "err" is non-nil if there was some I/O error, in which case "ok" is always false
	Store(key string, schemaType string, schemaVersion int, value []byte, expectedVersion int64) (ok bool, err error)

	// Deletes a record as long as the expectedVersion matches.
	//
	// "ok" indicates whether the expectedVersion matched and the value was deleted or not
	// "err" is non-nil if there was some I/O error, in which case "ok" is always false
	Delete(key string, schemaType string, expectedVersion int64) (ok bool, err error)

	// Gets a value from the map
	//
	// value is nil if no record existed. "err" is non-nil if there was some I/O error
	Get(key string, schemaType string) (value []byte, version int64, schemaVersion int, err error)
}

type PersistenceService struct {
	keyValueStore KeyValueStore
}

func New(kvs KeyValueStore) *PersistenceService {
	return &PersistenceService{kvs}
}

var schemaTypeOrderDraft string = "orderdraft"

func (p *PersistenceService) StoreOrder(key string, o v1.OrderDraft, expectedVersion int64) (ok bool, err error) {
	b, err := json.Marshal(o)
	if (err != nil) {
		return false, err
	}
	return p.keyValueStore.Store(key, schemaTypeOrderDraft, 1, b, expectedVersion)
}

func (p *PersistenceService) GetOrder(key string) (orderDraft *v1.OrderDraft, version int64, err error) {
	b, version, schemaVersion, err := p.keyValueStore.Get(key, schemaTypeOrderDraft)

	if err != nil || b == nil {
		return nil, 0, err
	}

	if schemaVersion != 1 {
		panic("Record in database is not at expected schemaVersion")
	}

	o := new(v1.OrderDraft)
	err = json.Unmarshal(b, o)
	if (err != nil) {
		return nil, 0, err
	}
	return o, version, nil
}

func (p *PersistenceService) DeleteOrder(key string, expectedVersion int64) (bool, error) {
	return p.keyValueStore.Delete(schemaTypeOrderDraft, key, expectedVersion)
}
