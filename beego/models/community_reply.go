package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"nxxlzx/lib"

	"github.com/astaxie/beego/orm"
)

type CommunityReply struct {
	Id      int       `orm:"column(id);auto"`
	Cid     int       `orm:"column(cid);size(11)"`
	UserId  int       `orm:"column(user_id);size(11)"`
	Content string    `orm:"column(content);size(255)"`
	Time    lib.Time `orm:"column(time);type(timestamp);auto_now_add"`
}

func (t *CommunityReply) TableName() string {
	return "community_reply"
}

func init() {
	orm.RegisterModel(new(CommunityReply))
}

// AddCommunityReply insert a new CommunityReply into database and returns
// last inserted Id on success.
func AddCommunityReply(m *CommunityReply) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCommunityReplyById retrieves CommunityReply by Id. Returns error if
// Id doesn't exist
func GetCommunityReplyById(id int) (v *CommunityReply, err error) {
	o := orm.NewOrm()
	v = &CommunityReply{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCommunityReply retrieves all CommunityReply matches certain condition. Returns empty list if
// no records exist
// 如果reType为1则是统计当前查询条件下的记录数量，为0则返回相应记录
func GetAllCommunityReply(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, reType int) (ml []interface{}, num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CommunityReply))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
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
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}

	var l []CommunityReply
	qs = qs.OrderBy(sortFields...)
	if reType == 1 {
		num, _ = qs.Count()
		return nil, num, nil
	}
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				var expertInfoTmp Expert
				expertInfo, err := GetExpertById(v.UserId)
				if err != nil {
					expertInfo = &expertInfoTmp
					expertInfo.Icon = ""
					expertInfo.Name = ""
					expertInfo.PhoneNum = ""
				}
				ml = append(
					ml,
					map[string]interface{}{
						"id":		 v.Id,
						"user_id":   v.UserId,
						"icon":      expertInfo.Icon,
						"name":      expertInfo.Name,
						"content":   v.Content,
						"time":      v.Time,
						"telephone": expertInfo.PhoneNum,
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
		return ml, 0, nil
	}
	return nil, 0, err
}

// UpdateCommunityReply updates CommunityReply by Id and returns error if
// the record to be updated doesn't exist
func UpdateCommunityReplyById(m *CommunityReply) (err error) {
	o := orm.NewOrm()
	v := CommunityReply{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCommunityReply deletes CommunityReply by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCommunityReply(id int) (err error) {
	o := orm.NewOrm()
	v := CommunityReply{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CommunityReply{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
