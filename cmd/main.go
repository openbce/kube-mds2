package main

import (
	"openbce.io/kmds/pkg/apiserver"

	_ "openbce.io/kmds/pkg/storage/engine"
)

func main() {
	bridge, err := apiserver.NewMdsBridage(nil)
	if err != nil {
		panic(err)
	}

	if err := bridge.Run(); err != nil {
		panic(err)
	}
}
