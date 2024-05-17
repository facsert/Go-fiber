package messages

import (
	"archive/zip"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"io"

	"github.com/gofiber/fiber/v2/log"
)


type Log struct {
	timestamp time.Time
	module string
}

type Frequent struct {
	name string
	filename string
	fileType string
	time_slice_second time.Duration
    threshold_count int
	lines *strings.Builder
	logGroup  map[string][]*Log
	count int
}


func NewFrequent(filename string) Rule {
	builder := strings.Builder{}
	builder.Grow(100000)
	builder.WriteString("filename,error,module,type,count,start,close,log\n")
	return &Frequent{
		filename: filename,
        name:     "Frequent",
        fileType: "messages",
        time_slice_second: 3 * time.Second,
		threshold_count: 6,
		lines: &builder,
		logGroup: make(map[string][]*Log, 100000),
		count: 0,
	}
}

func (f *Frequent) Scan(line *Line) error {
    lineTime, err := time.Parse(time.RFC3339, line.Timestamp)
	if err != nil {
		log.Error(fmt.Sprintf("error timestamp: %s", line.Timestamp))
		return err
	}

	content := strings.Join(line.LineList[4:], " ")
    mu.Lock()
	f.logGroup[content] = append(f.logGroup[content], &Log{
		timestamp: lineTime,
        module: line.Module,
	})
	mu.Unlock()
	return nil
}

func (f *Frequent) FilterFrequent() {
	for content, logs := range f.logGroup {
		if len(logs) < f.threshold_count { continue }

		sort.Slice(logs, func(x, y int) bool {
            return logs[x].timestamp.Before(logs[y].timestamp)
		})

		left, right, length := 0, f.threshold_count-1, len(logs)
		for right < length {
            if logs[right].timestamp.Sub(logs[left].timestamp) < f.time_slice_second {
				f.lines.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,\"%s\"\n",
					filepath.Base(f.filename),
					f.name,
					logs[left].module,
					f.fileType,
					strconv.Itoa(length),
					logs[0].timestamp.Format(time.RFC3339),
					logs[length-1].timestamp.Format(time.RFC3339),
					content,
				))
				f.count += 1
				break
			}
			left, right = left + 1, right + 1
		}
	} 
}

func (f *Frequent) Write(zip *zip.Writer) error {
	f.FilterFrequent()

    log.Infof("%s: Frequent=%d\n", f.filename, f.count)
	if f.count == 0 { return nil }

	file, err := zip.Create(fmt.Sprintf("%s_%s.csv", filepath.Base(f.filename), f.name))
    if err != nil {return err}
    
	if _, err := io.Copy(file, strings.NewReader(f.lines.String())); err != nil {
		return err
	}
	return nil
}