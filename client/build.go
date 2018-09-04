/*
 * CircleCI API
 *
 * The CircleCI API is a fully-featured JSON API that allows you to access all information and trigger all actions in CircleCI.
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

type Build struct {
	AllCommitDetails []Commit `json:"all_commit_details,omitempty"`

	AllCommitDetailsTruncated bool `json:"all_commit_details_truncated,omitempty"`

	AuthorDate time.Time `json:"author_date,omitempty"`

	AuthorEmail string `json:"author_email,omitempty"`

	AuthorName string `json:"author_name,omitempty"`

	Body string `json:"body,omitempty"`

	Branch string `json:"branch,omitempty"`

	BuildNum int32 `json:"build_num,omitempty"`

	BuildTimeMillis int32 `json:"build_time_millis,omitempty"`

	BuildUrl string `json:"build_url,omitempty"`

	Canceled bool `json:"canceled,omitempty"`

	CommitterDate time.Time `json:"committer_date,omitempty"`

	CommitterEmail string `json:"committer_email,omitempty"`

	CommitterName string `json:"committer_name,omitempty"`

	Compare string `json:"compare,omitempty"`

	DontBuild string `json:"dont_build,omitempty"`

	Failed bool `json:"failed,omitempty"`

	HasArtifacts bool `json:"has_artifacts,omitempty"`

	InfrastructureFail bool `json:"infrastructure_fail,omitempty"`

	IsFirstGreenBuild bool `json:"is_first_green_build,omitempty"`

	Lifecycle string `json:"lifecycle,omitempty"`

	NoDependencyCache bool `json:"no_dependency_cache,omitempty"`

	Oss bool `json:"oss,omitempty"`

	Outcome string `json:"outcome,omitempty"`

	Parallel int32 `json:"parallel,omitempty"`

	Platform string `json:"platform,omitempty"`

	Previous *BuildRef `json:"previous,omitempty"`

	PreviousSuccessfulBuild *BuildRef `json:"previous_successful_build,omitempty"`

	QueuedAt time.Time `json:"queued_at,omitempty"`

	Reponame string `json:"reponame,omitempty"`

	SshEnabled bool `json:"ssh_enabled,omitempty"`

	StartTime time.Time `json:"start_time,omitempty"`

	Status string `json:"status,omitempty"`

	StopTime time.Time `json:"stop_time,omitempty"`

	Subject string `json:"subject,omitempty"`

	Timedout bool `json:"timedout,omitempty"`

	UsageQueuedAt time.Time `json:"usage_queued_at,omitempty"`

	User *UserShort `json:"user,omitempty"`

	Username string `json:"username,omitempty"`

	VcsRevision string `json:"vcs_revision,omitempty"`

	VcsTag string `json:"vcs_tag,omitempty"`

	VcsType string `json:"vcs_type,omitempty"`

	VcsUrl string `json:"vcs_url,omitempty"`

	Why string `json:"why,omitempty"`
}