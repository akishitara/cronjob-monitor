package client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"regexp"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func podsList() *apiv1.PodList {
	podsClient := kubeClient().CoreV1().Pods(apiv1.NamespaceDefault)
	podsList, err := podsClient.List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	return podsList
}

func podListFilterUID(uid string) (pods []apiv1.Pod) {
	allPods := podsList()
	for _, item := range allPods.Items {
		for _, owner := range item.OwnerReferences {
			if string(owner.UID) == uid {
				pods = append(pods, item)
			}
		}
	}
	return pods
}

func podsToString(pods []apiv1.Pod)[]string{
	var res []string
  for _, pod := range pods {
    res = append(res, pod.Name)
	}
	return res
}

func makePod(image string) *apiv1.Pod {
	return &apiv1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod"},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{{Name: "nginx", Image: image}},
		},
	}
}

func podCreate() {
	podsClient := kubeClient().CoreV1().Pods(apiv1.NamespaceDefault)
	podsClient.Create(makePod("nginx"))
}

// PodNameList podの一覧
func PodNameList() []string {
	var res []string
	for _, v := range podsList().Items {
		res = append(res, v.Name)
	}
	return res
}

// PodLogs podのログを取得(500行まで)
func PodLogs(name string) []string {
	var logLineLimit = int64(500)

	podsClient := kubeClient().CoreV1().Pods(apiv1.NamespaceDefault)
	podLogOpts := apiv1.PodLogOptions{TailLines: &logLineLimit}
	podLogs, err := podsClient.GetLogs(name, &podLogOpts).Stream()
	if err != nil {
		log.Println(err)
		return []string{""}
	}
	defer podLogs.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		log.Println(err)
	}
	return stringToArray(buf.String())
	// strings.Split(buf.String(), "\n")
}

func stringToArray(data string) []string {
	var res []string
	for _, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(data, -1) {
		res = append(res, v)
	}
	return res
}
