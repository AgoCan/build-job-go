package git

import (
	"build-job-go/config"
	"fmt"

	. "build-job-go/utils"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func Git() {
	url := config.GitURL
	workspace := fmt.Sprintf("%s/%s/%s/%s",
		config.Workspace,
		config.Username,
		config.Project,
		config.CodeDir,
	)

	// Clone the given repository to the given directory
	Info("git clone %s %s --recursive", url, workspace)

	r, err := git.PlainClone(workspace, false, &git.CloneOptions{
		URL: url,
		// 只拉取最新代码
		RecurseSubmodules: git.NoRecurseSubmodules,
		Auth: &http.BasicAuth{
			Username: config.GitUsername,
			Password: config.GitPassword},
	})

	CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
