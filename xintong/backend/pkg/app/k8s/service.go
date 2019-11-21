package k8s

import (
	"fmt"
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dlservice"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func ExecService(serviceVo dlservice.DlserviceVo) error {

	service, deployment, ingress, err := newService(serviceVo)
	if err != nil {
		return err
	}
	deploymentClient := connect.K8s().Clientset.AppsV1().Deployments("modelzoo")
	deploymentClient.Delete(deployment.Name, nil)
	_, err = deploymentClient.Create(deployment)
	if err != nil {
		fmt.Print("deployment error %s", err.Error())
		return err
	}

	serviceClient := connect.K8s().Clientset.CoreV1().Services("modelzoo")
	serviceClient.Delete(service.Name, nil)
	_, err = serviceClient.Create(service)
	if err != nil {
		fmt.Print("service error %s", err.Error())
		return err
	}

	ingressClient := connect.K8s().Clientset.ExtensionsV1beta1().Ingresses("modelzoo")
	ingressClient.Delete(ingress.Name, nil)
	_, err = ingressClient.Create(ingress)
	if err != nil {
		fmt.Print("ingress error %s", err.Error())
		return err
	}

	return nil
}

func newService(serviceVo dlservice.DlserviceVo) (*corev1.Service, *appsv1.Deployment, *v1beta1.Ingress, error) {

	conf := connect.Conf
	nfsConf := conf.GetStringMapString("nfs")
	nfsHost := nfsConf["host"]
	nfsPath := nfsConf["path"]
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceVo.Name + "-service",
			Namespace: "modelzoo",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": serviceVo.Name + "-app",
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					TargetPort: intstr.IntOrString{IntVal: 8501},
					Port:       8501,
				},
			},
		},
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceVo.Name + "deploy",
			Namespace: "modelzoo",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"release": "canary",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     serviceVo.Name + "-app",
						"release": "canary",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name: serviceVo.Name,
							//Image: "core.harbor.domain/modelzoo/" + serviceVo.Image.Name + ":" + serviceVo.Image.Version,
							Image: "core.harbor.domain/modelzoo/tensorflow/serving:1.0",
							VolumeMounts: []corev1.VolumeMount{
								{Name: "model-persistent-storage", MountPath: "/models/" + serviceVo.Name},
							},
							Args: []string{"rest_api_port=0.0.0.0:8501"},
							Env:  []corev1.EnvVar{{Name: "MODEL_NAME", Value: serviceVo.Name}},
							Ports: []corev1.ContainerPort{
								{
									HostPort:      8501,
									ContainerPort: 8501,
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{Name: "model-persistent-storage",
							VolumeSource: corev1.VolumeSource{
								NFS: &corev1.NFSVolumeSource{
									Path:   nfsPath + serviceVo.Dlmodel.Url,
									Server: nfsHost,
								},
							},
						},
					},
				},
			},
		},
	}

	ingress := &v1beta1.Ingress{

		ObjectMeta: metav1.ObjectMeta{
			Name:      "ingress-" + serviceVo.Name,
			Namespace: "modelzoo",
			Annotations: map[string]string{
				"kubernetes.io/ingress.class": "nginx",
			},
		},
		Spec: v1beta1.IngressSpec{
			Rules: []v1beta1.IngressRule{
				{
					Host: serviceVo.Name + ".com",
					IngressRuleValue: v1beta1.IngressRuleValue{
						HTTP: &v1beta1.HTTPIngressRuleValue{
							Paths: []v1beta1.HTTPIngressPath{
								{
									Path: "/",
									Backend: v1beta1.IngressBackend{
										ServiceName: serviceVo.Name + "-service",
										ServicePort: intstr.IntOrString{IntVal: 8501}}},
							},
						},
					}},
			},
		},
	}
	return service, deployment, ingress, nil
}
