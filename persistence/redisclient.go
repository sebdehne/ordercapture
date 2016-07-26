package persistence

type RedisClient struct {

}

func (r *RedisClient) Store(key string, schemaType string, schemaVersion int, value []byte, expectedVersion int64) (ok bool, err error) {
	panic("Not implemented")
}

func (r *RedisClient) Delete(key string, schemaType string, expectedVersion int64) (ok bool, err error) {
	panic("Not implemented")
}

func (r *RedisClient) Get(key string, schemaType string) (value []byte, version int64, schemaVersion int, err error) {
	panic("Not implemented")
}