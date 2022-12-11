package pkg

import (
	"gosca/pkg/gosca"
	"testing"
)

func TestScaffolding(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		var s gosca.Scaffolding
		if jerr := s.FromJsonBytes([]byte(`{
			"scaffolding":{
				"name":"test",
				"files": ["test.txt"],
				"dirs": [{
					"name": "sub-test",
					"files": ["sub-text.txt"]
				}]
			}
		}`)); jerr != nil {
			t.Fatalf("fail at fill the Scaffolding from json %s", jerr)
		}

		if s.Scaffolding.Name != "test" {
			t.Fatalf("the name of the scaffolding expected to be 'test' but got '%s'", s.Scaffolding.Name)
		}

		if l := len(s.Scaffolding.Files); l != 1 {
			t.Fatalf("the files's length expted to be 1 but it's %v", l)
		}

		if f := s.Scaffolding.Files[0]; f != "test.txt" {
			t.Fatalf("the file to create expected to be 'test.txt' but got '%s'", f)
		}

		if l := len(s.Scaffolding.Dirs); l != 1 {
			t.Fatalf("the dirs's length expected to be 1 but got %v", l)
		}
	})

	t.Run("yaml", func(t *testing.T) {
		var s gosca.Scaffolding
		if yerr := s.FromYamlBytes([]byte(`---
scaffolding:
  name: test
  files:
  - test.txt
  dirs:
  - name: sub-test
    files:
    - sub-text.txt
`)); yerr != nil {
			t.Fatalf("fail at fill the Scaffolding from yaml %s", yerr)
		}

		if s.Scaffolding.Name != "test" {
			t.Fatalf("the name of the scaffolding expected to be 'test' but got '%s'", s.Scaffolding.Name)
		}

		if l := len(s.Scaffolding.Files); l != 1 {
			t.Fatalf("the files's length expted to be 1 but it's %v", l)
		}

		if f := s.Scaffolding.Files[0]; f != "test.txt" {
			t.Fatalf("the file to create expected to be 'test.txt' but got '%s'", f)
		}

		if l := len(s.Scaffolding.Dirs); l != 1 {
			t.Fatalf("the dirs's length expected to be 1 but got %v", l)
		}
	})
}
