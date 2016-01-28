package tree_socket

var (
	global_socket_events = make(map[string][]func(interface{}))
)

type TreeSockEvent struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func init() {
	global_socket_events["connection"] = make([]func(interface{}), 0)
	global_socket_events["error"] = make([]func(interface{}), 0)
}

func On(name string, f func(interface{})) {
	if _, ok := global_socket_events[name]; !ok {
		global_socket_events[name] = make([]func(interface{}), 0)
	}
	global_socket_events[name] = append(global_socket_events[name], f)
}

func Trigger(name string, data interface{}) {
	if callbacks, ok := global_socket_events[name]; ok {
		for _, cb := range callbacks {
			go cb(data)
		}
	}
}

func (sock *Socket) On(name string, f func(interface{})) {
	if sock.events == nil {
		sock.events = make(map[string][]func(interface{}))
	}

	if _, ok := sock.events[name]; !ok {
		sock.events[name] = make([]func(interface{}), 0)
	}
	sock.events[name] = append(sock.events[name], f)
}

func (sock *Socket) Trigger(name string, data interface{}) {
	if callbacks, ok := sock.events[name]; ok {
		for _, cb := range callbacks {
			go cb(data)
		}
	}
}
