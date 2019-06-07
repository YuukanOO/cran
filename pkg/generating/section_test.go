package generating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSection(t *testing.T) {
	s := NewSection("1", "A section", 1)

	assert.Equal(t, "1", s.ID())
	assert.Equal(t, SectionType, s.Type())
}

func TestAddChildrenToSection(t *testing.T) {
	s1 := NewSection("1", "A section", 1)
	s2 := NewSection("2", "Another section", 2)
	s3 := NewSection("3", "Another section", 2)

	s1.Append(s2, s3)

	assert.Len(t, s1.Children(), 2)
	assert.Equal(t, s2, s1.Children()[0])
	assert.Equal(t, s3, s1.Children()[1])
}
