package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/anhoder/foxful-cli/model"
	"github.com/anhoder/foxful-cli/util"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

var (
	mainMenu = NewMainMenu()
	subMenu  = NewSubMenu()
)

type MainMenu struct {
	model.DefaultMenu
	menus []model.MenuItem
}

func NewMainMenu() *MainMenu {
	m := &MainMenu{}
	m.menus = []model.MenuItem{
		{Title: "Main Menu 1", Subtitle: "Sub Title"},
		{Title: "Main Menu 2", Subtitle: "very long long long long long long long long long long long long subtitle"},
		{Title: "Main Menu 3"},
		{Title: "Main Menu 4"},
		{Title: "Main Menu 5"},
	}

	return m
}

func (m *MainMenu) GetMenuKey() string {
	return "main_menu"
}

func (m *MainMenu) MenuViews() []model.MenuItem {
	return m.menus
}

func (m *MainMenu) SubMenu(_ *model.App, index int) model.Menu {
	if index >= len(m.menus) {
		return nil
	}
	return subMenu
}

type SubMenu struct {
	model.DefaultMenu
	menus []model.MenuItem
}

func NewSubMenu() *SubMenu {
	m := &SubMenu{}
	m.menus = []model.MenuItem{
		{Title: "Sub Menu 1"},
		{Title: "Sub Menu 2"},
		{Title: "Sub Menu 3"},
	}

	return m
}

func (m *SubMenu) GetMenuKey() string {
	return "sub_menu"
}

func (m *SubMenu) MenuViews() []model.MenuItem {
	return m.menus
}

func (m *SubMenu) SubMenu(_ *model.App, _ int) model.Menu {
	return nil
}

func (m *SubMenu) BeforeEnterMenuHook() model.Hook {
	return func(main *model.Main) bool {
		time.Sleep(time.Millisecond * 200)
		return true
	}
}

type Component1 struct {
	app *model.App
}

func NewComponent1(app *model.App) *Component1 {
	p := &Component1{
		app: app,
	}

	return p
}

func (p *Component1) Update(_ tea.Msg, _ *model.App) {
	return
}

func (p *Component1) View(_ *model.App, main *model.Main) (string, int) {
	var (
		builder strings.Builder
		t       = time.Now()
	)
	builder.WriteString(strings.Repeat(" ", main.MenuStartColumn()))
	builder.WriteString(util.SetFgStyle("line1: "+strconv.Itoa(t.Hour())+"h\n", termenv.ANSIBrightBlue))
	builder.WriteString(strings.Repeat(" ", main.MenuStartColumn()))
	builder.WriteString(util.SetFgStyle("line2: "+strconv.Itoa(t.Minute())+"m\n", termenv.ANSIBrightCyan))
	builder.WriteString(strings.Repeat(" ", main.MenuStartColumn()))
	builder.WriteString(util.SetFgStyle("line3: "+strconv.Itoa(t.Second())+"s", termenv.ANSIBrightYellow))

	return builder.String(), 3
}

func main() {
	var (
		ops       = model.DefaultOptions()
		app       = model.NewApp(ops)
		component = NewComponent1(app)
	)
	ops.MainMenu = mainMenu
	ops.Components = []model.Component{component}
	ops.Ticker = model.DefaultTicker(time.Millisecond * 500)

	fmt.Println(app.Run())
}
