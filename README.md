# decorator-go

Here is an example of pattern decorator helping to extend some capabilities. 

First we need an interface 
```go
type Book interface {
	display() string
}
```

Then a default implementation.
We don't want this implementation to log something, it must be pure.
```go
type DefaultBook struct {
	name   string
	author string
}

func (b DefaultBook) display() string {
	return fmt.Sprintf("%s - %s", b.author, b.name)
}
```

Finally a decorator to log something
```go
type VerboseBook struct {
	original Book
}

func (b VerboseBook) display() string {
	fmt.Println("Book display")
	return b.original.display()
}
```

