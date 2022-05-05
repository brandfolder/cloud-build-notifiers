package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/slack-go/slack"
	cbpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)

func TestGetStoragePath(t *testing.T) {
	result := getStoragePath("build-id")
	expected := "messages/build-id"
	if result != expected {
		t.Errorf("Unexpected storage path: %q, expected: %q", result, expected)
	}
	return
}

func TestSingleBuildAttachmentMessageOption(t *testing.T) {
	b := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id",
		Status:    cbpb.Build_SUCCESS,
		Substitutions: map[string]string{
			repoNameSub:       "test-repo",
			branchNameSub:     "test-branch",
			commitShortShaSub: "1234fakesha",
			commitMsgSub:      "fake message",
			commitURLSub:      "https://some.example.com/fakecommit",
			commitAuthorSub:   "jdoe",
		},
		LogUrl: "https://some.example.com/log/url?foo=bar",
	}

	sb := storedBuild{
		Build: map[string]*cbpb.Build{
			b.Id: b,
		},
	}

	got := buildAttachmentMessageOption(sb)

	want := slack.MsgOptionAttachments(
		slack.Attachment{
			Text:  "SUCCESS: :test-repo: test-repo (my-project-id) \u003chttps://some.example.com/log/url?foo=bar\u0026utm_campaign=google-cloud-build-notifiers\u0026utm_medium=chat\u0026utm_source=google-cloud-build|View Build\u003e\n*Branch*: test-branch *Author*: jdoe \n\u003chttps://some.example.com/fakecommit|Commit\u003e *1234fakesha*: fake message",
			Color: "good",
		},
	)

	_, gotValues, err := slack.UnsafeApplyMsgOptions("fake-token", "fake-channel", "https://fake.com/", *got)
	if err != nil {
		t.Errorf("Unable to build message: %s", err.Error())
	}
	_, wantValues, err := slack.UnsafeApplyMsgOptions("fake-token", "fake-channel", "https://fake.com/", want)
	if diff := cmp.Diff(gotValues, wantValues); diff != "" {
		t.Logf("full message: %+v", gotValues)
		t.Errorf("writeMessage got unexpected diff: %s", diff)
	}
	return
}

func TestMultiRegionSuccessBuildAttachmentMessageOption(t *testing.T) {
	b1 := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id1",
		Status:    cbpb.Build_SUCCESS,
		Substitutions: map[string]string{
			repoNameSub:       "test-repo",
			branchNameSub:     "test-branch",
			commitShortShaSub: "1234fakesha1",
			commitMsgSub:      "fake message1",
			commitURLSub:      "https://some.example.com/fakecommit1",
			commitAuthorSub:   "jdoe1",
		},
		LogUrl: "https://some.example.com/log/url?foo=bar1",
	}

	b2 := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id2",
		Status:    cbpb.Build_SUCCESS,
		Substitutions: map[string]string{
			repoNameSub:       "test-repo",
			branchNameSub:     "test-branch",
			commitShortShaSub: "1234fakesha2",
			commitMsgSub:      "fake message2",
			commitURLSub:      "https://some.example.com/fakecommit2",
			commitAuthorSub:   "jdoe2",
		},
		LogUrl: "https://some.example.com/log/url?foo=bar2",
	}

	sb := storedBuild{
		Build: map[string]*cbpb.Build{
			b1.Id: b1,
			b2.Id: b2,
		},
	}

	got := buildAttachmentMessageOption(sb)

	want1 := slack.MsgOptionAttachments(
		slack.Attachment{
			Text:  "SUCCESS: :test-repo: test-repo (my-project-id) \u003chttps://some.example.com/log/url?foo=bar1\u0026utm_campaign=google-cloud-build-notifiers\u0026utm_medium=chat\u0026utm_source=google-cloud-build|View Build\u003e\n*Branch*: test-branch *Author*: jdoe1 \n\u003chttps://some.example.com/fakecommit1|Commit\u003e *1234fakesha1*: fake message1",
			Color: "good",
		},
	)

	want2 := slack.MsgOptionAttachments(
		slack.Attachment{
			Text:  "SUCCESS: :test-repo: test-repo (my-project-id) \u003chttps://some.example.com/log/url?foo=bar2\u0026utm_campaign=google-cloud-build-notifiers\u0026utm_medium=chat\u0026utm_source=google-cloud-build|View Build\u003e\n*Branch*: test-branch *Author*: jdoe2 \n\u003chttps://some.example.com/fakecommit2|Commit\u003e *1234fakesha2*: fake message2",
			Color: "good",
		},
	)

	_, gotValues, err := slack.UnsafeApplyMsgOptions("fake-token", "fake-channel", "https://fake.com/", *got)
	if err != nil {
		t.Errorf("Unable to build message: %s", err.Error())
	}
	_, wantValues1, err := slack.UnsafeApplyMsgOptions("fake-token", "fake-channel", "https://fake.com/", want1)
	diff1 := cmp.Diff(gotValues, wantValues1)
	_, wantValues2, err := slack.UnsafeApplyMsgOptions("fake-token", "fake-channel", "https://fake.com/", want2)
	diff2 := cmp.Diff(gotValues, wantValues2)
	if diff1 != "" && diff2 != "" {
		t.Logf("full message: %+v", gotValues)
		if diff1 != "" {
			t.Errorf("writeMessage got unexpected diff1: %s", diff1)
			return
		}
		t.Errorf("writeMessage got unexpected diff2: %s", diff2)
	}
	return
}

func TestMultiRegionFailBuildAttachmentMessageOption(t *testing.T) {
	b1 := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id1",
		Status:    cbpb.Build_FAILURE,
		Substitutions: map[string]string{
			repoNameSub:       "test-repo",
			branchNameSub:     "test-branch",
			commitShortShaSub: "1234failsha",
			commitMsgSub:      "fail message",
			commitURLSub:      "https://some.example.com/failcommit",
			commitAuthorSub:   "failperson",
		},
		LogUrl: "https://some.example.com/log/url?foo=fail",
	}

	b2 := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id2",
		Status:    cbpb.Build_SUCCESS,
		Substitutions: map[string]string{
			repoNameSub:       "test-repo",
			branchNameSub:     "test-branch",
			commitShortShaSub: "1234successsha",
			commitMsgSub:      "success message",
			commitURLSub:      "https://some.example.com/successcommit",
			commitAuthorSub:   "successperson",
		},
		LogUrl: "https://some.example.com/log/url?foo=success",
	}

	sb := storedBuild{
		Build: map[string]*cbpb.Build{
			b1.Id: b1,
			b2.Id: b2,
		},
	}

	got := buildAttachmentMessageOption(sb)

	want := slack.MsgOptionAttachments(
		slack.Attachment{
			Text:  "FAILURE: :test-repo: test-repo (my-project-id) \u003chttps://some.example.com/log/url?foo=fail\u0026utm_campaign=google-cloud-build-notifiers\u0026utm_medium=chat\u0026utm_source=google-cloud-build|View Build\u003e\n*Branch*: test-branch *Author*: failperson \n\u003chttps://some.example.com/failcommit|Commit\u003e *1234failsha*: fail message",
			Color: "danger",
		},
	)

	_, gotValues, err := slack.UnsafeApplyMsgOptions("fake-token", "fake-channel", "https://fake.com/", *got)
	if err != nil {
		t.Errorf("Unable to build message: %s", err.Error())
	}
	_, wantValues, err := slack.UnsafeApplyMsgOptions("fake-token", "fake-channel", "https://fake.com/", want)
	if diff := cmp.Diff(gotValues, wantValues); diff != "" {
		t.Logf("full message: %+v", gotValues)
		t.Errorf("writeMessage got unexpected diff: %s", diff)
	}
	return
}

func TestShouldDeleteBuildFileTrue(t *testing.T) {
	b1 := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id1",
		Status:    cbpb.Build_FAILURE,
	}

	b2 := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id2",
		Status:    cbpb.Build_SUCCESS,
	}

	sb := storedBuild{
		Build: map[string]*cbpb.Build{
			b1.Id: b1,
			b2.Id: b2,
		},
	}

	if shouldDeleteBuildFile(sb) == false {
		t.Errorf("no waiting message, it should delete: %+v", sb)
	}
}

func TestShouldDeleteBuildFileFalse(t *testing.T) {
	b1 := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id1",
		Status:    cbpb.Build_WORKING,
	}

	b2 := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id2",
		Status:    cbpb.Build_SUCCESS,
	}

	sb := storedBuild{
		Build: map[string]*cbpb.Build{
			b1.Id: b1,
			b2.Id: b2,
		},
	}

	if shouldDeleteBuildFile(sb) == true {
		t.Errorf("still working, it should not be deleted: %+v", sb)
	}
}
