package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/qyh794/go-paas/volume/domain/model"
)

type IVolumeRepository interface {
	InitTable() error
	QueryVolumeByID(int64) (*model.Volume, error)
	CreateVolume(*model.Volume) (int64, error)
	DeleteVolumeByID(int64) error
	QueryAllVolume() ([]model.Volume, error)
	UpdateVolume(*model.Volume) error
}

type VolumeRepository struct {
	mysqlDB *gorm.DB
}

func NewVolumeRepository(db *gorm.DB) IVolumeRepository {
	return &VolumeRepository{mysqlDB: db}
}

func (v *VolumeRepository) InitTable() error {
	return v.mysqlDB.CreateTable(&model.Volume{}).Error
}

func (v *VolumeRepository) QueryVolumeByID(id int64) (*model.Volume, error) {
	// select * from volume where id = ?;
	volume := &model.Volume{}
	return volume, v.mysqlDB.First(volume, id).Error
}

func (v *VolumeRepository) CreateVolume(volume *model.Volume) (int64, error) {
	// insert into volume values (xxx, xxx)
	return volume.ID, v.mysqlDB.Create(volume).Error
}

func (v *VolumeRepository) DeleteVolumeByID(id int64) error {
	// delete from volume where id = id
	return v.mysqlDB.Where("id = ?", id).Delete(&model.Volume{}).Error
}

func (v *VolumeRepository) QueryAllVolume() ([]model.Volume, error) {
	var volumes []model.Volume
	// select * from volume
	return volumes, v.mysqlDB.Find(&model.Volume{}).Error
}

func (v *VolumeRepository) UpdateVolume(volume *model.Volume) error {
	// update volume set xxx = xxx
	return v.mysqlDB.Model(&model.Volume{}).Update(volume).Error
}
