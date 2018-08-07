package pkg

import (
	"testing"

	"github.com/magicmatatjahu/milv/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Run("Config File", func(t *testing.T) {
		commands := cli.Commands{
			ConfigFile: "test-markdowns/milv-test.config.yaml",
		}

		expected := &Config{
			Files: []File{
				File{
					RelPath: "./src/foo.md",
					Config: &FileConfig{
						WhiteListExt: []string{"github.com"},
						WhiteListInt: []string{"#contributing"},
					},
				},
			},
			WhiteListExt: []string{"localhost", "abc.com"},
			WhiteListInt: []string{"LICENSE"},
			BlackList:    []string{"./README.md"},
		}

		result, err := NewConfig(commands)

		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
