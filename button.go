package slackbuildkit

import "github.com/slack-go/slack"

type Button struct {
	btn *slack.ButtonBlockElement
}

func NewButton(id, value, text string) *Button {
	return &Button{
		btn: slack.NewButtonBlockElement(
			id,
			value,
			slack.NewTextBlockObject(slack.PlainTextType, text, false, false),
		),
	}
}

func (b *Button) WithURL(url string) *Button {
	b.btn = b.btn.WithURL(url)
	return b
}

func (b *Button) AsBlockElement() slack.BlockElement {
	return b.btn
}
