package backup

import (
	"fmt"
	"path/filepath"
	"time"
)

// will have map of paths with associated hashes
type Monitor struct {
	Paths       map[string]string
	Archiver    Archiver
	Destination string
}

//iterates over every path and generates latest hash of that folder
func (m *Monitor) Now() (int, error) {	
	var counter int

	for path, lastHash := range m.Paths {
		newHash, err := DirHash(path)
		if err != nil {
			return counter, err
		}

		if newHash != lastHash {
			err := m.act(path)
			if err != nil {
				return counter, err
			}

			m.Paths[path] = newHash //update hash
			counter++
		}
	}

	return counter, nil
}

func (m *Monitor) act(path string) error {
	dirName := filepath.Base(path)	//returns last name of path	
	fileName := fmt.Sprintf(m.Archiver.DestFmt(), time.Now().UnixNano())	//time.Now().UnixNano() is used to generate a timestamp filename

	// decidng where archive will go and callin Archive method
	return m.Archiver.Archive(path, filepath.Join(m.Destination, dirName, fileName))
}
