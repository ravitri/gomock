package action

type Action interface {
	DoAction(int, string) error
}
