package generating

// Speaker represents a speaker in a particular report.
type Speaker struct {
	ID         string
	Name       string
	ProfileURL string
	PictureURL string
	Location   string
	Side       string
}

// NewSpeaker instantiates a speaker. The ID is basically the short name of the speaker.
func NewSpeaker(id, url string) *Speaker {
	return &Speaker{
		ID:         id,
		ProfileURL: url,
	}
}
