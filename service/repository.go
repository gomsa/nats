package service

import (
	"fmt"
	// 公共引入

	pb "github.com/gomsa/nats/proto/template"
	"github.com/micro/go-micro/util/log"

	"github.com/jinzhu/gorm"
)

//Repository 仓库接口
type Repository interface {
	Create(role *pb.Template) (*pb.Template, error)
	Delete(role *pb.Template) (bool, error)
	Update(role *pb.Template) (bool, error)
	Get(role *pb.Template) (*pb.Template, error)
	List(req *pb.ListQuery) ([]*pb.Template, error)
	Total(req *pb.ListQuery) (int64, error)
}

// TemplateRepository 消息事件仓库
type TemplateRepository struct {
	DB *gorm.DB
}

// List 获取所有消息事件信息
func (repo *TemplateRepository) List(req *pb.ListQuery) (roles []*pb.Template, err error) {
	db := repo.DB
	// 分页
	var limit, offset int64
	if req.Limit > 0 {
		limit = req.Limit
	} else {
		limit = 10
	}
	if req.Page > 1 {
		offset = (req.Page - 1) * limit
	} else {
		offset = -1
	}

	// 排序
	var sort string
	if req.Sort != "" {
		sort = req.Sort
	} else {
		sort = "created_at desc"
	}
	// 查询条件
	if req.Label != "" {
		db = db.Where("label like ?", "%"+req.Label+"%")
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&roles).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return roles, nil
}

// Total 获取所有消息事件查询总量
func (repo *TemplateRepository) Total(req *pb.ListQuery) (total int64, err error) {
	roles := []pb.Template{}
	db := repo.DB
	// 查询条件
	if req.Label != "" {
		db = db.Where("label like ?", "%"+req.Label+"%")
	}
	if err := db.Find(&roles).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取消息事件信息
func (repo *TemplateRepository) Get(r *pb.Template) (*pb.Template, error) {
	if r.Id != 0 {
		if err := repo.DB.Model(&r).Where("id = ?", r.Id).Find(&r).Error; err != nil {
			return nil, err
		}
	}
	if r.Event != "" {
		if err := repo.DB.Model(&r).Where("event = ?", r.Event).Find(&r).Error; err != nil {
			return nil, err
		}
	}
	if r.Name != "" {
		if err := repo.DB.Model(&r).Where("name = ?", r.Name).Find(&r).Error; err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Create 创建消息事件
// bug 无消息事件名创建消息事件可能引起 bug
func (repo *TemplateRepository) Create(r *pb.Template) (*pb.Template, error) {
	err := repo.DB.Create(r).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return r, fmt.Errorf("添加消息事件失败")
	}
	return r, nil
}

// Update 更新消息事件
func (repo *TemplateRepository) Update(r *pb.Template) (bool, error) {
	if r.Id == 0 {
		return false, fmt.Errorf("请传入更新id")
	}
	id := &pb.Template{
		Id: r.Id,
	}
	err := repo.DB.Model(id).Updates(r).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除消息事件
func (repo *TemplateRepository) Delete(r *pb.Template) (bool, error) {
	err := repo.DB.Delete(r).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
