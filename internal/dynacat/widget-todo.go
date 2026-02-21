package dynacat

import (
	"fmt"
	"html/template"
)

var todoWidgetTemplate = mustParseTemplate("todo.html", "widget-base.html")

type todoWidget struct {
	widgetBase `yaml:",inline"`
	cachedHTML template.HTML `yaml:"-"`
	TodoID     string        `yaml:"id"`
	Storage    string        `yaml:"storage"`
}

func (widget *todoWidget) initialize() error {
	widget.withTitle("To-do").withError(nil)

	if widget.Storage != "" && widget.Storage != "local" && widget.Storage != "server" {
		return fmt.Errorf("storage must be either \"local\" or \"server\", got %q", widget.Storage)
	}

	if widget.Storage == "server" && widget.TodoID == "" {
		return fmt.Errorf("storage \"server\" requires an \"id\" to be set")
	}

	widget.cachedHTML = widget.renderTemplate(widget, todoWidgetTemplate)
	return nil
}

func (widget *todoWidget) Render() template.HTML {
	return widget.cachedHTML
}
