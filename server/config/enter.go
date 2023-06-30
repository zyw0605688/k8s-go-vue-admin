package config

import (
	"flag"
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
)

var Config *rest.Config
var RestClient *rest.RESTClient
var ClientSet *kubernetes.Clientset

func InitK8s() {
	GetConfig()
	GetRestClient()
	GetClientSet()
}

func GetConfig() {
	var kubeConfig *string
	kubeConfig = flag.String("kubeConfig", filepath.Join("./config/145.yaml"), "file path")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	Config = config
}

func GetRestClient() {
	Config.APIPath = "api"
	Config.GroupVersion = &v1.SchemeGroupVersion
	Config.NegotiatedSerializer = scheme.Codecs
	restClient, err := rest.RESTClientFor(Config)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	RestClient = restClient
}

func GetClientSet() {
	clientSet, err := kubernetes.NewForConfig(Config)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	ClientSet = clientSet
}
