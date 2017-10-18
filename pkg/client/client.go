package client

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	log "github.com/Sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Run(config string, url string) {
	cfg, err := clientcmd.BuildConfigFromFlags(url, config)
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
				"pod": pod.Name,
				"namespace": pod.Namespace,
			}).Info("pod")
	}
}
