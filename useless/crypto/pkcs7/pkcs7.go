// PKCS7 implements PKCS#7 padding, described at http://tools.ietf.org/html/rfc5652#section-6.3.
package pkcs7

import (
	"bytes"
	"errors"
)

// Pad adds padding, with each padded byte being the total number of bytes added.
func Pad(b []byte, size int) (out []byte, err error) {
	if size < 1 || size >= 256 {
		return nil, errors.New("useless/crypto/pkcs7: block size is out of bounds")
	}

	padding := size - len(b)%size
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(b, padtext...), nil
}

// UnPad removes padding.
func UnPad(b []byte, size int) (out []byte, err error) {
	if size < 1 || size >= 256 {
		return nil, errors.New("useless/crypto/pkcs7: block size is out of bounds")
	}

	if len(b)%size != 0 {
		return nil, errors.New("useless/crypto/pkcs7: byte length is not a multiple of the block size")
	}
	return b[:len(b)-1], nil
}
