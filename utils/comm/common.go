package comm

import (
	"fmt"
	"os"
	"log/slog"
	"path/filepath"
	"strings"
)

var (
	ROOT_PATH string
)

func Init() {
    var err error
    ROOT_PATH, err = os.Getwd()
    if err != nil { panic("Failed to get current path")}
}

// 基于根目录的绝对路径
func AbsPath(elems ...string) string {
    for i := (len(elems)-1); i > 0; i-- {
		if filepath.IsAbs(elems[i]) {
            elems = elems[i:]
            break
        }
    }
    path := filepath.Join(elems...)
    if filepath.IsAbs(path) {
        return path
    }
    return filepath.Join(ROOT_PATH, path)
}

func MakeDirs(path string) error {
    if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("create %s failed: %v", path, err)
		}
    } else if err != nil {
		return fmt.Errorf("check %s exist failed: %v", path, err)
    }
    return nil
}

// 标题打印
func Title(title string, level int) string {
	separator := [...]string{"#", "=", "*", "-"}[level % 3]
	space := [...]string{"\n\n", "\n", "", ""}[level % 3]
	line := strings.Repeat(separator, 80)
	slog.Info(fmt.Sprintf("%s%s %s %s\n", space, line, title, line))
	return title
}

// 阶段性结果打印
func Display(msg string, success bool) string {
	length, chars := 80, ""
	if success {
		slog.Info(fmt.Sprintf("%-" + fmt.Sprintf("%d", length) + "s  [PASS]\n", msg))
		return msg
	}
	if len(msg) > length {
		chars = strings.Repeat(">", length - len(msg))
	}
	fmt.Println(msg + " " + chars + " [FAIL]")
	return msg
}


// 获取指定路径下的所有文件
func ListDir(root string) []string {
    files := []string{}
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil { return err }
		if !info.IsDir() { files = append(files, path) }
		return nil
    })
    if err != nil { panic(fmt.Sprintf("Error: %v\n", err)) }
    return files
}