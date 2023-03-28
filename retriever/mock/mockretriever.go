package mock

type Retriever struct {
	Contents string
}

//func (r *Retriever) String() string {
//	return fmt.Sprintf(
//		"Retriever: {Contents=%s}", r.Contents)
//}

func (r *Retriever) Post(url string, form map[string]string) {
	r.Contents = form["content"]
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
