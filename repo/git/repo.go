package git

import (
	"path/filepath"

	"time"

	"github.com/fbsb/dotf/repo"
	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/plumbing/format/gitignore"
	"gopkg.in/src-d/go-git.v4/plumbing/format/index"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

const DOTFILES = ".dotfiles"

type gitRepo struct {
	g *git.Repository
}

var _ repo.Repository = &gitRepo{}

func (r gitRepo) Add(path string) error {
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	_, err = w.Add(path)
	if err != nil {
		return err
	}

	return nil
}

func (r gitRepo) Remove(path string) error {
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	idx, err := r.g.Storer.Index()
	if err != nil {
		return err
	}

	err = r.doRemoveRecursively(w, idx, path)
	if err != nil {
		return err
	}

	return r.g.Storer.SetIndex(idx)
}

func (r gitRepo) doRemoveRecursively(w *git.Worktree, idx *index.Index, path string) error {

	fi, err := w.Filesystem.Lstat(path)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		_, err = idx.Remove(path)
		return err
	}

	files, err := w.Filesystem.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		name := filepath.Join(path, file.Name())

		if file.IsDir() {
			err = r.doRemoveRecursively(w, idx, name)
		} else {
			_, err = idx.Remove(name)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (r gitRepo) Commit(message, name, email string) error {
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	author := object.Signature{When: time.Now(), Name: name, Email: email}

	_, err = w.Commit(message, &git.CommitOptions{
		All:       true,
		Author:    &author,
		Committer: &author,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r gitRepo) Status() (repo.Status, error) {
	w, err := r.Worktree()
	if err != nil {
		return nil, err
	}

	s, err := w.Status()
	if err != nil {
		return nil, err
	}

	ts := make(repoStatus)

	// Convert git statuses to repo statuses
	for path, fs := range s {
		ts.add(path, *fs)
	}

	return ts, nil

}

func (r gitRepo) Worktree() (*git.Worktree, error) {
	w, err := r.g.Worktree()
	if err != nil {
		return nil, err
	}

	w.Excludes = append(w.Excludes, gitignore.ParsePattern(filepath.Join("./", DOTFILES), nil))

	return w, nil
}

func Init(path string) (repo.Repository, error) {
	s, err := newStorage(path)
	if err != nil {
		return nil, err
	}

	// We provide nil as worktree in order to initialize as a bare repository.
	// By doing so we prevent git from creating the .git file pointing to the .dotfiles directory.
	g, err := git.Init(s, nil)
	if err != nil {
		return nil, err
	}

	return &gitRepo{g}, nil
}

func Open(path string) (repo.Repository, error) {
	s, err := newStorage(path)
	if err != nil {
		return nil, err
	}

	w, err := newWorktree(path)
	if err != nil {
		return nil, err
	}

	g, err := git.Open(s, w)
	if err != nil {
		return nil, err
	}
	return &gitRepo{g}, nil
}

func newWorktree(path string) (billy.Filesystem, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	return osfs.New(absPath), nil
}

func newStorage(path string) (storage.Storer, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	fs := osfs.New(filepath.Join(absPath, DOTFILES))
	c := cache.NewObjectLRUDefault()

	return filesystem.NewStorage(fs, c), nil
}
