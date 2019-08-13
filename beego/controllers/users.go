package controllers

import (
	"nxxlzx/models"
	"strconv"
	"fmt"
)

// UsersController operations for Users
type UsersController struct {
	BaseController
}

// URLMapping ...
func (c *UsersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	//c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	//c.Mapping("Delete", c.Delete)
	c.Mapping("Login", c.Login)
	c.Mapping("SignUp", c.SignUp)
}

// GetOne 获取前台用户的详细信息
// @Title 获取前台用户的详细信息
// @Description 获取前台用户的详细信息
// @Param	id		path 	string	true		"用户ID"
// @Param	token	query 	string	true		"令牌"
// @Success 0 ...
// @Failure -1 ...
// @router /:id [get]
func (c *UsersController) GetOne() {
	c.ParseToken(c.GetString("token"))

	uId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	userInfo, err := models.GetUsersById(uId)
	if err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	} else {
		c.ReturnMsg(0, map[string]interface{}{"user_name": userInfo.UserName, "portrait": userInfo.Portrait, "year": userInfo.Year, "telephone": userInfo.Telephone, "address": userInfo.Address, "gender": userInfo.Gender}, "")
	}
}

/*
// GetAll ...
// @Title Get All
// @Description get Users
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Users
// @Failure 403
// @router / [get]
func (c *UsersController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}
*/

// Put 更新会员资料
// @Title 更新会员资料
// @Description 更新会员资料
// @Description 用户注册，在登录失败的时候后台数据库已经记录该用户，注册实际上是填写该用户的详细信息
// @Param	token			query 		string	true		"令牌"
// @Param	id			path 		string true			"ID"
// @Param	token		formData	string true			"令牌"
// @Param	icon	    formData	string true			"头像"
// @Param	user_name	formData	string true			"用户名"
// @Param	age			formData	string true			"年龄"
// @Param	sex			formData	string true			"性别"
// @Param	address 	formData	string true			"现地址"
// @Param	phoneNum	formData	string true			"手机号"
// @Success 200 {object} models.Users
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsersController) Put() {
	uIdStr := c.ParseToken(c.GetString("token"))
	uId, _ := strconv.Atoi(uIdStr)

	v := models.Users{Id: uId}
	v.Address = c.GetString("address")
	v.Gender = c.GetString("sex")
	v.Year = c.GetString("age")
	v.UserName = c.GetString("user_name")
	v.Portrait = c.GetString("icon")
	v.Telephone = c.GetString("phoneNum")

	fmt.Println(v)

	// 为用户更新IM信息
	if err := c.PutIm(uId, v.UserName, v.Gender, v.Address, v.Portrait, 1); err != nil{
		c.ReturnMsg(-1,nil,"IM账号信息修改失败，报错信息：" + err.Error())
	}

	if err := models.UpdateUsersById(&v, "portrait", "user_name", "year", "gender", "address", "telephone"); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Delete 删除用户
// @Title 删除用户
// @Description 删除用户
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUsers(id); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Login 前台会员登录
// @Title 前台会员登录
// @Description 若登录成功则返回用户详细信息，失败则返回用户的session_key，方便前端与微信交换用户详细信息
// @Param	user_code		formData 	string	true		"微信提供的用户身份代码"
// @Success 0 ...
// @Failure -1 ...
// @router /login [post]
func (c *UsersController) Login() {
	userCode := c.GetString("user_code")
	if userCode == "" {
		c.ReturnMsg(-1, nil, "user_code为空")
	}

	openid, session, err := c.GetWXopenid(userCode, 1)
	if err != nil {
		c.ReturnMsg(-1, nil, "换取openid失败")
	}

	// 需要根据用户是否注册执行不同的操作，未注册则注册并返回session_key给前端方便前端和微信交换用户详细信息，已注册则直接返回用户详细信息
	userInfo, err := models.GetUsersByOpenID(openid)
	if err != nil {
		var userInfo models.Users
		userInfo.Openid = openid
		uId, _ := models.AddUsers(&userInfo)
		c.ReturnMsg(9001, map[string]interface{}{"token": c.CreateToken(uId), "session_key": session}, err.Error())
	}
	if userInfo.UserName == "" {
		c.ReturnMsg(9001, map[string]interface{}{"token": c.CreateToken(int64(userInfo.Id)), "session_key": session}, "用户信息未填写")
	}
	c.ReturnMsg(0, map[string]interface{}{"id": userInfo.Id, "user_name": userInfo.UserName, "portrait": userInfo.Portrait, "token": c.CreateToken(int64(userInfo.Id)), "is_super": userInfo.IsSuper, "login_code": "e"+ strconv.Itoa(userInfo.Id)}, "")
}

// SignUp 前台会员注册
// @Title 前台用户注册
// @Description 用户注册，在登录失败的时候后台数据库已经记录该用户，注册实际上是填写该用户的详细信息
// @Param	token			query 		string	true		"令牌"
// @Param	user_name		formData 	string	true		"用户名"
// @Param	telephone		formData 	string	true		"电话"
// @Param	address			formData 	string	true		"地址"
// @Param	gender			formData 	string	true		"0 未知 1 男 2 女"
// @Param	portrait			formData 	string	true		"头像"
// @Success 0 ...
// @Failure -1 ...
// @router /signup [post]
func (c *UsersController) SignUp() {
	uId, _ := strconv.Atoi(c.ParseToken(c.GetString("token")))

	userInfo := models.Users{Id: uId}
	userInfo.UserName = c.GetString("user_name")
	userInfo.Telephone = c.GetString("telephone")
	userInfo.Address = c.GetString("address")
	userInfo.Portrait = c.GetString("portrait")
	userInfo.Gender = c.GetString("gender")

	// 为用户注册IM账号
	if err := c.PostIm(uId, userInfo.UserName, userInfo.Gender, userInfo.Address, userInfo.Portrait, 1); err != nil{
		c.ReturnMsg(-1,nil,"IM注册失败，报错信息："+err.Error())
	}

	// 将用户详细信息更新到数据库中
	if err := models.UpdateUsersById(&userInfo, "UserName", "Telephone", "Address", "Portrait", "Gender"); err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	}

	c.ReturnMsg(0, map[string]interface{}{"user_name": userInfo.UserName, "portrait": userInfo.Portrait}, "")
}

// SessionKey 获取用户在微信中的session_key
// @Title 获取用户在微信中的Session_key
// @Description 获取用户在微信中的Session_key，用以执行兑换用户手机号等操作
// @Param	token			query 		string	true		"令牌"
// @Param	user_code			formData 	string	true		"用户的code"
// @Success 0 ...
// @Failure -1 ...
// @router /session_key [post]
func (c *UsersController) SessionKey() {
	c.ParseToken(c.GetString("token"))

	userCode := c.GetString("user_code")
	if userCode == "" {
		c.ReturnMsg(-1, nil, "user_code为空")
	}

	_, session, err := c.GetWXopenid(userCode, 1)
	if err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	}

	c.ReturnMsg(0, map[string]interface{}{"session": session}, "")
}
