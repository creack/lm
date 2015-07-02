package log

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
