package controllers

import (
	"encoding/json"
	"nxxlzx/models"
	"strconv"
	"strings"
)

// BannerController operations for Banner
type BannerController struct {
	BaseController
}

// URLMapping ...
func (c *BannerController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post 添加banner条
// @Title 添加banner条
// @Description 添加banner条，只需要提交 banner图片的url、类型、超链接即可
// @Param	token	query 		string	true		"令牌"
// @Param	body		body 	models.Banner	true		"body for Banner content"
// @Success 201 {int} models.Banner
// @Failure 403 body is empty
// @router / [post]
func (c *BannerController) Post() {
	c.ParseToken(c.GetString("token"))
	var v models.Banner
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddBanner(&v); err == nil {
			c.ReturnMsg(0, v, "")
		} else {
			c.ReturnMsg(-1, nil, err.Error())
		}
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// GetAll 获取所有的banner条
// @Title 获取所有的banner条
// @Description 获取所有的banner条
// @Param	token	query 		string	true		"令牌"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Banner
// @Failure 403
// @router / [get]
func (c *BannerController) GetAll() {
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

	l, err := models.GetAllBanner(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ReturnMsg(-1, nil, err.Error())
	} else {
		c.ReturnMsg(0, l, "")
	}
}

// Put 更新banner条
// @Title 更新banner条
// @Description 更新banner条
// @Param	token	query 		string	true		"令牌"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Banner	true		"body for Banner content"
// @Success 200 {object} models.Banner
// @Failure 403 :id is not int
// @router /:id [put]
func (c *BannerController) Put() {
	c.ParseToken(c.GetString("token"))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Banner{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateBannerById(&v); err == nil {
			c.ReturnMsg(0, nil, "")
		} else {
			c.ReturnMsg(-1, nil, err.Error())
		}
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Delete 删除banner条
// @Title 删除banner条
// @Description 删除banner条
// @Param	token	query 		string	true		"令牌"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *BannerController) Delete() {
	c.ParseToken(c.GetString("token"))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteBanner(id); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}
