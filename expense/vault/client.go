package vault

import (
	"encoding/base64"
	"fmt"

	vaultclient "github.com/hashicorp/vault/api"
)

type Client struct {
	client *vaultclient.Client
}

func New(addr, token string) (*Client, error) {
	config := vaultclient.DefaultConfig()
	config.Address = addr

	client, err := vaultclient.NewClient(config)
	if err != nil {
		return nil, err
	}

	client.SetToken(token)

	return &Client{client}, nil
}

type EncryptRequest struct {
	PlainText string `json:"plaintext"`
}

type EncryptResponse struct {
	RequestID string      `json:"request_id"`
	Data      EncryptData `json:"data"`
}

type EncryptData struct {
	CipherText string `json:"ciphertext"`
}

// EncyptData with Vaults transit secret engine
func (c *Client) EncryptData(data, key string) (string, error) {

	//base64 the data
	data = base64.StdEncoding.EncodeToString([]byte(data))

	req := c.client.NewRequest("POST", fmt.Sprintf("/v1/transit/encrypt/%s", key))
	err := req.SetJSONBody(&EncryptRequest{data})
	if err != nil {
		return "", fmt.Errorf("unable to set json payload for encyrpt request data: %s", err)
	}

	resp, err := c.client.RawRequest(req)
	if err != nil {
		return "", fmt.Errorf("unable to execute request for encrypt data: %s", err)
	}

	er := &EncryptResponse{}
	err = resp.DecodeJSON(er)
	if err != nil {
		return "", fmt.Errorf("unable to decode json payload for encyrpt data response: %s", err)
	}

	return er.Data.CipherText, nil
}

type DecryptRequest struct {
	CipherText string `json:"ciphertext"`
}

type DecryptResponse struct {
	Data DecryptData `json:"data"`
}
type DecryptData struct {
	PlainText string `json:"plaintext"`
}

// DecryptData with Vaults transit secret engine
func (c *Client) DecryptData(data, key string) (string, error) {
	if data == "" {
		return "", nil
	}

	req := c.client.NewRequest("POST", fmt.Sprintf("/v1/transit/decrypt/%s", key))
	err := req.SetJSONBody(&DecryptRequest{data})
	if err != nil {
		return "", fmt.Errorf("unable to set json payload for decyrpt request data: %s", err)
	}

	resp, err := c.client.RawRequest(req)
	if err != nil {
		return "", fmt.Errorf("unable to execute request for decrypt data: %s", err)
	}

	dr := &DecryptResponse{}
	err = resp.DecodeJSON(dr)
	if err != nil {
		return "", fmt.Errorf("unable to decode json payload for encyrpt data response: %s", err)
	}

	d, err := base64.StdEncoding.DecodeString(dr.Data.PlainText)

	return string(d), nil
}
