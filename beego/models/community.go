package models

import (
	"errors"
	"fmt"
	"nxxlzx/lib"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Community struct {
	Id         int      `orm:"column(id);auto"`
	UserId     int      `orm:"column(userId);size(255)"`
	ClassifyId int      `orm:"column(classifyId);size(11)"`
	Content    string   `orm:"column(content);size(255)"`
	Time       lib.Time `orm:"column(time);type(timestamp);auto_now_add"`
	See        int      `orm:"column(see)"`
}

func (t *Community) TableName() string {
	return "community"
}

func init() {
	orm.RegisterModel(new(Community))
}

// AddCommunity insert a new Community into database and returns
// last inserted Id on success.
func AddCommunity(m *Community) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCommunityById retrieves Community by Id. Returns error if
// Id doesn't exist
func GetCommunityById(id int, offset int64) (communityInfo map[string]interface{}, err error) {
	o := orm.NewOrm()
	v := &Community{Id: id}
	if err = o.Read(v); err == nil {
		userInfo, _ := GetUsersById(v.UserId)
		communityClassSub, _ := GetCommunityClassSubById(v.ClassifyId)
		_, messageNum, _ := GetAllCommunityReply(map[string]string{"Cid": strconv.Itoa(v.Id)}, nil, nil, nil, 0, 10, 1)
		messageList, _, _ := GetAllCommunityReply(map[string]string{"Cid": strconv.Itoa(v.Id)}, nil, nil, nil, offset, 10, 0)
		communityInfo = map[string]interface{}{
			"id":           v.Id,
			"user_id":      v.UserId,
			"icon":         userInfo.Portrait,
			"name":         userInfo.UserName,
			"phone_num":    userInfo.Telephone,
			"classify":     communityClassSub.Title,
			"classify_id":  v.ClassifyId,
			"content":      v.Content,
			"time":         v.Time,
			"see":          v.See,
			"message_num":  messageNum,
			"message_list": messageList,
		}

		// 查询后将帖子浏览量+1
		v.See += 1
		_ = UpdateCommunityById(v)

		return communityInfo, nil
	}
	return nil, err
}

// GetAllCommunity retrieves all Community matches certain condition. Returns empty list if
// no records exist
func GetAllCommunity(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Community))
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

	var l []Community
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				userInfo, _ := GetUsersById(v.UserId)
				communityClassSub, _ := GetCommunityClassSubById(v.ClassifyId)
				_, messageNum, _ := GetAllCommunityReply(map[string]string{"Cid": strconv.Itoa(v.Id)}, nil, nil, nil, 0, 10, 1)
				message, _, err := GetAllCommunityReply(map[string]string{"Cid": strconv.Itoa(v.Id)}, nil, []string{"Id"}, []string{"desc"}, 0, 1, 0)
				if err != nil {
					message = nil
				}

				ml = append(
					ml,
					map[string]interface{}{
						"id":          v.Id,
						"user_id":     v.UserId,
						"icon":        userInfo.Portrait,
						"name":        userInfo.UserName,
						"phone_num":   userInfo.Telephone,
						"classify":    communityClassSub.Title,
						"classify_id": v.ClassifyId,
						"content":     v.Content,
						"time":        v.Time,
						"see":         v.See,
						"message_num": messageNum,
						"message":     message,
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

// UpdateCommunity updates Community by Id and returns error if
// the record to be updated doesn't exist
func UpdateCommunityById(m *Community) (err error) {
	o := orm.NewOrm()
	v := Community{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCommunity deletes Community by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCommunity(id int) (err error) {
	o := orm.NewOrm()
	v := Community{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Community{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
