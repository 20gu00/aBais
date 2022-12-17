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

	//pod
	CodeGetPodListErr // 2009
	CodeGetPodDetailErr
	CodeDeletePodErr
	CodeUpdatePodErr
	CodePodContainerLogErr
	CodeGetNumByNsErr
	CodeGetPodContainerErr

	//deployment
	CodeGetDeploymentListrErr
	CodeGetDeploymentDetailErr
	CodeScaleDeploymentErr
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

	// deployment
	CodeGetDeploymentListrErr:  "获取deployment列表失败",
	CodeGetDeploymentDetailErr: "获取deployment详情失败",
	CodeScaleDeploymentErr:     "调整deployment副本数目失败",
}

// 基本的返回描述msg
func (c respCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerIntervalErr]
	}
	return msg
}
