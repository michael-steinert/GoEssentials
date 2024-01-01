package network

import (
	"io/fs"
	"os"
)

func ReadHosts(filesystem fs.FS) (string, error) {
	content, err := fs.ReadFile(filesystem, "hosts")
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ReadHostsFromFile(hostFile string) (string, error) {
	content, err := os.ReadFile(hostFile)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
