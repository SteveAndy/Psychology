package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CommunityClass struct {
	Id        int    `orm:"column(id);auto"`
	ClassName string `orm:"column(class_name);size(50)"`
	Color     string `orm:"column(color);size(20)"`
}

func (t *CommunityClass) TableName() string {
	return "community_class"
}

func init() {
	orm.RegisterModel(new(CommunityClass))
}

// AddCommunityClass insert a new CommunityClass into database and returns
// last inserted Id on success.
func AddCommunityClass(m *CommunityClass) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCommunityClassById retrieves CommunityClass by Id. Returns error if
// Id doesn't exist
func GetCommunityClassById(id int) (v *CommunityClass, err error) {
	o := orm.NewOrm()
	v = &CommunityClass{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCommunityClass retrieves all CommunityClass matches certain condition. Returns empty list if
// no records exist
func GetAllCommunityClass(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CommunityClass))
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

	var l []CommunityClass
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				sub, _ := GetAllCommunityClassSub(map[string]string{"Father": strconv.Itoa(v.Id)}, nil, nil, nil, 0, 100)
				ml = append(
					ml,
					map[string]interface{}{
						"name":  v.ClassName,
						"color": v.Color,
						"sub":   sub,
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

// UpdateCommunityClass updates CommunityClass by Id and returns error if
// the record to be updated doesn't exist
func UpdateCommunityClassById(m *CommunityClass) (err error) {
	o := orm.NewOrm()
	v := CommunityClass{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCommunityClass deletes CommunityClass by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCommunityClass(id int) (err error) {
	o := orm.NewOrm()
	v := CommunityClass{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CommunityClass{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
