package io

import (
	"reflect"
	"testing"
)

func TestFileMapperFactory(t *testing.T) {

	type TestStruct struct {
		Name     string         `yaml:"name"  toml:"name" hcl:"name"`
		Includes string         `yaml:"includes" toml:"includes" hcl:"includes"`
		Tasks    map[string]any `yaml:"tasks" toml:"tasks" hcl:"tasks"`
	}

	result := TestStruct{
		Name:     "test",
		Includes: "{{ eq '1' '1' }}",
		Tasks:    map[string]any{"judo.echo": "el nuevo path es {{ .base_path }}"},
	}

	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want TestStruct
		wantErr bool
	}{
		//{"Load hcl content", args{path: "test/resources/test_load.hcl"}, result, false},
		{"Load toml content", args{path: "resources/test_load.toml"}, result, false},
		{"Load json content", args{path: "resources/test_load.json"}, result, false},
		{"Load yaml content", args{path: "resources/test_load.yaml"}, result, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var testValue TestStruct
			err := FileMapperFactory(tt.args.path).Mapper(&testValue)
			if ! tt.wantErr && err != nil {
				t.Errorf("FileMapperFactory() error: %v", err)
			} 
			if !reflect.DeepEqual(testValue, tt.want) {
				t.Errorf("FileMapperFactory() = %v, want %v", testValue, tt.want)
			}
		})
	}
}
