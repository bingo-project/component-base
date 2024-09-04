package signer

type Client struct {
	KeySK     string
	Algorithm Algorithm
}

type Algorithm string

const (
	KeySkDefault = "secret_key"

	AlgorithmMd5    = "md5"
	AlgorithmSha256 = "sha256"
)

// New default tls.Config{InsecureSkipVerify: true}
func New() (client *Client) {
	client = &Client{
		KeySK:     "secret_key",
		Algorithm: AlgorithmMd5,
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
