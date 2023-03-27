package mock

type Retriever struct {
	Container string
}

func (r Retriever) Get(url string) string {
	return r.Container
}
