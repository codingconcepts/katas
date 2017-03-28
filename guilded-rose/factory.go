package main

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

// ItemMappings contains a mapping between Item type and
// Item name, allowing for the dynamic update of item
// types without modification to the code.
type ItemMappings map[string]string

// LoadMappings loads ItemMappings from an io.Reader (be
// that a JSON file, or a strings.Reader for testing.
func LoadMappings(reader io.Reader) (mappings ItemMappings, err error) {
	bytes := new(bytes.Buffer)
	if _, err = io.Copy(bytes, reader); err != nil {
		return
	}

	if err = json.Unmarshal(bytes.Bytes(), &mappings); err != nil {
		return
	}

	return
}

// CategoriseItem find the mapping type in the map and
// fall back to the Standard Updater type if one can't
// be found.
func (mappings ItemMappings) CategoriseItem(item *Item) (updater Updater) {
	mappingType, ok := mappings[item.name]
	if !ok {
		return &Standard{Item: item}
	}

	switch strings.ToLower(mappingType) {
	case "aged":
		return &Aged{Item: item}
	case "legendary":
		return &Legendary{Item: item}
	case "backstage pass":
		return &BackstagePass{Item: item}
	case "conjured":
		return &Conjured{Item: item}
	default:
		return &Standard{Item: item}
	}
}
