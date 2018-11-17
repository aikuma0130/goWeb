package backup

type Archiver interface {
	Archive(src, dst string) error
}

type zipper struct{}

var ZIP Archiver = (*zipper)(nil)
