package storage

type Storage interface {
	Write(dir string, fileName string, data []byte) error
}
