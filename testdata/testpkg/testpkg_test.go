package testpkg_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/sqlcgetters/testdata/testpkg"
)

func Test_TestMe(t *testing.T) {
	want := testpkg.TestStruct{
		A: 1,
		B: "2",
	}
	require.Equal(t, want, testpkg.TestMe())
}
