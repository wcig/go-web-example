package file

import "os"

func IsFileExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func MakeDir(f string) error {
	if IsFileExist(f) {
		return nil
	}
	return os.MkdirAll(f, os.ModePerm)
}
