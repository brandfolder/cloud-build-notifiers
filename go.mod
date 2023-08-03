module github.com/GoogleCloudPlatform/cloud-build-notifiers

go 1.16

replace github.com/GoogleCloudPlatform/cloud-build-notifiers/lib/notifiers => ./lib/notifiers

require (
	cloud.google.com/go v0.110.0
	cloud.google.com/go/bigquery v1.50.0
	cloud.google.com/go/secretmanager v1.10.0
	cloud.google.com/go/storage v1.29.0
	github.com/antlr/antlr4 v0.0.0-20210404160547-4dfacf63e228 // indirect
	github.com/golang/glog v1.1.0
	github.com/golang/protobuf v1.5.3
	github.com/google/cel-go v0.7.3
	github.com/google/go-cmp v0.5.9
	github.com/google/go-containerregistry v0.16.1
	github.com/slack-go/slack v0.8.2
	google.golang.org/api v0.122.0
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/protobuf v1.30.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/client-go v0.20.5
)
