/*
 * CircleCI API
 *
 * The CircleCI API is a fully-featured JSON API that allows you to access all information and trigger all actions in CircleCI.
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type GitHubIdentity struct {
	Authorized bool `json:"authorized?,omitempty"`

	AvatarUrl string `json:"avatar_url,omitempty"`

	Domain string `json:"domain,omitempty"`

	ExternalId int32 `json:"external_id,omitempty"`

	Id int32 `json:"id,omitempty"`

	Login string `json:"login,omitempty"`

	Name string `json:"name,omitempty"`

	ProviderId string `json:"provider_id,omitempty"`

	Type_ string `json:"type,omitempty"`

	User bool `json:"user?,omitempty"`
}
