package three

import (
	. "github.com/myesui/uuid"
	"time"
	"github.com/myesui/uuid/kit"
)

var _ Service = &instrumentingMiddleware{}

type instrumentingMiddleware struct {
	kit.InstrumentingMiddleware
}

func (o instrumentingMiddleware) UUID(namespace Implementation, names ...interface{}) (id UUID) {
	defer func(begin time.Time) {
		o.Log(begin, "method", "uuid", "error", "false")
	}(time.Now())
	id = o.Next().(Service).UUID(namespace, names...)
	return
}

func (o *instrumentingMiddleware) Add(service kit.Kit) Service {
	return kit.AddMiddleware(service, o).(Service)
}

func (instrumentingMiddleware) String() string {
	return "uuid-three-instrumenting"
}

func NewInstrumentingMiddleware() kit.Kit {
	namespace := "uuid"
	subsystem := "three"

	return &instrumentingMiddleware{
		kit.NewInstrumentingMiddleware(namespace, subsystem),
	}
}
