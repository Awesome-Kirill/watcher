package file

import (
	"bufio"
	"context"
	"log"
	"os"
)

// todo может как-то заюсать io.Reader / как на это писать тесты // иожет юзать one
type File struct {
	path  string
	hosts []string
}

func (s *File) Hosts(ctx context.Context) ([]string, error) {
	return s.hosts, nil
}

func Load(path string) *File {
	log.Println("DO IT ONE!!!")
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hosts []string
	for scanner.Scan() {
		hosts = append(hosts, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &File{
		path:  path,
		hosts: hosts,
	}
}
