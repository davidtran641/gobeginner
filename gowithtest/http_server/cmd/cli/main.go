package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/pocker"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play pocker")

	store, close, err := pocker.NewFileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("create player store err %v", err)
	}

	defer close()

	game := pocker.NewCLI(store, os.Stdin)

	game.PlayPocker()

}
