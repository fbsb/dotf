package repo

import "gopkg.in/src-d/go-git.v4"

type Repository interface {
	Add(path string) error
	Status() (git.Status, error)
	Commit(message, name, email string) error
}
