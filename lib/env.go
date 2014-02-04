package envy

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// Env represents the values parsed from the .env file
type Env map[string]string

// Bootstrap loads a .env file into the current environment using envy.Load
func Bootstrap() (Env, error) {
	file, err := os.Open(".env")
	if err != nil {
		return nil, err
	}

	return Load(file)
}

// Load parses lines of a reader in the .env format.
func Load(reader io.Reader) (Env, error) {
	r := bufio.NewReader(reader)
	env := make(map[string]string)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		key, val, err := parseln(string(line))
		if err != nil {
			return env, err
		}

		env[key] = val
		os.Setenv(key, val)
	}

	return env, nil
}

func parseln(line string) (key string, val string, err error) {
	splits := strings.Split(line, "=")

	if len(splits) < 2 {
		return "", "", errors.New("missing delimiter '='")
	}

	key = strings.Trim(splits[0], " ")
	val = strings.Trim(splits[1], ` "'`)

	return key, val, nil
}
