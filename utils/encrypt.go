package utils

import "encoding/base64"

const offset = byte(100)

func DSMEncrypt(text string) string {
	textBytes := []byte(text)
	for i := 0; i < len(textBytes)-4; i += 4 {
		textBytes[i], textBytes[i+2] = textBytes[i+2], textBytes[i]
		textBytes[i+1], textBytes[i+3] = textBytes[i+3], textBytes[i+1]
		for j := i; j < i+3; j++ {
			if textBytes[j] <= (255 - offset) {
				textBytes[j] += offset
			} else {
				// 254 + 3 => 254-256+3
				textBytes[j] = textBytes[j] - byte(int(256)-int(offset))
			}
		}
	}
	l := len(textBytes)
	h := len(textBytes)/2
	tmp := make([]byte, h)
	copy(tmp, textBytes[:h])
	copy(textBytes[:h], textBytes[l-h:])
	copy(textBytes[l-h:], tmp)
	return base64.StdEncoding.EncodeToString(textBytes)
}

func DSMDecrypt(text string) string {
	textBytes, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}
	l := len(textBytes)
	h := len(textBytes)/2
	tmp := make([]byte, h)
	copy(tmp, textBytes[:h])
	copy(textBytes[:h], textBytes[l-h:])
	copy(textBytes[l-h:], tmp)

	for i := 0; i < len(textBytes)-4; i += 4 {
		for j := i; j < i+3; j++ {
			if textBytes[j] >= offset {
				textBytes[j] -= offset
			} else {
				textBytes[j] = byte(int(256) + int(textBytes[j]) - int(offset))
			}
		}
		textBytes[i], textBytes[i+2] = textBytes[i+2], textBytes[i]
		textBytes[i+1], textBytes[i+3] = textBytes[i+3], textBytes[i+1]
	}
	return string(textBytes)
}
