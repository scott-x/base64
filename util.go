/*
* @Author: scottxiong
* @Date:   2020-06-29 20:46:03
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-06-29 20:46:33
 */
package base64

import (
	"crypto/rand"
	"fmt"
	"io"
)

func uuid() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x%x%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
