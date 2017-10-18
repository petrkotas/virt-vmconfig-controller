package client

import (
	log "github.com/Sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Run(config string) {
	cfg, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		log.Error("Cannot connect to kubernetes.")
	}

	cl, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.WithField("Err", err).Error("Cannot create client")
	}

	pods, err := cl.CoreV1().Pods("kube-system").List(v1.ListOptions{})

	for _, pod := range pods.Items {
		log.WithFields(
			log.Fields{
				"pod":       pod.Name,
				"namespace": pod.Namespace,
			}).Info("pod")
	}
}
