package mqparams

// req: mq --> pss
// ntf: pss --> mq
const (
	InvalidMqType = "InvalidDaemonSysMqType"

	ReqTaskReqMq = "backStageToDispatchingTaskQueue"
	//数据传输Dt(机器人固件升级)队列
	ReqDtRobotsMq    = "backStageToDispatchRobotUpdate"
	ReqBrakeRobotsMq = "backagetoDispTiaoshiOp"

	DownloadRobotLogMq = "backStageToDispatchDownloadRobotLogQueue"

	//调度升级队列
	ReqPssUpgradeMq = "backStageToDispatchDispatchUpdate"
	//后台要求停、充电
	ReqChargeRobotsMq              = "robotManagmentToRobotDispatchChargeQueue"
	ReqOperateRobotNtfMq           = "robotManagmentToRobotDispatchRobotStatusQueue"
	ReqUpdateRobotPowerNtfMq       = "robotManagementToDispachingUpdatePowerQueue"
	ReqChgGatewayTypeMq            = "passagewayToDispachingUpTypeTaskQueue"
	ReqChgGatewayStatusMq          = "passagewayToDispachingUpStatusTaskQueue"
	ReqChgNodeTypeOrStatusMq       = "backStageToChangeMapDetailStatusOrtypeQueue"
	ReqPauseStartInitPssMq         = "backStageToDispatchingHallSuspension"
	ReqFineRobotPositionMq         = "backStageToDispatchingrobotFineTuning"
	ReqAsk4AddingRobotOrTrayMq     = "backStageToDispatchingAskAddRobotOrTray"
	ReqCancelOfAddingRobotOrTrayMq = "backStageToDispatchingNotAddRobotOrTray"
	ReqTransChargePileMsg2RobotMq  = "BackStagedToDispatchingChargingInformationQueue"
	//后台告知停车场状况
	ReqTellPssConditionMq = "backStageToDispatchingOperatePointQueue"
	// 后台回复调度: xxx入口的空托盘当前是空闲的（没有车主在往上面放置车辆)
	RspEnNotBeingPlaced = "backToDispatchingFetchHaveEntranceQueue"

	//后台通知调度已将机器人放回路径

	/*规划全局取车路线给后台*/
	NtfRouteMq = "dispatchingToBackStageTaskQueue"

	/*车辆经过的点实时传输给后台*/
	NtfRobotPassedPointMq = "dispatchingToBackStageTaskPointQueue"

	/*机器人故障，通知后台并弹出故障警告*/
	NtfRobotAccMq = "robotDispatchToRobotManagmentRobotFaultQueue"

	/*充电桩将信息发送给后台，后台转发数据给调度*/
	//BackStagedToPssInfoMq = "BackStagedToDispatchingChargingInformationQueue"

	//调度发给后台确认充电
	NtfRobotChargeStartEndMq = "robotDispatchToRobotManagmentRobotChargeStartOrEndQueue"

	/*机器人来到入口，对接装载完成之后，托盘和车离开入口，通知辅助系统*/
	NtfCarLeaveEntranceOfParkMq = "dispatchingToAuxiliaryConfirmParkingLeaveTaskQueue"

	//调度向后台发送任务失败警告队列
	NtfStageTaskErrorAlarmMq = "dispatchingToBackStageTaskErrorAlarmQueue"

	NtfdispatchToBackStageFaileOrSucMq = "dispatchToBackStageFaileOrSuc"
	//调度内部使用
	NtfInnerTaskErrorMq = "innerDispatchErrorTaskQueue"

	// 调度询问后台: 下面这些入口中，选一个当前空闲的（没有车主在往上面放置车辆)
	EnNotBeingPlacedAsk = "dispatchingToBackFetchHaveEntranceQueue"

	//故障处理队列
	ErrorHandlingReqTaskQueue = "backStageToDispatchingErrorDealQueue"

	//暂停调度系统（非必须）
	SuspendSystemTaskQueue = "backStageToDispatchingPauseSystemQueue"

	NtfChgNodeSituationMq         = "dispatchingOperatePointTobackStageQueue"
	NtfCarCancelParkMq            = "dispatchingToAuxiliaryCancelParkingUnloadTaskQueue"
	NtfCarUnloadedMq              = "dispatchingToAuxiliaryPickUpCarUnloadTaskQueue"
	NtfPickUpCarTimeoutMq         = "dispatchingToAuxiliaryPickUpCarLeaveTaskQueue"
	NtfPickupItemsUnloadedMq      = "dispatchingToAuxiliaryPickUpFetchUnloadTaskQueue"
	NtfPickupItemsLoadedMq        = "dispatchingToAuxiliaryPickUpFetchLoadingTaskQueue"
	NtfPickupItemsLeaveMq         = "dispatchingToAuxiliaryPickUpFetchLeaveTaskQueue"
	Ntf2InitGatewayMq             = "dispatchingToAuxiliaryMissionCompleteTaskQueue"
	NtfGatewayTypeChgedMq         = "dispatchingToPassagewayUpTypeTaskQueue"
	NtfGatewayStatusChgedMq       = "dispatchingToPassagewayUpStatusTaskQueue"
	NtfAllRobotsPausedMq          = "dispatchingToBackStageAtOnceStopQueue"
	NtfRobotTuneFinedMq           = "dispatchingToPassagewayRobotFineTuning"
	NtfRsp4AddingRobotOrTrayAskMq = "dispatchingToPassagewayAddRobotOrTray"
	NtfTransBatteryInfoMq         = "dispatchingToBackStageChargingInformationQueue"
	//故障队列
	NtfRobotFaultDetailsMq = "dispatchingToRobotFaultDetailsQueue"

	// 故障后，调度通知后台机器人自检结果
	NtfBackRobotSelfCheckMq = "dispatchingSelfCheckingResultQueue"

	//todo: 消息队列的名字还未确认
	//AskBackIfRobotRepairedMq = "dispatchingToRobotFaultDetailsQueue"
	NtfSystemSuspendedDoneMq = "dispatchingToBackStageSuspendedDoneQueue"

	//数据传输升级的结果队列
	NtfBackRobotDtResultMq = "dispatchingToBackStageDataTransferResultQueue"

	// 后台通知调度升级
	NtfPss4UpgradeMq = "dispatchToBackageDisUpdate"
	//通知后台机器人状态
	NtfRobotStatusMq = "dispToBackageTiaoshiAns"

	//回给回台全场暂停或恢复的结果
	NtfStagePauseStartInitPssResultMq = "dispatchingToBackStageHallSuspensionResult"

	//调度挂掉之后通知后台消息队列
	NtfToBackstageDspDown = "dispatchToBackstageDspDown"
)

//type TaskReqType int
//
//const (
//	InvalidTaskReq                 TaskReqType = iota
//	ParkTaskReq
//	PickUpCarTaskReq
//	CancelParkTaskReq
//	CleanBlockCarTaskReq
//	MoveEmptyTrayTaskReq
//	ChargeTaskReq
//	StopChargeTaskReq
//	PickUpItemsInCarTaskReq
//	PickUpCarTimeOutTaskReq
//	PickUpItemsInCarTimeOutTaskReq
//)
