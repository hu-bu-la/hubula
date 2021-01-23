package dao

import (
	"fmt"
	"github.com/hu-bu-la/hubula/datasource"
	"github.com/hu-bu-la/hubula/models"
	"github.com/hu-bu-la/hubula/util"
	"github.com/xormplus/xorm"
)

type AdminRoleAccessDao struct {
	//数据库相关的操作 xorm引擎
	engine *xorm.Engine
}

//New Dao 实例化公共方法
func NewAdminRoleAccessDao(engine *xorm.Engine) *AdminRoleAccessDao {
	return &AdminRoleAccessDao{
		engine: engine,
	}
}

// NewAdmin 初始化
func (d *AdminRoleAccessDao) newModel() *models.AdminRoleAccess {
	return new(models.AdminRoleAccess)
}

// newMakeDataArr 初始化列表
func (d *AdminRoleAccessDao) newMakeDataArr() []models.AdminRoleAccess {
	return make([]models.AdminRoleAccess, 0)
}

// GetAll 列表查询
//条件 fields字段常和更新一起使用为0查询或更新所有字段 排序 页数 每页条数 返回分页内容 err
//q, fields, orderBy, page, limit) (*Paginator, error)
func (d *AdminRoleAccessDao) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*datasource.Paginator, error) {
	session := datasource.Filter(q)
	defer session.Close()
	count, err := session.Count(&models.AdminRoleAccess{})
	if err != nil {
		fmt.Println(err)
		return nil, util.NewError(err.Error())
	}

	Query := datasource.Pagination(int(count), page, limit)
	if count == 0 {
		return Query, nil
	}

	session = datasource.Filter(q)
	if orderBy != "" {
		session.OrderBy(orderBy)
	}
	session.Limit(limit, Query.Offset)
	if len(fields) == 0 {
		//更新所有字段
		session.AllCols()
	}
	data := d.newMakeDataArr()
	err = session.Find(&data)
	if err != nil {
		fmt.Println(err)
		return nil, util.NewError(err.Error())
	}
	Query.Data = make([]interface{}, len(data))
	for y, x := range data {
		Query.Data[y] = x
	}
	return Query, nil
}

// GetById 获取单条记录
func (d *AdminRoleAccessDao) GetById(id int) (*models.AdminRoleAccess, error) {
	m := d.newModel()
	m.Aid = id

	_, err := d.engine.Get(&m)
	if err == nil {
		return m,nil
	}

	return nil, err
}

// CountAll 统计
func (d *AdminRoleAccessDao) CountAll() int64 {
	m := d.newModel()
	num, err := d.engine.Count(&m)
	if err != nil {
		return 0
	} else {
		return num
	}
}

// Create 添加单条记录
func (d *AdminRoleAccessDao) Create(data *models.AdminRoleAccess) (int64,error) {
	num, err := d.engine.InsertOne(data)
	return num,err
}

// Update 修改单条记录
func (d *AdminRoleAccessDao) Update(data *models.AdminRoleAccess, columns []string) (int64,error) {
	num, err := d.engine.ID(data.Aid).MustCols(columns...).Update(data)
	return num,err
}

// RuanDelete 软删除单条记录
func (d *AdminRoleAccessDao) RuanDelete(id int) (int64, error) {
	m := d.newModel()
	m.Aid = id
	m.RoleId = 0

	num, err := d.engine.ID(&m.Aid).Update(&m)
	if err == nil {
		return num, nil
	}
	return num, err
}

// Delete 删除单条记录
func (d *AdminRoleAccessDao) Delete(id int) (int64, error) {
	m := d.newModel()
	m.Aid = id

	num, err := d.engine.Delete(&m)
	if err == nil {
		return num, nil
	}
	return num, err
}

// GetWhere Sql语句
func (d *AdminRoleAccessDao) GetWhere(sql string) []models.AdminRoleAccess {
	datalist := d.newMakeDataArr()
	err := d.engine.SQL(sql).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}