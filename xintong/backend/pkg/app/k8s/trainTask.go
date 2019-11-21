package k8s

import (
	"fmt"
	"github.com/igm/sockjs-go/sockjs"
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/train"
	"github.com/xxmyjk/xintong/backend/pkg/app/util"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	"strconv"
)

const (
	ResourceGPU      = "nvidia.com/gpu"
	ResourceRDMA_HCA = "rdma/hca"
)

// TODO: tmp
func Bool(b bool) *bool {
	return &b
}

func ExecStateSet(trainTaskVo train_task.Train_taskVo) error {

	statefulSet, service, err := newStatefulSet(trainTaskVo)
	if err != nil {
		return err
	}
	fmt.Printf(statefulSet.String())
	statefulSetClient := connect.K8s().Clientset.AppsV1().StatefulSets("modelzoo")
	statefulSetClient.Delete(statefulSet.Name, nil)
	_, err = statefulSetClient.Create(statefulSet)
	if err != nil {
		fmt.Print("StatefulSet error %s", err.Error())
		return err
	}

	serviceClient := connect.K8s().Clientset.CoreV1().Services("modelzoo")
	serviceClient.Delete(service.Name, nil)
	_, err = serviceClient.Create(service)
	if err != nil {
		fmt.Print("service error %s", err.Error())
		return err
	}

	return nil
}

func Release() {
	conf := connect.K8s()
	c := conf.Clientset

	vols, err := c.CoreV1().PersistentVolumeClaims("default").List(metav1.ListOptions{
		//LabelSelector:        "app=test-job-app",
	})

	fmt.Println(err)
	fmt.Println(vols)
	if err != nil {
		fmt.Println(err)
	}

	if vols == nil {
		fmt.Println("pods not found")
	}
	for _, vol := range vols.Items {
		fmt.Println("-----------")
		name := vol.Spec.VolumeName
		fmt.Println("pvc", name)
		pv, _ := c.CoreV1().PersistentVolumes().Get(name, metav1.GetOptions{})
		fmt.Println(pv.Spec.NFS)
		fmt.Println("-----------")
	}
}

func Attach(taskName string, podInx string, session sockjs.Session) error {
	fmt.Println("+++++++")
	fmt.Println(podInx)
	fmt.Println("+++++++")
	conf := connect.K8s()
	c := conf.Clientset
	cfg := conf.Config
	podName := taskName + "-deployment-" + podInx
	// TODO: shall be change after RBAC ready
	req := c.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace("modelzoo").
		SubResource("exec")

	req.VersionedParams(&corev1.PodExecOptions{
		// 需要跟 Containers: []corev1.Container{ 中 保持一致
		// #Line 261
		Container: taskName,
		Command:   []string{"sh"},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(cfg, "POST", req.URL())
	if err != nil {
		return err
	}

	ptyHandler := util.TerminalSession{
		podName,
		make(chan error),
		session,
		make(chan remotecommand.TerminalSize),
		make(chan struct{}),
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:             ptyHandler,
		Stdout:            ptyHandler,
		Stderr:            ptyHandler,
		TerminalSizeQueue: ptyHandler,
		Tty:               true,
	})
	if err != nil {
		return err
	}

	return nil
}

func newStatefulSet(trainTask train_task.Train_taskVo) (*appsv1.StatefulSet, *corev1.Service, error) {
	var (
		statefulset *appsv1.StatefulSet
		service     *corev1.Service
		err         error
	)
	fmt.Println("image=" + "core.harbor.domain/modelzoo/" + trainTask.ImageUrl )
	// TODO: volume list here
	storageClassName := "modelzoo-nfs"
	// TODO
	resourceAlloc := corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.FormatInt(trainTask.CpuNum, 10)),
		corev1.ResourceMemory: resource.MustParse(strconv.FormatInt(trainTask.MemoryNum, 10) + "Gi"),
		ResourceGPU:           resource.MustParse(strconv.FormatInt(trainTask.GpuNum, 10)),
		ResourceRDMA_HCA:      resource.MustParse(strconv.FormatInt(trainTask.Rdma, 10)),
	}

	statefulset = &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			// TODO
			Name:      trainTask.Name + "-deployment",
			Namespace: "modelzoo",
		},
		Spec: appsv1.StatefulSetSpec{
			// TODO: should be point
			Replicas:    &trainTask.PodNum,
			ServiceName: "service-" + trainTask.Name,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": trainTask.Name + "-app"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": trainTask.Name + "-app"},
				},
				Spec: corev1.PodSpec{
					Subdomain: "busybox-subdomain",
					Containers: []corev1.Container{
						{
							Name:       trainTask.Name,
							Image:      "core.harbor.domain/modelzoo/" + trainTask.ImageUrl,
							WorkingDir: "",
							Command:    []string{"sh", "-c"},
							Args:       []string{trainTask.Cmd},
							//Args: []string{"sleep 3600"},
							Resources: corev1.ResourceRequirements{
								Limits:   resourceAlloc,
								Requests: resourceAlloc,
							},
							ImagePullPolicy: corev1.PullIfNotPresent,
							SecurityContext: &corev1.SecurityContext{
								// TODO: should be point
								AllowPrivilegeEscalation: Bool(true),
								Capabilities: &corev1.Capabilities{
									Add: []corev1.Capability{"IPC_LOCK"},
								},
							},

							//Env: []corev1.EnvVar{
							//	{
							//		Name:  "",
							//		Value: "",
							//		// ValueFrom: "",
							//	},
							//},
							VolumeMounts: []corev1.VolumeMount{
								// TTTT
								{Name: "dataset-persistent-storage", MountPath: "/modelzoo/dataset"},
								{Name: "log-persistend-storage", MountPath: "/modelzoo/log"},
								{Name: "model-persistend-storage", MountPath: "/modelzoo/model"},
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyAlways,

					Volumes: []corev1.Volume{
						{Name: "dataset-persistent-storage",
							VolumeSource: corev1.VolumeSource{
								NFS: &corev1.NFSVolumeSource{
									Path:   "/dataset/kubernetes/dataset/" + trainTask.DataSet.Url,
									Server: "192.168.1.16",
								},
							},
						},
					},
				},
			},
			VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "log-persistend-storage"},
					Spec: corev1.PersistentVolumeClaimSpec{
						AccessModes:      []corev1.PersistentVolumeAccessMode{"ReadWriteMany"},
						StorageClassName: &storageClassName,
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceStorage: resource.MustParse("5Gi"),
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "model-persistend-storage"},
					Spec: corev1.PersistentVolumeClaimSpec{
						AccessModes:      []corev1.PersistentVolumeAccessMode{"ReadWriteMany"},
						StorageClassName: &storageClassName,
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceStorage: resource.MustParse("5Gi"),
							},
						},
					},
				},
			},
		},
	}

	service = &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-" + trainTask.Name,
			Labels: map[string]string{
				"app": "service-" + trainTask.Name},
		},
		Spec: corev1.ServiceSpec{
			// TODO
			ClusterIP: "None",
			Selector: map[string]string{
				"app": trainTask.Name + "-app"},
		},
	}
	err = nil
	return statefulset, service, err
}
