package comm

import (
	"fmt"
	"io"
	"os"
	"archive/zip"
	"path/filepath"
)

func Extract(file, dir string) error { return nil }

func Compress(files []string, output string) error { return nil }

func CompressZip(files []string, output string) error {
	tar, err := os.Create(output)
	if err != nil { return fmt.Errorf("create %s failed: %v", output, err) }
	defer tar.Close()
    
	writer := zip.NewWriter(tar)
	defer writer.Close()
    
    for _, filename := range files {
		err := func(filename string) error {
			fmt.Printf("compress: %s\n", filename)
			
			w, err := writer.Create(filepath.Base(filename))
			if err != nil {
				return fmt.Errorf("create %s in archive failed: %v", filename, err)
			}
		
			file, err := os.Open(filename)
			if err != nil { 
				return fmt.Errorf("open %s failed: %v", filename, err) 
			}
			defer file.Close()
	
			_, err = io.Copy(w, file)
			if err != nil {
				return fmt.Errorf("copy %s content to archive failed: %v", filename, err) 
			}
			return nil
		}(filename)
        
		if err != nil { return err }
	}
	return nil
}

func CompressGz() {}



