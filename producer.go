package lm

import (
	"fmt"
	"strings"
)

// DefaultProducer used by global functions.
var DefaultProducer = &MessageProducer{
	Formatter:      DefaultFormatter,
	StringFormat:   DefaultStringFormat,
	ExtraCalldepth: 1,
	DefaultLevel:   LevelInfo,
}

// MessageProducer produces messages
// based on the given configuration.
type MessageProducer struct {
	Formatter      func(*Message) string
	StringFormat   string
	ExtraCalldepth int
	DefaultLevel   Level
	DefaultContext string
}

// NewProducer creates a new Message producer.
func NewProducer() *MessageProducer {
	return &MessageProducer{
		Formatter:      DefaultFormatter,
		StringFormat:   DefaultStringFormat,
		ExtraCalldepth: 0,
		DefaultLevel:   LevelDebug,
	}
}

// SetFormatter sets the producer's formatter.
func (p *MessageProducer) SetFormatter(formatter func(*Message) string) *MessageProducer {
	p.Formatter = formatter
	return p
}

// SetStringFormat sets the producer's string format used by the default Formatter.
func (p *MessageProducer) SetStringFormat(stringFormat string) *MessageProducer {
	p.StringFormat = stringFormat
	return p
}

// SetExtraCalldepth sets the producer's extraCalldepth.
func (p *MessageProducer) SetExtraCalldepth(depth int) *MessageProducer {
	p.ExtraCalldepth = depth
	return p
}

// SetDefaultLevel sets the default level to be applied to produced messages.
func (p *MessageProducer) SetDefaultLevel(level Level) *MessageProducer {
	p.DefaultLevel = level
	return p
}

// SetDefaultContext sets the default context to be applied to produced mesages.
func (p *MessageProducer) SetDefaultContext(context string) *MessageProducer {
	p.DefaultContext = context
	return p
}

func (p *MessageProducer) newMessagef(value string) *Message {
	return &Message{
		Value:          value,
		Formatter:      p.Formatter,
		StringFormat:   p.StringFormat,
		ExtraCalldepth: p.ExtraCalldepth,
		Level:          p.DefaultLevel,
		Context:        p.DefaultContext,
		producer:       p,
	}
}

// NewMessage creates a new message.
func (p *MessageProducer) NewMessage(args ...interface{}) *Message {
	m := p.newMessagef(fmt.Sprint(args...))
	m.ShortFile = lookupCaller(p.ExtraCalldepth)
	return m
}

// NewMessagef creates a new message.
func (p *MessageProducer) NewMessagef(f string, args ...interface{}) *Message {
	m := p.newMessagef(fmt.Sprintf(f, args...))
	m.ShortFile = lookupCaller(p.ExtraCalldepth)
	return m
}

// NewError creates a new message with the given error.
// If the error is a message, it becomes the child of a new Message.
// NOTE: disregard the DefaultLevel and sets LevelError unless error is a message.
func (p *MessageProducer) NewError(err error) *Message {
	m := p.newMessagef("")
	m.ShortFile = lookupCaller(p.ExtraCalldepth)
	if msg, ok := err.(*Message); ok {
		m.Level = msg.Level
		m.Child = msg
		if msg.Context != "" {
			m.Context = msg.Context
		} else {
			m.Context = p.DefaultContext
		}
		m.Set(msg)
	} else {
		m.Level = LevelError
		m.Context = p.DefaultContext
		m.Value = err.Error()
	}
	return m
}

// Default format string
const DefaultStringFormat = `%s | %s | %d | %s | %s | ["%s"]`

// DefaultFormatter formats the message for output.
// Containers:
// - the current message LogLevel
// - the current shortFile info
// - the first non empty code in the tree
// - the first non empty name in the tree
// - the first non empty context int the tree
// - An array of all the values in the tree
func DefaultFormatter(m *Message) string {
	var (
		code         int
		name         string
		context      string
		stringFormat string
		values       = []string{}
	)
	for msg := m; msg != nil; msg = msg.Child {
		if code == 0 {
			code = msg.Code
		}
		if name == "" {
			name = msg.Name
		}
		if context == "" {
			context = msg.Context
		}
		if msg.Value != "" {
			values = append(values, msg.Value)
		}
	}
	if m.StringFormat == "" {
		stringFormat = DefaultStringFormat
	} else {
		stringFormat = m.StringFormat
	}
	return fmt.Sprintf(stringFormat,
		m.Level, m.ShortFile, code, name, context, strings.Join(values, `","`))
}
