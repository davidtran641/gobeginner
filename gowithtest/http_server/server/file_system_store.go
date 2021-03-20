package server

import (
	"encoding/json"
	"io"
	"log"
)

// FileSystemPlayStore is an implementation of PlayStore using file system
type FileSystemPlayStore struct {
	database io.ReadWriteSeeker
}

// GetLeague returns list of players with score
func (f *FileSystemPlayStore) GetLeague() League {
	_, err := f.database.Seek(0, 0)
	if err != nil {
		log.Printf("Seek database error: %v", err)
		return nil
	}

	league, err := NewLeague(f.database)
	if err != nil {
		log.Printf("Read database error: %v", err)
	}
	return league
}

// NewLeague read players from a reader
func NewLeague(reader io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(reader).Decode(&league)
	if err != nil {
		return nil, err
	}
	return league, nil
}

// GetPlayerScore return player score given player name
func (f *FileSystemPlayStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

// RecordWin record user has won
func (f *FileSystemPlayStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
