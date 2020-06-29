/*
* @Author: scottxiong
* @Date:   2020-06-29 20:02:41
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-06-29 21:01:49
* ref: https://stackoverflow.com/questions/38648512/go-saving-base64-string-to-file?rq=1
 */

package base64

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	ErrSize         = errors.New("Invalid size!")
	ErrInvalidImage = errors.New("Invalid image!")
	ErrUUID         = errors.New("UUID error")
)

func SaveImageToDisk(fileNameBase, data string) (string, error) {
	index := strings.Index(data, ";base64,")
	if index < 0 {
		return "", ErrInvalidImage
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[index+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}
	_, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return "", err
	}

	// if imgCfg.Width != 750 || imgCfg.Height != 685 {
	// 	return "", ErrSize
	// }

	random, err := uuid()
	if err != nil {
		return "", ErrUUID
	}

	fileName := path.Join(fileNameBase, random+"."+fm)

	//create folder if not exists
	if _, err = os.Stat(fileNameBase); os.IsNotExist(err) {
		err = os.MkdirAll(fileNameBase, 0755)
		if err != nil {
			panic(err)
		}
	}

	//write
	ioutil.WriteFile(fileName, buff.Bytes(), 0644)
	return fileName, err
}
