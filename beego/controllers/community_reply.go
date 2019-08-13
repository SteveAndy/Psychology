package controllers

import (
	"nxxlzx/models"
	"strconv"
)

// CommunityReplyController operations for CommunityReply
type CommunityReplyController struct {
	BaseController
}

// URLMapping ...
func (c *CommunityReplyController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post 添加帖子回复
// @Title 添加帖子回复
// @Description 添加帖子回复，只需要提交 留言人ID、留言内容 即可
// @Param	token		query 		string	true		"令牌"
// @Param	Cid			formData 	string	true		"回复所属的帖子的ID"
// @Param	UserId		formData 	string	true		"添加本回复的用户的ID"
// @Param	Content		formData 	string	true		"回复的内容"
// @Success 201 {int} models.CommunityReply
// @Failure 403 body is empty
// @router / [post]
func (c *CommunityReplyController) Post() {
	c.ParseToken(c.GetString("token"))
	var v models.CommunityReply
	v.Cid, _ = strconv.Atoi(c.GetString("Cid"))
	v.UserId, _ = strconv.Atoi(c.GetString("UserId"))
	v.Content = c.GetString("Content")
	if _, err := models.AddCommunityReply(&v); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Put 更新回复
// @Title 更新回复
// @Description 更新回复
// @Param	token	query 		string	true		"令牌"
// @Param	id		path 	string	true		"要更新的回复的ID"
// @Param	Cid			formData 	string	true		"回复所属的帖子的ID"
// @Param	UserId		formData 	string	true		"添加本回复的用户的ID"
// @Param	Content		formData 	string	true		"回复的内容"
// @Success 200 {object} models.CommunityReply
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CommunityReplyController) Put() {
	c.ParseToken(c.GetString("token"))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.CommunityReply{Id: id}
	v.Cid, _ = strconv.Atoi(c.GetString("Cid"))
	v.UserId, _ = strconv.Atoi(c.GetString("UserId"))
	v.Content = c.GetString("Content")
	if err := models.UpdateCommunityReplyById(&v); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}

// Delete 删除回复
// @Title 删除回复
// @Description 删除回复
// @Param	token	query 		string	true		"令牌"
// @Param	id		path 	string	true		"要删除的回复的ID"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CommunityReplyController) Delete() {
	c.ParseToken(c.GetString("token"))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCommunityReply(id); err == nil {
		c.ReturnMsg(0, nil, "")
	} else {
		c.ReturnMsg(-1, nil, err.Error())
	}
}
