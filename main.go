package main

import (
	"github.com/grilix/chest-go/chest"
	"github.com/grilix/chest-go/chest-ui"
)

func main() {
	client := chest.NewClient(nil)

	chest_ui.Loop(client)
}
