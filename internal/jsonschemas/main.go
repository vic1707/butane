package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/coreos/butane/base/v0_1"
	"github.com/coreos/butane/base/v0_2"
	"github.com/coreos/butane/base/v0_3"
	"github.com/coreos/butane/base/v0_4"
	"github.com/coreos/butane/base/v0_5"
	"github.com/coreos/butane/base/v0_6"
	"github.com/coreos/butane/base/v0_7_exp"
	"github.com/invopop/jsonschema"
)

func main() {
	r := jsonschema.Reflector{
		FieldNameTag: "yaml",
	}

	schemas := []struct {
		Version string
		Config  interface{}
	}{
		{"v0_1", v0_1.Config{}},
		{"v0_2", v0_2.Config{}},
		{"v0_3", v0_3.Config{}},
		{"v0_4", v0_4.Config{}},
		{"v0_5", v0_5.Config{}},
		{"v0_6", v0_6.Config{}},
		{"v0_7_exp", v0_7_exp.Config{}},
	}

	for _, entry := range schemas {
		schema := r.Reflect(entry.Config)

		fileName := fmt.Sprintf("%s_schema.json", entry.Version)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Error creating file %s: %v", fileName, err)
			continue
		}

		err = json.NewEncoder(file).Encode(schema)
		if err != nil {
			fmt.Printf("Error writing schema for %s: %v", entry.Version, err)
		}

		file.Close()
	}
}
