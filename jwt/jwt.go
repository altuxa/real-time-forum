package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type jwt struct {
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
	Iat  int    `json:"iat"`
}

func NewJWT() *jwt {
	jwt := &jwt{
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

func (j *jwt) NewPayload(sub, name string, admin int) {
	j.payload.Sub = sub
	j.payload.Name = name
	j.payload.Iat = admin
}

func (j *jwt) UnsignedToken() (string, error) {
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

func (j *jwt) SignedToken(secret string) (string, error) {
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
