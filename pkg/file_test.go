package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFile(t *testing.T) {
	t.Run("File Exists", func(t *testing.T) {
		_, err := NewFile("test-markdowns/external_links.md", &FileConfig{})
		assert.NoError(t, err)
	})

	t.Run("File Not Exists", func(t *testing.T) {
		_, err := NewFile("test-markdowns/not_exist_file.md", &FileConfig{})
		assert.Error(t, err, "The specified file isn't a markdown file")
	})

	t.Run("Extract Links", func(t *testing.T) {
		file, err := NewFile("test-markdowns/external_links.md", &FileConfig{})
		require.NoError(t, err)

		expected := Links{
			Link{
				AbsPath: "https://twitter.com",
				TypeOf:  ExternalLink,
			},
			Link{
				AbsPath: "https://github.com",
				TypeOf:  ExternalLink,
			},
			Link{
				AbsPath: "http://dont.exist.link.com",
				TypeOf:  ExternalLink,
			},
		}

		file.ExtractLinks()
		assert.Equal(t, expected, file.Links)
	})

	t.Run("Extract Headers", func(t *testing.T) {
		file, err := NewFile("test-markdowns/hash_internal_links.md", &FileConfig{})
		require.NoError(t, err)

		expected := Headers{
			"First Header",
			"Second Header",
			"Third Header",
			"Links",
		}

		file.ExtractHeaders()
		assert.Equal(t, expected, file.Headers)
	})

	t.Run("Validate Links", func(t *testing.T) {
		file, err := NewFile("test-markdowns/external_links.md", &FileConfig{})
		require.NoError(t, err)

		expected := Links{
			Link{
				AbsPath: "https://twitter.com",
				TypeOf:  ExternalLink,
				Result: LinkResult{
					Status: true,
				},
			},
			Link{
				AbsPath: "https://github.com",
				TypeOf:  ExternalLink,
				Result: LinkResult{
					Status: true,
				},
			},
			Link{
				AbsPath: "http://dont.exist.link.com",
				TypeOf:  ExternalLink,
				Result: LinkResult{
					Status: false,
				},
			},
		}

		file.ExtractLinks()
		file.ValidateLinks()
		assert.Equal(t, expected, file.Links)
	})
}
