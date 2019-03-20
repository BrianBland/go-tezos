# goTezos: A Tezos Go Library

The purpose of this library is to allow developers to build go driven applications for Tezos of the Tezos RPC. This library is a work in progress, and not complete. 

## Installation

**Without wallet support:**

```
go get github.com/BrianBland/go-tezos
```

**With wallet support:**

1. Install pkg\_config (debian example below):
```
sudo apt-get install pkg_config
```

2. Install [libsoidum](https://libsodium.gitbook.io/doc/installation)

3. Get goTezos:

```
go get --tags wallet github.com/BrianBland/goTezos
```

## goTezos Documentation

[GoDoc](https://godoc.org/github.com/BrianBland/goTezos)

The goTezos Library requires you to set the RPC URL for a node to query. 


Usage:

```
package main

import (
	"fmt"
	goTezos "github.com/BrianBland/go-tezos"
)

func main() {
	gt := goTezos.NewGoTezos()
	gt.AddNewClient(goTezos.NewTezosRPCClient("localhost",":8732"))

	block,_ := gt.GetBlockAtLevel(1000)
	fmt.Println(block.Hash)
}
```

Fork of [**DefinitelyNotAGoat**](https://github.com/DefinitelyNotAGoat)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
