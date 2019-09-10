package trakt

import (
	"encoding/json"
	"strconv"
	"time"
)

// Start with declaring the Data Structures
// Make a Storage
// You can Insert a trakt
// List trakts

// Tag struct
type Tag struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// Trakt struct
type Trakt struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Tags      []Tag     `json:"tags"`
}

// Storage interface
type Storage interface {
	AddTrakt(trakt Trakt) bool
	GetTraktsByDate(from time.Time, to time.Time) []Trakt
	// 	getTraktsByTag(tag Tag) []Trakt
	// 	AddTag(tag Tag) bool
	// 	getTags(name string)
}

var storageInstance Storage

func getStorage() Storage {
	if storageInstance == nil {
		storageInstance = &RAMStorage{}
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
func GetTraktsByDate(from string, to string) (res []byte) {
	storage := getStorage()
	fromInt, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		panic(err)
	}
	toInt, err := strconv.ParseInt(to, 10, 64)
	if err != nil {
		panic(err)
	}
	fromUnix := time.Unix(fromInt, 0)
	toUnix := time.Unix(toInt, 0)
	trakts := storage.GetTraktsByDate(fromUnix, toUnix)
	res, err = json.Marshal(trakts)
	if err != nil {
		panic(err)
	}
	return res
}

// AddTag : Public API for Adding JSON based Tag
// func AddTag(tagData []byte, storage Storage) (bool, error) {
// 	// Deserialize Tag
// 	var tag Tag
// 	err := json.Unmarshal(tagData, &tag)
// 	if err != nil {
// 		return false, err
// 	}

// 	return storage.AddTag(tag), nil
// }

// (Serialization) JSON API where the Trakt package response to commands with JSON formatted Trakt or Tags
