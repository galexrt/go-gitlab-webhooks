package gitlab

type (
	// GitLab hook structure
	GitLab struct {
		Head
		ObjectKind        string                 `json:"object_kind"`
		CheckoutSHA       string                 `json:"checkout_sha"`
		UserID            int                    `json:"user_id"`
		UserName          string                 `json:"user_name"`
		UserEmail         string                 `json:"user_email"`
		UserAvatar        string                 `json:"user_avatar"`
		ProjectID         int                    `json:"project_id"`
		Project           gitLabProject          `json:"project"`
		Repository        gitLabRepository       `json:"repository"`
		Commits           []Commit               `json:"commits"`
		ObjectAttributes  gitLabObjectAttributes `json:"object_attributes"`
		TotalCommitsCount int                    `json:"total_commits_count"`
	}

	gitLabDescriptor struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	}

	// project
	gitLabProject struct {
		gitLabDescriptor
		Namespace         string `json:"namespace"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
		WebURL            string `json:"web_url"`
		AvatarURL         string `json:"avatar_url"`
		HTTPURL           string `json:"http_url"`
		SSHURL            string `json:"ssh_url"`
	}

	// repository
	gitLabRepository struct {
		gitLabDescriptor
	}

	// object_attributes
	gitLabObjectAttributes struct {
		ID              int           `json:"id"`
		Title           string        `json:"title"`
		AssigneeIDs     []string      `json:"assignee_ids"`
		AssigneeID      int           `json:"assignee_id"`
		AuthorID        int           `json:"author_id"`
		ProjectID       string        `json:"project_id"`
		CreatedAt       string        `json:"created_at"`
		UpdatedAt       string        `json:"updated_at"`
		Note            string        `json:"note"`
		NoteableType    string        `json:"noteable_type"`
		Action          string        `json:"action"`
		Description     string        `json:"description"`
		Iid             int           `json:"iid"`
		LastCommit      Commit        `json:"last_commit"`
		MergeStatus     string        `json:"merge_status"`
		MilestoneID     int           `json:"milestone_id"`
		Source          gitLabProject `json:"source"`
		SourceBranch    string        `json:"source_branch"`
		SourceProjectID int           `json:"source_project_id"`
		StCommits       string        `json:"st_commits"`
		StDiffs         gitLabStDiff  `json:"st_diffs"`
		State           string        `json:"state"`
		Target          gitLabProject `json:"target"`
		TargetBranch    string        `json:"target_branch"`
		TargetProjectID int           `json:"target_project_id"`
		URL             string        `json:"url"`
	}

	gitLabStDiff struct {
		Diff        string `json:"diff"`
		NewPath     string `json:"new_path"`
		OldPath     string `json:"old_path"`
		AMode       string `json:"a_mode"`
		BMode       string `json:"b_mode"`
		NewFile     bool   `json:"new_file"`
		RenamedFile bool   `json:"renamed_file"`
		DeletedFile bool   `json:"deleted_file"`
	}
)
