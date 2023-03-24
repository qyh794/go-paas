package service

import (
	"context"
	"errors"
	"github.com/micro/micro/v3/service/logger"
	"github.com/qyh794/go-paas/volume/domain/model"
	"github.com/qyh794/go-paas/volume/domain/repository"
	"github.com/qyh794/go-paas/volume/proto/volume"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
)

type IVolumeDataService interface {
	AddVolume(*model.Volume) (int64, error)
	DeleteVolumeByID(int64) error
	UpdateVolume(*model.Volume) error
	QueryVolumeByID(int64) (*model.Volume, error)
	QueryAllVolume() ([]model.Volume, error)
	CreateVolumeToK8s(info *volume.RVolumeInfo) error
	DeleteVolumeFromK8s(*model.Volume) error
}

type VolumeDataService struct {
	VolumeRepository repository.IVolumeRepository
	K8sClientSet     *kubernetes.Clientset
	Deployment       *v1.Deployment
}

func NewVolumeDataService(k8sClientSet *kubernetes.Clientset, volumeRepository repository.IVolumeRepository) IVolumeDataService {
	return &VolumeDataService{
		VolumeRepository: volumeRepository,
		K8sClientSet:     k8sClientSet,
	}
}

func (v *VolumeDataService) AddVolume(volume *model.Volume) (int64, error) {
	return v.VolumeRepository.CreateVolume(volume)
}

func (v *VolumeDataService) DeleteVolumeByID(id int64) error {
	return v.VolumeRepository.DeleteVolumeByID(id)
}

func (v *VolumeDataService) UpdateVolume(volume *model.Volume) error {
	return v.VolumeRepository.UpdateVolume(volume)
}

func (v *VolumeDataService) QueryVolumeByID(id int64) (*model.Volume, error) {
	return v.VolumeRepository.QueryVolumeByID(id)
}

func (v *VolumeDataService) QueryAllVolume() ([]model.Volume, error) {
	return v.VolumeRepository.QueryAllVolume()
}

func (v *VolumeDataService) CreateVolumeToK8s(info *volume.RVolumeInfo) error {
	pvc := v.setVolume(info)
	// 查询k8s
	_, err := v.K8sClientSet.CoreV1().PersistentVolumeClaims(info.VolumeNamespace).Get(context.TODO(), info.VolumeName, v13.GetOptions{})
	if err != nil { // pvc不存在
		if _, err = v.K8sClientSet.CoreV1().PersistentVolumeClaims(info.VolumeNamespace).Create(context.TODO(), pvc, v13.CreateOptions{}); err != nil {
			// pvc创建失败
			logger.Error(err)
			return err
		}
		logger.Info("pvc创建成功")
		return nil
	} else { // pvc已存在, err == nil
		logger.Error("pvc :" + info.VolumeName + "已存在,重复创建")
		return errors.New("pvc :" + info.VolumeName + "已存在,重复创建")
	}
}

func (v *VolumeDataService) DeleteVolumeFromK8s(volume *model.Volume) error {
	if err := v.K8sClientSet.CoreV1().PersistentVolumeClaims(volume.VolumeNamespace).Delete(context.TODO(), volume.VolumeName, v13.DeleteOptions{}); err != nil {
		// 删除pvc失败
		logger.Error("删除pvc失败", err)
		return err
	}
	// k8s删除pvc成功
	logger.Info("k8s删除pvc成功")
	return nil
}

func (v *VolumeDataService) setVolume(info *volume.RVolumeInfo) *v12.PersistentVolumeClaim {
	pvc := &v12.PersistentVolumeClaim{}
	pvc.TypeMeta = v13.TypeMeta{
		Kind:       "PersistentVolumeClaim",
		APIVersion: "v1",
	}
	pvc.ObjectMeta = v13.ObjectMeta{
		Name:      info.VolumeName,
		Namespace: info.VolumeNamespace,
	}
	pvc.Spec = v12.PersistentVolumeClaimSpec{
		AccessModes:      v.getAccessModes(info),
		StorageClassName: &info.VolumeStorageClassName,
		Resources:        v.getResource(info),
		VolumeMode:       v.getVolumeMode(info),
	}
	return pvc
}

func (v *VolumeDataService) getVolumeMode(info *volume.RVolumeInfo) *v12.PersistentVolumeMode {
	var pvm v12.PersistentVolumeMode
	switch info.VolumePersistentVolumeMode {
	case "Block":
		pvm = v12.PersistentVolumeBlock
	case "Filesystem":
		pvm = v12.PersistentVolumeFilesystem
	default:
		pvm = v12.PersistentVolumeFilesystem
	}
	return &pvm
}

func (v *VolumeDataService) getResource(info *volume.RVolumeInfo) v12.ResourceRequirements {
	return v12.ResourceRequirements{
		Requests: v12.ResourceList{
			"storage": resource.MustParse(strconv.FormatFloat(float64(info.VolumeRequest), 'f', 6, 64) + "Gi"),
		},
	}
}

func (v *VolumeDataService) getAccessModes(info *volume.RVolumeInfo) []v12.PersistentVolumeAccessMode {
	var pvam []v12.PersistentVolumeAccessMode
	var pm v12.PersistentVolumeAccessMode
	switch info.VolumeAccessMode {
	case "ReadWriteOnce":
		pm = v12.ReadWriteOnce
	case "ReadOnlyMany":
		pm = v12.ReadOnlyMany
	case "ReadWriteMany":
		pm = v12.ReadWriteMany
	case "ReadWriteOncePod":
		pm = v12.ReadWriteOncePod
	default:
		pm = v12.ReadWriteOnce
	}
	pvam = append(pvam, pm)
	return pvam
}
