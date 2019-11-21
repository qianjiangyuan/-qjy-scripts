package k8s

import (
	"fmt"
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/train"
	batch "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

func ExecJob(trainTask train_task.Train_taskVo) error {
	conf := connect.Conf
	nfsConf := conf.GetStringMapString("nfs")
	nfsHost := nfsConf["host"]
	nfsPath := nfsConf["path"]
        fmt.Printf("nfs server %s",nfsPath + trainTask.DataSet.Url)
	fmt.Printf("nfs host %s",nfsHost)

	// TODO
	resourceAlloc := corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.FormatInt(trainTask.CpuNum, 10)),
		corev1.ResourceMemory: resource.MustParse(strconv.FormatInt(trainTask.MemoryNum, 10) + "Gi"),
		ResourceGPU:           resource.MustParse(strconv.FormatInt(trainTask.GpuNum, 10)),
	//	ResourceRDMA_HCA:      resource.MustParse(strconv.FormatInt(trainTask.Rdma, 10)),
	}
	//strorageClassName := "modelzoo-nfs"
	pvcName := trainTask.Name + "-pvc"
        fmt.Printf("get pvc %s",pvcName)

	pvcClient := connect.K8s().Clientset.CoreV1().PersistentVolumeClaims("modelzoo")
	pvc, err := pvcClient.Get(pvcName, metav1.GetOptions{})
//	if (err != nil ||pvc. {
                fmt.Printf("create pvc")
		pvc = &corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				// TODO
				Name:      pvcName,
				Namespace: "modelzoo",
			},
			Spec: corev1.PersistentVolumeClaimSpec{
				AccessModes: []corev1.PersistentVolumeAccessMode{
					corev1.ReadWriteOnce,
				},
				//StorageClassName: &strorageClassName,
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: resource.MustParse(strconv.FormatInt(1, 10) + "Gi"),
					},
				},
			},
		}
		_, err = pvcClient.Create(pvc)
		//if err != nil {
		//	return err
		//}
                fmt.Printf("create pvc2")
//	}
        fmt.Printf("pvc= %s",pvc.ObjectMeta.Name)
	job := &batch.Job{
		ObjectMeta: metav1.ObjectMeta{
			// TODO
			Name:      trainTask.Name + "-job",
			Namespace: "modelzoo",
		},
		Spec: batch.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": trainTask.Name + "-app"},
				},
				Spec: corev1.PodSpec{
					Subdomain: "busybox-subdomain",
					Containers: []corev1.Container{
						{
							Name:       trainTask.Name,
							Image:      trainTask.ImageUrl,
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
							VolumeMounts: []corev1.VolumeMount{
								{Name: "dataset-persistent-storage", MountPath: "/modelzoo/dataset"},
								// {Name: "code-persistend-storage", MountPath: "/modelzoo/code"},
								//{Name: "log-persistend-storage", MountPath: "/modelzoo/log"},
								//{Name: "model-persistend-storage", MountPath: "/modelzoo/model"},
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
					Volumes: []corev1.Volume{
						{Name: "dataset-persistent-storage",
							VolumeSource: corev1.VolumeSource{
								NFS: &corev1.NFSVolumeSource{
									Path:   nfsPath + trainTask.DataSet.Url,
									Server: nfsHost,
								},
							},
						},
						/*
						   {Name: "code-persistend-storage",
						     VolumeSource: corev1.VolumeSource{
						       NFS: &corev1.NFSVolumeSource{
						         Path:   "/dataset/kubernetes/dataset/" + trainTask.CodePath,
						         Server: "192.168.1.16",
						       },
						     },
						   },

						   {Name: "log-persistend-storage",
						     VolumeSource: corev1.VolumeSource{
						       PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
						         ClaimName: "modelzoodir",
						         ReadOnly:  false,
						       },
						     },
						   },
						{Name: "model-persistend-storage",
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: trainTask.Name + "-pvc",
									ReadOnly:  false,
								},
							},
						},*/
					},
				},
			},
		},
	}

	jobClient := connect.K8s().Clientset.BatchV1().Jobs("modelzoo")
	jobClient.Delete(job.Name, nil)
	_, err = jobClient.Create(job)
	if err != nil {
		fmt.Print("job error %s", err.Error())
		return err
	}

	return nil
}

func DetailJob(name string) (string, error) {
	jobClient := connect.K8s().Clientset.BatchV1().Jobs("modelzoo")
	job, err := jobClient.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", nil
	}
	if job.Status.Succeeded > 0 {
		return "success", nil
	}
	if job.Status.Failed > 0 {
		return "failed", nil
	}
	if job.Status.Active > 0 {
		return "active", nil
	}
	return "", nil
}

func GetPV(name string) (string, error) {
	pvcClient := connect.K8s().Clientset.CoreV1().PersistentVolumeClaims("modelzoo")
	pvc, err := pvcClient.Get(name+"-pvc", metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	fmt.Println(pvc)
	pvClient := connect.K8s().Clientset.CoreV1().PersistentVolumes()
	pv, err := pvClient.Get(pvc.Spec.VolumeName, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	fmt.Println(pv.Spec.PersistentVolumeSource.NFS.Path)
	return pv.Spec.PersistentVolumeSource.NFS.Path, nil

}
func StopJob(name string) error {
	jobClient := connect.K8s().Clientset.BatchV1().Jobs("modelzoo")
	options := metav1.DeleteOptions{}
	var seconds int64 = 100
	options.GracePeriodSeconds = &seconds
	err := jobClient.Delete(name, &options)
	return err
}
