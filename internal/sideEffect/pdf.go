package sideEffect

import (
	"bytes"
	"github.com/ledongthuc/pdf"
)

type PdfService interface {
	ParsePdfContent(path string) (string, error)
}

type pdfService struct{}

func (pdfService *pdfService) ParsePdfContent(path string) (string, error) {
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

func NewPdfService() PdfService {
	return &pdfService{}
}
