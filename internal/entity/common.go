package entity

type Response struct {
	Code int    `default:"0"`
	Msg  string `default:"ok"`
}
