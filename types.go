package jsonschema

import (
	"encoding/json"
	"fmt"
)

type stringOrArray []string

func StringOrArray(values ...string) stringOrArray {
	return stringOrArray(values)
}

func (s stringOrArray) With(v string) stringOrArray {
	r := make([]string, len(s))
	for i, vv := range s {
		if vv == v {
			return s
		}
		r[i] = vv
	}
	r = append(r, v)

	return stringOrArray(r)
}

func (s stringOrArray) Has(v string) bool {
	for _, entry := range s {
		if v == entry {
			return true
		}
	}
	return false
}

func (s stringOrArray) MarshalJSON() ([]byte, error) {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = v
	}

	switch len(r) {
	case 0:
		return nil, nil
	case 1:
		return json.Marshal(r[0])
	default:
		return json.Marshal(r)
	}
}

func (s *stringOrArray) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return nil
	}

	switch input[0] {
	case '"':
		var parsed string

		err := json.Unmarshal(input, &parsed)
		if err != nil {
			return err
		}

		*s = StringOrArray(parsed)

	case '[':
		var parsed []string

		err := json.Unmarshal(input, &parsed)
		if err != nil {
			return err
		}

		*s = StringOrArray(parsed...)

	default:
		return fmt.Errorf("Can't parse string or array")
	}

	return nil
}
