package customjson

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/jeffail/gabs"
)

type (
	container struct {
		data interface{}
	}
)

func (c *container) UnmarshalJSON(b []byte) error {
	gC, err := gabs.ParseJSON(b)
	if err != nil {
		return err
	}
	// Here's where the magic happens
	dt := reflect.TypeOf(c.data).Elem()

	for i := 0; i < dt.NumField(); i++ {
		f := dt.Field(i)
		tags := StructTagMap(f.Tag)
		// TODO: Fix to allow for field named/non-json-tagged
		jsonField := tags["json"]
		for tagK, tagV := range tags {
			if v, ok := registry[tagK]; ok {
				err = v(dt.String(), f.Name, tagV, gC.S(jsonField).Data())
				if err != nil {
					return err
				}
			}
		}
	}

	return json.Unmarshal(b, &c.data)
}

func Unmarshal(b []byte, i interface{}) error {
	c := &container{
		data: i,
	}
	return json.Unmarshal(b, c)
}

// StructTagMap will convert a StructTag to a map of maps for the values in a Field's tags
func StructTagMap(tag reflect.StructTag) map[string]string {
	// This code was taken directly from the reflect library
	tags := map[string]string{}
	for tag != "" {
		// Skip leading space.
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := string(tag[:i])
		tag = tag[i+1:]

		// Scan quoted string to find value.
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		qvalue := string(tag[:i+1])
		tag = tag[i+1:]

		value, err := strconv.Unquote(qvalue)
		if err != nil {
			break
		}
		tags[name] = value
	}
	return tags
}
