package jokeapi

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"sync"
)

var _ JokeAdder = (*FileStore)(nil)
var _ JokeGetter = (*FileStore)(nil)

type FileStore struct {
	fileName string

	updateLock *sync.Mutex
}

func NewFileStore(fileName string) *FileStore {
	return &FileStore{
		fileName:   fileName,
		updateLock: &sync.Mutex{},
	}
}

// Get returns a random joke from the file.
func (fs *FileStore) Get(ctx context.Context) (*Joke, error) {
	jokes, err := fs.loadFile()
	if err != nil {
		return nil, err
	}

	index := rand.Intn(len(jokes))
	joke := jokes[index]
	return &joke, nil
}

func (fs *FileStore) loadFile() ([]Joke, error) {
	data, err := os.ReadFile(fs.fileName)
	if err != nil {
		return nil, err
	}

	jokes := []Joke{}

	err = json.Unmarshal(data, &jokes)
	if err != nil {
		return nil, err
	}
	return jokes, nil
}

// Adds a joke to the file.
func (fs *FileStore) Add(ctx context.Context, joke *Joke) error {
	jokes, err := fs.loadFile()
	if err != nil {
		return err
	}

	jokes = append(jokes, *joke)

	return fs.saveToFile(jokes)
}

func (fs *FileStore) saveToFile(jokes []Joke) error {
	data, err := json.Marshal(jokes)
	if err != nil {
		return err
	}

	tmpFile, err := os.CreateTemp("", "joke-api-save*")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(data)
	if err != nil {
		return err
	}

	os.Rename(tmpFile.Name(), fs.fileName)
	return nil
}
