package controllers

import (
	"encoding/json"
	"nxxlzx/models"
	"strconv"
	"strings"
)

// AdminController operations for Admin
type AdminController struct {
	BaseController
}

// URLMapping ...
func (c *AdminController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Login", c.Login)
}

// GetAll 获取管理员列表
// @Title 获取管理员列表
// @Description 获取管理员列表
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Admin
// @Failure 403
// @router / [get]
func (c *AdminController) GetAll() {
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
				c.ReturnMsg(-1, nil, "Error: invalid query key/value pair")
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllAdmin(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	} else {
		c.ReturnMsg(0, l, "")
	}
}

// Put 更新管理员信息
// @Title 更新管理员信息
// @Description 更新管理员信息
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Admin	true		"body for Admin content"
// @Success 200 {object} models.Admin
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AdminController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Admin{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		adminInfo, _ := models.GetAdminById(id)
		v.Openid = adminInfo.Openid
		if err := models.UpdateAdminById(&v); err == nil {
			c.ReturnMsg(0, map[string]interface{}{
				"id":       v.Id,
				"name":     v.UserName,
				"portrait": v.Portrait,
				"phoneNum": v.Telephone,
			}, "")
		} else {
			c.ReturnMsg(-1, nil, err.Error())
		}
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Delete 删除管理员
// @Title 删除管理员
// @Description 删除管理员
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AdminController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAdmin(id); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Login 管理员登录
// @Title 管理员登录
// @Description 若登录成功则返回管理员详细信息
// @Param	admin_code		formData 	string	true		"微信提供的用户身份代码"
// @Success 0 ...
// @Failure -1 ...
// @router /login [post]
func (c *AdminController) Login() {
	adminCode := c.GetString("admin_code")
	if adminCode == "" {
		c.ReturnMsg(-1, nil, "admin_code为空")
	}

	openid, _, err := c.GetWXopenid(adminCode, 2)
	if err != nil {
		c.ReturnMsg(-1, nil, "换取openid失败")
	}

	adminInfo, err := models.GetAdminByOpenID(openid)
	if err != nil {
		c.ReturnMsg(-1, nil, "非法用户:"+openid)
	}
	c.ReturnMsg(0, map[string]interface{}{
		"id":       adminInfo.Id,
		"name":     adminInfo.UserName,
		"portrait": adminInfo.Portrait,
		"phoneNum": adminInfo.Telephone,
		"token":    c.CreateToken(int64(adminInfo.Id)),
	}, "")
}
