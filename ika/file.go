package ika

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type DecoderCreator func(r io.Reader) Decoder

type File struct {
	path string
	dec  DecoderCreator
}

type Decoder interface {
	Decode(interface{}) error
}

func FileSource(path string) File {
	return File{path: path}
}

func FileWithDecoder(path string, dec DecoderCreator) File {
	return File{path: path, dec: dec}
}

func (fs File) Load(cfg interface{}) error {
	// open the configuration file
	f, err := os.OpenFile(fs.path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	// custom decoder
	if fs.dec != nil {
		return fs.dec(f).Decode(cfg)
	}

	decoder, err := DetectDecoder(fs.path, f)
	if err != nil {
		return err
	}

	return decoder.Decode(cfg)
}

// DetectDecoder returns file decoder according to it's extension
//
// Currently following file extensions are supported:
//
// - yaml
//
// - json
//
func DetectDecoder(path string, r io.Reader) (Decoder, error) {
	// parse the file depending on the file type
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".yaml", ".yml":
		return yaml.NewDecoder(r), nil
	case ".json":
		return json.NewDecoder(r), nil
	default:
		return nil, fmt.Errorf("file format '%s' doesn't supported by the parser", ext)
	}
}
