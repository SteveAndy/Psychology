package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Expert struct {
	Id         int    `orm:"column(id);pk" description:"专家的ID"`
	Openid     string `orm:"column(openid)"`
	Photo      string `orm:"column(photo);size(255)"`
	Icon       string `orm:"column(icon);size(255)"`
	Name       string `orm:"column(name);size(255)"`
	PhoneNum   string `orm:"column(phone_num);size(11)"`
	ClassifyId int    `orm:"column(classify_id)"`
	Address    string `orm:"column(address);size(255)"`
	Info       string `orm:"column(info);size(255)"`
	Gender     string `orm:"column(gender);size(1);"`
	WorkAge    string `orm:"column(workAge);size(2);"`
}

func (t *Expert) TableName() string {
	return "expert"
}

func init() {
	orm.RegisterModel(new(Expert))
}

// AddExpert insert a new Expert into database and returns
// last inserted Id on success.
func AddExpert(m *Expert) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetExpertById retrieves Expert by Id. Returns error if
// Id doesn't exist
func GetExpertById(id int) (v *Expert, err error) {
	o := orm.NewOrm()
	v = &Expert{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetExpertByOpenID retrieves Expert by OpenID. Returns error if
// Id doesn't exist
func GetExpertByOpenID(openid string) (v *Expert, err error) {
	o := orm.NewOrm()
	v = &Expert{Openid: openid}
	if err = o.Read(v, "Openid"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllExpert retrieves all Expert matches certain condition. Returns empty list if
// no records exist
func GetAllExpert(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Expert))
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

	var l []Expert
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
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

// UpdateExpert updates Expert by Id and returns error if
// the record to be updated doesn't exist
func UpdateExpertById(m *Expert, cols ...string) (err error) {
	o := orm.NewOrm()
	v := Expert{Id: m.Id}
	// ascertain id exists in the database
	fmt.Println(v.Id)
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, cols...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteExpert deletes Expert by Id and returns error if
// the record to be deleted doesn't exist
func DeleteExpert(id int) (err error) {
	o := orm.NewOrm()
	v := Expert{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Expert{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
