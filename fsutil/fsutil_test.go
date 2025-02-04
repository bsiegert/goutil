package fsutil_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMimeType(t *testing.T) {
	assert.Eq(t, "", fsutil.MimeType(""))
	assert.Eq(t, "", fsutil.MimeType("not-exist"))
	assert.Eq(t, "image/jpeg", fsutil.MimeType("testdata/test.jpg"))

	buf := new(bytes.Buffer)
	buf.Write([]byte("\xFF\xD8\xFF"))
	assert.Eq(t, "image/jpeg", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte("text"))
	assert.Eq(t, "text/plain; charset=utf-8", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte(""))
	assert.Eq(t, "", fsutil.ReaderMimeType(buf))
	buf.Reset()

	assert.True(t, fsutil.IsImageFile("testdata/test.jpg"))
	assert.False(t, fsutil.IsImageFile("testdata/not-exists"))
}

func TestTempDir(t *testing.T) {
	dir, err := fsutil.TempDir("testdata", "temp.*")
	assert.NoErr(t, err)
	assert.True(t, fsutil.IsDir(dir))
	assert.NoErr(t, fsutil.Remove(dir))
}

func TestSplitPath(t *testing.T) {
	dir, file := fsutil.SplitPath("/path/to/dir/some.txt")
	assert.Eq(t, "/path/to/dir/", dir)
	assert.Eq(t, "some.txt", file)
}

func TestToAbsPath(t *testing.T) {
	assert.Eq(t, "", fsutil.ToAbsPath(""))
	assert.Eq(t, "/path/to/dir/", fsutil.ToAbsPath("/path/to/dir/"))
	assert.Neq(t, "~/path/to/dir", fsutil.ToAbsPath("~/path/to/dir"))
	assert.Neq(t, ".", fsutil.ToAbsPath("."))
	assert.Neq(t, "..", fsutil.ToAbsPath(".."))
	assert.Neq(t, "./", fsutil.ToAbsPath("./"))
	assert.Neq(t, "../", fsutil.ToAbsPath("../"))
}

func TestSlashPath(t *testing.T) {
	assert.Eq(t, "/path/to/dir", fsutil.SlashPath("/path/to/dir"))
	assert.Eq(t, "/path/to/dir", fsutil.UnixPath("/path/to/dir"))
	assert.Eq(t, "/path/to/dir", fsutil.UnixPath("\\path\\to\\dir"))
}
