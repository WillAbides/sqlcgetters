package sqlcgetters

import (
	"bytes"
	"flag"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -write-golden

var writeGolden bool

func TestMain(m *testing.M) {
	flag.BoolVar(&writeGolden, "write-golden", false, "write golden files")
	flag.Parse()
	os.Exit(m.Run())
}

func TestGenerateGetters(t *testing.T) {
	fsys := os.DirFS(".")

	t.Run("nogofiles", func(t *testing.T) {
		var buf bytes.Buffer
		_, err := GenerateGetters(&buf, fsys)
		require.EqualError(t, err, "no sources provided")
		require.Equal(t, "", buf.String())
	})

	t.Run("nongofile", func(t *testing.T) {
		var buf bytes.Buffer
		_, err := GenerateGetters(&buf, fsys, "testdata/nongofile/foo.go", "testdata/nongofile/nongofile.txt")
		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "could not parse go source"))
	})

	t.Run("alpha", func(t *testing.T) {
		goldenTest(t, fsys, "testdata/alpha")
	})

	t.Run("nostructs", func(t *testing.T) {
		goldenTest(t, fsys, "testdata/nostructs")
	})

	t.Run("testpkg", func(t *testing.T) {
		var buf bytes.Buffer
		_, err := GenerateGetters(&buf, fsys, "testdata/testpkg/testpkg.go", "testdata/testpkg/testpkg_test.go")
		require.EqualError(t, err, "could not parse go source: package name mismatch: testpkg != testpkg_test")
	})

	t.Run("nostructs", func(t *testing.T) {
		goldenTest(t, fsys, "testdata/nostructs")
	})

	t.Run("sqlc", func(t *testing.T) {
		tdMap := map[string]bool{}
		err := fs.WalkDir(fsys, "testdata/sqlc", func(p string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if path.Ext(info.Name()) == ".go" {
				tdMap[path.Dir(p)] = true
			}
			return nil
		})
		require.NoError(t, err)
		testDirs := make([]string, 0, len(tdMap))
		for k := range tdMap {
			testDirs = append(testDirs, k)
		}
		sort.Strings(testDirs)
		for _, testDir := range testDirs {
			t.Run(testDir, func(t *testing.T) {
				goldenTest(t, fsys, testDir)
			})
		}
	})
}

func goldenTest(t *testing.T, fsys fs.FS, dir string) {
	t.Helper()
	sources, err := fs.Glob(fsys, path.Join(dir, "*.go"))
	require.NoError(t, err)
	var buf bytes.Buffer
	_, err = GenerateGetters(&buf, fsys, sources...)
	require.NoError(t, err)
	gettersGo := path.Join(dir, "getters.go")
	if writeGolden {
		require.NoError(t, os.MkdirAll(filepath.Dir(gettersGo), 0o755))
		require.NoError(t, os.WriteFile(gettersGo, buf.Bytes(), 0o644))
		return
	}
	golden, err := fs.ReadFile(fsys, gettersGo)
	require.NoError(t, err)
	require.Equal(t, string(golden), buf.String())
}
