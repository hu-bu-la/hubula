package dao

import (
	"fmt"
	"github.com/hu-bu-la/hubula/datasource"
	"github.com/hu-bu-la/hubula/models"
	"github.com/hu-bu-la/hubula/util"
	"github.com/xormplus/xorm"
)

type TemplateDao struct {
	//数据库相关的操作 xorm引擎
	engine *xorm.Engine
}

//New Dao 实例化公共方法
func NewTemplateDao(engine *xorm.Engine) *TemplateDao {
	return &TemplateDao{
		engine: engine,
	}
}

// NewAdmin 初始化
func (d *TemplateDao) newModel() *models.Template {
	return new(models.Template)
}

// newMakeDataArr 初始化列表
func (d *TemplateDao) newMakeDataArr() []models.Template {
	return make([]models.Template, 0)
}

// GetAll 列表查询
//条件 fields字段常和更新一起使用为0查询或更新所有字段 排序 页数 每页条数 返回分页内容 err
//q, fields, orderBy, page, limit) (*Paginator, error)
func (d *TemplateDao) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*datasource.Paginator, error) {
	session := datasource.Filter(q)
	defer session.Close()
	count, err := session.Count(&models.Template{})
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
func (d *TemplateDao) GetById(id int) (*models.Template, error) {
	m := d.newModel()
	m.TemplateId = id

	_, err := d.engine.Get(&m)
	if err == nil {
		return m,nil
	}

	return nil, err
}

// CountAll 统计
func (d *TemplateDao) CountAll() int64 {
	m := d.newModel()
	num, err := d.engine.Count(&m)
	if err != nil {
		return 0
	} else {
		return num
	}
}

// Create 添加单条记录
func (d *TemplateDao) Create(data *models.Template) (int64,error) {
	num, err := d.engine.InsertOne(data)
	return num,err
}

// Update 修改单条记录
func (d *TemplateDao) Update(data *models.Template, columns []string) (int64,error) {
	num, err := d.engine.ID(data.TemplateId).MustCols(columns...).Update(data)
	return num,err
}

// RuanDelete 软删除单条记录
func (d *TemplateDao) RuanDelete(id int) (int64, error) {
	m := d.newModel()
	m.TemplateId = id
	m.Aid = 0

	num, err := d.engine.ID(&m.Aid).Update(&m)
	if err == nil {
		return num, nil
	}
	return num, err
}

// Delete 删除单条记录
func (d *TemplateDao) Delete(id int) (int64, error) {
	m := d.newModel()
	m.TemplateId = id

	num, err := d.engine.Delete(&m)
	if err == nil {
		return num, nil
	}
	return num, err
}

// GetWhere Sql语句
func (d *TemplateDao) GetWhere(sql string) []models.Template {
	datalist := d.newMakeDataArr()
	err := d.engine.SQL(sql).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}