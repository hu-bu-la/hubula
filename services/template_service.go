package services

//数据处理（包括数据库，也包括缓存等其他形式数据）

import (
	"github.com/hu-bu-la/hubula/dao"
	"github.com/hu-bu-la/hubula/datasource"
	"github.com/hu-bu-la/hubula/models"
)

type TemplateService interface {
	// GetAll 列表查询
	GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*datasource.Paginator, error)
	// GetById 获取单条记录
	GetById(id int) (*models.Template, error)
	// CountAll 统计
	CountAll() int64

	// Create 添加单条记录
	Create(data *models.Template) (int64,error)
	// Update 修改单条记录
	Update(data *models.Template, columns []string) (int64,error)
	// RuanDelete 软删除单条记录
	RuanDelete(id int) (int64, error)
	// Delete 删除单条记录
	Delete(id int) (int64, error)

	// GetWhere Sql语句
	GetWhere(sql string) []models.Template
}

type templateService struct {
	dao *dao.TemplateDao
}

func NewTemplateService() TemplateService {
	return &templateService{
		dao: dao.NewTemplateDao(datasource.InstanceDbMaster()),
	}
}

// GetAll 列表查询
func (s *templateService)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*datasource.Paginator, error)  {
	return s.dao.GetAll(q, fields, orderBy, page, limit)
}

// GetById 获取单条记录
func (s *templateService) GetById(id int) (*models.Template, error) {
	return s.dao.GetById(id)
}

// CountAll 统计
func (s *templateService) CountAll() int64 {
	return s.dao.CountAll()
}

// Create 添加单条记录
func (s *templateService) Create(data *models.Template) (int64,error) {
	// 先更新缓存
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.Create(data)
}

// Update 修改单条记录
func (s *templateService) Update(data *models.Template, columns []string) (int64,error) {
	// 先更新缓存
	//s.updateByCache(data, columns)
	// 再更新数据库
	return s.dao.Update(data, columns)
}

// RuanDelete 软删除单条记录
func (s *templateService) RuanDelete(id int) (int64, error) {
	// 先更新缓存
	//data := &models.Users{Id: id}
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.RuanDelete(id)
}

// Delete 删除单条记录
func (s *templateService) Delete(id int) (int64, error) {
	// 先更新缓存
	//data := &models.Users{Id: id}
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.Delete(id)
}

// GetWhere Sql语句
func (s *templateService) GetWhere(sql string) []models.Template {
	return s.dao.GetWhere(sql)
}