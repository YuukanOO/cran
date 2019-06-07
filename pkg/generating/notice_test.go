package generating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotice(t *testing.T) {
	n := NewNotice("1", "Some content")

	assert.Equal(t, "1", n.ID())
	assert.Equal(t, NoticeType, n.Type())
}

func TestAppendToNotice(t *testing.T) {
	n := NewNotice("", "")
	c := NewNotice("", "")

	n.Append(c)

	assert.Len(t, n.Children(), 0, "Notice should not have any children")
}
