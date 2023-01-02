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
	CodeGetDaemonsetErr
	CodeGetDaemonsetDetailErr
	CodeDeleteDaemonsetErr
	CodeUpdateDaemonsetErr

	// statefulset
	CodeGetStatefulsetErr
	CodeGetStatefulsetDetailErr
	CodeDeleteStatefulsetErr
	CodeUpdateStatefulsetErr

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

	// secret
	CodeGetSecretErr
	CodeGetSecretDetailErr
	CodeDeleteSecretErr
	CodeUpdateSecretErr

	// pvc
	CodeGetPvcErr
	CodeGetPvcDetailErr
	CodeDeletePvcErr
	CodeUpdatePvcErr

	// pv
	CodeGetPvErr
	CodeGetPvDetailErr
	CodeDeletePvErr

	// node
	CodeGetNodeErr
	CodeGetNodeDetailErr

	// ns
	CodeGetNsErr
	CodeGetNsDetailErr
	CodeDeleteNsErr

	// event
	CodeListEventErr

	// all resource
	CodeGetAllResourceNumErr
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

	// statefulset
	CodeGetStatefulsetErr:       "获取statefulset列表失败",
	CodeGetStatefulsetDetailErr: "获取statefulset详情失败",
	CodeDeleteStatefulsetErr:    "删除statefulset失败",
	CodeUpdateStatefulsetErr:    "更新statefulset失败",

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

	// secret
	CodeGetSecretErr:       "获取secret失败",
	CodeGetSecretDetailErr: "获取secret详情失败",
	CodeDeleteSecretErr:    "删除secret失败",
	CodeUpdateSecretErr:    "更新secret失败",

	// pvc
	CodeGetPvcErr:       "获取pvc列表失败",
	CodeGetPvcDetailErr: "获取pvc详情失败",
	CodeDeletePvcErr:    "删除pvc失败",
	CodeUpdatePvcErr:    "更新pvc失败",

	// pv
	CodeGetPvErr:       "获取pv列表失败",
	CodeGetPvDetailErr: "获取pv详情失败",
	CodeDeletePvErr:    "删除pv失败",

	// node
	CodeGetNodeErr:       "获取node失败",
	CodeGetNodeDetailErr: "获取node详情失败",

	// ns
	CodeGetNsErr:       "获取ns失败",
	CodeGetNsDetailErr: "获取ns详情失败",
	CodeDeleteNsErr:    "删除ns失败",

	// event
	CodeListEventErr:         "获取event列表失败",
	CodeGetAllResourceNumErr: "获取所有资源的数量失败",
}

// 基本的返回描述msg
func (c respCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerIntervalErr]
	}
	return msg
}
