package sqlcgetters

import (
	"os"
	"path"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListSQLCFiles(t *testing.T) {
	fsys := os.DirFS(".")
	dir := "testdata/sqlc/examples/authors/postgresql"
	var want []string
	for _, f := range []string{"db.go", "models.go", "query.sql.go"} {
		want = append(want, path.Join(dir, f))
	}
	files, err := ListSQLCFiles(fsys, dir)
	require.NoError(t, err)
	sort.Strings(files)
	sort.Strings(want)
	require.Equal(t, want, files)
}
