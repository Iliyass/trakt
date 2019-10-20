package trakt

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"time"
)

const (
	BASE_FOLDER     = "."
	TRAKTS_FILENAME = "trakt.json"
	TAGS_FILENAME   = "tags.json"
)

// FileStorage is an Implementation of the Storage interface
type FileStorage struct {
	trakts *os.File
	tags   *os.File
}

func getPath(fileName string) string {
	return path.Join(BASE_FOLDER, fileName)
}

func createFile(fileName string) bool {
	newFile, err := os.Create(getPath(fileName))
	if err != nil {
		panic(err)
	}
	_, err = newFile.WriteString("[]")
	if err != nil {
		panic(err)
	}
	newFile.Close()
	return true
}

func checkIfFileExist(fileName string) bool {
	_, err := os.Stat(path.Join(BASE_FOLDER, fileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

func readFile(fileName string) []byte {
	fileExist := checkIfFileExist(fileName)
	if !fileExist {
		createFile(fileName)
	}
	data, err := ioutil.ReadFile(getPath(fileName))
	if err != nil {
		panic(err)
	}
	return data
}

func writeFile(fileName string, data []byte) error {
	err := ioutil.WriteFile(getPath(fileName), data, 0666)
	if err != nil {
		panic(err)
	}
	return nil
}

// AddTrakt a Trakt in the Storage
func (s *FileStorage) AddTrakt(trakt Trakt) bool {
	fileName := TRAKTS_FILENAME
	data := readFile(fileName)

	var trakts []Trakt
	err := json.Unmarshal(data, &trakts)
	if err != nil {
		panic(err)
	}
	trakts = append(trakts, trakt)
	data, err = json.Marshal(trakts)

	err = writeFile(fileName, data)
	if err != nil {
		panic(err)
	}
	return true
}

// GetTraktsByDate to list trakts by date
func (s *FileStorage) GetTraktsByDate(from time.Time, to time.Time) (res []Trakt) {
	fileName := TRAKTS_FILENAME
	data := readFile(fileName)

	err := json.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}
	for _, trakt := range res {
		if trakt.CreatedAt >= from.Unix() && trakt.CreatedAt <= to.Unix() {
			res = append(res, trakt)
		}
	}
	return res
}

// GetTraktsByTag find Trakt by Tag Object
func (s *FileStorage) GetTraktsByTag(tag Tag) (res []Trakt) {
	// for _, trakt := range s.trakts {
	// 	for _, t := range trakt.Tags {
	// 		if t.Name == tag.Name {
	// 			res = append(res, trakt)
	// 		}
	// 	}
	// }
	return res
}

// AddTag add tag
func (s *FileStorage) AddTag(tag Tag) bool {
	fileName := TAGS_FILENAME
	data := readFile(fileName)

	var tags []Tag
	err := json.Unmarshal(data, &tags)
	if err != nil {
		panic(err)
	}
	isDuplicate := false
	for _, t := range tags {
		if t.Name == tag.Name {
			isDuplicate = true
		}
	}

	if isDuplicate {
		return true
	}
	tags = append(tags, tag)
	data, err = json.Marshal(tags)
	if err != nil {
		panic(err)
	}
	err = writeFile(fileName, data)
	if err != nil {
		panic(err)
	}
	return true
}

// GetTag get tag
func (s *FileStorage) GetTag(name string) (tag *Tag, err error) {
	// for _, t := range s.tags {
	// 	if t.Name == name {
	// 		return tag, nil
	// 	}
	// }
	// return nil, errors.New("Tag not found")
	return nil, errors.New("Tag not found")
}
