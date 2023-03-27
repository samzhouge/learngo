package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real2"
)

type Retriever interface {
	Get(string) string
}

func Download(r Retriever) string {
	data := r.Get("http://www.imooc.com")
	return data
}

func main() {
	var r Retriever
	r = mock.Retriever{Container: "Mock Retriever"}
	inspect(r)

	//r = &real2.Retriever{
	//	UserAgent: "Mozilla/5.0",
	//	TimeOut:   time.Minute,
	//}

	r = &real2.Retriever{}
	inspect(r)

	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Container)
	} else {
		fmt.Println("not a mock retriever")
	}
	//fmt.Println(Download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Container)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
