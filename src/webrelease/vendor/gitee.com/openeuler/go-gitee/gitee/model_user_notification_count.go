/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 获取授权用户的通知数
type UserNotificationCount struct {
	// 通知总数
	TotalCount int32 `json:"total_count,omitempty"`
	// 通知数量
	NotificationCount int32 `json:"notification_count,omitempty"`
	// 私信数量
	MessageCount int32 `json:"message_count,omitempty"`
}