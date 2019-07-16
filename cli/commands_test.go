package cli

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"regexp"
)

func TestValidCommands(t *testing.T) {
	for _, tc := range []struct {
		argv []string
		args Commands
		desc string
	}{
		{
			argv: []string{"some-file"},
			args: Commands{
				FileNames: []string{"some-file"},
			},
			desc: "Valid single filename",
		},
		{
			argv: []string{"some-file", "some-file-2"},
			args: Commands{
				FileNames: []string{"some-file", "some-file-2"},
			},
			desc: "Valid two filenames",
		},
		{
			argv: []string{},
			args: Commands{
				FileNames: []string{},
			},
			desc: "Valid empty filenames",
		},
		{
			argv: []string{"-c", "./some/milv.yaml", "some-file"},
			args: Commands{
				ConfigFile: "./some/milv.yaml",
				FileNames: []string{"some-file"},
			},
			desc: "Valid config file flag",
		},
		{
			argv: []string{"--config-file", "./some/milv.yaml", "some-file"},
			args: Commands{
				ConfigFile: "./some/milv.yaml",
				FileNames: []string{"some-file"},
			},
			desc: "Valid config file command",
		},
		{
			argv: []string{"-d", "./some-path", "some-file"},
			args: Commands{
				DocumentRoot: "./some-path",
				FileNames: []string{"some-file"},
			},
			desc: "Valid document root flag",
		},
		{
			argv: []string{"--document-root", "./some-path", "some-file"},
			args: Commands{
				DocumentRoot: "./some-path",
				FileNames: []string{"some-file"},
			},
			desc: "Valid document root command",
		},
		{
			argv: []string{"-e", "https://example.com", "some-file"},
			args: Commands{
				ExternalPath: "https://example.com",
				FileNames: []string{"some-file"},
			},
			desc: "Valid external path flag",
		},
		{
			argv: []string{"--external-path", "https://example.com", "some-file"},
			args: Commands{
				ExternalPath: "https://example.com",
				FileNames: []string{"some-file"},
			},
			desc: "Valid external path command",
		},
		{
			argv: []string{"-w", "^.*$,^abc", "some-file"},
			args: Commands{
				WhiteList: []*regexp.Regexp{regexp.MustCompile(`^.*$`), regexp.MustCompile(`^abc`)},
				FileNames: []string{"some-file"},
			},
			desc: "Valid white-list flag",
		},
		{
			argv: []string{"--white-list", "^.*$,^abc", "some-file"},
			args: Commands{
				WhiteList: []*regexp.Regexp{regexp.MustCompile(`^.*$`), regexp.MustCompile(`^abc`)},
				FileNames: []string{"some-file"},
			},
			desc: "Valid white-list command",
		},
		{
			argv: []string{"-b", "^.*$,^abc", "some-file"},
			args: Commands{
				BlackList: []*regexp.Regexp{regexp.MustCompile(`^.*$`), regexp.MustCompile(`^abc`)},
				FileNames: []string{"some-file"},
			},
			desc: "Valid black-list flag",
		},
		{
			argv: []string{"--black-list", "^.*$,^abc", "some-file"},
			args: Commands{
				BlackList: []*regexp.Regexp{regexp.MustCompile(`^.*$`), regexp.MustCompile(`^abc`)},
				FileNames: []string{"some-file"},
			},
			desc: "Valid black-list command",
		},
		{
			argv: []string{"-t", "42", "some-file"},
			args: Commands{
				Timeout: 42 * time.Second,
				FileNames: []string{"some-file"},
			},
			desc: "Valid timeout flag",
		},
		{
			argv: []string{"--timeout", "42", "some-file"},
			args: Commands{
				Timeout: 42 * time.Second,
				FileNames: []string{"some-file"},
			},
			desc: "Valid timeout command",
		},
		{
			argv: []string{"-r", "42", "some-file"},
			args: Commands{
				Repeats: 42,
				FileNames: []string{"some-file"},
			},
			desc: "Valid repeats flag",
		},
		{
			argv: []string{"--repeats", "42", "some-file"},
			args: Commands{
				Repeats: 42,
				FileNames: []string{"some-file"},
			},
			desc: "Valid repeats command",
		},
		{
			argv: []string{"-v", "some-file"},
			args: Commands{
				Verbose: true,
				FileNames: []string{"some-file"},
			},
			desc: "Valid verbose flag",
		},
		{
			argv: []string{"--verbose", "some-file"},
			args: Commands{
				Verbose: true,
				FileNames: []string{"some-file"},
			},
			desc: "Valid verbose command",
		},
	}{
		t.Run(tc.desc, func(t *testing.T) {
			args, err := parseCommands(tc.argv)

			assert.NoError(t, err)
			assert.Equal(t, tc.args, args)
		})
	}
}

func TestInvalidCommands(t *testing.T) {
	for _, tc := range []struct {
		argv []string
		desc string
	}{
		{
			argv: []string{"-t", "foo", "some-file"},
			desc: "Invalid timeout value flag",
		},
		{
			argv: []string{"--timeout", "foo", "some-file"},
			desc: "Invalid timeout value flag",
		},
		{
			argv: []string{"-r", "foo", "some-file"},
			desc: "Invalid repeats value flag",
		},
		{
			argv: []string{"--repeats", "foo", "some-file"},
			desc: "Invalid repeats value flag",
		},
	}{
		t.Run(tc.desc, func(t *testing.T) {
			_, err := parseCommands(tc.argv)
			assert.Error(t, err)
		})
	}
}