package service

import (
	"context"
	"errors"
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/pod/domain/model"
	"github.com/qyh794/go-paas/pod/domain/repository"
	"github.com/qyh794/go-paas/pod/proto/pod"
	v1 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
)

type IPodDataService interface {
	AddPod(pod *model.Pod) (int64, error)
	DeletePod(int64) error
	UpdatePod(pod *model.Pod) error
	QueryPodByID(int64) (*model.Pod, error)
	QueryAllPods() ([]model.Pod, error)
	CreateToK8s(info *pod.RPodInfo) error
	DeleteFromK8s(pod *model.Pod) error
	UpdateToK8s(info *pod.RPodInfo) error
}

type PodDateService struct {
	PodRepository repository.IPodRepository
	// k8s客户端
	K8sClientSet *kubernetes.Clientset
	deployment   *v1.Deployment
}

func NewPodDateService(podRepository repository.IPodRepository, client *kubernetes.Clientset) IPodDataService {
	return &PodDateService{
		PodRepository: podRepository,
		K8sClientSet:  client,
		deployment:    &v1.Deployment{},
	}
}

// AddPod 添加pod
func (p *PodDateService) AddPod(pod *model.Pod) (int64, error) {
	return p.PodRepository.CreatePod(pod)
}

// DeletePod 删除pod
func (p *PodDateService) DeletePod(id int64) error {
	return p.PodRepository.DeletePodById(id)
}

// UpdatePod 更新pod
func (p *PodDateService) UpdatePod(pod *model.Pod) error {
	return p.PodRepository.UpdatePod(pod)
}

// QueryPodByID 根据id查询pod信息
func (p *PodDateService) QueryPodByID(id int64) (*model.Pod, error) {
	return p.PodRepository.QueryPodByID(id)
}

// QueryAllPods 查询所有pod
func (p *PodDateService) QueryAllPods() ([]model.Pod, error) {
	return p.PodRepository.QueryAllPods()
}

// CreateToK8s 创建pod到k8s中
func (p *PodDateService) CreateToK8s(info *pod.RPodInfo) error {
	// 创建deployment
	p.SetDeployment(info)
	// AppsV1()返回一个AppsV1Client

	/*

		func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
			return c.appsV1
		}

		func (c *AppsV1Client) Deployments(namespace string) DeploymentInterface {
			return newDeployments(c, namespace)
		}

		func newDeployments(c *AppsV1Client, namespace string) *deployments {
			return &deployments{
				client: c.RESTClient(),
				ns:     namespace,
			}
		}

		func (c *deployments) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Deployment, err error) {
			result = &v1.Deployment{}
			err = c.client.Get().
				Namespace(c.ns).
				Resource("deployments").
				Name(name).
				VersionedParams(&options, scheme.ParameterCodec).
				Do(ctx).
				Into(result)
			return
		}

	*/
	// Deployments(info.PodNamespace) 通过命名空间查找deployment
	// Get() 获取deployment的名称，并返回相应的deployment对象，如果存在则返回错误, 有错误说明对象存在,没有错误说明对象不存在
	_, err := p.K8sClientSet.AppsV1().Deployments(info.PodNamespace).Get(context.TODO(), info.PodName, v12.GetOptions{})
	// err == nil 说明不存在deployment对象
	// 不存在就创建
	if err == nil {
		// 创建deployment对象
		_, err = p.K8sClientSet.AppsV1().Deployments(info.PodNamespace).Create(context.TODO(), p.deployment, v12.CreateOptions{})
		if err != nil {
			// 创建失败
			logger.Error(err)
			return err
		}
		logger.Infof("创建deployment成功!")
		return err
	} else { //存在deployment
		logger.Error("Pod " + info.PodName + "已存在!")
		return errors.New("Pod " + info.PodName + "已存在!")
	}
}

// DeleteFromK8s 从k8s删除pod
func (p *PodDateService) DeleteFromK8s(pod *model.Pod) error {
	// 获取并删除指定命名空间下的deployment
	err := p.K8sClientSet.AppsV1().Deployments(pod.PodNamespace).Delete(context.TODO(), pod.PodName, v12.DeleteOptions{})
	if err != nil { //deployment不存在
		logger.Error("deployment不存在", err)
		return err
	} else { // deployment存在并已删除
		// 删除deployment创建的pod信息
		if err = p.PodRepository.DeletePodById(pod.ID); err != nil {
			logger.Error("pod删除出现错误: ", err)
			return err
		}
		logger.Infof("删除pod ID: " + strconv.FormatInt(pod.ID, 10) + "成功")
	}
	return err
}

// UpdateToK8s 更新deployment、pod
func (p *PodDateService) UpdateToK8s(info *pod.RPodInfo) error {
	p.SetDeployment(info)
	if _, err := p.K8sClientSet.AppsV1().Deployments(info.PodNamespace).Get(context.TODO(), info.PodName, v12.GetOptions{}); err != nil {
		logger.Error(err)
		return errors.New("Pod " + info.PodName + " 不存在,请先创建")
	} else { // 存在
		if _, err = p.K8sClientSet.AppsV1().Deployments(info.PodNamespace).Update(context.TODO(), p.deployment, v12.UpdateOptions{}); err != nil {
			logger.Error(err)
			return err
		}
		logger.Infof(info.PodName + " 创建成功")
		return nil
	}
}

// SetDeployment 创建deployment控制器
func (p *PodDateService) SetDeployment(info *pod.RPodInfo) {
	// 创建deployment控制器
	deployment := &v1.Deployment{}
	// 指定pod控制器类型和版本号
	deployment.TypeMeta = v12.TypeMeta{
		// 控制器类型为deployment
		Kind:       "deployment",
		APIVersion: "v1",
	}
	// deployment元数据信息
	deployment.ObjectMeta = v12.ObjectMeta{
		// replicaSet名称
		Name: info.PodName,
		// 所属命名空间
		Namespace: info.PodNamespace,
		// 标签
		Labels: map[string]string{
			"app-name": info.PodName,
			"author":   "qyh",
		},
	}
	// deployment名称
	deployment.Name = info.PodName
	// deployment详细描述,spec
	deployment.Spec = v1.DeploymentSpec{
		// 副本个数
		Replicas: &info.PodReplicas,
		// 标签选择器
		Selector: &v12.LabelSelector{
			// Labels匹配规则
			MatchLabels: map[string]string{
				"app-name": info.PodName,
			},
			// Expressions匹配规则
			MatchExpressions: nil,
		},
		// pod 模板
		Template: v13.PodTemplateSpec{
			ObjectMeta: v12.ObjectMeta{
				Labels: map[string]string{
					"app-name": info.PodName,
				},
			},
			// pod 详细描述spec
			Spec: v13.PodSpec{
				Containers: []v13.Container{ // 要在pod中运行的单个应用程序容器
					{
						Name:            info.PodName,               // 容器名称
						Image:           info.PodImage,              // 容器镜像名称
						Ports:           p.getContainerPort(info),   // 需要暴露的端口号列表
						Env:             p.getEnv(info),             // 容器运行前需要设置的环境变量列表
						Resources:       p.getResources(info),       // 资源限制和请求设置
						ImagePullPolicy: p.getImagePullPolicy(info), // 镜像拉去策略
					},
				},
			},
		},
		// 策略
		Strategy:        v1.DeploymentStrategy{},
		MinReadySeconds: 0,
		// 保留历史版本
		RevisionHistoryLimit: nil,
		// 暂停部署
		Paused: false,
		// 部署超时时间
		ProgressDeadlineSeconds: nil,
	}
	p.deployment = deployment
}

// getContainerPort 获取pod容器端口
func (p *PodDateService) getContainerPort(info *pod.RPodInfo) (containerPort []v13.ContainerPort) {
	for _, port := range info.PodPort {
		containerPort = append(containerPort, v13.ContainerPort{
			Name:          "port-" + strconv.FormatInt(int64(port.ContainerPort), 10),
			ContainerPort: port.ContainerPort,
			Protocol:      p.getProtocol(port.Protocol),
		})
	}
	return
}

func (p *PodDateService) getProtocol(protocol string) v13.Protocol {
	switch protocol {
	case "TCP":
		return "TCP"
	case "UDP":
		return "UDP"
	default:
		return "TCP"
	}
}

func (p *PodDateService) getEnv(info *pod.RPodInfo) (envVar []v13.EnvVar) {
	for _, env := range info.PodEnv {
		envVar = append(envVar, v13.EnvVar{
			Name:  env.EnvKey,
			Value: env.EnvValue,
		})
	}
	return
}

func (p *PodDateService) getResources(info *pod.RPodInfo) (source v13.ResourceRequirements) {
	// Limits描述允许的最大能使用的资源
	source.Limits = v13.ResourceList{
		"cpu":    resource.MustParse(strconv.FormatFloat(float64(info.PodCpuMax), 'f', 6, 64)),
		"memory": resource.MustParse(strconv.FormatFloat(float64(info.PodMemoryMax), 'f', 6, 64)),
	}
	// Requests描述了所需的最小计算资源
	source.Requests = v13.ResourceList{
		"cpu":    resource.MustParse(strconv.FormatFloat(float64(info.PodCpuMin), 'f', 6, 64)),
		"memory": resource.MustParse(strconv.FormatFloat(float64(info.PodMemoryMin), 'f', 6, 64)),
	}
	return
}

func (p *PodDateService) getImagePullPolicy(info *pod.RPodInfo) v13.PullPolicy {
	switch info.PodPullPolicy {
	case "Always":
		return "Always"
	case "Never":
		return "Never"
	case "IfNotPresent":
		return "IfNotPresent"
	default:
		return "Always"
	}
}
