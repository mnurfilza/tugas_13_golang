package tugas13

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func Tugas13(name string) string {
	encodeString := base64.StdEncoding.EncodeToString([]byte(name))
	sha := sha1.New()
	sha.Write([]byte(encodeString))
	encrypt := sha.Sum(nil)
	encrypted := fmt.Sprintf("%x", encrypt)

	return encrypted
}
