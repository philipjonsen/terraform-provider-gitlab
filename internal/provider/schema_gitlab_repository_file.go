package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xanzy/go-gitlab"
)

func gitlabRepositoryFileGetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"project": {
			Description: "The name or ID of the project.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"file_path": {
			Description: "The full path of the file. It must be relative to the root of the project without a leading slash `/`.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"ref": {
			Description: "The name of branch, tag or commit.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"file_name": {
			Description: "The filename.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"size": {
			Description: "The file size.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"encoding": {
			Description: "The file content encoding.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"content": {
			Description:  "base64 encoded file content. No other encoding is currently supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validateBase64Content,
		},
		"content_sha256": {
			Description: "File content sha256 digest.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"blob_id": {
			Description: "The blob id.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"commit_id": {
			Description: "The commit id.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"last_commit_id": {
			Description: "The last known commit id.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func gitlabRepositoryFileToStateMap(project string, repositoryFile *gitlab.File) map[string]interface{} {
	stateMap := make(map[string]interface{})
	stateMap["project"] = project
	stateMap["file_name"] = repositoryFile.FileName
	stateMap["file_path"] = repositoryFile.FilePath
	stateMap["size"] = repositoryFile.Size
	stateMap["encoding"] = repositoryFile.Encoding
	stateMap["content"] = repositoryFile.Content
	stateMap["content_sha256"] = repositoryFile.SHA256
	stateMap["ref"] = repositoryFile.Ref
	stateMap["blob_id"] = repositoryFile.BlobID
	stateMap["commit_id"] = repositoryFile.CommitID
	stateMap["last_commit_id"] = repositoryFile.LastCommitID
	return stateMap
}
