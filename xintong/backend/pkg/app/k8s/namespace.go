package k8s

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/workspace"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

func ExecNameSpace(ws workspace.Workspace) error {
	resourceQuota, namespace, err := newNameSpace(ws)
	if err != nil {
		return err
	}
	namespaceClient := connect.K8s().Clientset.CoreV1().Namespaces()
	//namespaceClient.Delete(namespace.Name, nil)
	option :=metav1.GetOptions{}
	n,err := namespaceClient.Get(namespace.Name,option)
	if(len(n.Name)==0) {
		_, err = namespaceClient.Create(namespace)
		if err != nil {
			return err
		}
	}
	resourceQuotaClient := connect.K8s().Clientset.CoreV1().ResourceQuotas(namespace.Name)
	resourceQuotaClient.Delete(resourceQuota.Name, nil)
	_, err = resourceQuotaClient.Create(resourceQuota)
	if err != nil {
		return err
	}
	return nil

}

func newNameSpace(ws workspace.Workspace) (*corev1.ResourceQuota, *corev1.Namespace, error) {

	// TODO
	resourceAlloc := corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.FormatInt(ws.CpuNum, 10)),
		corev1.ResourceMemory: resource.MustParse(strconv.FormatInt(ws.MemoryNum, 10) + "Gi"),
		ResourceGPU:           resource.MustParse(strconv.FormatInt(ws.GpuNum, 10)),
		ResourceRDMA_HCA:      resource.MustParse(strconv.FormatInt(ws.Rdma, 10)),
	}
	resourceQuota := &corev1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Name: ws.Name + "-resourcequota",
		},
		Spec: corev1.ResourceQuotaSpec{
			Hard: resourceAlloc,
		},
	}
	nameSpace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: ws.Name,
		},
	}
	return resourceQuota, nameSpace, nil
}
