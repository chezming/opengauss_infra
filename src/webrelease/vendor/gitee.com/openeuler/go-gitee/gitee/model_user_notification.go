/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 获取一条通知
type UserNotification struct {
	Id        int32  `json:"id,omitempty"`
	Content   string `json:"content,omitempty"`
	Type_     string `json:"type,omitempty"`
	Unread    string `json:"unread,omitempty"`
	Mute      string `json:"mute,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Url       string `json:"url,omitempty"`
	HtmlUrl   string `json:"html_url,omitempty"`
	// 通知发送者
	Actor      *UserBasic    `json:"actor,omitempty"`
	Repository *ProjectBasic `json:"repository,omitempty"`
	// 通知直接关联对象
	Subject *UserNotificationSubject `json:"subject,omitempty"`
	// 通知次级关联对象
	Namespaces []UserNotificationNamespace `json:"namespaces,omitempty"`
}
