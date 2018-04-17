package services

type PathExtractor struct{}

func NewPathExtractor() *PathExtractor {
	return &PathExtractor{}
}

func (service *PathExtractor) Call(path string) string {
	return path[1:]
}
