package taskdmn

import (
	"encoding/base64"
	"encoding/json"

	"github.com/google/uuid"
)

type PageToken string

type PageTokenData struct {
	LastID uuid.UUID `json:"last_id"`
}

func PageTokenFromString(value string) (PageToken, error) {
	token := PageToken(value)

	if _, err := DecodePageToken(token); err != nil {
		return "", err
	}

	return token, nil
}

func EncodePageToken(data PageTokenData) (PageToken, error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return PageToken(base64.URLEncoding.EncodeToString(raw)), nil
}

func DecodePageToken(token PageToken) (*PageTokenData, error) {
	raw, err := base64.URLEncoding.DecodeString(string(token))
	if err != nil {
		return nil, err
	}

	var data PageTokenData
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
