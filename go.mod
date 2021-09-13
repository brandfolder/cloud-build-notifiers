module github.com/GoogleCloudPlatform/cloud-build-notifiers

go 1.16

replace github.com/GoogleCloudPlatform/cloud-build-notifiers/lib/notifiers => ./lib/notifiers

require (
	cloud.google.com/go v0.83.0
	cloud.google.com/go/bigquery v1.16.0
	cloud.google.com/go/storage v1.14.0
	github.com/antlr/antlr4 v0.0.0-20210404160547-4dfacf63e228 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.7.3
	github.com/google/go-cmp v0.5.6
	github.com/google/go-containerregistry v0.6.0
	github.com/slack-go/slack v0.8.2
	google.golang.org/api v0.47.0
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/client-go v0.20.6
)
