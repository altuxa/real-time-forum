package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type JWT struct {
	header  header
	payload payload
}

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type payload struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
	Exp  int    `json:"exp"`
}

func NewJWT() *JWT {
	jwt := &JWT{
		header: header{
			Alg: "HS256",
			Typ: "JWT",
		},
		payload: payload{},
	}
	return jwt
}

func EncodeBase64(data []byte) []byte {
	base := base64.RawURLEncoding.EncodeToString(data)
	return []byte(base)
}

func (j *JWT) NewPayload(sub, name string, exp int) {
	j.payload.Sub = sub
	j.payload.Name = name
	j.payload.Exp = exp
}

func (j *JWT) UnsignedToken() (string, error) {
	header, err := json.Marshal(j.header)
	if err != nil {
		return "", err
	}
	payload, err := json.Marshal(j.payload)
	if err != nil {
		return "", err
	}
	res := string(EncodeBase64(header)) + "." + string(EncodeBase64(payload))
	return res, nil
}

func (j *JWT) SignedToken(secret string) (string, error) {
	unsigned, err := j.UnsignedToken()
	if err != nil {
		return "", err
	}
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(unsigned))
	sign := EncodeBase64(hash.Sum(nil))
	signed := unsigned + "." + string(sign)
	return signed, nil
}

func (j *JWT) Verify(token, secret string) error {
	signedToken, err := j.SignedToken(secret)
	if err != nil {
		return err
	}
	tokenFromClient := strings.Split(token, ".")
	tokenFromServer := strings.Split(signedToken, ".")
	signFromClient := tokenFromClient[2]
	signFromServer := tokenFromServer[2]
	if signFromClient != signFromServer {
		return errors.New("invalid token")
	}
	return nil
}
