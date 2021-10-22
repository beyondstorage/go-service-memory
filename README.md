# go-service-memory

memory service support for [go-storage](https://github.com/beyondstorage/go-storage).

## Notes

**This package has been moved to [go-storage](https://github.com/beyondstorage/go-storage/tree/master/services/memory).**

```shell
go get go.beyondstorage.io/services/memory
```

## Install

```go
go get github.com/beyondstorage/go-service-memory
```

## Usage

```go
import (
	"log"

	_ "github.com/beyondstorage/go-service-memory"
	"github.com/beyondstorage/go-storage/v4/services"
)

func main() {
	store, err := services.NewStoragerFromString("memory:///path/to/workdir")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/beyondstorage/go-storage-example).
