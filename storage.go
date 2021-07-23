package memory

import (
	"context"
	"github.com/beyondstorage/go-storage/v4/services"
	. "github.com/beyondstorage/go-storage/v4/types"
	"io"
)

func (s *Storage) create(path string, opt pairStorageCreate) (o *Object) {
	o = NewObject(s, true)
	o.ID = s.absPath(path)
	o.Path = s.relPath(path)
	return o
}

func (s *Storage) delete(ctx context.Context, path string, opt pairStorageDelete) (err error) {
	name, child := s.root.getChildByPath(s.absPath(path))
	if child == nil {
		return nil
	}
	child.parent.removeChild(name)
	return nil
}

func (s *Storage) list(ctx context.Context, path string, opt pairStorageList) (oi *ObjectIterator, err error) {
	fn := NextObjectFunc(func(ctx context.Context, page *ObjectPage) error {
		_, o := s.root.getChildByPath(s.absPath(path))
		if o == nil {
			// If the object is not exist, we should return IterateDone instead.
			return IterateDone
		}
		if !o.mode.IsDir() {
			// If the object mode is not dir, we should return directly.
			return services.ErrObjectModeInvalid
		}

		o.mu.Lock()
		defer o.mu.Unlock()

		for k, v := range o.child {
			xo := NewObject(s, true)
			xo.ID = s.absPath(path + "/" + k)
			xo.Path = s.relPath(path + "/" + k)
			xo.Mode = v.mode
			xo.SetContentLength(v.length)

			page.Data = append(page.Data, xo)
		}
		return IterateDone
	})
	return NewObjectIterator(ctx, fn, nil), nil
}

func (s *Storage) metadata(opt pairStorageMetadata) (meta *StorageMeta) {
	return &StorageMeta{
		Name:    "memory",
		WorkDir: "/",
	}
}

func (s *Storage) read(ctx context.Context, path string, w io.Writer, opt pairStorageRead) (n int64, err error) {
	name, o := s.root.getChildByPath(s.absPath(path))
	if name == "" {
		return 0, services.ErrObjectNotExist
	}

	written, err := w.Write(o.data)
	if err != nil {
		return int64(written), err
	}
	return int64(written), nil
}

func (s *Storage) stat(ctx context.Context, path string, opt pairStorageStat) (o *Object, err error) {
	name, ro := s.root.getChildByPath(s.absPath(path))
	if name == "" {
		return nil, services.ErrObjectNotExist
	}

	o = NewObject(s, true)
	o.ID = s.absPath(path)
	o.Path = s.relPath(path)
	o.Mode = ro.mode
	o.SetContentLength(ro.length)
	return o, nil
}

func (s *Storage) write(ctx context.Context, path string, r io.Reader, size int64, opt pairStorageWrite) (n int64, err error) {
	o := s.root.insertChildByPath(s.absPath(path))
	if o == nil {
		return 0, services.ErrObjectModeInvalid
	}

	o.length = size
	o.mode = ModeRead

	o.data = make([]byte, size)
	read, err := r.Read(o.data)
	if err != nil {
		return int64(read), nil
	}

	return int64(read), nil
}
