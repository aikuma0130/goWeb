package backup

type Monitor struct {
	Paths       map[string]string
	Archiver    Archiver
	Destination string
}
