package global

import (
	"testing"

	"github.com/monkeswag33/noter-go/types"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hash, err := HashPass("password", &types.HashParams{
		Memory:      64 * 1024,
		Iterations:  4,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestVerify(t *testing.T) {
	var hashed string = "$argon2id$v=19$m=65536,t=4,p=2$3luS5dgOuFoEED4sov3Dag$vJT+rcXjMd9S/bkHkRBcUjejm4TlNIIY3p1cJgQ6d2o" // equals: "password"
	match, err := VerifyPass("password", hashed)
	assert.NoError(t, err)
	assert.True(t, match)
}
