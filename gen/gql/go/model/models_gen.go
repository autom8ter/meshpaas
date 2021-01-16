// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type App struct {
	Name      string                 `json:"name"`
	Namespace string                 `json:"namespace"`
	Image     string                 `json:"image"`
	Env       map[string]interface{} `json:"env"`
	Ports     map[string]interface{} `json:"ports"`
	Replicas  int                    `json:"replicas"`
	Status    *Status                `json:"status"`
}

type AppConstructor struct {
	Name      string                 `json:"name"`
	Namespace string                 `json:"namespace"`
	Image     string                 `json:"image"`
	Env       map[string]interface{} `json:"env"`
	Ports     map[string]interface{} `json:"ports"`
	Replicas  int                    `json:"replicas"`
}

type AppRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type AppUpdate struct {
	Name      string                 `json:"name"`
	Namespace string                 `json:"namespace"`
	DryRun    *bool                  `json:"dry_run"`
	Image     *string                `json:"image"`
	Env       map[string]interface{} `json:"env"`
	Ports     map[string]interface{} `json:"ports"`
	Replicas  *int                   `json:"replicas"`
}

type Log struct {
	Message string `json:"message"`
}

type LogFilter struct {
	App        *AppRef `json:"app"`
	Expression *string `json:"expression"`
}

type Replica struct {
	Phase     string `json:"phase"`
	Condition string `json:"condition"`
	Reason    string `json:"reason"`
}

type Status struct {
	Replicas []*Replica `json:"replicas"`
}
