package gocd

import (
	"context"
	"fmt"
)

// PipelinesService describes the HAL _link resource for the api response object for a pipelineconfig
type PipelinesService service

// Pipeline describes a pipeline object
type Pipeline struct {
	Name                  string     `json:"name"`
	LabelTemplate         string     `json:"label_template,omitempty"`
	EnablePipelineLocking bool       `json:"enable_pipeline_locking,omitempty"`
	Template              string     `json:"template,omitempty"`
	Materials             []Material `json:"materials,omitempty"`
	Stages                []Stage    `json:"stages"`
	Version               string
}

// Material describes an artifact dependency for a pipeline object.
type Material struct {
	Type       string `json:"type"`
	Attributes struct {
		URL             string      `json:"url"`
		Destination     string      `json:"destination,omitempty"`
		Filter          interface{} `json:"filter,omitempty"`
		InvertFilter    bool        `json:"invert_filter,omitempty"`
		Name            string      `json:"name,omitempty"`
		AutoUpdate      bool        `json:"auto_update,omitempty"`
		Branch          string      `json:"branch,omitempty"`
		SubmoduleFolder string      `json:"submodule_folder,omitempty"`
		ShallowClone    bool        `json:"shallow_clone,omitempty"`
	} `json:"attributes"`
}

// PipelineHistory describes the history of runs for a pipeline
type PipelineHistory struct {
	Pipelines []*PipelineInstance `json:"pipelines"`
}

// PipelineInstance describes a single pipeline run
type PipelineInstance struct {
	BuildCause   BuildCause `json:"build_cause"`
	CanRun       bool       `json:"can_run"`
	Name         string     `json:"name"`
	NaturalOrder int        `json:"natural_order"`
	Comment      string     `json:"comment"`
	Stages       []Stage    `json:"stages"`
}

// BuildCause describes the triggers which caused the build to start.
type BuildCause struct {
	Approver          string             `json:"approver,omitempty"`
	MaterialRevisions []MaterialRevision `json:"material_revisions"`
	TriggerForced     bool               `json:"trigger_forced"`
	TriggerMessage    string             `json:"trigger_message"`
}

// MaterialRevision describes the uniquely identifiable version for the material which was pulled for this build
type MaterialRevision struct {
	Modifications []Modification `json:"modifications"`
	Material      struct {
		Description string `json:"description"`
		Fingerprint string `json:"fingerprint"`
		Type        string `json:"type"`
		ID          int    `json:"id"`
	} `json:"material"`
	Changed bool `json:"changed"`
}

// Modification describes the commit/revision for the material which kicked off the build.
type Modification struct {
	EmailAddress string `json:"email_address"`
	ID           int    `json:"id"`
	ModifiedTime int    `json:"modified_time"`
	UserName     string `json:"user_name"`
	Comment      string `json:"comment"`
	Revision     string `json:"revision"`
}

// GetHistory returns a list of pipeline instanves describing the pipeline history.
func (pgs *PipelinesService) GetHistory(ctx context.Context, name string, offset int) (*PipelineHistory, *APIResponse, error) {
	stub := fmt.Sprintf("pipelines/%s/history", name)
	if offset > 0 {
		stub = fmt.Sprintf("%s/%d", stub, offset)
	}

	req, err := pgs.client.NewRequest("GET", stub, nil, "")
	if err != nil {
		return nil, nil, err
	}

	pt := PipelineHistory{}
	resp, err := pgs.client.Do(ctx, req, &pt)
	if err != nil {
		return nil, resp, err
	}

	return &pt, resp, nil
}
