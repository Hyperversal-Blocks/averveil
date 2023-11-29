package publish

type publisher struct {
}

type Publisher interface {
}

func New() Publisher {
	return &publisher{}
}
