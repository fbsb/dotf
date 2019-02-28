package repo

import (
	"bytes"

	"fmt"

	"gopkg.in/src-d/go-git.v4"
)

// The repoStatus map ignores untracked or unmodified files
type repoStatus map[string]FileStatus

func (s repoStatus) File(path string) FileStatus {
	fs, ok := s[path]
	if !ok {
		return StatusUnchanged
	}
	return fs
}

func (s repoStatus) String() string {
	output := bytes.NewBuffer(nil)

	for path, status := range s {
		_, _ = fmt.Fprintf(output, "%s %s\n", status, path)
	}

	return output.String()
}

func (s repoStatus) IsClean() bool {
	return len(s) == 0
}

func (s *repoStatus) add(path string, gs git.FileStatus) {
	// Ignore untracked files
	if gs.Worktree == git.Untracked {
		return
	}

	// Ignore unmodified files
	if gs.Staging == git.Unmodified && gs.Worktree == git.Unmodified {
		return
	}

	(*s)[path] = convertGitStatus(gs)
}

func convertGitStatus(s git.FileStatus) FileStatus {
	// TODO(fbsb): deal with Renamed, Copied and UpdatedButUnmerged edge cases

	switch s.Staging {
	case git.Added:
		return StatusAdded
	case git.Deleted:
		return StatusRemoved
	case git.Modified:
		return StatusModified
	}

	switch s.Worktree {
	case git.Deleted:
		return StatusRemoved
	case git.Modified:
		return StatusModified
	}

	return StatusUnchanged
}
