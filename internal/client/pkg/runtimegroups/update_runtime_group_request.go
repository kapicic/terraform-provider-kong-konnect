package runtimegroups


type UpdateRuntimeGroupRequest struct {
  Name *string `json:"name,omitempty"`
  Description *string `json:"description,omitempty"`
  AuthType *AuthType1 `json:"auth_type,omitempty"`
  Labels *Labels `json:"labels,omitempty"`
}

func (u *UpdateRuntimeGroupRequest) SetName(name string) {
  u.Name = &name
}

func (u *UpdateRuntimeGroupRequest) GetName() *string {
  if u == nil {
    return nil
  }
  return u.Name
}

func (u *UpdateRuntimeGroupRequest) SetDescription(description string) {
  u.Description = &description
}

func (u *UpdateRuntimeGroupRequest) GetDescription() *string {
  if u == nil {
    return nil
  }
  return u.Description
}

func (u *UpdateRuntimeGroupRequest) SetAuthType(authType AuthType1) {
  u.AuthType = &authType
}

func (u *UpdateRuntimeGroupRequest) GetAuthType() *AuthType1 {
  if u == nil {
    return nil
  }
  return u.AuthType
}

func (u *UpdateRuntimeGroupRequest) SetLabels(labels Labels) {
  u.Labels = &labels
}

func (u *UpdateRuntimeGroupRequest) GetLabels() *Labels {
  if u == nil {
    return nil
  }
  return u.Labels
}

type AuthType1 string

const (
  AUTH_TYPE1_PINNED_CLIENT_CERTS AuthType1 = "pinned_client_certs"
  AUTH_TYPE1_PKI_CLIENT_CERTS AuthType1 = "pki_client_certs"
)




