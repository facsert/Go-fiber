package messages

import (
	"fmt"
	"path/filepath"
	"strings"
	"io"
	"archive/zip"

	"github.com/gofiber/fiber/v2/log"
)

type BlackList struct {
	name string
	filename string
	fileType string
	keywords [3]string
	lines     *strings.Builder
	count    int
}

func NewBlackList(filename string) Rule {
	builder := strings.Builder{}
	builder.Grow(100000)
	builder.WriteString("filename,error,module,timestamp,log\n")
	return &BlackList{
		filename: filename,
        name:     "BlackList",
        fileType: "messages",
        keywords: [3]string{"timeout", "panic", "error"},
		lines: &builder,
		count: 0,
	}
}

func (b *BlackList) Scan(line *Line) error {
    for _, key := range b.keywords {
        if strings.Contains(line.Content, key) {
			mu.Lock()
            b.lines.WriteString(fmt.Sprintf("%s,%s,%s,%s,\"%s\"\n",
				filepath.Base(b.filename),
				key,
				line.Module,
				line.Timestamp,
				line.Content,
			))
			b.count += 1
			mu.Unlock()
			return nil
		}
	}
	return nil
}

func (b *BlackList) Write(zip *zip.Writer) error {
	log.Infof("%s: BlackList=%d\n", b.filename, b.count)
	if b.count == 0 { return nil }
    
	file, err := zip.Create(fmt.Sprintf("%s_%s.csv", filepath.Base(b.filename), b.name))
    if err != nil {return err}
    
	if _, err := io.Copy(file, strings.NewReader(b.lines.String())); err != nil {
		return err
	}
	return nil
}