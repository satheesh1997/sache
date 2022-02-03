package sache

type (
	HashMap struct {
		Data    map[string]string
		MaxSize int
		Size    int
	}
)

func (storage *HashMap) Insert(key string, value string) {
	storage.Data[key] = value
	storage.Size++
}

func (storage *HashMap) Read(key string) string {
	value := storage.Data[key]
	return value
}

func (storage *HashMap) Remove(key string) {
	delete(storage.Data, key)
	storage.Size--
}

func (storage *HashMap) IsFull() bool {
	return storage.MaxSize == storage.Size
}
