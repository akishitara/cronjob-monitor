package client

import (
	batchv1 "k8s.io/api/batch/v1"
)

// ActiveCronjob Cronjobの設定情報とjob情報を返す
type ActiveCronjob struct {
	// Cronjob Name
	Name string `json:"name"`
	// Cronjob UID
	UID string
	// Cronjob Schedule
	Schedule string
	// ContainerImage
	// 複数のcontainerは想定してないです。一番最初のImageのみ表示
	Image string
	// 実行されるJobのList
	Jobs []activejob
}

type activejob struct {
	Name      string `json:"name"`
	StartTime string //*metav1.Time
	EndTime   string //*metav1.Time
	Status    []batchv1.JobCondition
	PodsNames []string
}
