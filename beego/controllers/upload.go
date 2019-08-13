package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// UsersController operations for Users
type UploadController struct {
	BaseController
}

// URLMapping ...
func (c *UploadController) URLMapping() {
	c.Mapping("Upload", c.Upload)
}

// upload 上传文件
// @Title 上传文件
// @Description 接收上传的文件并返回资源URL
// @Param	token	query 		string	true		"令牌"
// @Param	file	formData 	file	false		"要上传的文件"
// @Param	fileurl	formData 	string	false		"要上传的文件的URL"
// @Success 0 ...
// @Failure -1 ...
// @router /upload [post]
func (c *UploadController) Upload() {
	//c.ParseToken(c.GetString("token"))

	fileUrl := c.GetString("fileurl")

	if fileUrl != "" {
		fileExten := "jpg"
		filePath := fmt.Sprintf("static/upload/image/%d/%d/%d/%s.%s",
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			c.CreateRandStr(10, 1),
			fileExten,
		)

		req := httplib.Get(fileUrl)
	 	err := req.ToFile(filePath)
		if err != nil {
			c.ReturnMsg(-1, nil, err.Error())
		}

		c.ReturnMsg(0, map[string]string{"fileUrl": beego.AppConfig.String("hostdomin") + filePath}, "")
	}

	fp, fileInfo, err := c.GetFile("file")
	if err != nil {
		beego.Error("上传失败，报错信息：" + err.Error())
		c.ReturnMsg(-1, nil, "上传失败：" + err.Error())
	}
	defer fp.Close()

	filePath, _, err := c.UploadFile(fileInfo, "file")
	if err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	}
	c.ReturnMsg(0, map[string]string{"fileUrl": beego.AppConfig.String("hostdomin") + filePath}, "")
}
