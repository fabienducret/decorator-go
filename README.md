# decorator-go

Here is an example of pattern decorator helping to extend some capabilities.

First we need an interface

```go
type Book interface {
	print() string
}
```

Then a default implementation.
We don't want this implementation to log something, it must be pure.

```go
type DefaultBook struct {
	name   string
	author string
}

func (b DefaultBook) print() string {
	return fmt.Sprintf("%s - %s", b.author, b.name)
}
```

Finally a decorator to log something

```go
type VerboseBook struct {
	original Book
}

func (b VerboseBook) print() string {
	fmt.Println("Book printing")
	return b.original.print()
}
```

Here are the instanciations

```go
func getBook() Book {
	return VerboseBook{
		original: DefaultBook{
			author: "Yegor Bugayenko",
			name:   "Elegant Objects, vol 2",
		},
	}
}
```
