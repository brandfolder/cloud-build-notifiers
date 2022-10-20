module github.com/GoogleCloudPlatform/cloud-build-notifiers

go 1.16

replace github.com/GoogleCloudPlatform/cloud-build-notifiers/lib/notifiers => ./lib/notifiers

require (
	cloud.google.com/go v0.104.0
	cloud.google.com/go/bigquery v1.16.0
	cloud.google.com/go/secretmanager v1.7.0
	cloud.google.com/go/storage v1.23.0
	github.com/antlr/antlr4 v0.0.0-20210404160547-4dfacf63e228 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.7.3
	github.com/google/go-cmp v0.5.9
	github.com/google/go-containerregistry v0.12.0
	github.com/kr/text v0.2.0 // indirect
	github.com/slack-go/slack v0.8.2
	google.golang.org/api v0.96.0
	google.golang.org/genproto v0.0.0-20220920201722-2b89144ce006
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/client-go v0.20.5
)
