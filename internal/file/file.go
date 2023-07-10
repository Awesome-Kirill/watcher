package file

import (
	"bufio"
	"context"
	"os"

	"github.com/rs/zerolog"
)

type File struct {
	path  string
	hosts []string
}

func (s *File) Host(_ context.Context) ([]string, error) {
	return s.hosts, nil
}

func Load(path string, logger *zerolog.Logger) *File {
	file, err := os.Open(path)
	if err != nil {
		logger.Fatal().Err(err).Msg("open file error")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hosts []string
	for scanner.Scan() {
		hosts = append(hosts, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		logger.Fatal().Err(err).Msg("read file error")
	}

	return &File{
		path:  path,
		hosts: hosts,
	}
}
