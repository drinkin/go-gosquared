package gosquared_test

import (
	"testing"

	"github.com/drinkin/go-gosquared/gosquared"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, 1)

	client := gosquared.NewClient("", "")
	client.Identify("1", &gosquared.Person{
		ID: "1",
	})
}
