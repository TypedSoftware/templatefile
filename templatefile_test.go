package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func TestTemplatefile(t *testing.T) {
	matches, err := filepath.Glob("testdata/*.tmpl")
	require.Nil(t, err)

	for _, match := range matches {
		t.Run(match, func(t *testing.T) {
			expected, err := ioutil.ReadFile(strings.TrimSuffix(match, filepath.Ext(match)) + ".txt")
			require.Nil(t, err)

			vars, err := ioutil.ReadFile(strings.TrimSuffix(match, filepath.Ext(match)) + ".yml")
			require.Nil(t, err)

			actual, err := templatefile(".", match, string(vars))
			require.Nil(t, err)
			require.EqualValues(t, expected, actual)
		})
	}
}

func TestTemplatefileRecursive(t *testing.T) {
	_, err := templatefile(".", "testdata/recursive.tmpl-fails", "k: v")
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "Call to unknown function")
}

func FuzzTemplatefile(f *testing.F) {
	f.Add("Hello, ${name}!", "name: world\n")
	f.Add("${val}", "val: works\n")
	f.Add(`The items are ${join(", ", list)}`, "list:\n  - foo\n  - bar\n")
	f.Add("%{ for x in list ~}\n- ${x}\n%{ endfor ~}", "list:\n  - foo\n")

	f.Fuzz(func(t *testing.T, tmpl, vars string) {
		tmp := t.TempDir()
		path := filepath.Join(tmp, "fuzz.tmpl")
		require.NoError(t, os.WriteFile(path, []byte(tmpl), 0o600))

		// Assert that the parser doesn't panic or hang, and produces valid UTF-8
		actual, err := templatefile(tmp, path, vars)
		require.True(t, err != nil || utf8.ValidString(actual), "output is invalid UTF-8")
	})
}
