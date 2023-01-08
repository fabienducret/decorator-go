package main

import (
	"fmt"
)

// interface
type Book interface {
	print() string
}

// livre par défaut
type DefaultBook struct {
	name   string
	author string
}

func (b DefaultBook) print() string {
	return fmt.Sprintf("%s - %s", b.author, b.name)
}

// décorateur avec logger
type VerboseBook struct {
	original Book
}

func (b VerboseBook) print() string {
	fmt.Println("Book printing")
	return b.original.print()
}

func getBook() Book {
	return VerboseBook{
		original: DefaultBook{
			author: "Yegor Bugayenko",
			name:   "Elegant Objects, vol 2",
		},
	}
}

func main() {
	book := getBook()
	fmt.Println(book.print())
}
