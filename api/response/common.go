package response

type (
	Res struct {
		Error int         `json:"error"`
		Msg   string      `json:"msg"`
		Data  interface{} `json:"data"`
	}
)
