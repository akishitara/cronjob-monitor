module github.com/akishitara/cronjob-monitor

go 1.12

require (
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.4.0
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/prometheus/client_golang v1.0.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20190920012459-5008bf6f8cd6 // indirect
	sigs.k8s.io/kind v0.5.1 // indirect
)
