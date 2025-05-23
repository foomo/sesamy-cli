package tag

type Type string

const (
	TypeAlias    Type = "alias"
	TypeAppend   Type = "append"
	TypeGroup    Type = "group"
	TypeIdentify Type = "identify"
	TypeReset    Type = "reset"
	TypeSet      Type = "set"
	TypeSetOnce  Type = "set-once"
	TypeTrack    Type = "track"
)

func (t Type) String() string {
	return string(t)
}
