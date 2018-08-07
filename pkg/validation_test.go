package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	t.Run("External Links", func(t *testing.T) {
		links := []Link{
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

		expected := []Link{
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

		valid := &Validation{}
		result := valid.Links(links)

		assert.Equal(t, expected, result)
	})

	t.Run("Internal Links", func(t *testing.T) {
		links := []Link{
			Link{
				AbsPath: "test-markdowns/external_links.md",
				TypeOf:  InternalLink,
			},
			Link{
				AbsPath: "test-markdowns/sub_path/sub_sub_path/without_links.md",
				TypeOf:  InternalLink,
			},
			Link{
				AbsPath: "test-markdowns/sub_path/absolute_path.md",
				TypeOf:  InternalLink,
			},
			Link{
				AbsPath: "test-markdowns/sub_path/invalid.md",
				TypeOf:  InternalLink,
			},
		}

		expected := []Link{
			Link{
				AbsPath: "test-markdowns/external_links.md",
				TypeOf:  InternalLink,
				Result: LinkResult{
					Status: true,
				},
			},
			Link{
				AbsPath: "test-markdowns/sub_path/sub_sub_path/without_links.md",
				TypeOf:  InternalLink,
				Result: LinkResult{
					Status: true,
				},
			},
			Link{
				AbsPath: "test-markdowns/sub_path/absolute_path.md",
				TypeOf:  InternalLink,
				Result: LinkResult{
					Status: true,
				},
			},
			Link{
				AbsPath: "test-markdowns/sub_path/invalid.md",
				TypeOf:  InternalLink,
				Result: LinkResult{
					Status:  false,
					Message: "The specified file doesn't exist",
				},
			},
		}

		valid := &Validation{}
		result := valid.Links(links)

		assert.Equal(t, expected, result)
	})

	t.Run("Hash Internal Links", func(t *testing.T) {
		existHeaders := Headers{"First Header", "Second Header", "Third Header"}
		links := []Link{
			Link{
				RelPath: "#first-header",
				TypeOf:  HashInternalLink,
			},
			Link{
				RelPath: "#second-header",
				TypeOf:  HashInternalLink,
			},
			Link{
				RelPath: "#third-header",
				TypeOf:  HashInternalLink,
			},
			Link{
				RelPath: "#header",
				TypeOf:  HashInternalLink,
			},
		}

		expected := []Link{
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
				RelPath: "#header",
				TypeOf:  HashInternalLink,
				Result: LinkResult{
					Status:  false,
					Message: "The specified header doesn't exist",
				},
			},
		}

		valid := &Validation{}
		result := valid.Links(links, existHeaders)

		assert.Equal(t, expected, result)
	})
}
