package services

type GetService_200Response struct {
	ConnectTimeout *int64  `json:"connect_timeout,omitempty"`
	CreatedAt      *int64  `json:"created_at,omitempty"`
	Enabled        *bool   `json:"enabled,omitempty"`
	Host           *string `json:"host,omitempty"`
	Id             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	Path           *string `json:"path,omitempty"`
	Port           *int64  `json:"port,omitempty"`
	Protocol       *string `json:"protocol,omitempty"`
	ReadTimeout    *int64  `json:"read_timeout,omitempty"`
	Retries        *int64  `json:"retries,omitempty"`
	UpdatedAt      *int64  `json:"updated_at,omitempty"`
	WriteTimeout   *int64  `json:"write_timeout,omitempty"`
}

func (g *GetService_200Response) SetConnectTimeout(connectTimeout int64) {
	g.ConnectTimeout = &connectTimeout
}

func (g *GetService_200Response) GetConnectTimeout() *int64 {
	if g == nil {
		return nil
	}
	return g.ConnectTimeout
}

func (g *GetService_200Response) SetCreatedAt(createdAt int64) {
	g.CreatedAt = &createdAt
}

func (g *GetService_200Response) GetCreatedAt() *int64 {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetService_200Response) SetEnabled(enabled bool) {
	g.Enabled = &enabled
}

func (g *GetService_200Response) GetEnabled() *bool {
	if g == nil {
		return nil
	}
	return g.Enabled
}

func (g *GetService_200Response) SetHost(host string) {
	g.Host = &host
}

func (g *GetService_200Response) GetHost() *string {
	if g == nil {
		return nil
	}
	return g.Host
}

func (g *GetService_200Response) SetId(id string) {
	g.Id = &id
}

func (g *GetService_200Response) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetService_200Response) SetName(name string) {
	g.Name = &name
}

func (g *GetService_200Response) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetService_200Response) SetPath(path string) {
	g.Path = &path
}

func (g *GetService_200Response) GetPath() *string {
	if g == nil {
		return nil
	}
	return g.Path
}

func (g *GetService_200Response) SetPort(port int64) {
	g.Port = &port
}

func (g *GetService_200Response) GetPort() *int64 {
	if g == nil {
		return nil
	}
	return g.Port
}

func (g *GetService_200Response) SetProtocol(protocol string) {
	g.Protocol = &protocol
}

func (g *GetService_200Response) GetProtocol() *string {
	if g == nil {
		return nil
	}
	return g.Protocol
}

func (g *GetService_200Response) SetReadTimeout(readTimeout int64) {
	g.ReadTimeout = &readTimeout
}

func (g *GetService_200Response) GetReadTimeout() *int64 {
	if g == nil {
		return nil
	}
	return g.ReadTimeout
}

func (g *GetService_200Response) SetRetries(retries int64) {
	g.Retries = &retries
}

func (g *GetService_200Response) GetRetries() *int64 {
	if g == nil {
		return nil
	}
	return g.Retries
}

func (g *GetService_200Response) SetUpdatedAt(updatedAt int64) {
	g.UpdatedAt = &updatedAt
}

func (g *GetService_200Response) GetUpdatedAt() *int64 {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetService_200Response) SetWriteTimeout(writeTimeout int64) {
	g.WriteTimeout = &writeTimeout
}

func (g *GetService_200Response) GetWriteTimeout() *int64 {
	if g == nil {
		return nil
	}
	return g.WriteTimeout
}
