package gosca

import (
	"gosca/utils"
	"os"
)

func findPlaceholders(bs []byte, p func(string) string) []byte {
	var buff []byte
	l := len(bs)

	for c := 0; c < l; {
		// All bytes that are not '$'
		if bs[c] != 0x24 {
			buff = append(buff, bs[c])
			c += 1
			continue
		}

		// When current byte is 0x24, skip it.
		c += 1

		// Validate the index range
		if c >= l {
			break
		}

		// The following byte of 0x24 is not a 0x7b '{' skip it.
		if bs[c] != 0x7b {
			buff = append(buff, bs[c-1])
			continue
		}

		// Next byte
		c += 1

		var sbuff []byte
		_c := 0

		// Collect the byte inside the ${ }
		for _c < l && bs[c+_c] != 0x7d {
			sbuff = append(sbuff, bs[c+_c])
			_c += 1
		}

		buff = append(buff, []byte(p(string(sbuff)))...)
		c += _c + 1
	}

	return buff

}

func PreprocessConfig(f string, p func(string) string) ([]byte, error) {
	if b, berr := os.ReadFile(f); berr != nil {
		return nil, utils.WrapError(berr, "cannot open the config %s file at preprocess", f)
	} else {
		return findPlaceholders(b, p), nil
	}
}
