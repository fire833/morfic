package persist

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

func Unwrap(file string, i interface{}) error {

	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	switch filepath.Ext(file) {
	case ".json":
		{
			return json.Unmarshal(data, i)
		}
	case ".yml", ".yaml":
		{
			return yaml.Unmarshal(data, i)
		}
	case ".toml":
		{
			return toml.Unmarshal(data, i)
		}
	default:
		{
			return errors.New("Unable to unmarshal this file (unknown extension).")
		}
	}

}
