package pkg

import (
	"gosca/pkg/gosca"
	"testing"
)

func TestPreprocess(t *testing.T) {
	if config, err := gosca.PreprocessConfig("../assets/test.txt", func(p string) string {
		return "World!"
	}); err != nil {
		t.Fatalf("an error ocurrs at preprocess configuration %s", err)
	} else if string(config) != "Hello World!" {
		t.Fatalf("the replaced value of the placeholder is not correct: %s", config)
	}
}
