module github.com/GoogleCloudPlatform/cloud-build-notifiers

go 1.16

replace github.com/GoogleCloudPlatform/cloud-build-notifiers/lib/notifiers => ./lib/notifiers

require (
	cloud.google.com/go v0.100.2
	cloud.google.com/go/bigquery v1.16.0
	cloud.google.com/go/secretmanager v1.4.0
	cloud.google.com/go/storage v1.14.0
	github.com/antlr/antlr4 v0.0.0-20210404160547-4dfacf63e228 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.7.3
	github.com/google/go-cmp v0.5.8
	github.com/google/go-containerregistry v0.9.0
	github.com/slack-go/slack v0.8.2
	google.golang.org/api v0.75.0
	google.golang.org/genproto v0.0.0-20220421151946-72621c1f0bd3
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/client-go v0.20.5
)
