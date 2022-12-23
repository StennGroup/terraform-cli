package jsonformat

import (
	"fmt"
	"strings"

	"github.com/mitchellh/colorstring"

	"github.com/hashicorp/terraform/internal/command/jsonplan"
	"github.com/hashicorp/terraform/internal/command/jsonprovider"
	"github.com/hashicorp/terraform/internal/terminal"
)

type Plan struct {
	OutputChanges   map[string]jsonplan.Change        `json:"output_changes"`
	ResourceChanges []jsonplan.ResourceChange         `json:"resource_changes"`
	ResourceDrift   []jsonplan.ResourceChange         `json:"resource_drift"`
	ProviderSchemas map[string]*jsonprovider.Provider `json:"provider_schemas"`
}

type Renderer struct {
	Streams  *terminal.Streams
	Colorize *colorstring.Colorize
}

type JSONLogType string
type JSONLog map[string]interface{}

const (
	LogVersion         JSONLogType = "version"
	LogPlannedChange   JSONLogType = "planned_change"
	LogRefreshStart    JSONLogType = "refresh_start"
	LogRefreshComplete JSONLogType = "refresh_complete"
	LogApplyStart      JSONLogType = "apply_start"
	LogApplyComplete   JSONLogType = "apply_complete"
	LogChangeSummary   JSONLogType = "change_summary"
	LogOutputs         JSONLogType = "outputs"
)

func (r Renderer) RenderPlan(plan Plan) {
	// panic("not implemented")
	r.Streams.Printf("boop renderered plan!")
}

func (r Renderer) RenderLog(log JSONLog) {
	msg, ok := log["@message"].(string)
	if !ok {
		return
	}

	switch JSONLogType(log["type"].(string)) {
	case LogApplyStart, LogApplyComplete, LogRefreshStart, LogRefreshComplete:
		s := strings.Split(msg, ":")
		msg = fmt.Sprintf("[bold]%s: %s", s[0], s[1])

		r.Streams.Print(r.Colorize.Color(msg))
		r.Streams.Print("\n")
	case LogChangeSummary:
		msg = fmt.Sprintf("[bold][green]%s", msg)

		r.Streams.Print("\n\n")
		r.Streams.Print(r.Colorize.Color(msg))
		r.Streams.Print("\n\n")
	}
}
