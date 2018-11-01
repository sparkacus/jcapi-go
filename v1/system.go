/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type System struct {
	Organization string `json:"organization,omitempty"`

	Created string `json:"created,omitempty"`

	LastContact string `json:"lastContact,omitempty"`

	Os string `json:"os,omitempty"`

	Version string `json:"version,omitempty"`

	Arch string `json:"arch,omitempty"`

	NetworkInterfaces []SystemNetworkInterfaces `json:"networkInterfaces,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	SystemTimezone int32 `json:"systemTimezone,omitempty"`

	TemplateName string `json:"templateName,omitempty"`

	RemoteIP string `json:"remoteIP,omitempty"`

	Active bool `json:"active,omitempty"`

	SshdParams []string `json:"sshdParams,omitempty"`

	AllowSshPasswordAuthentication bool `json:"allowSshPasswordAuthentication,omitempty"`

	AllowSshRootLogin bool `json:"allowSshRootLogin,omitempty"`

	AllowMultiFactorAuthentication bool `json:"allowMultiFactorAuthentication,omitempty"`

	AllowPublicKeyAuthentication bool `json:"allowPublicKeyAuthentication,omitempty"`

	ModifySSHDConfig bool `json:"modifySSHDConfig,omitempty"`

	AgentVersion string `json:"agentVersion,omitempty"`

	ConnectionHistory []interface{} `json:"connectionHistory,omitempty"`

	SshRootEnabled bool `json:"sshRootEnabled,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Id string `json:"_id,omitempty"`

	Fde *Fde `json:"fde,omitempty"`

	AmazonInstanceID string `json:"amazonInstanceID"`
}
