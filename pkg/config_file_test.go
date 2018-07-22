package pkg

import (
	"testing"
	"github.com/magicmatatjahu/milv/cli"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestConfigFile(t *testing.T) {
	t.Run("Config File", func(t *testing.T) {
		SetBasePath("test-markdowns", false)

		commands := cli.Commands{
			ConfigFile: "test-markdowns/milv-test.config.yaml",
		}

		expected := &FileConfig{
			WhiteListExt: []string{"localhost", "abc.com", "github.com"},
			WhiteListInt: []string{"LICENSE", "#contributing"},
		}

		config, err := NewConfig(commands)
		require.NoError(t, err)

		result := NewFileConfig("test-markdowns/src/foo.md", config)

		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}