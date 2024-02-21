package main

import (
	"context"
	"orca/internal/gen/pkl/orca"
	. "orca/internal/logger"
)

func main() {
	cfg, err := orca.LoadFromPath(context.Background(), "./assets/inputs/hello-world.pkl")
	if err != nil {
		Error(err.Error())
	}
}
