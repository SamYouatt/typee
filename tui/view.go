package tui

import (
	"github.com/SamYouatt/typee/features/stats"
	"github.com/SamYouatt/typee/features/menu"
	"github.com/SamYouatt/typee/util"
)

var colours = util.AppColours


func (m Model) View() string {
	switch m.state {
	case Ready:
		return menu.View(m.width, m.height)
	case InTest:
		return m.test.View(m.width, m.height)
	case TestComplete:
		return stats.View(m.result, m.width, m.height)
	}

	return ""
}
