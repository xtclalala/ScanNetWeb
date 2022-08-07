package constant

// path
const (
	Slash      = "/"
	Point      = "."
	Rung       = "-"
	ConfigType = "yaml"
	ConfigName = Point + Slash + "conf" + Slash + "config"
	ConfigPath = ConfigName + Point + ConfigType
)

type TaskState = int64

const (
	_ TaskState = iota
	Build
	Ready
	Doing
	Finish
)

type TaskType = int64

const (
	SSH TaskType = iota
)
