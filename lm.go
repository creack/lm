package lm

import "encoding/json"

// Message represent a logging message.
type Message struct {
	// Message content
	Code      int    `json:"code"`
	Name      string `json:"name"`
	Context   string `json:"context"`
	Value     string `json:"value"`
	Level     Level  `json:"level"`
	ShortFile string `json:"shortFile"`

	// Child message
	Child *Message `json:"child,omitempty"`

	// Extra data
	Formatter      func(*Message) string `json:"-"`
	StringFormat   string                `json:"-"`
	ExtraCalldepth int                   `json:"-"`

	producer *MessageProducer
}

// NewMessage instantiate a new message.
func NewMessage(args ...interface{}) *Message {
	return DefaultProducer.NewMessage(args...)
}

// NewMessagef instantiate a new message.
func NewMessagef(f string, args ...interface{}) *Message {
	return DefaultProducer.NewMessagef(f, args)
}

// NewError create a new message form the given error.
func NewError(err error) *Message {
	return DefaultProducer.NewError(err)
}

// Set updates all empty fields with the given message's ones.
func (m *Message) Set(in *Message) *Message {
	if m.Context == "" {
		m.Context = in.Context
	}
	if m.Code == 0 {
		m.Code = in.Code
	}
	if m.Name == "" {
		m.Name = in.Name
	}
	if m.Level == 0 {
		m.Level = in.Level
	}
	return m
}

// Error implement the Error interface to let the message to be passed via error return.
func (m *Message) Error() string {
	return m.String()
}

// String calls the user defined formatter or the default one if none.
func (m *Message) String() string {
	if m.Formatter == nil {
		return DefaultFormatter(m)
	}
	return m.Formatter(m)
}

// Dump creates a full json object from the message.
func (m *Message) Dump() ([]byte, error) {
	return json.MarshalIndent(m, "", "\t")
}

// SetName sets the message name.
func (m *Message) SetName(name string) *Message {
	m.Name = name
	return m
}

// SetCode sets the message code.
func (m *Message) SetCode(code int) *Message {
	m.Code = code
	return m
}

// SetContext sets the message context.
func (m *Message) SetContext(context string) *Message {
	m.Context = context
	return m
}

// SetLevel sets the message level.
func (m *Message) SetLevel(lvl Level) *Message {
	m.Level = lvl
	return m
}

// SetValue sets the message value.
func (m *Message) SetValue(value string) *Message {
	m.Value = value
	return m
}

// SetFormatter sets the formatter callback for the message.
func (m *Message) SetFormatter(formatter func(*Message) string) *Message {
	m.Formatter = formatter
	return m
}

// SetStringFormat sets the string format for the default froamtter.
func (m *Message) SetStringFormat(stringFormat string) *Message {
	m.StringFormat = stringFormat
	return m
}

// NewError wraps the producer's NewError.
func (m *Message) NewError(err error) *Message {
	if m.producer == nil {
		m.producer = DefaultProducer
	}
	m1 := m.producer.NewError(err)
	m1.ShortFile = lookupCaller(m.producer.ExtraCalldepth)
	m1.Child = m
	return m1.Set(m)
}

// NewMessage eraps the producer's NewMessage.
func (m *Message) NewMessage(args ...interface{}) *Message {
	if m.producer == nil {
		m.producer = DefaultProducer
	}
	m1 := m.producer.NewMessage(args...)
	m1.Child = m
	return m1.Set(m)
}

// NewMessagef eraps the producer's NewMessage.
func (m *Message) NewMessagef(f string, args ...interface{}) *Message {
	if m.producer == nil {
		m.producer = DefaultProducer
	}
	m1 := m.producer.NewMessagef(f, args...)
	m1.Child = m
	return m1.Set(m)
}
