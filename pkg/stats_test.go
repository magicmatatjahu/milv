package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStats(t *testing.T) {
	var links Links

	t.Run("External Links", func(t *testing.T) {
		file, err := NewFile("test-markdowns/external_links.md", links, nil)
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
							Message: "Get http://dont.exist.link.com: dial tcp: lookup dont.exist.link.com: no such host",
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
		file, err := NewFile("test-markdowns/sub_path/internal_links.md", links, nil)
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
		file, err := NewFile("test-markdowns/hash_internal_links.md", links, nil)
		require.NoError(t, err)

		expected := &FileStats{
			SuccessLinks: SuccessLinks{
				Count: 8,
				Links: []Link{
					Link{
						AbsPath: "https://github.com",
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
					Link{
						RelPath: "#header-with-block",
						TypeOf:  HashInternalLink,
						Result: LinkResult{
							Status: true,
						},
					},
					Link{
						RelPath: "#header-with-link",
						TypeOf:  HashInternalLink,
						Result: LinkResult{
							Status: true,
						},
					},
					Link{
						RelPath: "#very-strange-header-really-people-create-headers-look-like-this",
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
							Message: "The specified header doesn't exist in file",
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
		file, err := NewFile("test-markdowns/sub_path/absolute_path.md", links, nil)
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
