package mqparams

// req: mq --> pss
// ntf: pss --> mq
const (
	InvalidMqType              = "InvalidDaemonSysMqType"
	ReqTaskReqMq               = "backStageToDispatchingTaskQueue"
	ReqDtRobotsMq              = "backStageToDispatchRobotUpdate"
	ReqAsk4AddingRobotOrTrayMq = "backStageToDispatchingAskAddRobotOrTray"
	//调度挂掉之后通知后台消息队列
	NtfToBackstageDspDown = "dispatchToBackstageDspDown"
)
