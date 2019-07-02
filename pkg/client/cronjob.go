package client

import (
	batchv1 "k8s.io/api/batch/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// cronjob全取得
func cronjobList() (*batchv1.CronJobList, error) {
	cronjobClient := kubeClient().BatchV1beta1().CronJobs(apiv1.NamespaceDefault)
	return cronjobClient.List(metav1.ListOptions{})
}

// 名前からcronjobを取得
func cronjobGet(name string) *batchv1.CronJob {
	cronjobClient := kubeClient().BatchV1beta1().CronJobs(apiv1.NamespaceDefault)
	res, _ := cronjobClient.Get(name, metav1.GetOptions{})
	return res
}

// cronjob(公式)から内部typeに変換
func cronjobToActiveCronjob(cronjob batchv1.CronJob) ActiveCronjob {
	var res ActiveCronjob
	res = ActiveCronjob{
		Name:     cronjob.Name,
		UID:      string(cronjob.UID),
		Schedule: cronjob.Spec.Schedule,
		Image:    cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image,
		Jobs:     jobsToActivejobs(jobListFilterCronjobName(cronjob.Name)),
	}
	return res
}

// ActiveCronjobsAllGet 登録されてる全部のcronjobを内部Typeで取得
func ActiveCronjobsAllGet() []ActiveCronjob {
	var res []ActiveCronjob
	cronjobs, _ := cronjobList()
	for _, a := range cronjobs.Items {
		res = append(res, cronjobToActiveCronjob(a))
	}
	return res
}
