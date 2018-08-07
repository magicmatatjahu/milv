package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinks(t *testing.T) {
	t.Run("Remove White Links", func(t *testing.T) {
		whiteListExternal, whiteListInternal := []string{"github.com"}, []string{"../external_links.md"}
		links := Links{
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
			Link{
				AbsPath: "test-markdowns/external_links.md",
				RelPath: "../external_links.md",
				TypeOf:  InternalLink,
			},
			Link{
				AbsPath: "test-markdowns/sub_path/sub_sub_path/without_links.md",
				RelPath: "sub_sub_path/without_links.md",
				TypeOf:  InternalLink,
			},
			Link{
				AbsPath: "test-markdowns/sub_path/absolute_path.md",
				RelPath: "absolute_path.md",
				TypeOf:  InternalLink,
			},
			Link{
				AbsPath: "test-markdowns/sub_path/invalid.md",
				RelPath: "invalid.md",
				TypeOf:  InternalLink,
			},
		}

		expected := Links{
			Link{
				AbsPath: "https://twitter.com",
				TypeOf:  ExternalLink,
			},
			Link{
				AbsPath: "http://dont.exist.link.com",
				TypeOf:  ExternalLink,
			},
			Link{
				AbsPath: "test-markdowns/sub_path/sub_sub_path/without_links.md",
				RelPath: "sub_sub_path/without_links.md",
				TypeOf:  InternalLink,
			},
			Link{
				AbsPath: "test-markdowns/sub_path/absolute_path.md",
				RelPath: "absolute_path.md",
				TypeOf:  InternalLink,
			},
			Link{
				AbsPath: "test-markdowns/sub_path/invalid.md",
				RelPath: "invalid.md",
				TypeOf:  InternalLink,
			},
		}

		result := links.RemoveWhiteLinks(whiteListExternal, whiteListInternal)
		assert.Equal(t, expected, result)
	})

	t.Run("Filter", func(t *testing.T) {
		// given
		links := Links{
			Link{
				AbsPath: "testPath1",
				TypeOf:  InternalLink,
			},
			Link{
				AbsPath: "testPath2",
				TypeOf:  HashInternalLink,
			},
			Link{
				AbsPath: "testPath3",
				TypeOf:  ExternalLink,
			},
		}

		expected := Links{
			Link{
				AbsPath: "testPath3",
				TypeOf:  ExternalLink,
			},
		}

		// when
		result := links.Filter(func(link Link) bool {
			return link.TypeOf == ExternalLink
		})

		// then
		assert.ElementsMatch(t, expected, result)
	})
}
