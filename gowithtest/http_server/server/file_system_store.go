package server

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

// FileSystemPlayStore is an implementation of PlayStore using file system
type FileSystemPlayStore struct {
	database *json.Encoder
	league   League
}

// NewFileSystemPlayerStore return FileSystemPlayStore connect to given db
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayStore, error) {
	err := initPlayerDbFile(file)
	if err != nil {
		return nil, err
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("can't load player store from file %s, error: %v", file.Name(), err)
	}

	return &FileSystemPlayStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func initPlayerDbFile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("can't get info from file %v, %s", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}

// GetLeague returns list of players with score
func (f *FileSystemPlayStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
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
