package memory

import (
	"github.com/beyondstorage/go-storage/v4/services"
	"github.com/beyondstorage/go-storage/v4/types"
	"strings"
)

// Storage is the example client.
type Storage struct {
	defaultPairs DefaultStoragePairs
	features     StorageFeatures

	workDir string
	root    *object

	types.UnimplementedStorager
	types.UnimplementedAppender
	types.UnimplementedCopier
	types.UnimplementedDirer
	types.UnimplementedMover
}

// String implements Storager.String
func (s *Storage) String() string {
	return "memory"
}

// NewStorager will create Storager only.
func NewStorager(pairs ...types.Pair) (types.Storager, error) {
	root := newObject(nil, types.ModeDir)
	root.parent = root

	return &Storage{
		root:    root,
		workDir: "/",
	}, nil
}

// formatError converts errors returned by SDK into errors defined in go-storage and go-service-*.
// The original error SHOULD NOT be wrapped.
func (s *Storage) formatError(op string, err error, path ...string) error {
	if _, ok := err.(services.InternalError); ok {
		return err
	}

	return err
}

func (s *Storage) absPath(p string) string {
	return s.workDir + "/" + p
}

func (s *Storage) relPath(p string) string {
	return strings.TrimPrefix(p, s.workDir)
}
