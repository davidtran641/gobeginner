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

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Opening  error  %s, %v", dbFileName, err)
	}

	store, err := pocker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("create player store err %v", err)
	}

	game := pocker.NewCLI(store, os.Stdin)

	game.PlayPocker()

}
