package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"regexp"
)

func GetIgnoreUrls(client *http.Client, askedIgnores []string) ([]string, error) {
	ghClient := github.NewClient(client)
	_, dir, _, err := ghClient.Repositories.GetContents(context.Background(), "github", "gitignore", "", &github.RepositoryContentGetOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get github/gitignore content")
	}

	_, globalDir, _, err := ghClient.Repositories.GetContents(context.Background(), "github", "gitignore", "Global", &github.RepositoryContentGetOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get github/gitignore global content")
	}

	return intersect(askedIgnores, append(dir, globalDir...)), nil
}

func intersect(wants []string, files []*github.RepositoryContent) []string {
	urls := make([]string, 0, len(wants))

	for _, w := range wants {
		wantRegex := regexp.MustCompile(fmt.Sprintf(`(?i)%s.gitignore`, regexp.QuoteMeta(w)))
		found := false
		for _, f := range files {
			if wantRegex.MatchString(f.GetName()) {
				urls = append(urls, f.GetDownloadURL())
				found = true
				break
			}
		}

		if !found {
			fmt.Fprintln(os.Stderr, "Unable to generate ignores for", w)
		}
	}

	return urls
}
