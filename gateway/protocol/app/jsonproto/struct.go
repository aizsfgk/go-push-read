package jsonproto

type Req struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type Msg struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Data string `json:"data"`
}
