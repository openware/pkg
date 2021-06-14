package models

type IPage interface {
	List() []IPage
	GetPath() string
	GetBody() string
	GetTitle() string
	GetDescription() string
}
