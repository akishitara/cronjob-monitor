package client

import (
	"errors"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func jobList() (*batchv1.JobList, error) {
	jobClient := kubeClient().BatchV1().Jobs(apiv1.NamespaceDefault)
	return jobClient.List(metav1.ListOptions{})
}

func jobListFilterUID(uid string) (jobs []batchv1.Job) {
	alljobs, _ := jobList()
	for _, item := range alljobs.Items {
		for _, owner := range item.OwnerReferences {
			if string(owner.UID) == uid {
				jobs = append(jobs, item)
			}
		}
	}
	return jobs
}

func jobListFilterCronjobName(name string) (jobs []batchv1.Job) {
	cronjobUID := cronjobGet(name).UID
	return jobListFilterUID(string(cronjobUID))
}

func jobToActivejob(job batchv1.Job) (activejob, error) {
	layout := "2006-01-02T15:04:05Z"
	if job.Status.CompletionTime != nil {
		return activejob{
			Name:      job.Name,
			StartTime: job.Status.StartTime.Format(layout),
			EndTime:   job.Status.CompletionTime.Format(layout),
			Status:    job.Status.Conditions,
			PodsNames: podsToString(podListFilterUID(string(job.UID))),
		}, nil
	}
	return activejob{}, errors.New("not complete")
}

func jobsToActivejobs(jobs []batchv1.Job) (res []activejob) {
	for _, job := range jobs {
		jobStatus, err := jobToActivejob(job)
		if err == nil {
			res = append(res, jobStatus)
		}
	}
	return res
}
