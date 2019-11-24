package fetcher

import (
	"encoding/json"
	"os"
	"time"

	tree "github.com/twstokes/treedivided"
)

// LocalFetcher reads from a local file
type LocalFetcher struct {
	path string
}

// NewLocalFetcher creates a fetcher that reads from path
func NewLocalFetcher(path string) *LocalFetcher {
	return &LocalFetcher{path}
}

// GetLatestState gets the latest game state from the file
func (l *LocalFetcher) GetLatestState() (*tree.GameState, error) {
	f, err := os.Open(l.path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	score := tree.GameState{}
	decoder := json.NewDecoder(f)

	err = decoder.Decode(&score)
	if err != nil {
		return nil, err
	}

	return &score, nil
}

// WaitTime is how long the local fetcher waits before re-reading the file
func (l *LocalFetcher) WaitTime() time.Duration {
	return time.Second * 5
}
