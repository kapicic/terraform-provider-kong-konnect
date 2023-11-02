package runtimegroups

type CreateRuntimeGroupRequest struct {
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	ClusterType *ClusterType `json:"cluster_type,omitempty"`
	AuthType    *AuthType    `json:"auth_type,omitempty"`
	Labels      *Labels      `json:"labels,omitempty"`
}

func (c *CreateRuntimeGroupRequest) SetName(name string) {
	c.Name = &name
}

func (c *CreateRuntimeGroupRequest) SetDescription(description string) {
	c.Description = &description
}

func (c *CreateRuntimeGroupRequest) SetClusterType(clusterType ClusterType) {
	c.ClusterType = &clusterType
}

func (c *CreateRuntimeGroupRequest) SetAuthType(authType AuthType) {
	c.AuthType = &authType
}

func (c *CreateRuntimeGroupRequest) SetLabels(labels Labels) {
	c.Labels = &labels
}

func (c *CreateRuntimeGroupRequest) GetLabels() *Labels {
	if c.Labels == nil {
		return nil
	}
	return c.Labels
}

type ClusterType string

const (
	CLUSTER_TYPE_CLUSTER_TYPE_HYBRID                  ClusterType = "CLUSTER_TYPE_HYBRID"
	CLUSTER_TYPE_CLUSTER_TYPE_K8_S_INGRESS_CONTROLLER ClusterType = "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER"
	CLUSTER_TYPE_CLUSTER_TYPE_COMPOSITE               ClusterType = "CLUSTER_TYPE_COMPOSITE"
)

type AuthType string

const (
	AUTH_TYPE_PINNED_CLIENT_CERTS AuthType = "pinned_client_certs"
	AUTH_TYPE_PKI_CLIENT_CERTS    AuthType = "pki_client_certs"
)
