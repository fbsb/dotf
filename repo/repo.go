package repo

type FileStatus byte

const (
	StatusUnchanged FileStatus = ' '
	StatusAdded     FileStatus = '+'
	StatusRemoved   FileStatus = '-'
	StatusModified  FileStatus = '*'
)

func (fs FileStatus) String() string {
	return string(fs)
}

type Status interface {
	File(path string) FileStatus
	String() string
	IsClean() bool
}

type Repository interface {
	Add(path string) error
	Remove(path string) error
	Status() (Status, error)
	Commit(message, name, email string) error
}
