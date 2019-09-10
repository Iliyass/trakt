package trakt

import (
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
		if trakt.CreatedAt.Unix() >= from.Unix() && trakt.CreatedAt.Unix() <= to.Unix() {
			res = append(res, trakt)
		}
	}
	return res
}

// getTraktsByTag find Trakt by Tag Object
// func (s *RAMStorage) getTraktsByTag(tag Tag) (res []Trakt) {
// 	for _, trakt := range s.trakts {
// 		for _, t := range trakt.Tags {
// 			if t.Name == tag.Name {
// 				res = append(res, trakt)
// 			}
// 		}
// 	}
// 	return res
// }
