package connect

import (
	"github.com/spf13/viper"
)

var Conn *Connector
var Conf *viper.Viper

// kubernetes resource type
const (
	ResourceGPU      = "nvidia.com/gpu"
	ResourceRDMA_HCA = "rdma/hca"
)

type Connector struct {
	Mongo *MongoConnector
	Kube  *K8sConnector
}

func Init(conf *viper.Viper) {
	m := MongoConnect(conf)
	k := K8sConnect(conf)

	Conf = conf
	Conn = &Connector{
		m,
		k,
	}
}

func Mongo() *MongoConnector {
	return Conn.Mongo
}

func K8s() *K8sConnector {
	return Conn.Kube
}

