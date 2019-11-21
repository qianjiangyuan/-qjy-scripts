package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller"
	"github.com/xxmyjk/xintong/backend/pkg/app/model"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/upload"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var BasePath string
type Params struct {
	Url string `json:"url" `
}
type QueryParams struct {
	MParams  Params  `json:"Params"`
	PageNum  int64 `json:"pagenum" `
	PageSize int64 `json:"pagesize" `
}
func HandleUploadMutiFile(c *gin.Context) {
	cors(c)
	err := c.Request.ParseMultipartForm(4 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件太大"})
		return
	}
	v := c.Request.URL.Query()
	taskPath := v.Get("path")
	results := []string{}
	formdata := c.Request.MultipartForm
	relativePath := c.PostForm("relativePath")
	filePath, err := createPath(taskPath + "/" + relativePath)
	if err != nil {

	}
	fmt.Println("filePath=" + filePath)
	files := formdata.File["file"]
	for _, v := range files {
		file, err := v.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}
		fileName := filePath + v.Filename
		results = append(results, fileName)
		err = ioutil.WriteFile(fileName, content, 777)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件写入失败"})
			return
		}
		os.Chmod(fileName, 0777)
	}
	c.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": results,
	})
}

func GetBasePath() string {
	return connect.Conf.GetString("filepath.path")
}

func GetPath(url string) string {
	return GetBasePath() + "/" + url
}

func GetAllFile(path string, pageNum int64, pageSize int64) (*upload.Files,  error) {
	files:=&upload.Files{}
	path = GetPath(path)
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return files,  nil
	}
	files.Total=len(rd)
	var inx int64 = 0
	for _, fi := range rd {
		if inx < (pageNum-1)*pageSize {
			continue
		}
		if inx >= pageNum*pageSize {
			break
		}
		file := model.File{}
		file.Name = fi.Name()
		file.IsDir = fi.IsDir()
		files.Files = append(files.Files, file)
		inx++
	}
	return files, nil
}
func List(ctx *gin.Context) {
	cors(ctx)
	query := QueryParams{}
	err := ctx.Bind(&query)

	if controller.ErrCatch(err, ctx) {
		return
	}
	fmt.Printf("path="+query.MParams.Url)

	files,  err := GetAllFile(query.MParams.Url, query.PageNum, query.PageSize)
	if controller.ErrCatch(err, ctx) {
		return
	}
	ctx.JSON(200, &gin.H{
	"code": 0,
	"msg":  "OK",
	"data": files,
	})
}
func createPath(relativePath string) (string, error) {
	fmt.Println("relativePath=" + relativePath)
	paths := strings.Split(relativePath, "/")
	basepath := GetBasePath() + "/"
	for inx, path := range paths {
		if inx < len(paths)-1 {
			basepath += path + "/"
			if !Exists(path) {
				CreatePath(basepath)
			}
		}
	}
	return basepath, nil
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreatePath(path string) error {

	err := os.Mkdir(path, os.ModePerm)
	return err
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, Access-Control-Allow-Origin,x-token,Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
}
