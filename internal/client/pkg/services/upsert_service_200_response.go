package services

type UpsertService_200Response struct {
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

func (u *UpsertService_200Response) SetConnectTimeout(connectTimeout int64) {
	u.ConnectTimeout = &connectTimeout
}

func (u *UpsertService_200Response) SetCreatedAt(createdAt int64) {
	u.CreatedAt = &createdAt
}

func (u *UpsertService_200Response) SetEnabled(enabled bool) {
	u.Enabled = &enabled
}

func (u *UpsertService_200Response) SetHost(host string) {
	u.Host = &host
}

func (u *UpsertService_200Response) SetId(id string) {
	u.Id = &id
}

func (u *UpsertService_200Response) SetName(name string) {
	u.Name = &name
}

func (u *UpsertService_200Response) SetPath(path string) {
	u.Path = &path
}

func (u *UpsertService_200Response) SetPort(port int64) {
	u.Port = &port
}

func (u *UpsertService_200Response) SetProtocol(protocol string) {
	u.Protocol = &protocol
}

func (u *UpsertService_200Response) SetReadTimeout(readTimeout int64) {
	u.ReadTimeout = &readTimeout
}

func (u *UpsertService_200Response) SetRetries(retries int64) {
	u.Retries = &retries
}

func (u *UpsertService_200Response) SetUpdatedAt(updatedAt int64) {
	u.UpdatedAt = &updatedAt
}

func (u *UpsertService_200Response) SetWriteTimeout(writeTimeout int64) {
	u.WriteTimeout = &writeTimeout
}
