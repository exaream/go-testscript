// Sample using package testscirpt
package sample_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

var scriptDir = filepath.Join("testdata", "script")

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"mycat": mycat,
	}))
}

func TestSample(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir:         filepath.Join(scriptDir, t.Name()),
		WorkdirRoot: t.TempDir(),
	})
}

// mycat behaves like cat command.
// Do NOT use *testing.T as an argument because mycat is used in testscript.RunMain().
func mycat() int {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			return 1
		}
		return 0
	}

	for _, fname := range os.Args[1:] {
		f, err := os.Open(fname)
		if err != nil {
			return 1
		}

		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			return 1
		}
	}
	return 0
}
