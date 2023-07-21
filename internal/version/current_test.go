package version

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrent(t *testing.T) {
	v, err := Current("../..")
	assert.NoError(t, err)
	assert.NotNil(t, v)
	assert.Regexp(t, regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`), v.String())
}
