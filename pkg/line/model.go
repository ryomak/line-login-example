package line

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

type Event struct {
	*linebot.Event
}

func toEvent(es []*linebot.Event) []*Event {
	events := make([]*Event, len(es))
	for i, v := range es {
		events[i] = &Event{v}
	}
	return events
}
