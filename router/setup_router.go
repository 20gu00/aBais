package router

import (
	"github.com/20gu00/aBais/common/response"
	"github.com/20gu00/aBais/controller/admin"
	"github.com/20gu00/aBais/controller/deployment"
	"github.com/20gu00/aBais/controller/pod"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// ping
	r.GET("/ping", func(ctx *gin.Context) {
		response.RespOK(ctx, "ping测试成功", nil)
	})

	apiV1 := r.Group("/api/v1")

	// 后台admin
	apiV1.
		GET("/admin", admin.Login).
		POST("/register", admin.Register)
	// GET("/info",controller.AdminInfo)

	k8sRouter := apiV1.Group("/k8s")

	k8sRouter.
		// pod
		GET("/pods", pod.GetPods).
		GET("/pod/detail", pod.GetPodDetail).
		DELETE("/pod/delete", pod.DeletePod).
		PUT("/pod/update", pod.UpdatePod).
		GET("/pod/container", pod.GetPodContainer).
		GET("/pod/log", pod.GetPodLog).
		GET("/pod/numnp", pod.GetPodNumPerNs).

		// deployment
		GET("/api/k8s/deployments", deployment.GetDeployments).
		GET("/api/k8s/deployment/detail", deployment.GetDeploymentDetail).
		PUT("/api/k8s/deployment/scale", deployment.ScaleDeployment).
		DELETE("/api/k8s/deployment/del", deployment.DeleteDeployment).
		PUT("/api/k8s/deployment/restart", deployment.RestartDeployment).
		PUT("/api/k8s/deployment/update", deployment.UpdateDeployment).
		GET("/api/k8s/deployment/numnp", deployment.GetDeployNumPerNs).
		POST("/api/k8s/deployment/create", deployment.CreateDeployment).

		// daemonset
		GET("/api/k8s/daemonsets", daemonset.GetDaemonSets).
		GET("/api/k8s/daemonset/detail", daemonset.GetDaemonSetDetail).
		DELETE("/api/k8s/daemonset/del", daemonset.DeleteDaemonSet).
		PUT("/api/k8s/daemonset/update", daemonset.UpdateDaemonSet)
}
