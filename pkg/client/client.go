package client

//func init() {
//	prometheus.MustRegister(JobStatusMetrics)
//}

// Run Debug用
func Run() {
//	Cronjob := ActiveCronjobsAllGet()
//	debugger.YamlPrint(Cronjob)
//
//	for i, _ := range Cronjob {
//		debugger.YamlPrint(Cronjob[i].lastJobs())
//		for j, _ := range Cronjob[i].Jobs {
//			debugger.YamlPrint(Cronjob[i].Jobs[j].Status[0].Status)
//		}
//	}
//
//	go func() {
//		for {
//			JobStatusMetrics.With(prometheus.Labels{"cronjob": "crontest1"}).Set(1)
//			time.Sleep(30 * time.Second)
//		}
//	}()
//	http.Handle("/metrics", promhttp.Handler())
//	log.Fatal(http.ListenAndServe(":8080", nil))
}

