package file

import (
	"bufio"
	"context"
	"log"
	"os"
	"sync"
)

// todo может как-то заюсать io.Reader / как на это писать тесты // иожет юзать one
type File struct {
	once  *sync.Once
	path  string
	hosts []string
}

func (s *File) load() {
	s.once.Do(
		func() {
			log.Println("DO IT ONE!!!")
			file, err := os.Open(s.path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				s.hosts = append(s.hosts, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		})
}
func (s *File) Hosts(ctx context.Context) ([]string, error) {
	s.load()
	return s.hosts, nil
}

func New(path string) *File {
	return &File{
		once: &sync.Once{},
		path: path,
	}
}
