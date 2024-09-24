package slackbuildkit

import "github.com/slack-go/slack"

type MixedElement interface {
	AsMixedElement() slack.MixedElement
}

type BlockElement interface {
	AsBlockElement() slack.BlockElement
}
