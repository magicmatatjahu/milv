package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	t.Run("File Exists", func(t *testing.T) {
		err := fileExists("test-markdowns/external_links.md")
		assert.NoError(t, err)
	})

	t.Run("File Not Exists", func(t *testing.T) {
		err := fileExists("test-markdowns/foo_bar.md")
		assert.Error(t, err)
	})

	t.Run("Header Exists", func(t *testing.T) {
		header := "#first-header"
		existHeaders := Headers{"First Header", "Second Header", "Third Header"}

		result := headerExists(header, existHeaders)
		assert.Equal(t, true, result)
	})

	t.Run("Header Not Exists", func(t *testing.T) {
		header := "#non-exist-header"
		existHeaders := Headers{"First Header", "Second Header", "Third Header"}

		result := headerExists(header, existHeaders)
		assert.Equal(t, false, result)
	})

	t.Run("Unique Elements", func(t *testing.T) {
		elements := []string{"localhost", "abc.com", "localhost", "github.com", "github.com"}

		expected := []string{"abc.com", "github.com", "localhost"}
		result := unique(elements)

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("Remove Files From Black List", func(t *testing.T) {
		filePaths, blackList := []string{"./abc.md", "./foo/bar.md"}, []string{"foo"}

		expected := []string{"./abc.md"}
		result := removeBlackList(filePaths, blackList)

		assert.Equal(t, expected, result)
	})

	t.Run("Remove Code Blocks", func(t *testing.T) {
		filePaths, blackList := []string{"./abc.md", "./foo/bar.md"}, []string{"foo"}

		expected := []string{"./abc.md"}
		result := removeBlackList(filePaths, blackList)

		assert.Equal(t, expected, result)
	})
}
