package hw

type Map interface {
	Get(string) (interface{}, bool)
	Set(string, interface{})
	Del(string)
}

type kvPair struct {
	k string
	v interface{}
}

func NewMap() Map {
	return &MapImpl{}
}

type MapImpl struct {
}

func (m *MapImpl) Get(key string) (interface{}, bool) {
	return nil, false
}

func (m *MapImpl) Set(key string, val interface{}) {

}

func (m *MapImpl) Del(key string) {

}

// djb2 hash function
// https://stackoverflow.com/a/7666577
func stringHash(str string) uint {
	hash := uint(5381)

	for _, v := range str {
		hash = uint(v) + ((hash << uint(5)) + hash)
	}

	return hash
}
