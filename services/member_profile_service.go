package services

//数据处理（包括数据库，也包括缓存等其他形式数据）

import (
	"github.com/hu-bu-la/hubula/dao"
	"github.com/hu-bu-la/hubula/datasource"
	"github.com/hu-bu-la/hubula/models"
)

type MemberProfileService interface {
	// GetAll 列表查询
	GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*datasource.Paginator, error)
	// GetById 获取单条记录
	GetById(id int) (*models.MemberProfile, error)
	// CountAll 统计
	CountAll() int64

	// Create 添加单条记录
	Create(data *models.MemberProfile) (int64,error)
	// Update 修改单条记录
	Update(data *models.MemberProfile, columns []string) (int64,error)
	// RuanDelete 软删除单条记录
	RuanDelete(id int) (int64, error)
	// Delete 删除单条记录
	Delete(id int) (int64, error)

	// GetWhere Sql语句
	GetWhere(sql string) []models.MemberProfile
}

type memberProfileService struct {
	dao *dao.MemberProfileDao
}

func NewMemberProfileService() MemberProfileService {
	return &memberProfileService{
		dao: dao.NewMemberProfileDao(datasource.InstanceDbMaster()),
	}
}

// GetAll 列表查询
func (s *memberProfileService)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*datasource.Paginator, error)  {
	return s.dao.GetAll(q, fields, orderBy, page, limit)
}

// GetById 获取单条记录
func (s *memberProfileService) GetById(id int) (*models.MemberProfile, error) {
	return s.dao.GetById(id)
}

// CountAll 统计
func (s *memberProfileService) CountAll() int64 {
	return s.dao.CountAll()
}

// Create 添加单条记录
func (s *memberProfileService) Create(data *models.MemberProfile) (int64,error) {
	// 先更新缓存
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.Create(data)
}

// Update 修改单条记录
func (s *memberProfileService) Update(data *models.MemberProfile, columns []string) (int64,error) {
	// 先更新缓存
	//s.updateByCache(data, columns)
	// 再更新数据库
	return s.dao.Update(data, columns)
}

// RuanDelete 软删除单条记录
func (s *memberProfileService) RuanDelete(id int) (int64, error) {
	// 先更新缓存
	//data := &models.Users{Id: id}
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.RuanDelete(id)
}

// Delete 删除单条记录
func (s *memberProfileService) Delete(id int) (int64, error) {
	// 先更新缓存
	//data := &models.Users{Id: id}
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.Delete(id)
}

// GetWhere Sql语句
func (s *memberProfileService) GetWhere(sql string) []models.MemberProfile {
	return s.dao.GetWhere(sql)
}