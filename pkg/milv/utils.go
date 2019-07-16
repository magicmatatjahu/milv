package milv

import (
	"os"
	"io/ioutil"
)

func fileExists(file string) (bool, error) {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err

	}
	return true, nil
}

func readFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
