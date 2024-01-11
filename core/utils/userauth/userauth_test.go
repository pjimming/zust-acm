package userauth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var username = "panjiangming"

func TestToken(t *testing.T) {
	ast := assert.New(t)

	token, err := GenToken(username)
	ast.Nil(err)
	ast.NotEqual("", token)

	uc, err := ParseToken(token)
	ast.Nil(err)
	ast.NotNil(uc)
}
