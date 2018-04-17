package stores

type URLStore interface {
	Create(string, string) (error)
	Find(string) (string, error)
}
