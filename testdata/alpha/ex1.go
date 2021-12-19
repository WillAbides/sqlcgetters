package alpha

import (
	"io"
	"time"
	tm "time"
	tm2 "time"
)

type Foo struct {
	io.Reader
	Bar struct {
		Baz string `json:"baz"`
	}
	CreatedAt   tm.Time  `json:"created_at"`
	UpdatedAt   tm2.Time `json:"updated_at"`
	DeletedAt   time.Time
	Qux         *string
	GetQux      string
	GetQux_     int32
	GetQux__    int64
	GetGetQux   uint32
	GetGetQux_  uint64
	GetGetQux__ interface{}
	Func        func(string) func(string) func(string) func(string) (*string, error) // hi
}

var _ WantFoo = Foo{}

type WantFoo interface {
	GetBar() struct {
		Baz string `json:"baz"`
	}
	GetCreatedAt() tm.Time
	GetUpdatedAt() tm2.Time
	GetDeletedAt() time.Time
	GetQux___() *string
	GetGetQux___() string
	GetGetQux____() int32
	GetGetQux_____() int64
	GetGetGetQux() uint32
	GetGetGetQux_() uint64
	GetGetGetQux__() interface{}
	GetFunc() func(string) func(string) func(string) func(string) (*string, error)
}
