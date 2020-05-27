package internal

//DiskStore ...
type DiskStore struct {
	Status string
	Val    int
}

//PutObject ...
func (ds *DiskStore) PutObject(k string) string {
	return "shata"
}

//GetObject ...
func (ds *DiskStore) GetObject(k string) string {
	return "shata"
}
