package enum

type StatusTransaction int

const (
	approved StatusTransaction = iota + 1
	rejected
)

func (status StatusTransaction) String() string {
	return [...]string{"approved", "rejected"}[status-1]
}

func (status StatusTransaction) EnumIndex() int  {
	return int(status)
}