package generating

// Service provide an entry point to generate report based on an URL.
type Service interface {
	Generate(url string, callback DoneCallback)
}

type service struct{}

// NewService instantiates the default generating Service.
func NewService() Service {
	return &service{}
}

func (s *service) Generate(url string, callback DoneCallback) {
	p, err := GuessProvider(url)

	if err != nil {
		callback(nil, err)
		return
	}

	p.Fetch(url, callback)
}
