package parser

import "testing"

// Replace return the variable matching a word
func TestReplace(t *testing.T) {
	variables := make(map[string]interface{})
	variables["one"] = "uno"
	variables["two"] = "dos"
	variables["three"] = "tres"

	parser := Parser{
		Variables: variables,
		Delimeter: "$",
	}

	var testTable = []struct {
		Var      string
		Expected string
	}{
		{"$one$", "uno"},
		{"two", "dos"},
		{"three$", "tres"},
		{"$$one", "uno"},
		{"$$", ""},
	}

	for _, tt := range testTable {
		w, err := parser.Replace(tt.Var)
		if err != nil {
			if tt.Var != "$$" {
				t.Error(err)
			}
		}
		if tt.Expected != w {
			t.Fatalf("%s != %s", tt.Expected, w)
		}
	}
}

func TestRender(t *testing.T) {
	variables := make(map[string]interface{})
	variables["one"] = "uno"
	variables["two"] = "dos"
	variables["three"] = "tres"

	parser := Parser{
		Variables: variables,
		Delimeter: "$",
	}

	var testTable = []struct {
		Var      string
		Expected string
	}{
		{"$one$", "uno"},
		{"two$", "dos"},
		{"three$=drei!", "tres=drei!"},
		{"err$", ""},
		{"err$!", ""},
	}

	for k, tt := range testTable {
		w, err := parser.Render(tt.Var, k)
		if err != nil {
			if tt.Var != "err$" && tt.Var != "err$!" {
				t.Error(err)
			}
		}
		if tt.Expected != w {
			t.Fatalf("%s != %s", tt.Expected, w)
		}
	}

}
