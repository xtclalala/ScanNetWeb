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
	State  TaskState = 0
	Build            = State + 1
	Ready            = Build + 1
	Doing            = Ready + 1
	Finish           = Doing + 1
)
