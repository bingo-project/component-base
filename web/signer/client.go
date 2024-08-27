package signer

type Client struct {
	KeySK     string
	Algorithm Algorithm
}

type Algorithm string

const (
	AlgorithmMd5 = "md5"
)

// New default tls.Config{InsecureSkipVerify: true}
func New() (client *Client) {
	client = &Client{
		KeySK: "secret_key",
	}

	return client
}

func (c *Client) SetKeySK(data string) (client *Client) {
	c.KeySK = data

	return c
}

func (c *Client) SetAlgorithm(data Algorithm) (client *Client) {
	c.Algorithm = data

	return c
}
