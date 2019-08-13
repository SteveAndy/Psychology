package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

// BaseController API基类
type BaseController struct {
	beego.Controller
}

// ReturnMsg 格式化接口返回信息
func (c *BaseController) ReturnMsg(status int, data interface{}, msg string) {
	if msg == "" {
		msg = "success"
	}

	c.Data["json"] = map[string]interface{}{"status": status, "data": data, "msg": msg}
	c.ServeJSON()
	c.StopRun()
}

// CreateToken 创建令牌
func (c *BaseController) CreateToken(userID int64) string {
	type UserInfo map[string]interface{}

	t := time.Now()
	key := beego.AppConfig.String("tokenkey")
	userInfo := make(UserInfo)

	userInfo["exp"] = strconv.FormatInt(t.UTC().UnixNano(), 10)
	userInfo["iat"] = "0"
	userInfo["aud"] = strconv.FormatInt(userID, 10)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for index, val := range userInfo {
		claims[index] = val
	}
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

// ParseToken 检查令牌并返回用户的ID
func (c *BaseController) ParseToken(tokenString string) string {
	// 开发时指定的超级token
	if tokenString == "hehe" {
		return "102"
	}

	if tokenString == "" {
		c.ReturnMsg(-1, nil, "access_token为空！")
	}

	t := time.Now()
	key := beego.AppConfig.String("tokenkey")
	// 设置token生存时间为3天，单位为纳秒
	var expTime int64 = 259200000000000
	var userID int64 = -1

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		oldT, _ := strconv.ParseInt(claims["exp"].(string), 10, 64)
		userID, _ = strconv.ParseInt(claims["aud"].(string), 10, 64)
		ct := t.UTC().UnixNano()
		diff := ct - oldT
		if diff > expTime {
			// token已过期
			beego.Error("token已过期，当前系统时间:" + strconv.FormatInt(ct, 10) + ",token的过期时间：" + strconv.FormatInt(oldT, 10) + ",时间差：" + strconv.FormatInt(diff, 10) + ",token生存时间：" + strconv.FormatInt(expTime, 10))
			c.ReturnMsg(-1, nil, "token已过期")
		}
	} else {
		// token非法
		beego.Error("token非法，接口返回错误")
		c.ReturnMsg(-1, nil, "token非法")
	}
	// token正常
	return strconv.FormatInt(userID, 10)
}

// SendMailAuth 发送邮件验证码
func (c *BaseController) SendMailAuth(email string, activCode string) error {
	// 邮件发送
	htmlText := fmt.Sprintf(`<html><head></head><body><div>你的激活码为： %s</div></body></html>`, activCode)
	temail := utils.NewEMail(beego.AppConfig.String("mailconfig"))
	temail.To = []string{email}                      // 指定收件人邮箱地址
	temail.From = beego.AppConfig.String("sendname") // 指定发件人的邮箱地址
	temail.Subject = "cosplay作品分享平台"                 // 指定邮件的标题
	temail.HTML = htmlText                           // 指定邮件正文

	err := temail.Send()
	if err != nil {
		return errors.New("邮件发送失败：" + err.Error())
	}
	return nil
}

// CreateRandStr 生成随机字符串
// 参数 len：随机字符串的长度
// 参数 level：0、生成纯数字随机字符串，1、生成纯字母随机字符串
func (c *BaseController) CreateRandStr(strLen int, level int) (randStr string) {
	switch level {
	case 0:
		randStr = strconv.FormatInt(int64(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(c.Pow(10, strLen)))), 10)
	case 1:
		str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		bytes := []byte(str)
		result := []byte{}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < strLen; i++ {
			result = append(result, bytes[r.Intn(len(bytes))])
		}
		randStr = string(result)
	default:
		return ""
	}
	return randStr
}

// Pow 求x的n次幂
func (c *BaseController) Pow(x int, n int) int {
	ret := 1
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}

// GetNowTime 获取当前的北京时间，格式：2019-02-07 23:39:10
func (c *BaseController) GetNowTime() time.Time {
	// +28800是因为默认获取到的是UTC时间，需要+8个小时来土法转北京时间
	return time.Unix(time.Now().Unix()+28800, 0)
}

// GetWXopenid 获取微信小程序用户的openid及session
// level 1：获取普通用户的openid等信息，2：获取专家的，3：获取管理员的
func (c *BaseController) GetWXopenid(userCode string, level int) (openid string, session string, err error) {
	var appid string
	var appsecret string

	switch level {
	case 1:
		appid = beego.AppConfig.String("user_appid")
		appsecret = beego.AppConfig.String("user_appsecret")
	case 2:
		appid = beego.AppConfig.String("expert_appid")
		appsecret = beego.AppConfig.String("expert_appsecret")
	case 3:
		appid = beego.AppConfig.String("admin_appid")
		appsecret = beego.AppConfig.String("admin_appsecret")
	}

	wxAPIURL := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appid,
		appsecret,
		userCode,
	)

	req := httplib.Get(wxAPIURL)
	str, err := req.Bytes()
	if err != nil {
		return "", "", err
	}

	// 微信返回的数据是json格式的，需要从里面提取有用数据
	wxReq := make(map[string]interface{})
	json.Unmarshal(str, &wxReq)

	for k, v := range wxReq {
		if k == "openid" {
			openid, _ = v.(string)
		} else if k == "session_key" {
			session, _ = v.(string)
		} else if k == "errcode" {
			err, _ := v.(string)
			return "", "", errors.New(err)
		}
	}

	return openid, session, nil
}

// UploadFile 上传文件，注：仅mp4格式的视频
func (c *BaseController) UploadFile(fileInfo *multipart.FileHeader, fileStream string) (filePath string, fileType string, err error) {
	fmt.Println(fileInfo.Filename)
	fileExtenTmp := strings.Split(fileInfo.Filename, ".")
	fileExten := fileExtenTmp[len(fileExtenTmp)-1]
	if fileExten == "jpg" || fileExten == "png" || fileExten == "jpeg" || fileExten == "JPEG" || fileExten == "PNG" || fileExten == "JPG" {
		fileType = "image"
	} else if fileExten == "mp4" || fileExten == "MP4" {
		fileType = "video"
	} else {
		beego.Error("文件格式不受支持，格式为：" + fileExten)
		return "", "", errors.New("文件格式不受支持")
	}
	filePathHead := fmt.Sprintf("static/upload/%s/%d/%d/%d/",
		fileType,
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
	)
	fileName := fmt.Sprintf("%s.%s",
		c.CreateRandStr(10, 1),
		fileExten,
	)

	if _, err := os.Stat(filePathHead); err != nil {
		if err = os.MkdirAll(filePathHead, 0755); err != nil {
			beego.Error("文件上传失败：" + err.Error())
			return "", "", errors.New("文件上传失败：" + err.Error())
		}
	}
	filePath = filePathHead + fileName
	beego.Debug("保存的文件名：\n" + filePath)
	if err = c.SaveToFile(fileStream, filePath); err != nil {
		beego.Error("文件上传失败：" + err.Error())
		return "", "", errors.New("文件上传失败：" + err.Error())
	}
	return filePath, fileType, nil
}

// 注册极光IM账号
// level：1、普通用户 2、专家用户 3、管理员用户
func (c *BaseController) PostIm(uId int, nickname string, gender string, address string, avatar string, level int) (err error) {
	var imReq []map[string]interface{}
	// 为该用户注册IM账号

	// 从用户头像URL中解析出头像存储位置
	avatarArray := strings.Split(avatar, "/")

	avatar = ""
	for i := 0; i < len(avatarArray); i++ {
		if i < 3 {
			continue
		}
		if i != len(avatarArray)-1 {
			avatar = avatar + avatarArray[i] + "/"
		} else {
			avatar = avatar + avatarArray[i]
		}
	}

	// 上传用户头像
	mediaId, err := c.PostFileToIm(avatar)
	if err != nil {
		return err
	}

	// 调用接口注册IM账号
	var userName string
	if level == 1 {
		userName = fmt.Sprintf("u%d", uId)
	} else if level == 2 {
		userName = fmt.Sprintf("e%d", uId)
	}
	req := httplib.Post("https://api.im.jpush.cn/v1/users/")
	req.Header("Content-Type", "application/json")
	req.Header("Authorization", "Basic MmU0NTA0NDIwY2JmOTg1NjNlZjU0NmExOjAyYzBjNGU0OWNkMGVjYmY0NzUzODc2YQ==")
	req.JSONBody([]interface{}{map[string]interface{}{"username": userName, "password": "123456", "nickname": nickname, "avatar": mediaId, "gender": gender, "address": address}})

	fmt.Println([]interface{}{map[string]interface{}{"username": userName, "password": "123456", "nickname": nickname, "avatar": mediaId, "gender": gender, "address": address}})
	imReqJson, _ := req.Bytes()
	fmt.Println(string(imReqJson))
	fmt.Println("到我了")

	if err := json.Unmarshal(imReqJson, &imReq); err != nil {
		return errors.New("IM接口返回未知格式数据")
	}

	for _, v := range imReq {
		for k1, v1 := range v {
			if k1 == "error" {
				return errors.New(v1.(map[string]interface{})["message"].(string))
			}
		}
	}

	return nil
}

// 修改极光IM账号
// level：1、普通用户 2、专家用户 3、管理员用户
func (c *BaseController) PutIm(uId int, nickname string, gender string, address string, avatar string, level int) (err error) {
	genderInt, _ := strconv.Atoi(gender)

	var imReq []map[string]interface{}
	// 为该用户注册IM账号

	// 从用户头像URL中解析出头像存储位置
	avatarArray := strings.Split(avatar, "/")

	avatar = ""
	for i := 0; i < len(avatarArray); i++ {
		if i < 3 {
			continue
		}
		if i != len(avatarArray)-1 {
			avatar = avatar + avatarArray[i] + "/"
		} else {
			avatar = avatar + avatarArray[i]
		}
	}

	// 上传用户头像
	mediaId, err := c.PostFileToIm(avatar)
	if err != nil {
		return err
	}

	// 调用接口更新IM账号
	var userName string
	if level == 1 {
		userName = fmt.Sprintf("u%d", uId)
	}
	fmt.Println(userName)
	req := httplib.Put("https://api.im.jpush.cn/v1/users/" + userName)
	req.Header("Content-Type", "application/json")
	req.Header("Authorization", "Basic MmU0NTA0NDIwY2JmOTg1NjNlZjU0NmExOjAyYzBjNGU0OWNkMGVjYmY0NzUzODc2YQ==")
	req.JSONBody(map[string]interface{}{"nickname": nickname, "avatar": mediaId, "gender": genderInt, "address": address})

	fmt.Println(map[string]interface{}{"nickname": nickname, "avatar": mediaId, "gender": genderInt, "address": address})
	imReqJson, _ := req.Bytes()

	if err := json.Unmarshal(imReqJson, &imReq); err != nil && len(imReqJson) != 0 {
		return errors.New("IM接口返回未知格式数据")
	}

	for _, v := range imReq {
		for k1, v1 := range v {
			if k1 == "error" {
				return errors.New(v1.(map[string]interface{})["message"].(string))
			}
		}
	}

	return nil
}

// 上传文件到极光IM
func (c *BaseController) PostFileToIm(filePath string) (mediaId string, err error) {
	var imReq map[string]interface{}

	req := httplib.Post("https://api.im.jpush.cn/v1/resource?type=image")
	req.Header("Content-Type", "application/json")
	req.Header("Authorization", "Basic MmU0NTA0NDIwY2JmOTg1NjNlZjU0NmExOjAyYzBjNGU0OWNkMGVjYmY0NzUzODc2YQ==")
	req.PostFile("filename", filePath)
	imReqJson, _ := req.Bytes()

	if err := json.Unmarshal(imReqJson, &imReq); err != nil {
		return "", errors.New("文件上传失败:" + err.Error())
	}

	for k, v := range imReq {
		if k == "media_id" {
			mediaId = v.(string)
			break
		}
	}

	return mediaId, nil
}

// 调用腾讯云接口发送短信通知
// level 1:专家提交认证申请时给管理员发送短信 2:认证审核通过时给专家发送短信 3:认证审核未通过时给专家发送短信
func (c *BaseController) SendShortMessage(telephone string, level int, message string) (err error) {
	var tplId int // 腾讯云的短信模板ID
	var reqData map[string]interface{}
	var paramsArray []string

	switch level {
	case 1:
		tplId = 321355
		paramsArray = []string{}
	case 2:
		tplId = 321357
		paramsArray = []string{}
	case 3:
		tplId = 321358
		paramsArray = []string{message}
	}

	/*
		腾讯云接口请求格式
		{
			"ext": "",
			"extend": "",
			"params": [
				"验证码",
				"1234",
				"4"
			],
			"sig": "ecab4881ee80ad3d76bb1da68387428ca752eb885e52621a3129dcf4d9bc4fd4",
			"tel": {
				"mobile": "13788888888",
				"nationcode": "86"
			},
			"time": 1457336869,
			"tpl_id": 19
		}
	*/

	var random string
	var unixTime int64
	var sigStr string
	var unixTimeStr string

	random = c.CreateRandStr(6, 1)
	unixTime = time.Now().Unix()
	unixTimeStr = strconv.FormatInt(unixTime, 10)

	sigStr = fmt.Sprintf(
		"%x",
		sha256.Sum256(
			[]byte(
				"appkey="+beego.AppConfig.String("tx_sms_appkey")+
					"&random="+random+
					"&time="+unixTimeStr+
					"&mobile="+telephone,
			),
		),
	)

	fmt.Println(sigStr)
	fmt.Println(
		"appkey=" + beego.AppConfig.String("tx_sms_appkey") +
			"&random=" + random +
			"&time=" + unixTimeStr +
			"&mobile=" + telephone,
	)

	req := httplib.Post("https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=" + beego.AppConfig.String("tx_sms_sdkappid") + "&random=" + random)
	req.Header("Content-Type", "application/json")
	req.JSONBody(
		map[string]interface{}{
			"ext":    "",
			"extend": "",
			"sig":    sigStr,
			"time":   unixTime,
			"tpl_id": tplId,
			"params": paramsArray,
			"tel": map[string]string{
				"mobile":     telephone,
				"nationcode": "86",
			},
		},
	)

	ReqJson, _ := req.Bytes()

	if err := json.Unmarshal(ReqJson, &reqData); err != nil {
		return errors.New("腾讯云短信推送接口返回未知格式数据")
	}

	if reqData["result"].(float64) != 0 {
		return errors.New(reqData["errmsg"].(string))
	}

	return nil
}
