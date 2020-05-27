package internal

//DataStore ...
type DataStore interface {
	PutObject(data string) string
	GetObject(id string) string
}
