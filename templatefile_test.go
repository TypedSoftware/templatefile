package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

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
