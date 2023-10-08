package main

import (
	"fmt"
	"image"
	"strconv"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/kc8/active-reception/gui"
	_ "github.com/kc8/active-reception/requests"
	"github.com/kc8/active-reception/states"
)

var (
	showDemoWindow bool
	value1         int32
	value2         int32
	value3         int32
	values         [2]int32 = [2]int32{value1, value2}
	content        string   = "Let me try"
	r              float32
	g              float32
	b              float32
	a              float32
	color4         [4]float32 = [4]float32{r, g, b, a}
	selected       bool
	backend        imgui.Backend[imgui.GLFWWindowFlags]
	oldimg         *image.RGBA
	texture        *imgui.Texture
	barValues      []int64
)

var globalStateOfApp = states.InitState()

func callback(data imgui.InputTextCallbackData) int {
	fmt.Println("got call back")
	return 0
}

func showWidgetsDemo(state states.ProgramState) {
	imgui.SetNextWindowSizeV(imgui.NewVec2(300, 300), imgui.CondOnce)
	imgui.Begin("Window 1")
	if imgui.ButtonV("Click Me", imgui.NewVec2(80, 20)) {
		//w, h := state.Gui.Backend.DisplaySize()
		// fmt.Println(w, h)
		state.Clicks += 1
		state.DebugMessage(states.DEBUG, strconv.FormatInt(int64(state.Clicks), 10))
	}
	imgui.TextUnformatted("Unformatted text")
	imgui.Checkbox("Show demo window", &showDemoWindow)
	if imgui.BeginCombo("Combo", "Combo preview") {
		imgui.SelectableBoolPtr("Item 1", &selected)
		imgui.SelectableBool("Item 2")
		imgui.SelectableBool("Item 3")
		imgui.EndCombo()
	}

	if imgui.RadioButtonBool("Radio button1", selected) {
		selected = true
	}

	imgui.SameLine()

	if imgui.RadioButtonBool("Radio button2", !selected) {
		selected = false
	}

	// This is currently bugged in lastest version
	imgui.InputTextWithHint("Name", "write your name here", &content, 0, callback)
	imgui.InputTextWithHint("Name", "write your name here", &content, 0, callback)
	imgui.Text(content)
	imgui.SliderInt("Slider int", &value3, 0, 100)
	imgui.DragInt("Drag int", &value1)
	imgui.DragInt2("Drag int2", &values)
	value1 = values[0]
	imgui.ColorEdit4("Color Edit3", &color4)

	tableWindowsCordsPos := imgui.WindowPos()
	//printable := fmt.Sprintf("Window Pos: (%f, %f)", tableWindowsCordsPos.X, tableWindowsCordsPos.Y)

	state.WindowPositionX = tableWindowsCordsPos.X
	state.WindowPositionY = tableWindowsCordsPos.Y
	//state.DebugMessage(states.DEBUG, printable)

	imgui.End()
}

func drawTable(state states.ProgramState) {
	imgui.SetNextWindowSizeV(imgui.NewVec2(300, 300), imgui.CondOnce)
	//var names = []string {"One", "Two", "Three"}

	const tableFlags = imgui.TableFlagsBorders

	imgui.Begin("Table")
	if imgui.BeginTable("Table", 3) {
		imgui.TableSetupColumnV("One", tableFlags, 0.0, 0)
		imgui.TableSetupColumnV("Two", tableFlags, 0.0, 0)
		imgui.TableSetupColumnV("Three", tableFlags, 0.0, 0)
		imgui.TableHeadersRow()
		for r := 0; r < 3; r++ {
			imgui.TableNextRow()
			var c int32 = 0
			for c = 0; c < 3; c++ {
				imgui.TableSetColumnIndex(c)
				rowCol := fmt.Sprintf("(%d, %d)", r, c)
				imgui.Text(rowCol)
			}
		}
		imgui.EndTable()
	}

	//tableWindowsCordsPos := imgui.WindowPos()
	//printable := fmt.Sprintf("Table Pos: (%f, %f)", tableWindowsCordsPos.X, tableWindowsCordsPos.Y)
	//state.DebugMessage(states.DEBUG, printable)
	// TODO how can we draw a line between one 'window' to another
	// imgui.ForegroundDrawListViwewportPtr().AddLine(tableWindowsCordsPos, imgui.NewVec2(state.windowPositionX, state.windowPositionY), 1)

	imgui.End()
	//imgui.WindowDrawList().addL
}

func afterCreateContext() {
}

func loop() {
	showWidgetsDemo(globalStateOfApp)
	drawTable(globalStateOfApp)
	gui.BuildBasicRequestUI(globalStateOfApp, "test")
	gui.BuildBasicRequestUI(globalStateOfApp, "another")
}

func beforeDestroyContext() {
	//imgui.PlotDestroyContext()
}

func main() {
	globalStateOfApp.Gui = states.GuiState{
		Backend: imgui.CreateBackend(imgui.NewGLFWBackend()),
	}
	globalStateOfApp.Gui.Backend = imgui.CreateBackend(imgui.NewGLFWBackend())

	globalStateOfApp.Gui.Backend.SetAfterCreateContextHook(afterCreateContext)
	globalStateOfApp.Gui.Backend.SetBeforeDestroyContextHook(beforeDestroyContext)

	globalStateOfApp.Gui.Backend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))
	globalStateOfApp.Gui.Backend.CreateWindow("Active Reception", 1200, 900)

	globalStateOfApp.Gui.Backend.SetDropCallback(func(p []string) {
		//fmt.Printf("drop triggered: %v", p)
	})

	globalStateOfApp.Gui.Backend.SetCloseCallback(func(b imgui.Backend[imgui.GLFWWindowFlags]) {
		//	fmt.Println("window is closing")
	})
	gui.NewBasicRequestWindow(globalStateOfApp)

	globalStateOfApp.Gui.Backend.Run(loop)
}
