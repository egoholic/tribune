package actor

type Actor struct {
	messageChannel    chan interface{}
	finalizingChannel chan bool
	function          func(interface{}) error
}

func (a *Actor) Run() {
	go func() {
		var m interface{}

		select {
		case m = <-a.messageChannel:
			a.function(m)
		case <-a.finalizingChannel:
			return
		}
	}()
}
