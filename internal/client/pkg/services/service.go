package services

type Service struct {
	CaCertificates    *[]string          `json:"ca_certificates,omitempty"`
	ClientCertificate *ClientCertificate `json:"client_certificate,omitempty"`
	ConnectTimeout    *int64             `json:"connect_timeout,omitempty"`
	CreatedAt         *int64             `json:"created_at,omitempty"`
	Enabled           *bool              `json:"enabled,omitempty"`
	Host              *string            `json:"host,omitempty"`
	Id                *string            `json:"id,omitempty"`
	Name              *string            `json:"name,omitempty"`
	Path              *string            `json:"path,omitempty"`
	Port              *int64             `json:"port,omitempty"`
	Protocol          *string            `json:"protocol,omitempty"`
	ReadTimeout       *int64             `json:"read_timeout,omitempty"`
	Retries           *int64             `json:"retries,omitempty"`
	Tags              *[]string          `json:"tags,omitempty"`
	TlsVerify         *bool              `json:"tls_verify,omitempty"`
	TlsVerifyDepth    *int64             `json:"tls_verify_depth,omitempty"`
	UpdatedAt         *int64             `json:"updated_at,omitempty"`
	Url               *string            `json:"url,omitempty"`
	WriteTimeout      *int64             `json:"write_timeout,omitempty"`
}

func (s *Service) SetCaCertificates(caCertificates []string) {
	s.CaCertificates = &caCertificates
}

func (s *Service) SetClientCertificate(clientCertificate ClientCertificate) {
	s.ClientCertificate = &clientCertificate
}

func (s *Service) GetClientCertificate() *ClientCertificate {
	if s.ClientCertificate == nil {
		return nil
	}
	return s.ClientCertificate
}

func (s *Service) SetConnectTimeout(connectTimeout int64) {
	s.ConnectTimeout = &connectTimeout
}

func (s *Service) SetCreatedAt(createdAt int64) {
	s.CreatedAt = &createdAt
}

func (s *Service) SetEnabled(enabled bool) {
	s.Enabled = &enabled
}

func (s *Service) SetHost(host string) {
	s.Host = &host
}

func (s *Service) SetId(id string) {
	s.Id = &id
}

func (s *Service) SetName(name string) {
	s.Name = &name
}

func (s *Service) SetPath(path string) {
	s.Path = &path
}

func (s *Service) SetPort(port int64) {
	s.Port = &port
}

func (s *Service) SetProtocol(protocol string) {
	s.Protocol = &protocol
}

func (s *Service) SetReadTimeout(readTimeout int64) {
	s.ReadTimeout = &readTimeout
}

func (s *Service) SetRetries(retries int64) {
	s.Retries = &retries
}

func (s *Service) SetTags(tags []string) {
	s.Tags = &tags
}

func (s *Service) SetTlsVerify(tlsVerify bool) {
	s.TlsVerify = &tlsVerify
}

func (s *Service) SetTlsVerifyDepth(tlsVerifyDepth int64) {
	s.TlsVerifyDepth = &tlsVerifyDepth
}

func (s *Service) SetUpdatedAt(updatedAt int64) {
	s.UpdatedAt = &updatedAt
}

func (s *Service) SetUrl(url string) {
	s.Url = &url
}

func (s *Service) SetWriteTimeout(writeTimeout int64) {
	s.WriteTimeout = &writeTimeout
}

type ClientCertificate struct {
	Id *string `json:"id,omitempty"`
}

func (c *ClientCertificate) SetId(id string) {
	c.Id = &id
}
