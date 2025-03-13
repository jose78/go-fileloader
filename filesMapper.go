package io

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)


// Factory to provide the mapper based on the extension. By default, the Judofile will be parsed as a yaml file.
func FileMapperFactory(path string) FileMapper {
	extension := filepath.Ext(path)
	var mapper FileMapper
	switch extension {
	case ".yaml":
		mapper = Yaml{path: path}
	case ".yml":
		mapper = Yaml{path: path}
	case ".json":
		mapper = Json{path: path}
	case ".toml":
		mapper = Toml{path: path}
	case ".hcl":
		mapper = Hcl{path: path}
	}

	return mapper
}

type FileMapper interface {
	Mapper(any) error
}

type Json struct {
	path string
}

type Yaml struct {
	path string
}

type Toml struct {
	path string
}

type Hcl struct {
	path string
}

// Implementation the interface FileMapper
func (fileType Hcl) Mapper(data any) error {



	parser := hclparse.NewParser()
	f, parseDiags := parser.ParseHCLFile(fileType.path)
	if parseDiags.HasErrors() {
		return fmt.Errorf(parseDiags.Error())
	}

	decodeDiags := gohcl.DecodeBody(f.Body, nil, data)
	if decodeDiags.HasErrors() {
		return fmt.Errorf(parseDiags.Error())
	}

	return nil
}

// Implementation the interface FileMapper
func (fileType Json) Mapper(data any) error {

	file, err := os.ReadFile(fileType.path)
	if err != nil {
		return err
	}

	err2 := json.Unmarshal(file, &data)

	if err2 != nil {
		return err2
	}
	return nil
}

// Implementation the interface FileMapper
func (fileType Yaml) Mapper(data any) error {

	file, err := os.ReadFile(fileType.path)
	if err != nil {
		return err
	}

	err2 := yaml.Unmarshal(file, data)

	if err2 != nil {
		return err2
	}
	return nil
}

// Implementation the interface FileMapper
func (fileType Toml) Mapper(data any) error {

	value := data
	_, err := toml.DecodeFile(fileType.path, value)
	if err != nil {
		return err
	}
	return nil
}

