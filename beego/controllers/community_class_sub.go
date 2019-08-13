package controllers

import (
	"encoding/json"
	"nxxlzx/models"
	"strconv"
)

// CommunityClassSubController operations for CommunityClassSub
type CommunityClassSubController struct {
	BaseController
}

// URLMapping ...
func (c *CommunityClassSubController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post 创建子分类
// @Title 创建子分类
// @Description 创建子分类，只需要填写标题、图标、父分类ID即可。
// @Param	token	query 		string	true		"令牌"
// @Param	body		body 	models.CommunityClassSub	true		"body for CommunityClassSub content"
// @Success 201 {int} models.CommunityClassSub
// @Failure 403 body is empty
// @router / [post]
func (c *CommunityClassSubController) Post() {
	c.ParseToken(c.GetString("token"))
	var v models.CommunityClassSub
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddCommunityClassSub(&v); err == nil {
			c.ReturnMsg(0, v, "")
		} else {
			c.ReturnMsg(-1, nil, err.Error())
		}
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Put 更新子分类
// @Title 更新子分类
// @Description 更新子分类
// @Param	token	query 		string	true		"令牌"
// @Param	id		path 	string	true		"要更新的分类的ID"
// @Param	body		body 	models.CommunityClassSub	true		"body for CommunityClassSub content"
// @Success 200 {object} models.CommunityClassSub
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CommunityClassSubController) Put() {
	c.ParseToken(c.GetString("token"))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.CommunityClassSub{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateCommunityClassSubById(&v); err == nil {
			c.ReturnMsg(0, nil, "")
		} else {
			c.ReturnMsg(-1, nil, err.Error())
		}
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Delete 删除子分类
// @Title 删除子分类
// @Description 删除子分类
// @Param	token	query 		string	true		"令牌"
// @Param	id		path 	string	true		"要删除的子分类的ID"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CommunityClassSubController) Delete() {
	c.ParseToken(c.GetString("token"))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCommunityClassSub(id); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}
