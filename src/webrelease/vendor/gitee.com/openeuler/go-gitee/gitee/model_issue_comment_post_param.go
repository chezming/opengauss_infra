/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

type IssueCommentPostParam struct {
	// 用户授权码
	AccessToken string `json:"access_token,omitempty"`
	// The contents of the comment
	Body string `json:"body,omitempty"`
}
