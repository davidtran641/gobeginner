package server

import (
	"encoding/json"
	"io"
	"os"
)

// FileSystemPlayStore is an implementation of PlayStore using file system
type FileSystemPlayStore struct {
	database *json.Encoder
	league   League
}

// NewFileSystemPlayerStore return FileSystemPlayStore connect to given db
func NewFileSystemPlayerStore(db *os.File) *FileSystemPlayStore {
	db.Seek(0, 0)
	league, _ := NewLeague(db)

	return &FileSystemPlayStore{
		database: json.NewEncoder(&tape{db}),
		league:   league,
	}
}

// GetLeague returns list of players with score
func (f *FileSystemPlayStore) GetLeague() League {
	return f.league
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
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

// RecordScore record user has won
func (f *FileSystemPlayStore) RecordScore(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}
