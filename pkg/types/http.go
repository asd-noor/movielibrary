package types

type HttpRequest struct {
	Body        interface{}
	Headers     map[string]string
	QueryParams map[string]string
	Method      string
	Endpoint    string
	ContentType string
	AuthToken   string
}
