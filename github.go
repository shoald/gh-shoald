package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

const org = "shoald"

type generateRequest struct {
	Owner   string `json:"owner"`
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

type repoSettings struct {
	AllowSquashMerge         bool   `json:"allow_squash_merge"`
	AllowMergeCommit         bool   `json:"allow_merge_commit"`
	AllowRebaseMerge         bool   `json:"allow_rebase_merge"`
	DeleteBranchOnMerge      bool   `json:"delete_branch_on_merge"`
	SquashMergeCommitTitle   string `json:"squash_merge_commit_title"`
	SquashMergeCommitMessage string `json:"squash_merge_commit_message"`
}

func createFromTemplate(client *api.RESTClient, template, name string) error {
	body, err := json.Marshal(generateRequest{
		Owner:   org,
		Name:    name,
		Private: true,
	})
	if err != nil {
		return err
	}

	path := fmt.Sprintf("repos/%s/template-%s/generate", org, template)
	return client.Post(path, bytes.NewReader(body), nil)
}

func applySettings(client *api.RESTClient, name string) error {
	body, err := json.Marshal(repoSettings{
		AllowSquashMerge:         true,
		AllowMergeCommit:         false,
		AllowRebaseMerge:         false,
		DeleteBranchOnMerge:      true,
		SquashMergeCommitTitle:   "PR_TITLE",
		SquashMergeCommitMessage: "BLANK",
	})
	if err != nil {
		return err
	}

	path := fmt.Sprintf("repos/%s/%s", org, name)
	return client.Patch(path, bytes.NewReader(body), nil)
}
