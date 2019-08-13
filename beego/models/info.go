package models

import (
	"errors"
	"fmt"
	"nxxlzx/lib"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Info struct {
	Id         int      `orm:"column(id);auto"`
	ClassifyId int      `orm:"column(classifyId);size(11)" description:"所属的分类的ID"`
	Title      string   `orm:"column(title);size(255)" description:"标题"`
	Content    string   `orm:"column(content)" description:"内容"`
	Icon       string   `orm:"column(icon);size(255)" description:"缩略图"`
	Uid        int      `orm:"column(uid);size(11)" description:"作者"`
	Time       lib.Time `orm:"column(time);type(timestamp);auto_now_add"`
	AuthorType int      `orm:"column(author_type);size(2) description:"1专家，2管理员"`
}

func (t *Info) TableName() string {
	return "info"
}

func init() {
	orm.RegisterModel(new(Info))
}

// AddInfo insert a new Info into database and returns
// last inserted Id on success.
func AddInfo(m *Info) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInfoById retrieves Info by Id. Returns error if
// Id doesn't exist
func GetInfoById(id int) (info map[string]interface{}, err error) {
	o := orm.NewOrm()
	v := &Info{Id: id}
	if err = o.Read(v); err == nil {
		expertInfo, err := GetExpertById(v.Uid)
		if err != nil {
			expertInfo.Name = ""
		}
		info = map[string]interface{}{
			"title":      v.Title,
			"author":     expertInfo.Name,
			"time":       v.Time,
			"content":    v.Content,
			"icon":       v.Icon,
			"uid":        v.Uid,
			"AuthorType": v.AuthorType,
		}
		return info, nil
	}
	return nil, err
}

// GetAllInfo retrieves all Info matches certain condition. Returns empty list if
// no records exist
func GetAllInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Info))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			// 支持模糊搜索标题和内容
			if k == "Content" {
				cond := orm.NewCondition()
				cond1 := cond.Or("content__icontains", v).Or("title__icontains", v)
				qs = qs.SetCond(cond1)
			} else {
				qs = qs.Filter(k, v)
			}
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Info
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).Distinct().All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				var user_name string

				if v.AuthorType == 1 {
					userInfo, _ := GetExpertById(v.Uid)
					user_name = userInfo.Name
				} else {
					userInfo, _ := GetAdminById(v.Uid)
					user_name = userInfo.UserName
				}

				// 处理HTML内容，返回文章摘要
				var abs []rune // 文章的摘要
				var isTab int  // 标记当前游标是否处在html标签中

				// 去除字符串中的空格
				v.Content = strings.Replace(v.Content, " ", "", -1)
				// 去除字符串中的空格换行符
				v.Content = strings.Replace(v.Content, "\n", "", -1)
				// 将字符串转换为[]rune这样每个中文就变成一个元素，可以单独操作了
				content_tmp := []rune(v.Content)
				// 从内容中取出60个字符作为摘要，跳过HTML标签
				for i := 0; i < len(content_tmp); i++ {
					if content_tmp[i] == '<' || isTab == 1 {
						if i != 0 {
							if content_tmp[i] == '>' {
								isTab = 0
								continue
							}
						}
						isTab = 1
						continue
					}
					abs = append(abs, content_tmp[i])
					if len(abs) == 60 || content_tmp[i] == '\000' {
						break
					}
				}

				ml = append(
					ml,
					map[string]interface{}{
						"Id":         v.Id,
						"ClassifyId": v.ClassifyId,
						"Title":      v.Title,
						"Content":    string(abs),
						"Icon":       v.Icon,
						"Uid":        v.Uid,
						"UserName":   user_name,
						"Time":       v.Time,
					},
				)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateInfo updates Info by Id and returns error if
// the record to be updated doesn't exist
func UpdateInfoById(m *Info) (err error) {
	o := orm.NewOrm()
	v := Info{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInfo deletes Info by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInfo(id int) (err error) {
	o := orm.NewOrm()
	v := Info{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Info{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
