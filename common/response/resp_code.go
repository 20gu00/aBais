package response

type respCode int64

//自定义的错误码
const (
	CodeSuccess           respCode = iota        // 0
	CodeInvalidParam      respCode = 1999 + iota // 2000
	CodeUserExist                                // 2001
	CodeUserNotExist                             // 2002
	CodeInvalidPassword                          // 2003
	CodeServerIntervalErr                        // 2004
	CodeNeedLogin                                // 2005
	CodeInvalidToken                             // 2006
	CodeTwoDevice                                // 2007

	CodeGetK8sClientErr // 2008

	CodeActionConfigErr

	// pod
	CodeGetPodListErr // 2009
	CodeGetPodDetailErr
	CodeDeletePodErr
	CodeUpdatePodErr
	CodePodContainerLogErr
	CodeGetNumByNsErr
	CodeGetPodContainerErr
	CodeCreatePodErr

	//deployment
	CodeGetDeploymentListrErr
	CodeGetDeploymentDetailErr
	CodeScaleDeploymentErr
	CodeDeleteDeploymentErr
	CodeRestartDeploymentErr
	CodeUpdateDeploymentErr
	CodeGetDeploymentPerNsErr
	CodeCreateDeploymentErr

	// daemonset
	CodeCreateDaemonsetErr
	CodeGetDaemonsetErr
	CodeGetDaemonsetDetailErr
	CodeDeleteDaemonsetErr
	CodeUpdateDaemonsetErr

	// statefulset
	CodeGetStatefulsetErr
	CodeGetStatefulsetDetailErr
	CodeDeleteStatefulsetErr
	CodeUpdateStatefulsetErr
	CodeCreateStatefulsetErr

	// service
	CodeGetSvcErr
	CodeGetSvcDetailErr
	CodeUpdateSvcErr
	CodeDeleteSvcErr
	CodeCreateSvcErr

	// ingress
	CodeGetIngressErr
	CodeGetIngressDetailErr
	CodeDeleteIngressErr
	CodeUpdateIngressErr
	CodeCreateIngressErr

	// configmap
	CodeGetCmErr
	CodeGetCmDetailErr
	CodeDeleteCmErr
	CodeUpdateCmErr
	CodeCreateCmErr

	// secret
	CodeGetSecretErr
	CodeGetSecretDetailErr
	CodeDeleteSecretErr
	CodeUpdateSecretErr
	CodeCreateSecretErr

	// pvc
	CodeGetPvcErr
	CodeGetPvcDetailErr
	CodeDeletePvcErr
	CodeUpdatePvcErr
	CodeCreatePvcErr

	// pv
	CodeGetPvErr
	CodeGetPvDetailErr
	CodeDeletePvErr
	CodeCreatePvErr

	// node
	CodeGetNodeErr
	CodeGetNodeDetailErr
	//CodeCreateNodeErr

	// ns
	CodeGetNsErr
	CodeGetNsDetailErr
	CodeDeleteNsErr
	CodeCreateNsErr

	// event
	CodeListEventErr

	// all resource
	CodeGetAllResourceNumErr

	// job
	CodeCreateJobErr
	CodeDeleteJobErr
	CodeUpdateJobErr
	CodeGetJobErr
	CodeGetJobDetailErr

	// helm
	CodeListReleaseErr
	CodeDetailReleaseErr
	CodeInstallReleaseErr
	CodeUninstallReleaseErr
	CodeAddChartErr
	CodeDeleteChartErr
	CodeListChartErr
	CodeDeleteReleaseErr
	CodeUpdateChartErr
	CodeDeleteChartFileErr
	CodeUploadChartFileErr
	CodeGetUploadMessageErr
)

var codeMsgMap = map[respCode]string{
	CodeSuccess:           "resp success",
	CodeInvalidParam:      "请求参数错误",
	CodeServerIntervalErr: "服务器内部错误", // 不返回确切的server报错给前端
	CodeUserExist:         "用户已存在",
	CodeUserNotExist:      "用户不存在",
	CodeInvalidPassword:   "用户名或密码错误",
	CodeNeedLogin:         "需要登录",
	CodeInvalidToken:      "无效的token",

	CodeGetK8sClientErr: "获取操作k8s的client失败",
	CodeActionConfigErr: "获取action的config失败",
	//pod
	CodeGetPodListErr:      "获取pod列表失败",
	CodeGetPodDetailErr:    "获取pod详情失败",
	CodeDeletePodErr:       "删除pod失败",
	CodeUpdatePodErr:       "更新pod失败",
	CodePodContainerLogErr: "获取pod中的容器日志失败",
	CodeGetNumByNsErr:      "根据ns获取pod数目失败",
	CodeGetPodContainerErr: "获取pod中的容器失败",
	CodeCreatePodErr:       "创建pod失败",

	// deployment
	CodeGetDeploymentListrErr:  "获取deployment列表失败",
	CodeGetDeploymentDetailErr: "获取deployment详情失败",
	CodeScaleDeploymentErr:     "调整deployment副本数目失败",
	CodeDeleteDeploymentErr:    "删除deployment失败",
	CodeRestartDeploymentErr:   "重启deployment失败",
	CodeUpdateDeploymentErr:    "更新deployment失败",
	CodeGetDeploymentPerNsErr:  "根据ns获取deployment失败",
	CodeCreateDeploymentErr:    "创建deployment失败",

	// daemonset
	CodeGetDaemonsetErr:       "获取daemonset失败",
	CodeGetDaemonsetDetailErr: "获取daemonset详情失败",
	CodeDeleteDaemonsetErr:    "删除daemonset失败",
	CodeUpdateDaemonsetErr:    "更新daemonset失败",
	CodeCreateDaemonsetErr:    "创建daemonset失败",

	// statefulset
	CodeGetStatefulsetErr:       "获取statefulset列表失败",
	CodeGetStatefulsetDetailErr: "获取statefulset详情失败",
	CodeDeleteStatefulsetErr:    "删除statefulset失败",
	CodeUpdateStatefulsetErr:    "更新statefulset失败",
	CodeCreateStatefulsetErr:    "创建statefulset失败",

	// service
	CodeGetSvcErr:       "获取svc失败",
	CodeGetSvcDetailErr: "获取svc详情失败",
	CodeUpdateSvcErr:    "更新svc失败",
	CodeDeleteSvcErr:    "删除svc失败",
	CodeCreateSvcErr:    "创建svc失败",

	// ingress
	CodeGetIngressErr:       "获取ingress失败",
	CodeGetIngressDetailErr: "获取ingress详情失败",
	CodeDeleteIngressErr:    "删除ingress失败",
	CodeUpdateIngressErr:    "更新ingress失败",
	CodeCreateIngressErr:    "创建ingress失败",

	// configmap
	CodeGetCmErr:       "获取cm失败",
	CodeGetCmDetailErr: "获取cm详情失败",
	CodeDeleteCmErr:    "删除cm失败",
	CodeUpdateCmErr:    "更新cm失败",
	CodeCreateCmErr:    "创建cm失败",

	// secret
	CodeGetSecretErr:       "获取secret失败",
	CodeGetSecretDetailErr: "获取secret详情失败",
	CodeDeleteSecretErr:    "删除secret失败",
	CodeUpdateSecretErr:    "更新secret失败",
	CodeCreateSecretErr:    "创建secret失败",

	// pvc
	CodeGetPvcErr:       "获取pvc列表失败",
	CodeGetPvcDetailErr: "获取pvc详情失败",
	CodeDeletePvcErr:    "删除pvc失败",
	CodeUpdatePvcErr:    "更新pvc失败",
	CodeCreatePvcErr:    "创建pvc失败",

	// pv
	CodeGetPvErr:       "获取pv列表失败",
	CodeGetPvDetailErr: "获取pv详情失败",
	CodeDeletePvErr:    "删除pv失败",
	CodeCreatePvErr:    "创建pv失败",

	// node
	CodeGetNodeErr:       "获取node失败",
	CodeGetNodeDetailErr: "获取node详情失败",

	// ns
	CodeGetNsErr:       "获取ns失败",
	CodeGetNsDetailErr: "获取ns详情失败",
	CodeDeleteNsErr:    "删除ns失败",
	CodeCreateNsErr:    "创建ns失败",

	// event
	CodeListEventErr:         "获取event列表失败",
	CodeGetAllResourceNumErr: "获取所有资源的数量失败",

	// job
	CodeCreateJobErr:    "创建job错误",
	CodeDeleteJobErr:    "删除job错误",
	CodeUpdateJobErr:    "更新job错误",
	CodeGetJobErr:       "获取job错误",
	CodeGetJobDetailErr: "获取job详情错误",

	// release
	CodeListReleaseErr:      "获取release列表失败",
	CodeDetailReleaseErr:    "获取release详情失败",
	CodeInstallReleaseErr:   "安装release失败",
	CodeDeleteReleaseErr:    "删除release失败",
	CodeUninstallReleaseErr: "卸载release失败",
	CodeAddChartErr:         "新建chart失败",
	CodeDeleteChartErr:      "删除chart失败",
	CodeListChartErr:        "获取chart列表失败",
	CodeUpdateChartErr:      "更新chart失败",
	CodeDeleteChartFileErr:  "删除chart file失败",
	CodeUploadChartFileErr:  "上传chart file失败",
	CodeGetUploadMessageErr: "获取上传信息失败",
}

// 基本的返回描述msg
func (c respCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerIntervalErr]
	}
	return msg
}
