package association

type Associable interface {
	IsIterable() bool
	Value() interface{}
}

type Association struct {
	Name       string
	Left       Associable
	Right      Associable
	leftHooks  map[string]hook
	rightHooks map[string]hook
}

type hook func(Associable, Associable) bool

func (a *Association) LeftHook(k string) bool {
	return a.leftHooks[k](a.Left, a.Right)
}

func (a *Association) RightHook(k string) bool {
	return a.leftHooks[k](a.Left, a.Right)
}
