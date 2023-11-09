package runtimegroups

type Labels struct {
	Name *string `json:"name,omitempty"`
}

func (l *Labels) SetName(name string) {
	l.Name = &name
}

func (l *Labels) GetName() *string {
	if l == nil {
		return nil
	}
	return l.Name
}
