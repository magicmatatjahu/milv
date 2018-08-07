package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStats(t *testing.T) {
	t.Run("External Links", func(t *testing.T) {
		file, err := NewFile("test-markdowns/external_links.md", &FileConfig{})
		require.NoError(t, err)

		expected := &FileStats{
			SuccessLinks: SuccessLinks{
				Count: 2,
				Links: []Link{
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
				},
			},
			FailedLinks: FailedLinks{
				Count: 1,
				Links: []Link{
					Link{
						AbsPath: "http://dont.exist.link.com",
						TypeOf:  ExternalLink,
						Result: LinkResult{
							Status: false,
						},
					},
				},
			},
		}

		file.Run()

		require.NoError(t, err)
		assert.Equal(t, expected, file.Stats)
	})

	t.Run("Internal Links", func(t *testing.T) {
		file, err := NewFile("test-markdowns/sub_path/internal_links.md", &FileConfig{})
		require.NoError(t, err)

		expected := &FileStats{
			SuccessLinks: SuccessLinks{
				Count: 3,
				Links: []Link{
					Link{
						AbsPath: "test-markdowns/external_links.md",
						RelPath: "../external_links.md",
						TypeOf:  InternalLink,
						Result: LinkResult{
							Status: true,
						},
					},
					Link{
						AbsPath: "test-markdowns/sub_path/sub_sub_path/without_links.md",
						RelPath: "sub_sub_path/without_links.md",
						TypeOf:  InternalLink,
						Result: LinkResult{
							Status: true,
						},
					},
					Link{
						AbsPath: "test-markdowns/sub_path/absolute_path.md",
						RelPath: "absolute_path.md",
						TypeOf:  InternalLink,
						Result: LinkResult{
							Status: true,
						},
					},
				},
			},
			FailedLinks: FailedLinks{
				Count: 1,
				Links: []Link{
					Link{
						AbsPath: "test-markdowns/sub_path/invalid.md",
						RelPath: "invalid.md",
						TypeOf:  InternalLink,
						Result: LinkResult{
							Status:  false,
							Message: "The specified file doesn't exist",
						},
					},
				},
			},
		}

		file.Run()

		require.NoError(t, err)
		assert.Equal(t, expected, file.Stats)
	})

	t.Run("Hash Internal Links", func(t *testing.T) {
		file, err := NewFile("test-markdowns/hash_internal_links.md", &FileConfig{})
		require.NoError(t, err)

		expected := &FileStats{
			SuccessLinks: SuccessLinks{
				Count: 3,
				Links: []Link{
					Link{
						RelPath: "#first-header",
						TypeOf:  HashInternalLink,
						Result: LinkResult{
							Status: true,
						},
					},
					Link{
						RelPath: "#second-header",
						TypeOf:  HashInternalLink,
						Result: LinkResult{
							Status: true,
						},
					},
					Link{
						RelPath: "#third-header",
						TypeOf:  HashInternalLink,
						Result: LinkResult{
							Status: true,
						},
					},
				},
			},
			FailedLinks: FailedLinks{
				Count: 1,
				Links: []Link{
					Link{
						RelPath: "#header",
						TypeOf:  HashInternalLink,
						Result: LinkResult{
							Status:  false,
							Message: "The specified header doesn't exist",
						},
					},
				},
			},
		}

		file.Run()

		require.NoError(t, err)
		assert.Equal(t, expected, file.Stats)
	})

	t.Run("Absolute Internal Path", func(t *testing.T) {
		file, err := NewFile("test-markdowns/sub_path/absolute_path.md", &FileConfig{})
		require.NoError(t, err)

		expected := &FileStats{
			SuccessLinks: SuccessLinks{
				Count: 1,
				Links: []Link{
					Link{
						AbsPath: "test-markdowns/external_links.md",
						RelPath: "/external_links.md",
						TypeOf:  InternalLink,
						Result: LinkResult{
							Status: true,
						},
					},
				},
			},
			FailedLinks: FailedLinks{
				Count: 0,
				Links: nil,
			},
		}

		file.Run()

		require.NoError(t, err)
		assert.Equal(t, expected, file.Stats)
	})
}
