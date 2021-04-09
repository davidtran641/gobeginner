package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/poker"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")

	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("create player store err %v", err)
	}

	defer close()

	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)

	cli.PlayPoker()

}
