package pkg

import (
	"gosca/pkg/gosca"
	"gosca/utils"
	"testing"
)

func TestConstructScaffolding(t *testing.T) {
	s := gosca.Scaffolding{
		Scaffolding: gosca.Dir{
			Name:  "test",
			Files: []string{"test.txt"},
		},
	}

	if cerr := gosca.ConstructScaffolding("../assets", &s); cerr != nil {
		t.Fatalf("the construction of the scaffolding fail %s", cerr)
	}

	if !utils.FileExists("../assets/test/test.txt") {
		t.Fatalf("the test.txt file create by the construct scaffolding doesn't exists")
	}
}
