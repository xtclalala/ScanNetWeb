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

// 任务类型

const (
	LinuxScan = "LinuxScan"
)

const (
	Def     = "default"
	Error   = "error"
	Info    = "info"
	Success = "success"
	Warning = "warning"
)
