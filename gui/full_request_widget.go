package gui

import (
    "fmt"
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/kc8/active-reception/requests"
	"github.com/kc8/active-reception/states"
)

func BuildBasicRequestUI(state states.ProgramState, requestKey string) {
    request, found := state.Requests[requestKey]
    if found && request != nil { // && request.Show 
        imgui.SetNextWindowSizeV(imgui.NewVec2(300, 300), imgui.CondOnce)
        imgui.Begin(request.UiLabel)
        /*imgui.InputTextWithHint("URI", "", &request.RequestInfo.Uri, 0, nil)
        if imgui.ButtonV("Send", imgui.NewVec2(80, 20)) {
            state.Requests[requestKey].RequestInfo.Response = requests.SendGetRequset(request.RequestInfo.Uri)
        }
        if request.RequestInfo.Response != "" {
            imgui.Text(request.RequestInfo.Response)
        }*/
        imgui.End()
    }
}

func MakeBasicRequestUI(state states.ProgramState) bool {
    var name string
    valid := false

    imgui.SetNextWindowSizeV(imgui.NewVec2(300, 300), imgui.CondOnce)
    imgui.Begin("Add New Request Widget")
    callback := func(data imgui.InputTextCallbackData) int {
        name = data.Buf()
        fmt.Println(data.Buf())
        return 0
    }
    // TODO we are not getting the value out of our text block it seems to always be ""
    imgui.InputTextWithHint("Name", "", &name, 0, callback)
    if imgui.ButtonV("Create", imgui.NewVec2(80, 20)) {
        result := new(states.RequestState)
        result.RequestInfo = states.Request{
            Uri: "",
            Method: requests.GET,
        }
                                            // TODO the underlying type here is string
        result.RequestInfo.PerformRequest = requests.GetCorrectRequsetFunc(string(result.RequestInfo.Method)) 
        result.UiLabel = name
        // TODO we should not be indexing w/ a key like this from user input
        if result.UiLabel == "" {
            fmt.Println(name)
            result.UiLabel = "TEMP NAME"
        }
        state.Requests[name] = result
        valid = true
    }
    imgui.End()
    
    return valid
}

func NewBasicRequestWindow(state states.ProgramState) {
    name := "test"
    uri := "https://cat-fact.herokuapp.com/facts/"
    state.Requests[name] = &states.RequestState{
        UiLabel: name,
        RequestInfo: states.Request{
            Uri: uri,
            Method: requests.GET,
            PerformRequest: requests.GetCorrectRequsetFunc(requests.GET),
        },
    }
}
