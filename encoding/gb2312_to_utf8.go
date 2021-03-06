package encoding

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func Gb2312ToUtf8(raw string) (string, error) {
	var charDecoding transform.Transformer
	charDecoding = simplifiedchinese.HZGB2312.NewDecoder()

	result, _, err := transform.String(charDecoding, raw)
	if err != nil {
		return raw, err
	}
	return result, nil
}
