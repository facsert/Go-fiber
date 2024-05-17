package messages

import (
	"bufio"
	"io"
	"archive/zip"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

type Rule interface {
	Scan(line *Line) error
	Write(zip *zip.Writer) error
}


var mu sync.RWMutex

// frequent, blacklist
var NewRules = map[string](func(filename string) Rule){
	"frequent": NewFrequent,
	"blacklist": NewBlackList,
}

type Line struct {
	Content   string 
	LineList  []string 
	Timestamp string
	Module    string
}

type Messages struct {
	Filename   string
	Reader     io.Reader
	Rules      []Rule
	wg         sync.WaitGroup
}

func NewMessages(reader io.Reader, filename string, rules []string) *Messages {
    m := &Messages{
        Filename: filename,
		Reader: reader,
		Rules: make([]Rule, 0, len(NewRules)),
	}
	for _, name := range rules {
		if newFunc, ok := NewRules[name]; ok {
			m.Rules = append(m.Rules, newFunc(filename))
		}
	}
	return m
}

// 筛选符合格式的行
func (m *Messages) filterLine(line string) (*Line, bool) {
    lineList := strings.Fields(line)
	if len(lineList) < 8 { return &Line{}, false }
	if lineList[2] != "ecs_ua_server:" { return &Line{}, false }
	return &Line{
		Content: line,
		LineList: lineList,
		Timestamp: lineList[0],
		Module: lineList[5],
	}, true
}

// 读取内容写入管道
func (m *Messages) Read(ch chan *Line) {
	scanner, line := bufio.NewScanner(m.Reader), ""
	for scanner.Scan() {
		line = scanner.Text()
		l,  conform := m.filterLine(line)
		if !conform { continue }
		ch <- l
	}
	close(ch)
	m.wg.Done()
}

// 并发协程池子(100), 扫描日志
func (m *Messages) Scan() error {
    var tokenPool = make(chan struct{}, 100)
	for i := 0; i < 100; i++ {
		tokenPool <- struct{}{}
	}

	chLine := make(chan *Line, 100)
	m.wg.Add(1)
	go m.Read(chLine)
    
	for line := range chLine {
		for _, rule := range m.Rules {
            <- tokenPool
			m.wg.Add(1)
			go func(line *Line, rule Rule) {
				err := rule.Scan(line)
				if err != nil { log.Error(err) }
				tokenPool <- struct{}{}
				m.wg.Done()
			}(line, rule)
		}
	}
	m.wg.Wait()
	close(tokenPool)
	return nil
}

func (m *Messages)Write(zip *zip.Writer) error {
	for _, rule := range m.Rules {
		err := rule.Write(zip)
		if err != nil { return err }
	}
	return nil
}