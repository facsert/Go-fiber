package scan

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"bufio"
	"strings"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/log"

    "fibert/lib/comm"
    "fibert/utils/messages"
)

func Init(api fiber.Router) {

	router := api.Group("/scan")
	router.Post("/file", UploadFile)
	router.Post("/log", UploadLog)
	router.Post("/gz", UploadGz)
}

// @tags     logScan
// @summary  Upload file
// @Param    file formData file true "upload file"
// @Router   /scan/file  [post]
func UploadFile(c *fiber.Ctx) error {
	tempFile, err := c.FormFile("file")
	if err != nil { return err }
    
	file, err := tempFile.Open()
	if err != nil { return err }

	log.Info(c.FormValue("rules"))
    
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Info(scanner.Text())
	}
	return c.JSON(map[string]string{
		"filename": tempFile.Filename,
	})
}

// @tags     logScan
// @summary  Upload file
// @Param    file formData file true "upload log"
// @Param    rules formData string true "scan rules"
// @Router   /scan/log  [post]
func UploadLog(c *fiber.Ctx) error {
    tempFile, err := c.FormFile("file")
    if err != nil { return err }
    
    file, err := tempFile.Open()
    if err != nil { return err }
    
    rules := strings.Fields(strings.ReplaceAll(c.FormValue("rules"), ",", " "))
    log.Infof("%s:%v", tempFile.Filename, rules)

    if strings.HasSuffix(tempFile.Filename, ".gz") {
        return c.SendString(fmt.Sprintf("%s is package, use scan/gz api", tempFile.Filename))
    }

    s := messages.NewMessages(file, tempFile.Filename, rules)
    s.Scan()

    var buff bytes.Buffer
    zipWrite := zip.NewWriter(&buff)

    err = s.Write(zipWrite)
    if err != nil { log.Error(err) }

    if err := zipWrite.Close(); err != nil { log.Error(err) }

    c.Set("Content-Type", "application/zip")
    c.Set("Content-Disposition", "attachment; filename=result.zip")
    return c.Send(buff.Bytes())
}

// @tags     logScan
// @summary  Upload gz
// @Param    file formData file true "upload package"
// @Param    rules formData string true "scan rules"
// @Router   /scan/gz  [post]
func UploadGz(c *fiber.Ctx) error {
    tempFile, err := c.FormFile("file")
    if err != nil { return err }
    
    if !strings.HasSuffix(tempFile.Filename, ".gz") {
        return c.SendString(fmt.Sprintf("%s is not package, use scan/log api", tempFile.Filename))
    }

    file, err := tempFile.Open()
    if err != nil { return err }

    rules := strings.Fields(strings.ReplaceAll(c.FormValue("rules"), ",", " "))
    log.Infof("%s:%v", tempFile.Filename, rules)

    gzReader, err := gzip.NewReader(file)
    if err != nil {return err}


    s := messages.NewMessages(gzReader, tempFile.Filename, rules)
    s.Scan()

    var buff bytes.Buffer
    zipWrite := zip.NewWriter(&buff)
    err = s.Write(zipWrite)
    if err != nil { log.Error(err) }

    if err := zipWrite.Close(); err != nil { log.Error(err) }

    c.Set("Content-Type", "application/zip")
    c.Set("Content-Disposition", "attachment; filename=result.zip")
    return c.Send(buff.Bytes())
}



func Download(c *fiber.Ctx) error {
    return c.Download(comm.AbsPath("temp", "download.txt"))

	// @tags     logScan
	// @summary  Download file
	// @Success  200 {file} octet-stream "download"
	// @Router   /logScan  [get]
}

