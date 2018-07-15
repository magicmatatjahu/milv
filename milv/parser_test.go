package milv

import (
	"testing"
	"path/filepath"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestLinksParse(t *testing.T) {
	t.Run("From Url", func(t *testing.T) {
		url := "https://github.com/dkhamsing/awesome_bot/blob/master/README.md"
		links, err := GetLinks(url)

		require.NoError(t, err)
		assert.NotNil(t, links)
	})

	t.Run("From file", func(t *testing.T) {
		absFilePath, err := filepath.Abs("../README.md")
		require.NoError(t, err)

		links, err := GetLinks(absFilePath)

		require.NoError(t, err)
		assert.NotNil(t, links)
	})
}