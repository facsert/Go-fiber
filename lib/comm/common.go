 package comm

 import (
	 "fmt"
	 "io/fs"
	 "log/slog"
	 "path/filepath"
	 "runtime"
	 "strings"
 )
 
 var (
	 ROOT_PATH string
 )
 
 func init() {
	 _, file, _, ok := runtime.Caller(0)
	 if !ok { panic("Failed to get caller info") }
	 ROOT_PATH = filepath.Dir(filepath.Dir(filepath.Dir(file)))
 }
 
 // 基于根目录的绝对路径
 func AbsPath(elems ...string) string {
	 path := filepath.Join(elems...)
	 if filepath.IsAbs(path) {
		 return path
	 }
	 return filepath.Join(ROOT_PATH, path)
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
	 var files []string
	 var fn fs.WalkDirFunc = func(path string, d fs.DirEntry, err error) error {
		 if err != nil {
			 return err
		 }
		 if !d.IsDir() {
			 files = append(files, path)
		 }
		 return nil
	 }
	 filepath.WalkDir(root, fn)
	 return files
 }