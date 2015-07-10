package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func FuzzyQuery(path, file string) (paths []string, err error) {

	err = filepath.Walk(path, func(p string, info os.FileInfo, err2 error) error {

		if info == nil {
			return nil
		}

		if !info.IsDir() {
			ok, _ := filepath.Match("*"+file+"*", info.Name())
			if ok {
				querypath, err := filepath.Rel(path, p)
				if err != nil {
					return nil
				}

				paths = append(paths, querypath)

				err = nil
				//return errors.New("Found")
				return err
			}
		}

		return nil
	})

	if err == nil && len(paths) > 0 {
		return paths, nil
	}

	return paths, errors.New("Not found")
}
