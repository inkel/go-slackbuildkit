package slackbuildkit

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
)

type Text struct {
	strings.Builder

	elemType string
}

func PlainText() *Text { return &Text{elemType: slack.PlainTextType} }

func Markdown() *Text { return &Text{elemType: slack.MarkdownType} }

func Markdownf(format string, args ...any) *Text {
	return Markdown().Printf(format, args...)
}

func (t *Text) Printf(format string, args ...any) *Text {
	// strings.Builder never returns error
	_, _ = fmt.Fprintf(t, format, args...)
	return t
}

func (t *Text) AsMixedElement() slack.MixedElement {
	return slack.NewTextBlockObject(t.elemType, t.String(), false, false)
}
