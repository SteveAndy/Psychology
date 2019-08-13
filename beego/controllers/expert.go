package controllers

import (
	"encoding/json"
	"fmt"
	"nxxlzx/models"
	"strconv"
	"strings"
)

// ExpertController operations for Expert
type ExpertController struct {
	BaseController
}

// URLMapping ...
func (c *ExpertController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// GetOne 获取专家信息
// @Title 获取专家信息
// @Description 获取专家信息
// @Param	token			query 		string	true		"令牌"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Expert
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ExpertController) GetOne() {
	c.ParseToken(c.GetString("token"))

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetExpertById(id)
	if err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	} else {
		c.ReturnMsg(0, v, "")
	}
}

// GetAll 获取所有专家信息
// @Title 获取所有专家信息
// @Description 获取所有专家信息
// @Param	token			query 		string	true		"令牌"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Expert
// @Failure 403
// @router / [get]
func (c *ExpertController) GetAll() {
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

	l, err := models.GetAllExpert(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	} else {
		c.ReturnMsg(0, l, "")
	}
}

// Put 更新专家信息
// @Title 更新专家信息
// @Description 更新专家信息
// @Param	token			query 		string	true		"令牌"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Expert	true		"body for Expert content"
// @Success 200 {object} models.Expert
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ExpertController) Put() {
	c.ParseToken(c.GetString("token"))

	idStr := c.Ctx.Input.Param(":id")
	fmt.Println(idStr)
	fmt.Println(c.Ctx.Input.Param(":id"))
	id, _ := strconv.Atoi(idStr)
	fmt.Println(id)
	v := models.Expert{Id: id}
	fmt.Println(v)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println(v)
		if err := models.UpdateExpertById(&v, "Photo", "Icon", "Name", "PhoneNum", "ClassifyId", "Address", "Info", "Gender", "WorkAge"); err == nil {
			c.ReturnMsg(0, nil, "")
		} else {
			c.ReturnMsg(-1, nil, err.Error())
		}
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Delete 删除专家
// @Title 删除专家
// @Description 删除专家
// @Param	token			query 		string	true		"令牌"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ExpertController) Delete() {
	c.ParseToken(c.GetString("token"))

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteExpert(id); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Login 专家登录
// @Title 专家登录
// @Description 若登录成功则返回专家详细信息，失败则返回用户的session_key，方便前端与微信交换用户详细信息
// @Param	expert_code		formData 	string	true		"微信提供的用户身份代码"
// @Success 0 ...
// @Failure -1 ...
// @router /login [post]
func (c *ExpertController) Login() {
	expertCode := c.GetString("expert_code")
	if expertCode == "" {
		c.ReturnMsg(-1, nil, "expert_code为空")
	}

	openid, session, err := c.GetWXopenid(expertCode, 2)
	if err != nil {
		c.ReturnMsg(-1, nil, "换取openid失败")
	}

	// 需要根据用户是否注册执行不同的操作，未注册则注册并返回session_key给前端方便前端和微信交换用户详细信息，已注册则直接返回用户详细信息
	expertInfo, err := models.GetExpertByOpenID(openid)
	if err != nil {
		// 如果当前用户不是专家，则查询用户是否正在申请成为专家，再根据情况返回相应的值
		authInfo, err := models.GetExpertAuthByOpenID(openid)
		fmt.Println(openid)
		if err != nil {
			c.ReturnMsg(2, map[string]string{"openid": openid}, "未申请")
		}
		if authInfo.Status == 2 {
			//c.ReturnMsg(0, map[string]string{"status": "3", "reject": authInfo.Reject}, "认证审核未通过")
			c.ReturnMsg(3, authInfo.Reject, "认证审核未通过")
		}
		c.ReturnMsg(1, nil, "申请中")
	}
	if expertInfo.Name == "" {
		c.ReturnMsg(9001, map[string]interface{}{"token": c.CreateToken(int64(expertInfo.Id)), "session_key": session}, "用户信息未填写")
	}
	c.ReturnMsg(0, map[string]interface{}{"id": expertInfo.Id, "name": expertInfo.Name, "photo": expertInfo.Photo, "token": c.CreateToken(int64(expertInfo.Id))}, "")
}
