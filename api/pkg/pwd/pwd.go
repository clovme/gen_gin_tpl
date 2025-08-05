package pwd

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func _md5(value string) string {
	value = fmt.Sprintf("%s%s", value, "ViewsRegeditPostHandler-fm")
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}

func _base64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func Encryption(value string) string {
	value = _md5(value)
	value = _base64(value)
	value = _md5(value)
	return strings.ReplaceAll(_base64(value), "=", "")
}
