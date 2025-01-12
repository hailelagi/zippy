package zippy

/*
go-fuse specifies interfaces for fusing onto a tree-like fs.
*/

import (
	"archive/zip"
	"context"

	"github.com/hanwen/go-fuse/v2/fs"
)

type zippyRoot struct {
	fs.Inode

	zr *zip.ReadCloser
}

type zipFile struct {
	fs.Inode
	file *zip.File
	data []byte
}

var _ fs.NodeOnAdder = (*zippyRoot)(nil)

// onAdd defines when you add an inode to the filesystem tree
// populates metadata from the archive format into an inode
func (r *zippyRoot) OnAdd(ctx context.Context) {}

// Open and Read service the `open` and `read` syscalls
// Getattr is serviced by `stat` and kin

func NewArchiveFileSystem(name string) (root fs.InodeEmbedder, err error) {
	return NewZip(name)
}

func NewZip(name string) (fs.InodeEmbedder, error) {
	r, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}

	return &zippyRoot{zr: r}, nil
}
