package envy

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func parseln(line string) (key string, val string, err error) {
	splits := strings.Split(line, "=")

	if len(splits) < 2 {
		return "", "", errors.New("missing delimiter '='")
	}

	key = strings.Trim(splits[0], " ")
	val = strings.Trim(splits[1], ` "'`)

	return key, val, nil
}

// Bootstrap loads a .env file into the current environment using envy.Load
func Bootstrap() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}

	err = Load(file)
	if err != nil {
		return err
	}

	return nil
}

// Load parses lines of a reader in the .env format.
func Load(reader io.Reader) error {
	r := bufio.NewReader(reader)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		key, val, err := parseln(string(line))
		if err != nil {
			return err
		}

		os.Setenv(key, val)
	}

	return nil
}
