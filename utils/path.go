package common

import "os"

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}
