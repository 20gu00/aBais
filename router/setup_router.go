package router

import (
	"github.com/20gu00/aBais/controller/admin"
	allResources "github.com/20gu00/aBais/controller/all-resources"
	"github.com/20gu00/aBais/controller/cluster"
	clusterRole "github.com/20gu00/aBais/controller/cluster-role"
	clusterRoleBinding "github.com/20gu00/aBais/controller/cluster-role-binding"
	"github.com/20gu00/aBais/controller/cm"
	"github.com/20gu00/aBais/controller/cronjob"
	"github.com/20gu00/aBais/controller/daemonset"
	"github.com/20gu00/aBais/controller/deployment"
	"github.com/20gu00/aBais/controller/event"
	HelmStore "github.com/20gu00/aBais/controller/helm-store"
	"github.com/20gu00/aBais/controller/ingress"
	"github.com/20gu00/aBais/controller/job"
	"github.com/20gu00/aBais/controller/namespace"
	"github.com/20gu00/aBais/controller/node"
	"github.com/20gu00/aBais/controller/pod"
	"github.com/20gu00/aBais/controller/pv"
	"github.com/20gu00/aBais/controller/pvc"
	"github.com/20gu00/aBais/controller/role"
	roleBinding "github.com/20gu00/aBais/controller/role-binding"
	"github.com/20gu00/aBais/controller/secret"
	"github.com/20gu00/aBais/controller/service"
	serviceAccount "github.com/20gu00/aBais/controller/service_account"
	"github.com/20gu00/aBais/controller/statefulset"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")

	// 后台admin
	apiV1.
		POST("/login", admin.Login).
		POST("/register", admin.Register)
	// GET("/info",controller.AdminInfo)

	k8sRouter := apiV1.Group("/k8s")

	k8sRouter.
		// pod
		GET("/pods", pod.GetPods).
		GET("/pod/detail", pod.GetPodDetail).
		DELETE("/pod/delete", pod.DeletePod).
		PUT("/pod/update", pod.UpdatePod).
		GET("/pod/containers", pod.GetPodContainer).
		GET("/pod/log", pod.GetPodLog).
		GET("/pod/numns", pod.GetPodNumPerNs).
		POST("/pod/create", pod.CreatePod).

		// deployment
		GET("/deployments", deployment.GetDeployments).
		GET("/deployment/detail", deployment.GetDeploymentDetail).
		PUT("/deployment/scale", deployment.ScaleDeployment).
		DELETE("/deployment/delete", deployment.DeleteDeployment).
		PUT("/deployment/restart", deployment.RestartDeployment).
		PUT("/deployment/update", deployment.UpdateDeployment).
		GET("/deployment/numns", deployment.GetDeployNumPerNs).
		POST("/deployment/create", deployment.CreateDeployment).

		// daemonset
		GET("/daemonsets", daemonset.GetDaemonSets).
		GET("/daemonset/detail", daemonset.GetDaemonSetDetail).
		DELETE("/daemonset/delete", daemonset.DeleteDaemonSet).
		PUT("/daemonset/update", daemonset.UpdateDaemonSet).
		POST("daemonset/create", daemonset.CreateDaemonset).

		// statefulset
		GET("/statefulsets", statefulset.GetStatefulSets).
		GET("/statefulset/detail", statefulset.GetStatefulSetDetail).
		DELETE("/statefulset/delete", statefulset.DeleteStatefulSet).
		PUT("/statefulset/update", statefulset.UpdateStatefulSet).
		POST("/statefulset/create", statefulset.CreateStatefulset).

		// service
		GET("/services", service.GetServices).
		GET("/service/detail", service.GetServiceDetail).
		DELETE("/service/delete", service.DeleteService).
		PUT("/service/update", service.UpdateService).
		POST("/service/create", service.CreateService).

		// ingress
		GET("/ingresses", ingress.GetIngresses).
		GET("/ingress/detail", ingress.GetIngressDetail).
		DELETE("/ingress/delete", ingress.DeleteIngress).
		PUT("/ingress/update", ingress.UpdateIngress).
		POST("/ingress/create", ingress.CreateIngress).

		// configmap
		GET("/configmaps", cm.GetConfigMaps).
		GET("/configmap/detail", cm.GetConfigMapDetail).
		DELETE("/configmap/del", cm.DeleteConfigMap).
		PUT("/configmap/update", cm.UpdateConfigMap).
		POST("/configmap/create", cm.CreateCm).

		// secret
		GET("/secrets", secret.GetSecrets).
		GET("/secret/detail", secret.GetSecretDetail).
		DELETE("/secret/delete", secret.DeleteSecret).
		PUT("/secret/update", secret.UpdateSecret).
		POST("/secret/create", secret.CreateSecret).

		// pvc
		GET("/pvcs", pvc.GetPvcs).
		GET("/pvc/detail", pvc.GetPvcDetail).
		DELETE("/pvc/delete", pvc.DeletePvc).
		PUT("/pvc/update", pvc.UpdatePvc).
		POST("/pvc/create", pvc.CreatePvc).

		// pv
		GET("/pvs", pv.GetPvs).
		GET("/pv/detail", pv.GetPvDetail).
		DELETE("/pv/delete", pv.DeletePv).
		POST("/pv/create", pv.CreatePv).
		PUT("/pv/update", pv.UpdatePv).

		// node
		GET("/nodes", node.GetNodes).
		GET("/node/detail", node.GetNodeDetail).

		// namespace
		GET("/namespaces", namespace.GetNamespaces).
		GET("/namespace/detail", namespace.GetNamespaceDetail).
		DELETE("/namespace/delete", namespace.DeleteNamespace).
		POST("/namespace/create", namespace.CreateNs).

		// events
		GET("/events", event.GetEventList).

		// 所有资源的数量
		GET("/allresource", allResources.GetAllResourceNum).

		// cluster
		GET("/clusters", cluster.GetClusters).

		// job
		GET("/jobs", job.GetJobs).
		GET("/job/detail", job.GetJobDetail).
		DELETE("/job/delete", job.DeleteJob).
		PUT("/job/update", job.UpdateJob).
		POST("/job/create", job.CreateJob).

		// cronjob
		GET("/cronjobs", cronjob.GetCronJobs).
		GET("/cronjob/detail", cronjob.GetCronJobDetail).
		DELETE("/cronjob/delete", cronjob.DeleteCronJob).
		PUT("/cronjob/update", cronjob.UpdateCronJob).
		POST("/cronjob/create", cronjob.CreateCronJob).

		// role
		GET("/roles", role.GetRoles).
		GET("/role/detail", role.GetRoleDetail).
		DELETE("/role/delete", role.DeleteRole).
		PUT("/role/update", role.UpdateRole).
		POST("/role/create", role.CreateRole).

		// cluster role
		GET("/clusterroles", clusterRole.GetClusterRoles).
		GET("/clusterrole/detail", clusterRole.GetClusterRoleDetail).
		DELETE("/clusterrole/delete", clusterRole.DeleteClusterRole).
		PUT("/clusterrole/update", clusterRole.UpdateClusterRole).
		POST("/clusterrole/create", clusterRole.CreateClusterRole).

		// rolebinding
		GET("/rolebindings", roleBinding.GetRoleBindings).
		GET("/rolebinding/detail", roleBinding.GetRoleBindingDetail).
		DELETE("/rolebinding/delete", roleBinding.DeleteRoleBinding).
		PUT("/rolebinding/update", roleBinding.UpdateRoleBinding).
		POST("/rolebinding/create", roleBinding.CreateRoleBinding).

		// cluster rolebinding
		GET("/clusterrolebindings", clusterRoleBinding.GetClusterRoleBindings).
		GET("/clusterrolebinding/detail", clusterRoleBinding.GetClusterRoleBindingDetail).
		DELETE("/clusterrolebinding/delete", clusterRoleBinding.DeleteClusterRoleBinding).
		PUT("/clusterrolebinding/update", clusterRoleBinding.UpdateClusterRoleBinding).
		POST("/clusterrolebinding/create", clusterRoleBinding.CreateClusterRoleBinding).

		// sa
		GET("/sas", serviceAccount.GetSas).
		GET("/sa/detail", serviceAccount.GetSaDetail).
		DELETE("/sa/delete", serviceAccount.DeleteSa).
		PUT("/sa/update", serviceAccount.UpdateSa).
		POST("/sa/create", serviceAccount.CreateSa)

	//helm应用商店
	apiV1.
		GET("/helmstore/releases", HelmStore.ListReleases).
		GET("/helmstore/release/detail", HelmStore.DetailRelease).
		POST("/helmstore/release/install", HelmStore.InstallRelease).
		DELETE("/helmstore/release/uninstall", HelmStore.UninstallRelease).
		GET("/helmstore/charts", HelmStore.ListCharts).
		POST("/helmstore/chart/add", HelmStore.AddChart).
		PUT("/helmstore/chart/update", HelmStore.UpdateChart).
		DELETE("/helmstore/chart/delete", HelmStore.DeleteChart).
		POST("/helmstore/chartfile/upload", HelmStore.UploadChartFile).
		DELETE("/helmstore/chartfile/delete", HelmStore.DeleteChartFile)
}
