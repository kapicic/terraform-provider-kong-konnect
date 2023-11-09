package services

type CreateService_201Response struct {
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

func (c *CreateService_201Response) SetConnectTimeout(connectTimeout int64) {
	c.ConnectTimeout = &connectTimeout
}

func (c *CreateService_201Response) GetConnectTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ConnectTimeout
}

func (c *CreateService_201Response) SetCreatedAt(createdAt int64) {
	c.CreatedAt = &createdAt
}

func (c *CreateService_201Response) GetCreatedAt() *int64 {
	if c == nil {
		return nil
	}
	return c.CreatedAt
}

func (c *CreateService_201Response) SetEnabled(enabled bool) {
	c.Enabled = &enabled
}

func (c *CreateService_201Response) GetEnabled() *bool {
	if c == nil {
		return nil
	}
	return c.Enabled
}

func (c *CreateService_201Response) SetHost(host string) {
	c.Host = &host
}

func (c *CreateService_201Response) GetHost() *string {
	if c == nil {
		return nil
	}
	return c.Host
}

func (c *CreateService_201Response) SetId(id string) {
	c.Id = &id
}

func (c *CreateService_201Response) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *CreateService_201Response) SetName(name string) {
	c.Name = &name
}

func (c *CreateService_201Response) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateService_201Response) SetPath(path string) {
	c.Path = &path
}

func (c *CreateService_201Response) GetPath() *string {
	if c == nil {
		return nil
	}
	return c.Path
}

func (c *CreateService_201Response) SetPort(port int64) {
	c.Port = &port
}

func (c *CreateService_201Response) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *CreateService_201Response) SetProtocol(protocol string) {
	c.Protocol = &protocol
}

func (c *CreateService_201Response) GetProtocol() *string {
	if c == nil {
		return nil
	}
	return c.Protocol
}

func (c *CreateService_201Response) SetReadTimeout(readTimeout int64) {
	c.ReadTimeout = &readTimeout
}

func (c *CreateService_201Response) GetReadTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ReadTimeout
}

func (c *CreateService_201Response) SetRetries(retries int64) {
	c.Retries = &retries
}

func (c *CreateService_201Response) GetRetries() *int64 {
	if c == nil {
		return nil
	}
	return c.Retries
}

func (c *CreateService_201Response) SetUpdatedAt(updatedAt int64) {
	c.UpdatedAt = &updatedAt
}

func (c *CreateService_201Response) GetUpdatedAt() *int64 {
	if c == nil {
		return nil
	}
	return c.UpdatedAt
}

func (c *CreateService_201Response) SetWriteTimeout(writeTimeout int64) {
	c.WriteTimeout = &writeTimeout
}

func (c *CreateService_201Response) GetWriteTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.WriteTimeout
}
