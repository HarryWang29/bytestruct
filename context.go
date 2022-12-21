package bytestruct

import (
	"github.com/HarryWang29/bytestruct/types"
	"sync"
)

type RuntimeContext struct {
	Buf    []byte
	Option *types.Option
}

var (
	runtimeContextPool = sync.Pool{
		New: func() interface{} {
			return &RuntimeContext{
				Option: &types.Option{},
			}
		},
	}
)

func TakeRuntimeContext() *RuntimeContext {
	return runtimeContextPool.Get().(*RuntimeContext)
}

func ReleaseRuntimeContext(ctx *RuntimeContext) {
	runtimeContextPool.Put(ctx)
}
