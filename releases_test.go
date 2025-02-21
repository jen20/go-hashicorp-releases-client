package releases_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	releases "github.com/jen20/go-hashicorp-releases-client"
)

func TestClient_LatestRelease(t *testing.T) {
	server := httptest.NewServer(makeTestReleasesHandler(t))

	client, err := releases.New(releases.WithBaseURL(server.URL))
	requireNoError(t, err)

	release, err := client.LatestRelease(context.Background(), "waypoint", releases.LicenseClassOSS)
	requireNoError(t, err)

	requireEqual(t, waypoint_0_11_4, release)
}

func TestClient_Release(t *testing.T) {
	server := httptest.NewServer(makeTestReleasesHandler(t))

	client, err := releases.New(releases.WithBaseURL(server.URL))
	requireNoError(t, err)

	t.Run("0.1.0", func(t *testing.T) {
		release, err := client.Release(context.Background(), "waypoint", "0.1.0")
		requireNoError(t, err)
		requireEqual(t, waypoint_0_1_0, release)
	})

	t.Run("0.11.4", func(t *testing.T) {
		release, err := client.Release(context.Background(), "waypoint", "0.11.4")
		requireNoError(t, err)
		requireEqual(t, waypoint_0_11_4, release)
	})
}

func TestClient_Releases(t *testing.T) {
	server := httptest.NewServer(makeTestReleasesHandler(t))

	client, err := releases.New(releases.WithBaseURL(server.URL))
	requireNoError(t, err)

	releasesIterator, err := client.Releases(context.Background(), "waypoint", releases.LicenseClassOSS)
	requireNoError(t, err)

	items := collectResults(t, releasesIterator)
	requireEqual(t, 43, len(items))
	requireEqual(t, waypoint_0_11_4, items[0])
	requireEqual(t, waypoint_0_1_0, items[42])
}

func TestClient_ReleasesPaged(t *testing.T) {
	server := httptest.NewServer(makeTestReleasesHandler(t))

	client, err := releases.New(releases.WithBaseURL(server.URL))
	requireNoError(t, err)

	releasePagesIterator, err := client.ReleasesPaged(context.Background(), "waypoint", releases.LicenseClassOSS)
	requireNoError(t, err)

	pages := collectResults(t, releasePagesIterator)
	requireEqual(t, 3, len(pages))
	requireEqual(t, waypoint_0_11_4, pages[0][0])
	requireEqual(t, waypoint_0_1_0, pages[2][10])
}

func makeTestReleasesHandler(t *testing.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mustWriteFile := func(path string) {
			w.WriteHeader(http.StatusOK)
			data, err := os.ReadFile(path)
			requireNoError(t, err)
			n, err := w.Write(data)
			requireNoError(t, err)
			requireEqual(t, len(data), n)
		}

		mustWriteJSON := func(val any) {
			w.WriteHeader(http.StatusOK)
			enc := json.NewEncoder(w)
			requireNoError(t, enc.Encode(val))
		}

		switch r.URL.Path {
		case "/v1/releases/waypoint":
			if r.URL.Path != "/v1/releases/waypoint" {
				w.WriteHeader(http.StatusNotFound)
				t.Errorf("unexpected path accessed: %s", r.URL.Path)
			}

			if limit := r.URL.Query().Get("limit"); limit != "16" {
				w.WriteHeader(http.StatusBadRequest)
				t.Errorf("unexpected limit set: %s", limit)
			}

			after := r.URL.Query().Get("after")
			switch after {
			case "":
				mustWriteFile("testdata/releases/page1.json")
			case "2022-04-07T16:15:06Z":
				mustWriteFile("testdata/releases/page2.json")
			case "2021-04-08T18:56:58Z":
				mustWriteFile("testdata/releases/page3.json")
			case "2020-10-15T16:37:48Z":
				mustWriteFile("testdata/releases/page4.json")
			default:
				w.WriteHeader(http.StatusBadRequest)
				t.Errorf("unexpected after value: %s", after)
			}
		case "/v1/releases/waypoint/0.1.0":
			mustWriteJSON(waypoint_0_1_0)
		case "/v1/releases/waypoint/0.11.4":
			mustWriteJSON(waypoint_0_11_4)
		case "/v1/releases/waypoint/latest":
			mustWriteJSON(waypoint_0_11_4)
		default:
			t.Fatalf("unexpected HTTP URL Path %q", r.URL.Path)
		}

	})
}

var waypoint_0_11_4 = releases.ReleaseInfo{Builds: []releases.BuildInfo{
	{
		Arch:        "amd64",
		OS:          "darwin",
		Unsupported: false,
		URL:         "https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_darwin_amd64.zip",
	},
	{
		Arch:        "arm64",
		OS:          "darwin",
		Unsupported: false,
		URL:         "https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_darwin_arm64.zip",
	},
	{
		Arch:        "386",
		OS:          "linux",
		Unsupported: false,
		URL:         "https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_linux_386.zip",
	},
	{
		Arch:        "amd64",
		OS:          "linux",
		Unsupported: false,
		URL:         "https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_linux_amd64.zip",
	},
	{
		Arch:        "arm",
		OS:          "linux",
		Unsupported: false,
		URL:         "https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_linux_arm.zip",
	},
	{
		Arch:        "arm64",
		OS:          "linux",
		Unsupported: false,
		URL:         "https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_linux_arm64.zip",
	},
	{
		Arch:        "386",
		OS:          "windows",
		Unsupported: false,
		URL:         "https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_windows_386.zip",
	},
	{
		Arch:        "amd64",
		OS:          "windows",
		Unsupported: false,
		URL:         "https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_windows_amd64.zip",
	},
},
	DockerNameTag: "",
	IsPrerelease:  false,
	LicenseClass:  "oss",
	Name:          "waypoint",
	Status: releases.ReleaseStatus{
		Message: "",
		State:   "supported",
	},
	TimestampCreated:           time.Date(2023, time.August, 9, 18, 33, 15, 901000000, time.UTC),
	TimestampUpdated:           time.Date(2023, time.August, 9, 18, 33, 15, 901000000, time.UTC),
	URLBlogpost:                "",
	URLChangelog:               "https://github.com/hashicorp/waypoint/blob/release/0.11.x/CHANGELOG.md",
	URLDockerRegistryDockerhub: "https://hub.docker.com/r/hashicorp/waypoint",
	URLDockerRegistryECR:       "https://gallery.ecr.aws/hashicorp/waypoint",
	URLLicense:                 "https://github.com/hashicorp/waypoint/blob/main/LICENSE",
	URLProjectWebsite:          "https://www.waypointproject.io/",
	URLReleaseNotes:            "",
	URLSHASUMs:                 "https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_SHA256SUMS",
	URLSHASUMsSignatures: []string{
		"https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_SHA256SUMS.sig",
		"https://releases.hashicorp.com/waypoint/0.11.4/waypoint_0.11.4_SHA256SUMS.72D7468F.sig",
	},
	URLSourceRepository: "https://github.com/hashicorp/waypoint",
	Version:             "0.11.4",
}

var waypoint_0_1_0 = releases.ReleaseInfo{
	Builds: []releases.BuildInfo{
		{
			Arch:        "amd64",
			OS:          "darwin",
			Unsupported: false,
			URL:         "https://releases.hashicorp.com/waypoint/0.1.0/waypoint_0.1.0_darwin_amd64.zip",
		},
		{
			Arch:        "386",
			OS:          "linux",
			Unsupported: false,
			URL:         "https://releases.hashicorp.com/waypoint/0.1.0/waypoint_0.1.0_linux_386.zip",
		},
		{
			Arch:        "amd64",
			OS:          "linux",
			Unsupported: false,
			URL:         "https://releases.hashicorp.com/waypoint/0.1.0/waypoint_0.1.0_linux_amd64.zip",
		},
		{
			Arch:        "arm",
			OS:          "linux",
			Unsupported: false,
			URL:         "https://releases.hashicorp.com/waypoint/0.1.0/waypoint_0.1.0_linux_arm.zip",
		},
		{
			Arch:        "386",
			OS:          "windows",
			Unsupported: false,
			URL:         "https://releases.hashicorp.com/waypoint/0.1.0/waypoint_0.1.0_windows_386.zip",
		},
		{
			Arch:        "amd64",
			OS:          "windows",
			Unsupported: false,
			URL:         "https://releases.hashicorp.com/waypoint/0.1.0/waypoint_0.1.0_windows_amd64.zip",
		},
	},
	DockerNameTag: "",
	IsPrerelease:  false,
	LicenseClass:  "oss",
	Name:          "waypoint",
	Status: releases.ReleaseStatus{
		Message: "",
		State:   "supported",
	},
	TimestampCreated:           time.Date(2020, time.October, 15, 16, 37, 48, 0, time.UTC),
	TimestampUpdated:           time.Date(2020, time.October, 15, 16, 37, 48, 0, time.UTC),
	URLBlogpost:                "",
	URLChangelog:               "",
	URLDockerRegistryDockerhub: "",
	URLDockerRegistryECR:       "",
	URLLicense:                 "",
	URLProjectWebsite:          "",
	URLReleaseNotes:            "",
	URLSHASUMs:                 "https://releases.hashicorp.com/waypoint/0.1.0/waypoint_0.1.0_SHA256SUMS",
	URLSHASUMsSignatures: []string{
		"https://releases.hashicorp.com/waypoint/0.1.0/waypoint_0.1.0_SHA256SUMS.sig",
		"https://releases.hashicorp.com/waypoint/0.1.0/waypoint_0.1.0_SHA256SUMS.348FFC4C.sig",
		"https://releases.hashicorp.com/waypoint/0.1.0/waypoint_0.1.0_SHA256SUMS.72D7468F.sig",
	},
	URLSourceRepository: "https://github.com/hashicorp/waypoint",
	Version:             "0.1.0",
}
