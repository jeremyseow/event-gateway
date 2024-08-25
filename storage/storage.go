package storage

type Storage interface {
	Write(data []byte) (n int, err error)
}
