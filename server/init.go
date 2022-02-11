package server

func Init() *StationServer {
	return &StationServer{&InMemoryStationStore{}}
}

type InMemoryStationStore struct {
}

func (i *InMemoryStationStore) GetStationURL(name string) string {
	return "fakeURL"
}
