package argon2

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/monkeswag33/noter-go/errordef"
	"github.com/monkeswag33/noter-go/types"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/argon2"
)

func HashPass(password string, params *types.HashParams) (encodedHash string, err error) {
	salt, err := genSalt(params.SaltLength)
	if err != nil {
		return "", err
	}
	logrus.Tracef("Generated salt %d bytes long", params.SaltLength)

	var hash []byte = argon2.IDKey([]byte(password),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength)
	logrus.Trace("Hashed password using argon2 to bytes")
	var b64Salt string = base64.RawStdEncoding.EncodeToString(salt)
	var b64Hash string = base64.RawStdEncoding.EncodeToString(hash)
	logrus.Trace("Converted salt and hash to strings")
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, params.Memory, params.Iterations, params.Parallelism, b64Salt, b64Hash)
	logrus.Trace("Generated string hash")
	return encodedHash, nil
}

func genSalt(saltLength uint32) (bytes []byte, err error) {
	bytes = make([]byte, saltLength)
	_, err = rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func VerifyPass(password, encodedHash string) (match bool, err error) {
	params, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	var passwordHash []byte = argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	if subtle.ConstantTimeCompare(hash, passwordHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (params *types.HashParams, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errordef.ErrArgon2InvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errordef.ErrArgon2IncompatibleVersion
	}

	params = &types.HashParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	params.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	params.KeyLength = uint32(len(hash))

	return params, salt, hash, nil
}
