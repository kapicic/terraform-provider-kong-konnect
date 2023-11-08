package routes


type Route struct {
  CreatedAt *int64 `json:"created_at,omitempty"`
  Headers *Headers `json:"headers,omitempty"`
  Hosts *[]string `json:"hosts,omitempty"`
  HttpsRedirectStatusCode *int64 `json:"https_redirect_status_code,omitempty"`
  Id *string `json:"id,omitempty"`
  Methods *[]string `json:"methods,omitempty"`
  Name *string `json:"name,omitempty"`
  PathHandling *string `json:"path_handling,omitempty"`
  Paths *[]string `json:"paths,omitempty"`
  PreserveHost *bool `json:"preserve_host,omitempty"`
  Protocols *[]string `json:"protocols,omitempty"`
  RegexPriority *int64 `json:"regex_priority,omitempty"`
  RequestBuffering *bool `json:"request_buffering,omitempty"`
  ResponseBuffering *bool `json:"response_buffering,omitempty"`
  Service *Service `json:"service,omitempty"`
  Snis *[]string `json:"snis,omitempty"`
  StripPath *bool `json:"strip_path,omitempty"`
  Tags *[]string `json:"tags,omitempty"`
  UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (r *Route) SetCreatedAt(createdAt int64) {
  r.CreatedAt = &createdAt
}

func (r *Route) GetCreatedAt() *int64 {
  if r == nil {
    return nil
  }
  return r.CreatedAt
}

func (r *Route) SetHeaders(headers Headers) {
  r.Headers = &headers
}

func (r *Route) GetHeaders() *Headers {
  if r == nil {
    return nil
  }
  return r.Headers
}

func (r *Route) SetHosts(hosts []string) {
  r.Hosts = &hosts
}

func (r *Route) GetHosts() *[]string {
  if r == nil {
    return nil
  }
  return r.Hosts
}

func (r *Route) SetHttpsRedirectStatusCode(httpsRedirectStatusCode int64) {
  r.HttpsRedirectStatusCode = &httpsRedirectStatusCode
}

func (r *Route) GetHttpsRedirectStatusCode() *int64 {
  if r == nil {
    return nil
  }
  return r.HttpsRedirectStatusCode
}

func (r *Route) SetId(id string) {
  r.Id = &id
}

func (r *Route) GetId() *string {
  if r == nil {
    return nil
  }
  return r.Id
}

func (r *Route) SetMethods(methods []string) {
  r.Methods = &methods
}

func (r *Route) GetMethods() *[]string {
  if r == nil {
    return nil
  }
  return r.Methods
}

func (r *Route) SetName(name string) {
  r.Name = &name
}

func (r *Route) GetName() *string {
  if r == nil {
    return nil
  }
  return r.Name
}

func (r *Route) SetPathHandling(pathHandling string) {
  r.PathHandling = &pathHandling
}

func (r *Route) GetPathHandling() *string {
  if r == nil {
    return nil
  }
  return r.PathHandling
}

func (r *Route) SetPaths(paths []string) {
  r.Paths = &paths
}

func (r *Route) GetPaths() *[]string {
  if r == nil {
    return nil
  }
  return r.Paths
}

func (r *Route) SetPreserveHost(preserveHost bool) {
  r.PreserveHost = &preserveHost
}

func (r *Route) GetPreserveHost() *bool {
  if r == nil {
    return nil
  }
  return r.PreserveHost
}

func (r *Route) SetProtocols(protocols []string) {
  r.Protocols = &protocols
}

func (r *Route) GetProtocols() *[]string {
  if r == nil {
    return nil
  }
  return r.Protocols
}

func (r *Route) SetRegexPriority(regexPriority int64) {
  r.RegexPriority = &regexPriority
}

func (r *Route) GetRegexPriority() *int64 {
  if r == nil {
    return nil
  }
  return r.RegexPriority
}

func (r *Route) SetRequestBuffering(requestBuffering bool) {
  r.RequestBuffering = &requestBuffering
}

func (r *Route) GetRequestBuffering() *bool {
  if r == nil {
    return nil
  }
  return r.RequestBuffering
}

func (r *Route) SetResponseBuffering(responseBuffering bool) {
  r.ResponseBuffering = &responseBuffering
}

func (r *Route) GetResponseBuffering() *bool {
  if r == nil {
    return nil
  }
  return r.ResponseBuffering
}

func (r *Route) SetService(service Service) {
  r.Service = &service
}

func (r *Route) GetService() *Service {
  if r == nil {
    return nil
  }
  return r.Service
}

func (r *Route) SetSnis(snis []string) {
  r.Snis = &snis
}

func (r *Route) GetSnis() *[]string {
  if r == nil {
    return nil
  }
  return r.Snis
}

func (r *Route) SetStripPath(stripPath bool) {
  r.StripPath = &stripPath
}

func (r *Route) GetStripPath() *bool {
  if r == nil {
    return nil
  }
  return r.StripPath
}

func (r *Route) SetTags(tags []string) {
  r.Tags = &tags
}

func (r *Route) GetTags() *[]string {
  if r == nil {
    return nil
  }
  return r.Tags
}

func (r *Route) SetUpdatedAt(updatedAt int64) {
  r.UpdatedAt = &updatedAt
}

func (r *Route) GetUpdatedAt() *int64 {
  if r == nil {
    return nil
  }
  return r.UpdatedAt
}

type Headers struct {
  Key *string `json:"key,omitempty"`
}

func (h *Headers) SetKey(key string) {
  h.Key = &key
}

func (h *Headers) GetKey() *string {
  if h == nil {
    return nil
  }
  return h.Key
}



type Service struct {
  Id *string `json:"id,omitempty"`
}

func (s *Service) SetId(id string) {
  s.Id = &id
}

func (s *Service) GetId() *string {
  if s == nil {
    return nil
  }
  return s.Id
}





