package server

type taskType string

const (
	taskOnline  taskType = "online"  // 用户上线任务
	taskChannel taskType = "channel" // 频道任务
	taskP2p     taskType = "p2p"     // 单聊
	taskStats   taskType = "stats"   // 统计
)
