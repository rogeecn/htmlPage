package encoding

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/traditionalchinese"
)

func Big5ToUtf8(raw string) (string, error) {
	var charDecoding transform.Transformer
	charDecoding = traditionalchinese.Big5.NewDecoder()

	result, _, err := transform.String(charDecoding, raw)
	if err != nil {
		return raw, err
	}
	return result, nil
}
