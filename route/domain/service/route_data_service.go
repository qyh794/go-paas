package service

import (
	"context"
	"errors"
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/route/domain/model"
	"github.com/qyh794/go-paas/route/domain/repository"
	"github.com/qyh794/go-paas/route/proto/route"
	v1 "k8s.io/api/apps/v1"
	v2 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type IRouteDataService interface {
	AddRoute(*model.Route) (int64, error)
	DeleteRoute(int64) error
	UpdateRoute(*model.Route) error
	QueryRouteByID(int64) (*model.Route, error)
	QueryAllRoute() ([]model.Route, error)
	CreateRouteToK8s(*route.RRouteInfo) error
	DeleteRouteFromK8s(*model.Route) error
	UpdateRouteToK8s(*route.RRouteInfo) error
}

type RouteDataService struct {
	RouteRepository repository.IRouteRepository
	K8sClientSet    *kubernetes.Clientset
	deployment      *v1.Deployment
}

func NewRouteDataService(routeRepository repository.IRouteRepository, k8sClientSet *kubernetes.Clientset) IRouteDataService {
	return &RouteDataService{
		RouteRepository: routeRepository,
		K8sClientSet:    k8sClientSet,
		deployment:      &v1.Deployment{},
	}
}

func (r *RouteDataService) AddRoute(route *model.Route) (int64, error) {
	return r.RouteRepository.CreateRoute(route)
}

func (r *RouteDataService) DeleteRoute(id int64) error {
	return r.RouteRepository.DeleteRouteByID(id)
}

func (r *RouteDataService) UpdateRoute(route *model.Route) error {
	return r.RouteRepository.UpdateRoute(route)
}

func (r *RouteDataService) QueryRouteByID(id int64) (*model.Route, error) {
	return r.RouteRepository.QueryRouteByID(id)
}

func (r *RouteDataService) QueryAllRoute() ([]model.Route, error) {
	return r.RouteRepository.QueryAllRoute()
}

func (r *RouteDataService) CreateRouteToK8s(info *route.RRouteInfo) error {
	// 配置ingress
	ingress := r.setIngress(info)
	// 判断k8s中是否存在该路由, err == nil 说明存在
	if _, err := r.K8sClientSet.NetworkingV1().Ingresses(info.RouteNamespace).Get(context.TODO(), info.RouteName, metav1.GetOptions{}); err == nil {
		logger.Info("路由: " + info.RouteName + " 已存在")
		return errors.New("路由: " + info.RouteName + " 已存在")
	} else { // 不存在就创建
		if _, err = r.K8sClientSet.NetworkingV1().Ingresses(info.RouteNamespace).Create(context.TODO(), ingress, metav1.CreateOptions{}); err != nil {
			// err != nil,创建失败
			logger.Error(err)
			return err
		}
	}
	return nil
}

func (r *RouteDataService) DeleteRouteFromK8s(route *model.Route) error {
	// 根据ingress名称删除
	if err := r.K8sClientSet.NetworkingV1().Ingresses(route.RouteNamespace).Delete(context.TODO(), route.RouteName, metav1.DeleteOptions{}); err != nil {
		// 删除失败
		logger.Error("删除ingress失败, ", err)
		return err
	}
	// 删除成功
	return nil
}

func (r *RouteDataService) UpdateRouteToK8s(info *route.RRouteInfo) error {
	ingress := r.setIngress(info)
	// 更新
	if _, err := r.K8sClientSet.NetworkingV1().Ingresses(info.RouteNamespace).Update(context.TODO(), ingress, metav1.UpdateOptions{}); err != nil {
		// 更新失败
		logger.Error("更新ingress失败, ", err)
		return err
	}
	// 更新成功
	return nil
}

func (r *RouteDataService) setIngress(info *route.RRouteInfo) *v2.Ingress {
	routeObj := &v2.Ingress{}
	routeObj.TypeMeta = metav1.TypeMeta{
		Kind:       "Ingress",
		APIVersion: "v1",
	}
	// ObjectMeta 元数据
	routeObj.ObjectMeta = metav1.ObjectMeta{
		Name:      info.RouteName,
		Namespace: info.RouteNamespace,
		Labels: map[string]string{
			"app-name": info.RouteName,
		},
	}
	className := "nginx"
	routeObj.Spec = v2.IngressSpec{
		IngressClassName: &className,
		DefaultBackend:   nil,
		TLS:              nil,
		Rules:            r.getIngressPath(info),
	}
	return routeObj
}

func (r *RouteDataService) getIngressPath(info *route.RRouteInfo) []v2.IngressRule {
	var path []v2.IngressRule
	pathRule := v2.IngressRule{Host: info.RouteHost}
	var ingressPath []v2.HTTPIngressPath
	for i := range info.RoutePath {
		pathType := v2.PathTypePrefix
		ingressPath = append(ingressPath, v2.HTTPIngressPath{
			Path:     info.RoutePath[i].RoutePathName,
			PathType: &pathType,
			Backend: v2.IngressBackend{
				Service: &v2.IngressServiceBackend{
					Name: info.RoutePath[i].RouteBackendService,
					Port: v2.ServiceBackendPort{
						Number: info.RoutePath[i].RouteBackendServicePort,
					},
				},
			},
		})
	}
	pathRule.IngressRuleValue = v2.IngressRuleValue{HTTP: &v2.HTTPIngressRuleValue{Paths: ingressPath}}
	path = append(path, pathRule)
	return path
}
