package persistence

import "sync"

type InMemClient struct {
	data map[string]record
	lock sync.RWMutex
}

type record struct {
	version       int64
	schemaVersion int
	value         []byte
}

func NewInMemClient() *InMemClient {
	return &InMemClient{data:make(map[string]record)}
}

func (r *InMemClient) Get(key string, schemaType string) (value []byte, version int64, schemaVersion int, err error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	if er, ok := r.data[schemaType + ":" + key]; ok {
		return er.value, er.version, er.schemaVersion, nil
	} else {
		return nil, 0, 0, nil
	}
}

func (r *InMemClient) Store(key string, schemaType string, schemaVersion int, value []byte, expectedVersion int64) (ok bool, err error) {
	newRecord := record{
		version:expectedVersion + 1,
		schemaVersion:schemaVersion,
		value:value}

	r.lock.Lock()
	defer r.lock.Unlock()

	if existingRecord, ok := r.data[schemaType + ":" + key]; !ok && expectedVersion == 0 {
		r.data[schemaType + ":" + key] = newRecord
		return true, nil
	} else if existingRecord.version == expectedVersion {
		r.data[schemaType + ":" + key] = newRecord
		return true, nil
	}
	return false, nil
}

func (r *InMemClient) Delete(schemaType string, key string, expectedVersion int64) (ok bool, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if er, ok := r.data[schemaType + ":" + key]; ok && er.version == expectedVersion {
		delete(r.data, schemaType + ":" + key)
		return true, nil
	}
	return false, nil
}


