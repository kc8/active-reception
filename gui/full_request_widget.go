package gui

import (
	// "fmt"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/kc8/active-reception/requests"
	"github.com/kc8/active-reception/states"
)

func BuildBasicRequestUI(state states.ProgramState, requestKey string) {
    request, found := state.Requests[requestKey]
    if found { // && request.Show 
        imgui.SetNextWindowSizeV(imgui.NewVec2(300, 300), imgui.CondOnce)
        imgui.Begin(request.UiLabel)
        imgui.InputTextWithHint("URI", "", &request.RequestInfo.Uri, 0, nil)
            if imgui.ButtonV("Send", imgui.NewVec2(80, 20)) {
                state.Requests[requestKey].RequestInfo.Response = requests.SendGetRequset(request.RequestInfo.Uri)
            }
            if request.RequestInfo.Response != "" {
                imgui.Text(request.RequestInfo.Response)
            }
        imgui.End()
    } 
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
