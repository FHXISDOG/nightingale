// Package crawler provides ...
package crawler

type HttpBody interface{}

type HttpInitMsg struct {
	Url  string
	Body HttpBody
}
