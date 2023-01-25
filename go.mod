module github.com/GoogleCloudPlatform/cloud-build-notifiers

go 1.16

replace github.com/GoogleCloudPlatform/cloud-build-notifiers/lib/notifiers => ./lib/notifiers

require (
	cloud.google.com/go v0.105.0
	cloud.google.com/go/bigquery v1.43.0
	cloud.google.com/go/secretmanager v1.10.0
	cloud.google.com/go/storage v1.27.0
	github.com/antlr/antlr4 v0.0.0-20210404160547-4dfacf63e228 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.7.3
	github.com/google/go-cmp v0.5.9
	github.com/google/go-containerregistry v0.13.0
	github.com/kr/text v0.2.0 // indirect
	github.com/slack-go/slack v0.8.2
	google.golang.org/api v0.103.0
	google.golang.org/genproto v0.0.0-20221201164419-0e50fba7f41c
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/client-go v0.20.5
)
