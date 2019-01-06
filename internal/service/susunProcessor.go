package service

import (
	"fmt"
	"github.com/faizhasim/susunplease/internal/model"
	"github.com/faizhasim/susunplease/internal/sideEffect"
	"github.com/faizhasim/susunplease/internal/utils"
	"github.com/panjf2000/ants"
	"log"
	"path/filepath"
	"runtime"
	"sync"
)

type SusunProcessor interface {
	Process(pattern string, rules []model.Rule) error
}

type susunProcessor struct {
	sideEffect.PdfService
	sideEffect.FsFiling
	RulesParser
}

type job struct {
	filename string
}

func (susunProcessor *susunProcessor) Process(pattern string, rules []model.Rule) error {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	result := make(map[string]string)
	mutex := &sync.Mutex{}
	poolSize := runtime.NumCPU()*2 + 2

	pool, _ := ants.NewPoolWithFunc(poolSize, func(payload interface{}) {
		if job, ok := payload.(*job); ok {
			str, err := susunProcessor.ParsePdfContent(job.filename)
			if err != nil {
				log.Println(fmt.Sprintf("Unable to extract PDF content from: %s", job.filename))
			} else {
				mutex.Lock()
				if rule, hasRule := susunProcessor.MatchRule(rules, str); hasRule {

					tm, err := utils.ParseTimes(str)
					if err != nil {
						log.Println(fmt.Sprintf("Unable to extract PDF content from: %s", job.filename))
					}

					if err := susunProcessor.MoveFile(&tm, job.filename, &rule); err != nil {
						panic(err)
					}
				} else {
					fmt.Println(job.filename, str)
				}
				result[job.filename] = str
				mutex.Unlock()
			}
		}
		wg.Done()
	})
	defer pool.Release()

	for _, match := range matches {
		job := &job{filename: match}
		wg.Add(1)
		if err := pool.Serve(job); err != nil {
			log.Println("throttle limit error")
		}
	}
	wg.Wait()

	return nil
}

func NewSusunProcessor(pdfService sideEffect.PdfService, fsFiling sideEffect.FsFiling, rulesParser RulesParser) SusunProcessor {
	return &susunProcessor{pdfService, fsFiling, rulesParser}
}
