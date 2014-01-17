package envy

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func Parseln(line string) (key string, val string, err error) {
	splits := strings.Split(line, "=")

	if len(splits) < 2 {
		return "", "", errors.New("missing delimiter '='")
	}

	key = strings.Trim(splits[0], " ")
	val = strings.Trim(splits[1], ` "'`)

	return key, val, nil
}

// Loads a reader into the environment using Parseln
func Parse(reader io.Reader) error {
	// 1. read each line
	r := bufio.NewReader(reader)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		key, val, err := Parseln(string(line))
		if err != nil {
			return err
		}

		os.Setenv(key, val)
	}

	// 2. Invoke parseln on each line
	// 3. insert into the env AKA (os.Setenv(key, val))
	return nil
}
