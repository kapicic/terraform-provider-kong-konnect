package runtimegroups

type RuntimeGroup struct {
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Labels      *any    `json:"labels,omitempty"`
	Config      *Config `json:"config,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

func (r *RuntimeGroup) SetId(id string) {
	r.Id = &id
}

func (r *RuntimeGroup) SetName(name string) {
	r.Name = &name
}

func (r *RuntimeGroup) SetDescription(description string) {
	r.Description = &description
}

func (r *RuntimeGroup) SetLabels(labels any) {
	r.Labels = &labels
}

func (r *RuntimeGroup) SetConfig(config Config) {
	r.Config = &config
}

func (r *RuntimeGroup) GetConfig() *Config {
	if r.Config == nil {
		return nil
	}
	return r.Config
}

func (r *RuntimeGroup) SetCreatedAt(createdAt string) {
	r.CreatedAt = &createdAt
}

func (r *RuntimeGroup) SetUpdatedAt(updatedAt string) {
	r.UpdatedAt = &updatedAt
}

type Config struct {
	ControlPlaneEndpoint *string `json:"control_plane_endpoint,omitempty"`
	TelemetryEndpoint    *string `json:"telemetry_endpoint,omitempty"`
}

func (c *Config) SetControlPlaneEndpoint(controlPlaneEndpoint string) {
	c.ControlPlaneEndpoint = &controlPlaneEndpoint
}

func (c *Config) SetTelemetryEndpoint(telemetryEndpoint string) {
	c.TelemetryEndpoint = &telemetryEndpoint
}
