package route

import (
	"encoding/json"
	"fmt"
	"time"

	"code.cloudfoundry.org/eirini"
	"code.cloudfoundry.org/eirini/route"
	"code.cloudfoundry.org/lager"
	apps "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type InstanceChangeInformer struct {
	Cancel     <-chan struct{}
	Client     kubernetes.Interface
	SyncPeriod time.Duration
	Namespace  string
	Logger     lager.Logger
}

func NewInstanceChangeInformer(client kubernetes.Interface, syncPeriod time.Duration, namespace string) route.Informer {
	return &InstanceChangeInformer{
		Client:     client,
		SyncPeriod: syncPeriod,
		Namespace:  namespace,
		Cancel:     make(<-chan struct{}),
		Logger:     lager.NewLogger("instance-change-informer"),
	}
}

func (c *InstanceChangeInformer) Start(work chan<- *route.Message) {
	factory := informers.NewSharedInformerFactoryWithOptions(c.Client,
		c.SyncPeriod,
		informers.WithNamespace(c.Namespace))

	podInformer := factory.Core().V1().Pods().Informer()
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		UpdateFunc: func(_ interface{}, updatedObj interface{}) {
			c.onPodUpdate(updatedObj, work)
		},
		DeleteFunc: func(obj interface{}) {
			c.onPodDelete(obj, work)
		},
	})

	podInformer.Run(c.Cancel)
}

func (c *InstanceChangeInformer) onPodDelete(deletedObj interface{}, work chan<- *route.Message) {
	deletedPod := deletedObj.(*v1.Pod)
	userDefinedRoutes, err := c.getUserDefinedRoutes(deletedPod)
	if err != nil {
		c.logError("failed-to-get-user-defined-routes", err, deletedPod)
		return
	}

	routes := newRoutes(deletedPod)
	routes.UnregisteredRoutes = userDefinedRoutes
	work <- routes
}

func (c *InstanceChangeInformer) onPodUpdate(updatedObj interface{}, work chan<- *route.Message) {
	updatedPod := updatedObj.(*v1.Pod)
	userDefinedRoutes, err := c.getUserDefinedRoutes(updatedPod)
	if err != nil {
		c.logError("failed-to-get-user-defined-routes", err, updatedPod)
		return
	}

	routes := newRoutes(updatedPod)
	routes.Routes = userDefinedRoutes
	work <- routes
}

func (c *InstanceChangeInformer) getUserDefinedRoutes(pod *v1.Pod) ([]string, error) {
	owner, err := c.getOwner(pod)
	if err != nil {
		c.logError("unexpected-pod-owner", err, pod)
		return []string{}, err
	}

	return decodeRoutes(owner.Annotations[eirini.RegisteredRoutes])
}

func (c *InstanceChangeInformer) logError(message string, err error, pod *v1.Pod) {
	if c.Logger != nil {
		c.Logger.Error(message, err, lager.Data{"pod-name": pod.Name})
	}
}

func (c *InstanceChangeInformer) getOwner(pod *v1.Pod) (*apps.StatefulSet, error) {
	ownerReferences := pod.OwnerReferences

	if len(ownerReferences) != 1 {
		return nil, fmt.Errorf("unexpected owner count - expected 1, but got %d", len(ownerReferences))
	}

	ownerName := ownerReferences[0].Name
	return c.Client.Apps().StatefulSets(c.Namespace).Get(ownerName, meta.GetOptions{})
}

func newRoutes(pod *v1.Pod) *route.Message {
	return &route.Message{
		Name:       pod.Name,
		InstanceID: pod.Name,
		Address:    pod.Status.PodIP,
		Port:       getContainerPort(pod),
		TLSPort:    0,
	}
}

func getContainerPort(pod *v1.Pod) uint32 {
	port := pod.Spec.Containers[0].Ports[0].ContainerPort
	return uint32(port)
}

func decodeRoutes(s string) ([]string, error) {
	uris := []string{}
	err := json.Unmarshal([]byte(s), &uris)

	return uris, err
}