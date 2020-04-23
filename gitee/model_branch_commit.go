/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

type BranchCommit struct {
	Sha string `json:"sha,omitempty"`
	Url string `json:"url,omitempty"`
	Commit string `json:"commit,omitempty"`
	Author string `json:"author,omitempty"`
	Parents string `json:"parents,omitempty"`
	Committer string `json:"committer,omitempty"`
}
