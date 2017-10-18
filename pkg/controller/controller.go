package controller

import (
	"github.com/Sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/client-go/tools/cache"
)

type VMConfigController struct {
	logger *logrus.Entry
	clientset kubernetes.Interface
	queue workqueue.RateLimitingInterface
	informer cache.SharedIndexInformer
}
