package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func resolveConfigWithPath(path string, out interface{}) (err error) {
	var f *os.File
	defer func() {
		if f == nil {
			return
		}
		closeErr := f.Close()
		err = closeErr
	}()

	f, err = os.Open(path)
	if err != nil {
		return err
	}

	err = UnmarshalYAMLStrict(f, out)
	return
}

// UnmarshalYAMLStrict provides a YAML decoder that does not allow unknown
// fields.
func UnmarshalYAMLStrict(r io.Reader, out interface{}) error {
	d := yaml.NewDecoder(r)
	d.KnownFields(true)
	return d.Decode(out)
}
