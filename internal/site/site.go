package site

import (
	"context"
	"log"
	"os"
	"strings"
)

type Site struct {
	sits []string
}

func (s *Site) GetSites(ctx context.Context) ([]string, error) {
	return s.sits, nil
}
func New(path string) *Site {

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	//todo bufio.ScanLines()
	list := strings.Split(string(data), "\n")

	return &Site{sits: list}

}
