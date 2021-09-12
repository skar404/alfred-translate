package global

import aw "github.com/deanishe/awgo"

type AppFlags struct {
	SetKey  string
	GetKey  string
	Value   string
	Command string
}

var (
	WF   *aw.Workflow
	Flag AppFlags
)
