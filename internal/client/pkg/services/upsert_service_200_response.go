package services


type UpsertService_200Response struct {
  ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
  CreatedAt *int64 `json:"created_at,omitempty"`
  Enabled *bool `json:"enabled,omitempty"`
  Host *string `json:"host,omitempty"`
  Id *string `json:"id,omitempty"`
  Name *string `json:"name,omitempty"`
  Path *string `json:"path,omitempty"`
  Port *int64 `json:"port,omitempty"`
  Protocol *string `json:"protocol,omitempty"`
  ReadTimeout *int64 `json:"read_timeout,omitempty"`
  Retries *int64 `json:"retries,omitempty"`
  UpdatedAt *int64 `json:"updated_at,omitempty"`
  WriteTimeout *int64 `json:"write_timeout,omitempty"`
}

func (u *UpsertService_200Response) SetConnectTimeout(connectTimeout int64) {
  u.ConnectTimeout = &connectTimeout
}

func (u *UpsertService_200Response) GetConnectTimeout() *int64 {
  if u == nil {
    return nil
  }
  return u.ConnectTimeout
}

func (u *UpsertService_200Response) SetCreatedAt(createdAt int64) {
  u.CreatedAt = &createdAt
}

func (u *UpsertService_200Response) GetCreatedAt() *int64 {
  if u == nil {
    return nil
  }
  return u.CreatedAt
}

func (u *UpsertService_200Response) SetEnabled(enabled bool) {
  u.Enabled = &enabled
}

func (u *UpsertService_200Response) GetEnabled() *bool {
  if u == nil {
    return nil
  }
  return u.Enabled
}

func (u *UpsertService_200Response) SetHost(host string) {
  u.Host = &host
}

func (u *UpsertService_200Response) GetHost() *string {
  if u == nil {
    return nil
  }
  return u.Host
}

func (u *UpsertService_200Response) SetId(id string) {
  u.Id = &id
}

func (u *UpsertService_200Response) GetId() *string {
  if u == nil {
    return nil
  }
  return u.Id
}

func (u *UpsertService_200Response) SetName(name string) {
  u.Name = &name
}

func (u *UpsertService_200Response) GetName() *string {
  if u == nil {
    return nil
  }
  return u.Name
}

func (u *UpsertService_200Response) SetPath(path string) {
  u.Path = &path
}

func (u *UpsertService_200Response) GetPath() *string {
  if u == nil {
    return nil
  }
  return u.Path
}

func (u *UpsertService_200Response) SetPort(port int64) {
  u.Port = &port
}

func (u *UpsertService_200Response) GetPort() *int64 {
  if u == nil {
    return nil
  }
  return u.Port
}

func (u *UpsertService_200Response) SetProtocol(protocol string) {
  u.Protocol = &protocol
}

func (u *UpsertService_200Response) GetProtocol() *string {
  if u == nil {
    return nil
  }
  return u.Protocol
}

func (u *UpsertService_200Response) SetReadTimeout(readTimeout int64) {
  u.ReadTimeout = &readTimeout
}

func (u *UpsertService_200Response) GetReadTimeout() *int64 {
  if u == nil {
    return nil
  }
  return u.ReadTimeout
}

func (u *UpsertService_200Response) SetRetries(retries int64) {
  u.Retries = &retries
}

func (u *UpsertService_200Response) GetRetries() *int64 {
  if u == nil {
    return nil
  }
  return u.Retries
}

func (u *UpsertService_200Response) SetUpdatedAt(updatedAt int64) {
  u.UpdatedAt = &updatedAt
}

func (u *UpsertService_200Response) GetUpdatedAt() *int64 {
  if u == nil {
    return nil
  }
  return u.UpdatedAt
}

func (u *UpsertService_200Response) SetWriteTimeout(writeTimeout int64) {
  u.WriteTimeout = &writeTimeout
}

func (u *UpsertService_200Response) GetWriteTimeout() *int64 {
  if u == nil {
    return nil
  }
  return u.WriteTimeout
}



