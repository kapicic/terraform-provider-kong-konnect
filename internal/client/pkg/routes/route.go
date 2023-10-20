package routes

type Route struct {
	CreatedAt               *int64    `json:"created_at,omitempty"`
	Headers                 *any      `json:"headers,omitempty"`
	Hosts                   *[]string `json:"hosts,omitempty"`
	HttpsRedirectStatusCode *int64    `json:"https_redirect_status_code,omitempty"`
	Id                      *string   `json:"id,omitempty"`
	Methods                 *[]string `json:"methods,omitempty"`
	Name                    *string   `json:"name,omitempty"`
	PathHandling            *string   `json:"path_handling,omitempty"`
	Paths                   *[]string `json:"paths,omitempty"`
	PreserveHost            *bool     `json:"preserve_host,omitempty"`
	Protocols               *[]string `json:"protocols,omitempty"`
	RegexPriority           *int64    `json:"regex_priority,omitempty"`
	RequestBuffering        *bool     `json:"request_buffering,omitempty"`
	ResponseBuffering       *bool     `json:"response_buffering,omitempty"`
	Service                 *Service  `json:"service,omitempty"`
	Snis                    *[]string `json:"snis,omitempty"`
	StripPath               *bool     `json:"strip_path,omitempty"`
	Tags                    *[]string `json:"tags,omitempty"`
	UpdatedAt               *int64    `json:"updated_at,omitempty"`
}

func (r *Route) SetCreatedAt(createdAt int64) {
	r.CreatedAt = &createdAt
}

func (r *Route) SetHeaders(headers any) {
	r.Headers = &headers
}

func (r *Route) SetHosts(hosts []string) {
	r.Hosts = &hosts
}

func (r *Route) SetHttpsRedirectStatusCode(httpsRedirectStatusCode int64) {
	r.HttpsRedirectStatusCode = &httpsRedirectStatusCode
}

func (r *Route) SetId(id string) {
	r.Id = &id
}

func (r *Route) SetMethods(methods []string) {
	r.Methods = &methods
}

func (r *Route) SetName(name string) {
	r.Name = &name
}

func (r *Route) SetPathHandling(pathHandling string) {
	r.PathHandling = &pathHandling
}

func (r *Route) SetPaths(paths []string) {
	r.Paths = &paths
}

func (r *Route) SetPreserveHost(preserveHost bool) {
	r.PreserveHost = &preserveHost
}

func (r *Route) SetProtocols(protocols []string) {
	r.Protocols = &protocols
}

func (r *Route) SetRegexPriority(regexPriority int64) {
	r.RegexPriority = &regexPriority
}

func (r *Route) SetRequestBuffering(requestBuffering bool) {
	r.RequestBuffering = &requestBuffering
}

func (r *Route) SetResponseBuffering(responseBuffering bool) {
	r.ResponseBuffering = &responseBuffering
}

func (r *Route) SetService(service Service) {
	r.Service = &service
}

func (r *Route) GetService() *Service {
	if r.Service == nil {
		return nil
	}
	return r.Service
}

func (r *Route) SetSnis(snis []string) {
	r.Snis = &snis
}

func (r *Route) SetStripPath(stripPath bool) {
	r.StripPath = &stripPath
}

func (r *Route) SetTags(tags []string) {
	r.Tags = &tags
}

func (r *Route) SetUpdatedAt(updatedAt int64) {
	r.UpdatedAt = &updatedAt
}

type Service struct {
	Id *string `json:"id,omitempty"`
}

func (s *Service) SetId(id string) {
	s.Id = &id
}
