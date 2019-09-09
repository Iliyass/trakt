package main

import (
	"fmt"
	"time"
)

// Start with declaring the Data Structures
// Make a Storage
// You can Insert a trakt
// List trakts

// Tag struct
type Tag struct {
	Name      string
	CreatedAt time.Time
}

// Trakt struct
type Trakt struct {
	Text      string
	CreatedAt time.Time
	Tags      []Tag
}

// Storage interface
type Storage interface {
	Insert(trakt Trakt) bool
	GetTrakts(from time.Time, to time.Time) []Trakt
}

type RAMStorage struct {
	trakts []Trakt
}

func (s *RAMStorage) Insert(trakt Trakt) bool {
	s.trakts = append(s.trakts, trakt)
	return len(s.trakts) > 0
}

func (s *RAMStorage) GetTrakts(from time.Time, to time.Time) (res []Trakt) {
	for _, trakt := range s.trakts {
		if trakt.CreatedAt.Unix() >= from.Unix() && trakt.CreatedAt.Unix() <= to.Unix() {
			res = append(res, trakt)
		}
	}
	return res
}

func main() {
	from := time.Now()
	ex := Tag{Name: "Exersicing", CreatedAt: time.Now()}
	run := Trakt{Text: "Run 5k on 29m", CreatedAt: time.Now(), Tags: []Tag{ex}}
	to := time.Now()
	rmStorage := RAMStorage{}

	fmt.Println(rmStorage.Insert(run))
	fmt.Println(rmStorage.trakts)
	fmt.Println(rmStorage.GetTrakts(from, to))
}
