package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

// VERSION is the version number of github-release
const VERSION = "0.1.0"

// newClient creates a GitHub client
func newClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return github.NewClient(tc)
}

func main() {
	version := flag.String("version", "1.0.0", "version number for the release")
	owner := flag.String("owner", "hackebrot", "github project owner")
	repo := flag.String("repo", "github-release", "github project repo")
	draft := flag.Bool("draft", false, "true to create a draft (unpublished) release")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf("github-release - v%v\n", VERSION))
		flag.PrintDefaults()
	}
	flag.Parse()

	token := os.Getenv("GITHUB_RELEASE_TOKEN")
	if token == "" {
		fmt.Fprintf(os.Stderr, "environment variable GITHUB_RELEASE_TOKEN is required")
		os.Exit(1)
	}
	client := newClient(token)

	releaseBody := fmt.Sprintf("Hello world - %v", *version)

	release, _, err := client.Repositories.CreateRelease(
		*owner,
		*repo,
		&github.RepositoryRelease{
			TagName: version,
			Draft:   draft,
			Name:    version,
			Body:    &releaseBody,
		},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create release%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("release %+v\n", release)
}
