package client

import (
	"log"
	"path/filepath"

	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	k8scmd "k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func KubeClient() *kubernetes.Clientset {
	return kubeClient()
}

// KUBE_CONFIGがあったらそのまま
// なかったら~/.kube/configを返す
func kubeConfigPath() string {
	var confPath string
	if os.Getenv("KUBE_CONFIG") != "" {
		return os.Getenv("KUBE_CONFIG")
	}
	confPath = filepath.Join(homedir.HomeDir(), ".kube", "config")
	_, err := os.Stat(confPath)
	if err != nil {
		return ""
	}
	return confPath
}

// confファイルがあったらconfから
// なかったらInClusterということにしとく
func kubeConfig() (config *rest.Config) {
	var err error
	if kubeConfigPath() == "" {
		config, err = rest.InClusterConfig()
	} else {
		config, err = k8scmd.BuildConfigFromFlags("", kubeConfigPath())
	}
	if err != nil {
		log.Println(err)
	}
	return
}

func kubeClient() *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(kubeConfig())
	if err != nil {
		log.Println(err)
	}
	return clientset
}
