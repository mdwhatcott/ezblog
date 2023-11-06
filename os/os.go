package os

import "os"

var ReadFile = func(path string) ([]byte, error) {
	return os.ReadFile(path)
}
var WriteFile = func(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}
