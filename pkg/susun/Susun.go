package susun

import (
	"fmt"
	"github.com/panjf2000/ants"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Job struct {
	Filename string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "0123456789abcdef"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func randStringBytesMask() string {
	n := 5
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func move(content, src, destRootDir string, rule Rule) error {
	_, srcFile := path.Split(src)
	if tm, err := ParseTimes(content); err == nil {
		srcFile = tm.Format("2006-01-02") + " " + rule.DocumentType + " " + randStringBytesMask() + path.Ext(srcFile)
	} else {
		srcFile = "unknowndate " + strings.TrimSuffix(srcFile, path.Ext(srcFile)) + " " + rule.DocumentType + " " + randStringBytesMask() + path.Ext(srcFile)
	}
	destDir := path.Join(destRootDir, rule.TargetDir)
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Println("Moving from", src, "to", path.Join(destDir, srcFile))
	if err := os.Rename(src, path.Join(destDir, srcFile)); err != nil {
		return err
	}
	return nil
}

func Process(pattern, destRootDir string, rules []Rule) error {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	result := make(map[string]string)
	mutex := &sync.Mutex{}
	poolSize := runtime.NumCPU()*2 + 2

	pool, _ := ants.NewPoolWithFunc(poolSize, func(payload interface{}) {
		if job, ok := payload.(*Job); ok {
			str, err := ExtractPdfContent(job.Filename)
			if err != nil {
				log.Println(fmt.Sprintf("Unable to extract PDF content from: %s", job.Filename))
			} else {
				fmt.Println(path.Split(job.Filename))

				if err != nil {
					panic(err)
				}

				mutex.Lock()
				if rule, hasRule := MatchRule(rules, str); hasRule {
					if err := move(str, job.Filename, destRootDir, rule); err != nil {
						panic(err)
					}
				} else {
					fmt.Println(job.Filename, str)
				}
				result[job.Filename] = str
				mutex.Unlock()
			}
		}
		wg.Done()
	})
	defer pool.Release()

	for _, match := range matches {
		job := &Job{Filename: match}
		wg.Add(1)
		if err := pool.Serve(job); err != nil {
			log.Println("throttle limit error")
		}
	}
	wg.Wait()

	return nil

}
