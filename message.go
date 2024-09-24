package slackbuildkit

import (
	"context"

	"github.com/slack-go/slack"
)

type Message struct {
	options []slack.MsgOption
	blocks  []slack.Block
}

func (m *Message) InThread(ts string) *Message {
	m.options = append(m.options, slack.MsgOptionTS(ts))
	return m
}

func (m *Message) Broadcast() *Message {
	m.options = append(m.options, slack.MsgOptionBroadcast())
	return m
}

func (m *Message) AddContext(id string, elems ...MixedElement) {
	es := make([]slack.MixedElement, len(elems))
	for i, e := range elems {
		es[i] = e.AsMixedElement()
	}
	m.blocks = append(m.blocks, slack.NewContextBlock(id, es...))
}

func (m *Message) AddSection(t *Text) *Message {
	m.blocks = append(m.blocks, slack.NewSectionBlock(t.AsMixedElement().(*slack.TextBlockObject), nil, nil))
	return m
}

func (m *Message) AddActions(id string, actions ...BlockElement) *Message {
	as := make([]slack.BlockElement, len(actions))
	for i, a := range actions {
		as[i] = a.AsBlockElement()
	}
	m.blocks = append(m.blocks, slack.NewActionBlock(id, as...))
	return m
}

func (m *Message) Send(ctx context.Context, api *slack.Client, channel string) (string, error) {
	options := append(m.options, slack.MsgOptionBlocks(m.blocks...))
	_, ts, _, err := api.SendMessageContext(ctx, channel, options...)

	return ts, err
}
