package utils

import (
	"bufio"
	"fmt"
	"os"
)

func WrapError(e error, f string, a ...any) error {
	return fmt.Errorf("%s -> %w", fmt.Sprintf(f, a...), e)
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Input(f string, a ...any) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(f, a...)
	r, _ := reader.ReadString('\n')
	return r
}
