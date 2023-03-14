package service

import (
	"context"
	"errors"
	"github.com/asim/go-micro/v3/logger"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"svc/domain/model"
	"svc/domain/repository"
	"svc/proto/svc"
)

type IServiceDataService interface {
	AddService(*model.Svc) (int64, error)
	DeleteServiceByID(int64) error
	UpdateService(*model.Svc) error
	QueryServiceByID(int64) (*model.Svc, error)
	QueryAllService() ([]model.Svc, error)
	CreateServiceToK8s(*svc.RSvcInfo) error
	UpdateServiceToK8s(*svc.RSvcInfo) error
	DeleteFromK8s(*model.Svc) error
}

type ScvDataService struct {
	ServiceRepository repository.ISvcRepository
	K8sClientSet      *kubernetes.Clientset
}

func NewServiceDateService(serviceRepository repository.ISvcRepository, clientSet *kubernetes.Clientset) IServiceDataService {
	return &ScvDataService{
		ServiceRepository: serviceRepository,
		K8sClientSet:      clientSet,
	}
}

func (s *ScvDataService) AddService(svc *model.Svc) (int64, error) {
	return s.ServiceRepository.CreateService(svc)
}

func (s *ScvDataService) DeleteServiceByID(id int64) error {
	return s.ServiceRepository.DeleteServiceByID(id)
}

func (s *ScvDataService) UpdateService(svc *model.Svc) error {
	return s.ServiceRepository.UpdateService(svc)
}

func (s *ScvDataService) QueryServiceByID(id int64) (*model.Svc, error) {
	return s.ServiceRepository.QueryServiceByID(id)
}

func (s *ScvDataService) QueryAllService() ([]model.Svc, error) {
	return s.ServiceRepository.QueryAllService()
}

func (s *ScvDataService) CreateServiceToK8s(svcInfo *svc.RSvcInfo) error {
	service := s.setService(svcInfo)
	// 查询k8s中是否存在该service
	// Get获取服务的名称，并返回相应的服务对象，如果存在，则返回错误, err == nil对象不存在, err != nil 对象存在
	if _, err := s.K8sClientSet.CoreV1().Services(svcInfo.SvcNamespace).Get(context.TODO(), svcInfo.SvcName, v12.GetOptions{}); err == nil {
		// service不存在
		if _, err = s.K8sClientSet.CoreV1().Services(svcInfo.SvcNamespace).Create(context.TODO(), service, v12.CreateOptions{}); err != nil {
			// 创建失败
			logger.Error(err)
			return err
		}
	} else { // err != nil,存在对象
		logger.Info("service " + svcInfo.SvcName + "已存在")
		return errors.New("service " + svcInfo.SvcName + "已存在")
	}
	return nil
}

func (s *ScvDataService) UpdateServiceToK8s(svcInfo *svc.RSvcInfo) error {
	service := s.setService(svcInfo)
	// 查询, GGet获取服务的名称，并返回相应的服务对象，如果存在，则返回错误, err == nil对象不存在, err != nil 对象存在
	if _, err := s.K8sClientSet.CoreV1().Services(svcInfo.SvcNamespace).Get(context.TODO(), svcInfo.SvcName, v12.GetOptions{}); err != nil {
		// 存在就更新
		if _, err = s.K8sClientSet.CoreV1().Services(svcInfo.SvcNamespace).Update(context.TODO(), service, v12.UpdateOptions{}); err != nil {
			// 更新失败
			logger.Error(err)
			return err
		}
		logger.Info("service " + svcInfo.SvcName + "更新成功")
		return nil
	} else {
		logger.Error(err)
		return errors.New("service " + svcInfo.SvcName + "不存在, 请先创建")
	}
}

func (s *ScvDataService) DeleteFromK8s(svc *model.Svc) error {
	// Delete 接受服务的名称并删除它。如果发生错误，则返回错误
	if err := s.K8sClientSet.CoreV1().Services(svc.SvcNamespace).Delete(context.TODO(), svc.SvcName, v12.DeleteOptions{}); err != nil {
		// 删除出错
		logger.Error(err)
		return err
	}
	return nil
}

func (s *ScvDataService) setService(svcInfo *svc.RSvcInfo) *v1.Service {
	// 创建service
	service := &v1.Service{}
	service.TypeMeta = v12.TypeMeta{
		APIVersion: "v1",      // 资源版本
		Kind:       "service", // 资源类型
	}
	service.ObjectMeta = v12.ObjectMeta{
		Name:      svcInfo.SvcName,
		Namespace: svcInfo.SvcNamespace,
		Labels: map[string]string{
			"app-name": svcInfo.SvcPodName,
		},
	}
	service.Spec = v1.ServiceSpec{
		Ports: s.getServicePort(svcInfo),
		Selector: map[string]string{
			"app-name": svcInfo.SvcPodName,
		},
		Type: "ClusterIP",
	}
	return service
}

func (s *ScvDataService) getServicePort(svcInfo *svc.RSvcInfo) (servicePort []v1.ServicePort) {
	for i := range svcInfo.SvcPort {
		// 包含有关service端口的信息
		servicePort = append(servicePort, v1.ServicePort{
			Name:       "port-" + strconv.FormatInt(int64(svcInfo.SvcPort[i].SvcPort), 10),
			Protocol:   v1.Protocol(svcInfo.SvcPort[i].SvcPortProtocol),
			Port:       svcInfo.SvcPort[i].SvcPort,
			TargetPort: intstr.FromInt(int(svcInfo.SvcPort[i].SvcTargetPort)),
		})
	}
	return
}
