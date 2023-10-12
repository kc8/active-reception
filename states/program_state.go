package states

import (
	"fmt"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/kc8/active-reception/requests"
)

type DEBUG_TYPE int32

const (
	DEBUG = 0
	ERROR = 1
	VERBOSE = 2
)

type debug func(debugType DEBUG_TYPE, msg string)

// That state of the GUI
type GuiState struct {
	Backend        imgui.Backend[imgui.GLFWWindowFlags]
}

type ProgramState struct {
	Clicks       int
	WriteToDebug debug
	// TODO better data structure
	WindowPositionX float32
	WindowPositionY float32

    Gui GuiState

    Requests map[string]*RequestState
}

// A single requests state
type Request struct {
    Uri string
    RequestType string // TODO type this and make more robust
    Response string // This might have to carry more info in like code, headesr etc body
    Method requests.Method
    PerformRequest requests.PerformRequest
}

// A rolled up request state that contains requestInfo GUI info etc
type RequestState struct {
    UiLabel string
    ShowUi string
    RequestInfo Request
}

func InitState() ProgramState {
    requests := make(map[string]*RequestState)
    result := ProgramState{
        Requests: requests,
    }
    return result;
}

func (ps ProgramState) DebugMessage(debugType DEBUG_TYPE, msg string) {
	if ps.WriteToDebug != nil {
		ps.WriteToDebug(debugType, msg)
	} else {
        switch debugType {
        case DEBUG: 
            fmt.Printf("[DEBUG], %s\n", msg);
        case ERROR:
            fmt.Printf("[ERROR], %s\n", msg);
        case VERBOSE:
            fmt.Printf("[VERBOSE], %s\n", msg);
        }
	}
}
