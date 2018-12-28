package susun

import (
	"bytes"
	"fmt"
	"github.com/ledongthuc/pdf"
	"log"
	"path/filepath"
)

func ExtractPdfContent(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer

	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}

	_, err = buf.ReadFrom(b)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ExtractPdfFromGlob(pattern string) (map[string]string, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, match := range matches {
		str, err := ExtractPdfContent(match)
		if err != nil {
			log.Println(fmt.Sprintf("Unable to extract PDF content from: %s", str))
		} else {
			result[match] = str
		}
	}
	return result, nil
}
