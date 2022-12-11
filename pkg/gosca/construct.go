package gosca

import (
	"fmt"
	"gosca/utils"
	"os"
)

func constructDirectory(p string, d Dir) error {
	p = fmt.Sprintf("%s/%s", p, d.Name)

	if merr := os.MkdirAll(p, 0777); merr != nil {
		return utils.WrapError(merr, "cannot create the directory %s at directory construction", p)
	}

	for _, file := range d.Files {
		if ferr := os.WriteFile(fmt.Sprintf("%s/%s", p, file), []byte{}, 0777); ferr != nil {
			return utils.WrapError(ferr, "cannot create the file %s at directory construction")
		}
	}

	for _, dir := range d.Dirs {
		if cerr := constructDirectory(p, dir); cerr != nil {
			return cerr
		}
	}

	return nil
}

func ConstructScaffolding(p string, s *Scaffolding) error {
	return constructDirectory(p, s.Scaffolding)
}
