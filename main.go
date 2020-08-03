package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/arham09/go-di-example/manual"
	// "github.com/krmahadevan/di/manual"
	// "github.com/krmahadevan/di/wire"
)

func main() {
	impl := flag.String("type", "manual", "Which type of Dependency Injection to use (manual/dig/wire)")

	flag.Parse()

	fmt.Println("type:", *impl)
	switch *impl {
	case "manual":
		manual.Main()
	// case "wire":
	// 	wire.Main()
	default:
		panic(errors.New("unknown Dependency Injection model chosen"))
	}
}
