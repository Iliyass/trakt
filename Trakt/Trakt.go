package trakt

import (
	"encoding/json"
	"time"
)

// Start with declaring the Data Structures
// Make a Storage
// You can Insert a trakt
// List trakts

// Tag struct
type Tag struct {
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}

// UnmarshallJSON unmarshalling json of Tag
func (tag *Tag) UnmarshallJSON(data []byte) (Tag, error) {
	var t Tag
	err := json.Unmarshal(data, &t)
	if err != nil {
		panic(err)
	}
	return NewTag(t.Name, t.CreatedAt), nil
}

// NewTag Factory func to create Tag
func NewTag(name string, createdAt int64) (tag Tag) {
	if name == "" {
		panic("Name is required")
	}
	if createdAt == 0 {
		panic("CreatedAt is required")
	}
	return Tag{Name: name, CreatedAt: createdAt}
}

// Trakt struct
type Trakt struct {
	Text      string `json:"text"`
	CreatedAt int64  `json:"created_at"`
	Tags      []Tag  `json:"tags"`
}

// Storage interface
type Storage interface {
	AddTrakt(trakt Trakt) bool
	GetTraktsByDate(from time.Time, to time.Time) []Trakt
	GetTraktsByTag(tag Tag) []Trakt
	AddTag(tag Tag) bool
	// GetTag(name string) (*Tag, error)
	ListTags() ([]Tag, error)
}

var storageInstance Storage

func getStorage() Storage {
	if storageInstance == nil {
		storageInstance = &FileStorage{}
	}
	return storageInstance
}

// (Deserialization) JSON API where the Trakt package gets JSON of Trakt or Tag to parse it

// AddTrakt : Public API for Adding JSON based Trakt
func AddTrakt(traktData []byte) (bool, error) {
	storage := getStorage()
	// Deserialize Trakt
	var trakt Trakt
	err := json.Unmarshal(traktData, &trakt)
	if err != nil {
		return false, err
	}

	// add trakt to storage
	return storage.AddTrakt(trakt), nil
}

// GetTraktsByDate returning takts by date
func GetTraktsByDate(from int64, to int64) (res []byte) {
	storage := getStorage()
	fromUnix := time.Unix(from, 0)
	toUnix := time.Unix(to, 0)
	trakts := storage.GetTraktsByDate(fromUnix, toUnix)
	res, err := json.Marshal(trakts)
	if err != nil {
		panic(err)
	}
	return res
}

// AddTag : Public API for Adding JSON based Tag
func AddTag(tagData []byte) (bool, error) {
	storage := getStorage()
	// Deserialize Tag
	var tag Tag
	tag, err := tag.UnmarshallJSON(tagData)
	if err != nil {
		return false, err
	}

	return storage.AddTag(tag), nil
}

// List Tags: Public Api
func ListTags() (tags []Tag, err error) {
	storage := getStorage()
	return storage.ListTags()
}

// (Serialization) JSON API where the Trakt package response to commands with JSON formatted Trakt or Tags
