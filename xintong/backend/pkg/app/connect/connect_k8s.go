package connect

import (
	"fmt"
	"github.com/spf13/viper"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"strconv"
)

type K8sConnector struct {
	Clientset *kubernetes.Clientset
	Config    *rest.Config
}

var k8sConn *K8sConnector

func K8sConnect(conf *viper.Viper) *K8sConnector {
	kubeconfigPath := conf.GetString("kubectl.config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)

	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err)
	}

	fmt.Println("kubernetes connect ok ...")

	k8sConn = &K8sConnector{
		clientset,
		config,
	}

	return k8sConn
}

// TODO: tmp
func Bool(b bool) *bool {
	return &b
}

func Int32(i int32) *int32 {
	return &i
}

func (k *K8sConnector) NewStatefulSet(param *interface{}) (*appsv1.StatefulSet, *corev1.Service, error) {
	var (
		statefulset *appsv1.StatefulSet
		service     *corev1.Service
		err         error
	)

	// TODO: volume list here

	// TODO
	resourceAlloc := corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.FormatInt(10, 10)),
		corev1.ResourceMemory: resource.MustParse(strconv.FormatInt(10, 10) + "Mi"),
		ResourceGPU:           resource.MustParse(strconv.FormatInt(10, 10)),
		ResourceRDMA_HCA:      resource.MustParse(strconv.FormatInt(10, 10)),
	}

	statefulset = &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			// TODO
			Name:      "",
			Namespace: "",
		},
		Spec: appsv1.StatefulSetSpec{
			// TODO: should be point
			Replicas:    Int32(1),
			ServiceName: "",
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"k1": "v1",
					"k2": "v2",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:       "",
							Image:      "",
							WorkingDir: "",
							Command:    []string{"", ""},
							Args:       []string{"", ""},
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
							Ports: []corev1.ContainerPort{
								{
									Name:          "",
									ContainerPort: 2222,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "",
									Value: "",
									// ValueFrom: "",
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								// TODO
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyAlways,
					Volumes:       []corev1.Volume{
						// TODO
					},
				},
			},
		},
	}

	service = &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			// TODO
		},
		Spec: corev1.ServiceSpec{
			// TODO
		},
	}

	err = nil
	return statefulset, service, err
}

/**

func (k *K8sConnector) Deploy(statefulset *appsv1.StatefulSet, service *corev1.Service) error {
	c := k.clientset

	var err error

	// TODO: service statefulset catch
	_, err = c.CoreV1().Services("").Create(service)
	_, err = c.AppsV1().StatefulSets("").Create(statefulset)

	return err
}

func (k *K8sConnector) Attach(id string, session sockjs.Session) error {
	c := k.clientset
	cfg := k.config

	req := c.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(id).
		Namespace("").
		SubResource("exec")

	req.VersionedParams(&corev1.PodExecOptions{
		Container: "",
		Command:   []string{"bash"},
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
		id,
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

func (k *K8sConnector) Log(id string) (string, error) {
	c := k.clientset

	opts := &corev1.PodLogOptions{}
	req := c.CoreV1().Pods("").GetLogs(
		fmt.Sprintf("format something, %s", id),
		opts,
	)

	readCloser, err := req.Stream()

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	defer readCloser.Close()

	r := bufio.NewReader(readCloser)

	var log string

	for {
		str, err := r.ReadString('\n')
		log += str

		if err != nil {
			break
		}

	}

	return log, err
}

func (k *K8sConnector) Status(id string) (int32, int32, error) {
	c := k.clientset

	var err error

	rs, err := c.AppsV1().StatefulSets("").Get(
		fmt.Sprintf("format something, %s", id),
		metav1.GetOptions{},
	)

	if err != nil {
		fmt.Println(err.Error())
		return 0, 0, err
	}

	if rs != nil {
		return rs.Status.ReadyReplicas, rs.Status.Replicas, err
	}

	return 0, 0, errors.New("xxxx")
}

func (k *K8sConnector) Del(id string) error {
	c := k.clientset

	var err error

	err = c.AppsV1().StatefulSets("").Delete(
		fmt.Sprintf("format something, %s", id),
		&metav1.DeleteOptions{},
	)

	err = c.CoreV1().Services("").Delete(
		fmt.Sprintf("format something, %s", id),
		&metav1.DeleteOptions{},
	)

	return err
}

*/
