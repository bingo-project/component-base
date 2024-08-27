package signer

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func (c *Client) Sign(params map[string]any, sk string) (sign string) {
	data := FormatURLParam(params)
	data = fmt.Sprintf("%s&%s=%s", data, c.KeySK, sk)

	var hash []byte
	if c.Algorithm == AlgorithmMd5 {
		h := md5.New()
		h.Write([]byte(data))
		hash = h.Sum(nil)
	}

	return hex.EncodeToString(hash)
}

func (c *Client) VerifySign(params map[string]any, secretKey string, sign string) bool {
	resign := c.Sign(params, secretKey)

	return sign == resign
}

func FormatURLParam(body map[string]any) (urlParam string) {
	var (
		buf  strings.Builder
		keys []string
	)

	for k := range body {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		v, ok := body[k].(string)
		if !ok {
			v = convertToString(body[k])
		}
		if v != "" {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}

	if buf.Len() <= 0 {
		return ""
	}

	return buf.String()[:buf.Len()-1]
}

func convertToString(v any) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)

	if bs, err = json.Marshal(v); err != nil {
		return ""
	}

	str = string(bs)

	return
}
