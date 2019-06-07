package generating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntervention(t *testing.T) {
	i := NewIntervention("1", "John Doe", "Say yes!")

	assert.Equal(t, "1", i.ID())
	assert.Equal(t, InterventionType, i.Type())
}

func TestAppendToIntervention(t *testing.T) {
	i := NewIntervention("", "", "")
	c := NewIntervention("", "", "")

	i.Append(c)

	assert.Len(t, i.Children(), 0, "Intervention should not have any children")
}
