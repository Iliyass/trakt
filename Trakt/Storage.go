package trakt

import (
	"errors"
	"time"
)

// RAMStorage is an Implementation of the Storage interface
type RAMStorage struct {
	trakts []Trakt
	tags   []Tag
}

// AddTrakt a Trakt in the Storage
func (s *RAMStorage) AddTrakt(trakt Trakt) bool {
	s.trakts = append(s.trakts, trakt)
	return len(s.trakts) > 0
}

//GetTraktsByDate to list trakts by date
func (s *RAMStorage) GetTraktsByDate(from time.Time, to time.Time) (res []Trakt) {
	for _, trakt := range s.trakts {
		if trakt.CreatedAt >= from.Unix() && trakt.CreatedAt <= to.Unix() {
			res = append(res, trakt)
		}
	}
	return res
}

// GetTraktsByTag find Trakt by Tag Object
func (s *RAMStorage) GetTraktsByTag(tag Tag) (res []Trakt) {
	for _, trakt := range s.trakts {
		for _, t := range trakt.Tags {
			if t.Name == tag.Name {
				res = append(res, trakt)
			}
		}
	}
	return res
}

// AddTag add tag
func (s *RAMStorage) AddTag(tag Tag) bool {
	prevLen := len(s.tags)
	s.tags = append(s.tags, tag)
	return len(s.tags) > prevLen
}

// GetTag get tag
func (s *RAMStorage) GetTag(name string) (tag *Tag, err error) {
	for _, t := range s.tags {
		if t.Name == name {
			return tag, nil
		}
	}
	return nil, errors.New("Tag not found")
}
