package jsonschema

import (
	"encoding/json"
	"testing"
)

func TestStringOrArrayMarshalJSON(t *testing.T) {
	type testHarness struct {
		Entry stringOrArray `json:",omitempty"`
	}

	testCases := []struct {
		name   string
		input  testHarness
		output string
	}{
		{
			name:   "multiple",
			input:  testHarness{Entry: StringOrArray("string", "null")},
			output: `{"Entry":["string","null"]}`,
		},
		{
			name:   "simple",
			input:  testHarness{Entry: StringOrArray("string")},
			output: `{"Entry":"string"}`,
		},
		{
			name:   "empty",
			input:  testHarness{},
			output: `{}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := json.Marshal(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			got := string(b)

			if got != tc.output {
				t.Errorf("Expected %q, got %q", tc.output, got)
			}
		})
	}
}

func TestStringOrArrayWith(t *testing.T) {
	s := stringOrArray(nil).With("string")

	if len(s) != 1 {
		t.Errorf("Expected length 1, got %d", len(s))
	}

	b := s.With("null")
	if len(s) != 1 {
		t.Errorf("Expected length 1, got %d", len(s))
	}

	if len(b) != 2 {
		t.Errorf("Expected length 2, got %d", len(b))
	}
}
