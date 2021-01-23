package services

//数据处理（包括数据库，也包括缓存等其他形式数据）

import (
	"github.com/hu-bu-la/hubula/dao"
	"github.com/hu-bu-la/hubula/datasource"
	"github.com/hu-bu-la/hubula/models"
)

type MemberStatusService interface {
	// GetAll 列表查询
	GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*datasource.Paginator, error)
	// GetById 获取单条记录
	GetById(id int) (*models.MemberStatus, error)
	// CountAll 统计
	CountAll() int64

	// Create 添加单条记录
	Create(data *models.MemberStatus) (int64,error)
	// Update 修改单条记录
	Update(data *models.MemberStatus, columns []string) (int64,error)
	// RuanDelete 软删除单条记录
	RuanDelete(id int) (int64, error)
	// Delete 删除单条记录
	Delete(id int) (int64, error)

	// GetWhere Sql语句
	GetWhere(sql string) []models.MemberStatus
}

type memberStatusService struct {
	dao *dao.MemberStatusDao
}

func NewMemberStatusService() MemberStatusService {
	return &memberStatusService{
		dao: dao.NewMemberStatusDao(datasource.InstanceDbMaster()),
	}
}

// GetAll 列表查询
func (s *memberStatusService)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*datasource.Paginator, error)  {
	return s.dao.GetAll(q, fields, orderBy, page, limit)
}

// GetById 获取单条记录
func (s *memberStatusService) GetById(id int) (*models.MemberStatus, error) {
	return s.dao.GetById(id)
}

// CountAll 统计
func (s *memberStatusService) CountAll() int64 {
	return s.dao.CountAll()
}

// Create 添加单条记录
func (s *memberStatusService) Create(data *models.MemberStatus) (int64,error) {
	// 先更新缓存
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.Create(data)
}

// Update 修改单条记录
func (s *memberStatusService) Update(data *models.MemberStatus, columns []string) (int64,error) {
	// 先更新缓存
	//s.updateByCache(data, columns)
	// 再更新数据库
	return s.dao.Update(data, columns)
}

// RuanDelete 软删除单条记录
func (s *memberStatusService) RuanDelete(id int) (int64, error) {
	// 先更新缓存
	//data := &models.Users{Id: id}
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.RuanDelete(id)
}

// Delete 删除单条记录
func (s *memberStatusService) Delete(id int) (int64, error) {
	// 先更新缓存
	//data := &models.Users{Id: id}
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.Delete(id)
}

// GetWhere Sql语句
func (s *memberStatusService) GetWhere(sql string) []models.MemberStatus {
	return s.dao.GetWhere(sql)
}