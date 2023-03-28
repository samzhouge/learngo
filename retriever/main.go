package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real2"
	"time"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string)
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	data := r.Get(url)
	return data
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"content": "another fake imooc",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "Mock Retriever"}
	r = &retriever
	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("try a session")
	fmt.Println(session(&retriever))
	//fmt.Println(Download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println(" > Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println(" > UserAgent:", v.UserAgent)
	}
}
