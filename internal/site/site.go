package site

import (
	"bufio"
	"context"
	"log"
	"os"
)

type Site struct {
	sits []string
}

func (s *Site) GetSites(ctx context.Context) ([]string, error) {
	return s.sits, nil
}
func New(path string) *Site {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var site Site
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		site.sits = append(site.sits, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &site

}
