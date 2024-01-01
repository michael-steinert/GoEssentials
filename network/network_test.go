package network_test

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/michael-steinert/GoEssentials/network"
	"github.com/stretchr/testify/assert"
)

func TestReadHosts(t *testing.T) {
	exceptedContent := "127.0.0.1 localhost\n"
	// Create temporary Filesystem
	filesystem := fstest.MapFS{
		"hosts": {
			Data: []byte(exceptedContent),
		},
	}
	// Test Function
	content, err := network.ReadHosts(filesystem)
	assert.Equal(t, exceptedContent, content)
	assert.NoError(t, err)
}

func TestReadHostsFromFile(t *testing.T) {
	exceptedContent := "127.0.0.1 localhost\n"
	// Create temporary Filesystem
	temporaryDirectory, err := os.MkdirTemp("", "test-")
	assert.NoError(t, err)
	// Cleanup temporary Filesystem after Test
	defer os.RemoveAll(temporaryDirectory)
	// Create File in temporary Filesystem
	hostFile := filepath.Join(temporaryDirectory, "hosts")
	err = os.WriteFile(hostFile, []byte(exceptedContent), 0644)
	assert.NoError(t, err)
	// Test Function
	content, err := network.ReadHostsFromFile(hostFile)
	assert.Equal(t, exceptedContent, content)
	assert.NoError(t, err)
}
