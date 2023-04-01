package service

import (
	"context"
	"errors"
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/middleware/domain/model"
	"github.com/qyh794/go-paas/middleware/domain/repository"
	"github.com/qyh794/go-paas/middleware/proto/middleware"
	v1 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
)

type IMiddlewareDataService interface {
	AddMiddleware(*model.Middleware) (int64, error)
	DeleteMiddlewareByID(int64) error
	UpdateMiddleware(*model.Middleware) error
	QueryMiddlewareByID(int64) (*model.Middleware, error)
	QueryAllMiddleware() ([]model.Middleware, error)
	QueryAllMiddlewareByTypeID(int64) ([]model.Middleware, error)
	CreateToK8s(*middleware.RMiddlewareInfo) error
	DeleteFromK8s(*model.Middleware) error
	UpdateToK8s(*middleware.RMiddlewareInfo) error
}

type MiddlewareDataService struct {
	MiddlewareRepository repository.IMiddlewareRepository
	K8sClientSet         *kubernetes.Clientset
}

func NewMiddlewareDataService(middlewareRepository repository.IMiddlewareRepository, clientSet *kubernetes.Clientset) IMiddlewareDataService {
	return &MiddlewareDataService{
		MiddlewareRepository: middlewareRepository,
		K8sClientSet:         clientSet,
	}
}

func (m *MiddlewareDataService) AddMiddleware(middlewareObj *model.Middleware) (int64, error) {
	return m.MiddlewareRepository.CreateMiddleware(middlewareObj)
}

func (m *MiddlewareDataService) DeleteMiddlewareByID(middlewareID int64) error {
	return m.MiddlewareRepository.DeleteMiddlewareByID(middlewareID)
}

func (m *MiddlewareDataService) UpdateMiddleware(middlewareObj *model.Middleware) error {
	return m.MiddlewareRepository.UpdateMiddleware(middlewareObj)
}

func (m *MiddlewareDataService) QueryMiddlewareByID(middlewareID int64) (*model.Middleware, error) {
	return m.MiddlewareRepository.QueryMiddlewareByID(middlewareID)
}

func (m *MiddlewareDataService) QueryAllMiddleware() ([]model.Middleware, error) {
	return m.MiddlewareRepository.QueryAllMiddleware()
}

func (m *MiddlewareDataService) QueryAllMiddlewareByTypeID(middlewareTypeID int64) ([]model.Middleware, error) {
	return m.MiddlewareRepository.QueryAllMiddlewareByType(middlewareTypeID)
}

func (m *MiddlewareDataService) CreateToK8s(middlewareInfo *middleware.RMiddlewareInfo) error {
	statefulSet := m.setStatefulSet(middlewareInfo)
	if _, err := m.K8sClientSet.AppsV1().StatefulSets(middlewareInfo.MiddleNamespace).Get(context.TODO(), middlewareInfo.MiddleName, v12.GetOptions{}); err != nil {
		// 不存在
		if _, err = m.K8sClientSet.AppsV1().StatefulSets(middlewareInfo.MiddleNamespace).Create(context.TODO(), statefulSet, v12.CreateOptions{}); err != nil {
			// 创建失败
			logger.Error(err)
			return err
		}
		logger.Info("中间件: " + middlewareInfo.MiddleName + "创建成功")
		return nil
	} else { // k8s存在已存在该资源
		logger.Info("中间件: " + middlewareInfo.MiddleName + "已存在,请勿重复创建")
		return errors.New("中间件: " + middlewareInfo.MiddleName + "已存在,请勿重复创建")
	}
}

func (m *MiddlewareDataService) DeleteFromK8s(middlewareObj *model.Middleware) error {
	if err := m.K8sClientSet.AppsV1().StatefulSets(middlewareObj.MiddleNamespace).Delete(context.TODO(), middlewareObj.MiddleName, v12.DeleteOptions{}); err != nil {
		// statefulSet删除失败
		logger.Error(err)
		return err
	}
	return nil
}

func (m *MiddlewareDataService) UpdateToK8s(middlewareInfo *middleware.RMiddlewareInfo) error {
	statefulSet := m.setStatefulSet(middlewareInfo)
	if _, err := m.K8sClientSet.AppsV1().StatefulSets(middlewareInfo.MiddleNamespace).Get(context.TODO(), middlewareInfo.MiddleName, v12.GetOptions{}); err != nil {
		// k8s中不存在该资源
		logger.Error(err)
		return errors.New("中间件 " + middlewareInfo.MiddleName + " 不存在,请先创建")
	} else {
		// k8s中存在该资源
		if _, err = m.K8sClientSet.AppsV1().StatefulSets(middlewareInfo.MiddleNamespace).Update(context.TODO(), statefulSet, v12.UpdateOptions{}); err != nil {
			// 更新失败
			logger.Error(err)
			return err
		}
		logger.Info("中间件 " + middlewareInfo.MiddleName + " 更新成功")
		return nil
	}
}

func (m *MiddlewareDataService) setStatefulSet(middlewareInfo *middleware.RMiddlewareInfo) *v1.StatefulSet {
	statefulSet := &v1.StatefulSet{}
	statefulSet.TypeMeta = v12.TypeMeta{
		Kind:       "StatefulSet",
		APIVersion: "v1",
	}
	statefulSet.ObjectMeta = v12.ObjectMeta{
		Name:      middlewareInfo.MiddleName,
		Namespace: middlewareInfo.MiddleNamespace,
		Labels: map[string]string{
			"app-name": middlewareInfo.MiddleName,
		},
	}
	statefulSet.Spec = v1.StatefulSetSpec{
		Replicas: &middlewareInfo.MiddleReplicas,
		Selector: &v12.LabelSelector{
			MatchLabels: map[string]string{
				"app-name": middlewareInfo.MiddleName,
			},
		},
		// Template在检测到复制副本不足时将创建的pod的对象
		//由StatefulSet创建出来的每个pod都将实现该模板
		Template: v13.PodTemplateSpec{
			ObjectMeta: v12.ObjectMeta{
				Labels: map[string]string{
					"app-name": middlewareInfo.MiddleName,
				},
			},
			// pod 详细描述spec
			Spec: v13.PodSpec{
				Containers: []v13.Container{ // 要在pod中运行的单个应用程序容器
					{
						Name:      middlewareInfo.MiddleName,
						Image:     middlewareInfo.MiddleDockerImageVersion,
						Ports:     m.getPort(middlewareInfo.MiddlePort),
						Env:       m.getEnv(middlewareInfo.MiddleEnv),
						Resources: m.getResources(middlewareInfo.MiddleCpu, middlewareInfo.MiddleMemory),
						// 挂载目录
						VolumeMounts: m.getVolumeMounts(middlewareInfo.MiddleStorage),
					},
				},
				TerminationGracePeriodSeconds: m.getTime("10"),
				ImagePullSecrets:              nil,
			},
		},
		VolumeClaimTemplates: m.getPVC(middlewareInfo.MiddleStorage, middlewareInfo.MiddleNamespace),
		ServiceName:          middlewareInfo.MiddleName,
	}
	return statefulSet
}

func (m *MiddlewareDataService) getPVC(middlewareStorageInfo []*middleware.MiddleStorage, namespace string) []v13.PersistentVolumeClaim {
	if len(middlewareStorageInfo) == 0 {
		return nil
	}
	pvcArr := make([]v13.PersistentVolumeClaim, 0, len(middlewareStorageInfo))
	for i := 0; i < len(middlewareStorageInfo); i++ {
		pvc := v13.PersistentVolumeClaim{
			TypeMeta: v12.TypeMeta{
				Kind:       "PersistentVolumeClaim",
				APIVersion: "v1",
			},
			ObjectMeta: v12.ObjectMeta{
				Name:      middlewareStorageInfo[i].MiddleStorageName,
				Namespace: namespace,
			},
			Spec: v13.PersistentVolumeClaimSpec{
				AccessModes:      m.getAccessModes(middlewareStorageInfo[i].MiddleStorageAccessMode),
				Resources:        m.getPVCResources(middlewareStorageInfo[i].MiddleStorageSize),
				VolumeName:       middlewareStorageInfo[i].MiddleStorageName,
				StorageClassName: &middlewareStorageInfo[i].MiddleStorageClass,
			},
		}
		pvcArr = append(pvcArr, pvc)
	}
	return pvcArr
}

// 获取pvc访问方式
func (m *MiddlewareDataService) getAccessModes(modes string) []v13.PersistentVolumeAccessMode {
	var accessModeArr []v13.PersistentVolumeAccessMode
	var accessMode v13.PersistentVolumeAccessMode
	switch modes {
	/*
		// can be mounted in read/write mode to exactly 1 host
		ReadWriteOnce PersistentVolumeAccessMode = "ReadWriteOnce"
		// can be mounted in read-only mode to many hosts
		ReadOnlyMany PersistentVolumeAccessMode = "ReadOnlyMany"
		// can be mounted in read/write mode to many hosts
		ReadWriteMany PersistentVolumeAccessMode = "ReadWriteMany"
		// can be mounted in read/write mode to exactly 1 pod
		// cannot be used in combination with other access modes
		ReadWriteOncePod PersistentVolumeAccessMode = "ReadWriteOncePod"
	*/
	case "ReadWriteOnce":
		accessMode = v13.ReadWriteOnce
	case "ReadOnlyMany":
		accessMode = v13.ReadOnlyMany
	case "ReadWriteMany":
		accessMode = v13.ReadWriteMany
	case "ReadWriteOncePod":
		accessMode = v13.ReadWriteOncePod
	default:
		accessMode = v13.ReadWriteOnce
	}
	accessModeArr = append(accessModeArr, accessMode)
	return accessModeArr
}

func (m *MiddlewareDataService) getPVCResources(middleStorageSize float32) v13.ResourceRequirements {
	var resourceRequirements v13.ResourceRequirements
	resourceRequirements.Requests = v13.ResourceList{
		"storage": resource.MustParse(strconv.FormatFloat(float64(middleStorageSize), 'f', 6, 64) + "Gi"),
	}
	return resourceRequirements
}

func (m *MiddlewareDataService) getTime(time string) *int64 {
	i, err := strconv.ParseInt(time, 10, 64)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return &i
}

func (m *MiddlewareDataService) getVolumeMounts(volume []*middleware.MiddleStorage) []v13.VolumeMount {
	if len(volume) == 0 {
		return nil
	}
	var volumeMountArr []v13.VolumeMount
	for i := 0; i < len(volume); i++ {
		volumeMount := v13.VolumeMount{
			Name:      volume[i].MiddleStorageName,
			MountPath: volume[i].MiddleStoragePath,
		}
		volumeMountArr = append(volumeMountArr, volumeMount)
	}
	return volumeMountArr
}

func (m *MiddlewareDataService) getResources(middlewareCpu float32, middlewareMemory float32) v13.ResourceRequirements {
	var resourceRequirements v13.ResourceRequirements
	resourceRequirements.Limits = v13.ResourceList{
		"cpu":    resource.MustParse(strconv.FormatFloat(float64(middlewareCpu), 'f', 6, 64)),
		"memory": resource.MustParse(strconv.FormatFloat(float64(middlewareMemory), 'f', 6, 64)),
	}
	resourceRequirements.Requests = v13.ResourceList{
		"cpu":    resource.MustParse(strconv.FormatFloat(float64(middlewareCpu), 'f', 6, 64)),
		"memory": resource.MustParse(strconv.FormatFloat(float64(middlewareMemory), 'f', 6, 64)),
	}
	return resourceRequirements
}

// 调用函数时，只传递 info.MiddleEnv，而不是整个 info 对象，这样可以避免创建和传递不必要的对象,
// 避免了在函数内部访问不必要的字段和分配内存空间，提高性能。
func (m *MiddlewareDataService) getEnv(middleEnvs []*middleware.MiddleEnv) []v13.EnvVar {
	var envVarArr []v13.EnvVar
	for i := 0; i < len(middleEnvs); i++ {
		envVar := v13.EnvVar{
			Name:      middleEnvs[i].EnvKey,
			Value:     middleEnvs[i].EnvValue,
			ValueFrom: nil,
		}
		envVarArr = append(envVarArr, envVar)
	}
	return envVarArr
}

func (m *MiddlewareDataService) getPort(middlewarePort []*middleware.MiddlePort) []v13.ContainerPort {
	var ContainerPortArr []v13.ContainerPort
	for i := 0; i < len(middlewarePort); i++ { //int32
		containerPort := v13.ContainerPort{
			Name:          "middleware-port-" + strconv.FormatInt(int64(middlewarePort[i].MiddlePort), 10),
			ContainerPort: middlewarePort[i].MiddlePort,
			Protocol:      m.getProtocol(middlewarePort[i].MiddleProtocol),
		}
		ContainerPortArr = append(ContainerPortArr, containerPort)
	}
	return ContainerPortArr
}

func (m *MiddlewareDataService) getProtocol(middlewareProtocol string) v13.Protocol {
	switch middlewareProtocol {
	case "TCP":
		return "TCP"
	case "UDP":
		return "UDP"
	case "SCTP":
		return "SCTP"
	default:
		return "TCP"
	}
}
