package controllers

import (
	"encoding/json"
	"fmt"
	"nxxlzx/models"
	"strconv"
	"strings"
)

// ExpertAuthController operations for ExpertAuth
type ExpertAuthController struct {
	BaseController
}

// URLMapping ...
func (c *ExpertAuthController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post 提交认证申请
// @Title 提交认证申请
// @Description 提交认证申请
// @Param	login_code	query 		string	true		"登录代码"
// @Param	body		body 	models.ExpertAuth	true		"除时间和ID外所有字段都需要提交，openid字段填写用户code"
// @Success 201 {int} models.ExpertAuth
// @Failure 403 body is empty
// @router / [post]
func (c *ExpertAuthController) Post() {
	loginCode, _ := strconv.Atoi(c.GetString("login_code")[1:])

	_, err := models.GetUsersById(loginCode)
	if err != nil {
		c.ReturnMsg(-1, nil, "登录码错误")
	}

	var v models.ExpertAuth
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddExpertAuth(&v); err == nil {
			// 成功提交申请后需要给管理员发送短信通知
			if err := c.SendShortMessage(v.PhoneNum, 1, "a"); err != nil {
				c.ReturnMsg(0, nil, "短信发送失败："+err.Error())
			}
			c.ReturnMsg(0, nil, "")
		} else {
			c.ReturnMsg(-1, nil, err.Error())
		}
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// GetAll 获取所有的专家认证列表
// @Title 获取所有的专家认证列表
// @Description 获取所有的专家认证列表
// @Param	token	query 		string	true		"令牌"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ExpertAuth
// @Failure 403
// @router / [get]
func (c *ExpertAuthController) GetAll() {
	c.ParseToken(c.GetString("token"))

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

	l, err := models.GetAllExpertAuth(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	} else {
		c.ReturnMsg(0, l, "")
	}
}

// Put 审核专家认证申请
// @Title 审核专家认证申请
// @Description 审核专家认证申请
// @Param	token	query 		string	true		"令牌"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	status	query 	string	true		"是否审核通过，1 通过，0拒绝，若拒绝则需要填写驳回原因"
// @Param	reject	query 	string	false		"驳回原因"
// @Success 200 {object} models.ExpertAuth
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ExpertAuthController) Put() {
	c.ParseToken(c.GetString("token"))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ExpertAuth{Id: id}
	authInfo, err := models.GetExpertAuthById(id)
	if err != nil {
		c.ReturnMsg(-1, nil, "不存在本条认证条目")
	}

	if c.GetString("status") == "1" {
		v.Status = 1
		// 将用户提交的认证信息加入到专家表中
		expertInfo := models.Expert{}
		expertInfo.Photo = authInfo.Photo
		expertInfo.Icon = authInfo.Photo
		expertInfo.Name = authInfo.Name
		expertInfo.PhoneNum = authInfo.PhoneNum
		expertInfo.ClassifyId = authInfo.ClassifyId
		expertInfo.Address = authInfo.Address
		expertInfo.Info = authInfo.Info
		expertInfo.Openid = authInfo.Openid
		expertInfo.WorkAge = authInfo.Age

		expertId, err := models.AddExpert(&expertInfo)
		if err != nil {
			c.ReturnMsg(-1, nil, err.Error())
		}

		if err := c.PostIm(int(expertId), expertInfo.Name, "2", expertInfo.Address, expertInfo.Photo, 2); err != nil {
			c.ReturnMsg(-1, nil, "IM注册失败，报错信息："+err.Error())
		}

		// 发送短信通知
		if err := c.SendShortMessage(authInfo.PhoneNum, 2, ""); err != nil {
			c.ReturnMsg(0, nil, "短信发送失败："+err.Error())
		}
	} else if c.GetString("status") == "0" {
		v.Reject = c.GetString("reject")
		v.Status = 2
		// 发送短信通知
		fmt.Println(authInfo.PhoneNum)
		if err := c.SendShortMessage(authInfo.PhoneNum, 3, v.Reject); err != nil {
			c.ReturnMsg(0, nil, "短信发送失败："+err.Error())
		}

	}
	if err := models.UpdateExpertAuthById(&v, "status", "reject"); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Delete 删除认证信息
// @Title 删除认证信息
// @Description 删除认证信息
// @Param	token	query 		string	true		"令牌"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ExpertAuthController) Delete() {
	c.ParseToken(c.GetString("token"))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteExpertAuth(id); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}
