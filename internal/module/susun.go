package module

import (
	"github.com/faizhasim/susunplease/internal/service"
	"github.com/faizhasim/susunplease/internal/sideEffect"
)

type SusunPdf struct {
	sideEffect.PdfService
	service.RulesParser
	sideEffect.FsFiling
	service.SusunProcessor
}
