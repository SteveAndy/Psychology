package controllers

import (
	"encoding/json"
	"nxxlzx/models"
	"strconv"
	"strings"
)

// CommunityController operations for Community
type CommunityController struct {
	BaseController
}

// URLMapping ...
func (c *CommunityController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post 发布帖子
// @Title 发布帖子
// @Description 发布帖子，只需要填写 分类ID、内容、作者ID即可
// @Param	token	query 		string	true		"令牌"
// @Param	body		body 	models.Community	true		"body for Community content"
// @Success 201 {int} models.Community
// @Failure 403 body is empty
// @router / [post]
func (c *CommunityController) Post() {
	c.ParseToken(c.GetString("token"))
	var v models.Community
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddCommunity(&v); err == nil {
			c.ReturnMsg(0, v, "")
		} else {
			c.ReturnMsg(-1, nil, err.Error())
		}
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// GetOne 获取帖子详情
// @Title 获取帖子详情
// @Description 获取帖子详情
// @Param	id		path 	string	true		"帖子ID"
// @Param	token			query 		string	true		"令牌"
// @Param	offset			query 		string	true		"从第几条记录开始查询帖子回复"
// @Success 200 {object} models.Community
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CommunityController) GetOne() {
	c.ParseToken(c.GetString("token"))
	offset, _ := strconv.Atoi(c.GetString("offset"))

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCommunityById(id, int64(offset))
	if err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	} else {
		c.ReturnMsg(0, v, "")
	}
}

// GetAll 获取帖子列表
// @Title 获取帖子列表
// @Description 获取帖子列表
// @Param	token	query 		string	true		"令牌"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Community
// @Failure 403
// @router / [get]
func (c *CommunityController) GetAll() {
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

	l, err := models.GetAllCommunity(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	} else {
		c.ReturnMsg(0, l, "")
	}
}

// Put 更新帖子内容
// @Title 更新帖子内容
// @Description 更新帖子内容
// @Param	token	query 		string	true		"令牌"
// @Param	id		path 	string	true		"要更新的帖子的ID"
// @Param	body		body 	models.Community	true		"body for Community content"
// @Success 200 {object} models.Community
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CommunityController) Put() {
	c.ParseToken(c.GetString("token"))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Community{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateCommunityById(&v); err == nil {
			c.ReturnMsg(0, nil, "")
		} else {
			c.ReturnMsg(-1, nil, err.Error())
		}
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Delete 删除帖子
// @Title 删除帖子
// @Description 删除帖子
// @Param	token	query 		string	true		"令牌"
// @Param	id		path 	string	true		"要删除的帖子的ID"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CommunityController) Delete() {
	c.ParseToken(c.GetString("token"))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCommunity(id); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}
