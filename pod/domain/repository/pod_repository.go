/*
	被查考的表为主表, 参考别的表的为从表
	疑问:  1.多表删除,先删从表再删主表有什么问题
          2.如果在
*/

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/qyh794/go-paas/pod/domain/model"
)

type IPodRepository interface {
	// 初始化表
	InitTable() error
	// 根据ID查询
	QueryPodByID(int64) (*model.Pod, error)
	// 创建Pod
	CreatePod(*model.Pod) (int64, error)
	// 删除Pod
	DeletePodById(int64) error
	// 修改Pod
	UpdatePod(*model.Pod) error
	// 查询所有Pod
	QueryAllPods() ([]model.Pod, error)
}

/*
	repository层直接对数据库进行操作
*/

type PodRepository struct {
	mysqlDB *gorm.DB
}

func NewPodRepository(db *gorm.DB) *PodRepository {
	return &PodRepository{mysqlDB: db}
}

// InitTable 初始化表
func (p *PodRepository) InitTable() error {
	return p.mysqlDB.CreateTable(model.Pod{}, model.PodEnv{}, model.PodPort{}).Error
	//TODO implement me
	panic("implement me")
}

// QueryPodByID 根据id查询pod信息
func (p *PodRepository) QueryPodByID(id int64) (*model.Pod, error) {
	pod := &model.Pod{}
	// 使用preload函数进行多表关联查询
	return pod, p.mysqlDB.Preload("PodEnv").Preload("PodPort").First(pod, id).Error
	//TODO implement me
	panic("implement me")
}

// CreatePod 创建一条pod数据
func (p *PodRepository) CreatePod(pod *model.Pod) (int64, error) {
	return pod.ID, p.mysqlDB.Create(pod).Error
	//TODO implement me
	panic("implement me")
}

// DeletePodById 根据ID删除pod
func (p *PodRepository) DeletePodById(id int64) error {
	// 删除pod信息涉及多表,需要使用事务保证数据一致性
	tx := p.mysqlDB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	// 多表删除时,先删除从表(参考别的表的表)再删除主表(被参考的表)
	// 使用软删除
	// 从pod表中删除pod信息
	// delete from pod where id = ?
	if err := p.mysqlDB.Where("id = ?", id).Delete(&model.Pod{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 从pod_Env表中删除pod信息
	// delete from pod_env where pod_id = ?
	if err := p.mysqlDB.Where("pod_id = ?", id).Delete(&model.PodEnv{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 从pod_Port表中删除pod信息
	// delete from pod_port where pod_id = ?
	if err := p.mysqlDB.Where("pod_id = ?", id).Delete(&model.PodPort{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 事务全部成功,提交
	return tx.Commit().Error
	//TODO implement me
	panic("implement me")
}

// updatePod 更新pod信息
func (p *PodRepository) UpdatePod(pod *model.Pod) error {
	// gorm更新语句会根据传入的模型有哪些字段进行更新
	// 例如db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false}),就只会更新Name和Age字段,并且会
	// update pod set xxx = pod.xxx, xxx = pod.xxx... where id = pod.id
	return p.mysqlDB.Model(&model.Pod{}).Update(pod).Error
	//TODO implement me
	panic("implement me")
}

// QueryAllPods 查询所有pod信息
func (p *PodRepository) QueryAllPods() ([]model.Pod, error) {
	allPods := make([]model.Pod, 0)
	// 获取第一条记录（主键升序） --> db.First(&user) --> select * from users order by id limit 1
	// 获取全部记录 --> db.Find(&users), var users []User
	// select * from pod
	return allPods, p.mysqlDB.Find(&allPods).Error
	//TODO implement me
	panic("implement me")
}
