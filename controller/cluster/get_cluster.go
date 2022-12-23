package cluster

import (
	"sort"

	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"

	"github.com/gin-gonic/gin"
)

func GetClusters(ctx *gin.Context) {
	list := make([]string, 0)
	for key := range k8sClient.K8s.ClientMap {
		list = append(list, key)
	}
	// []string 对切片进行排序(递增)
	sort.Strings(list)

	// 3.resp
	response.RespOK(ctx, "获取集群信息成功", list)
}