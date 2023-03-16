module github.com/GoogleCloudPlatform/cloud-build-notifiers

go 1.16

replace github.com/GoogleCloudPlatform/cloud-build-notifiers/lib/notifiers => ./lib/notifiers

require (
	cloud.google.com/go v0.107.0
	cloud.google.com/go/bigquery v1.44.0
	cloud.google.com/go/secretmanager v1.9.0
	cloud.google.com/go/storage v1.27.0
	github.com/antlr/antlr4 v0.0.0-20210404160547-4dfacf63e228 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3
	github.com/google/cel-go v0.7.3
	github.com/google/go-cmp v0.5.9
	github.com/google/go-containerregistry v0.14.0
	github.com/kr/text v0.2.0 // indirect
	github.com/slack-go/slack v0.8.2
	google.golang.org/api v0.108.0
	google.golang.org/genproto v0.0.0-20230124163310-31e0e69b6fc2
	google.golang.org/protobuf v1.29.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/client-go v0.20.5
)
