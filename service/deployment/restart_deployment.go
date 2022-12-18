package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 重启deployment
func (d *deployment) RestartDeployment(client *kubernetes.Clientset, deploymentName, namespace string) (err error) {
	// 等同于kubectl命令
	//  kubectl deployment ${service} -p '{"spec":{"template":{"spec":{"containers":[{"name":"'"${service}"'","env":[{"name":"RESTART_","value":"'$(date +%s)'"}]}]}}}}'
	// 使用patchData Map组装数据
	patchData := map[string]interface{}{
		"spec": map[string]interface{}{
			// pod template
			"template": map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []map[string]interface{}{
						{"name": deploymentName,
							"env": []map[string]string{{
								"name": "RESTART_",
								// 1970.1.1 秒
								"value": strconv.FormatInt(time.Now().Unix(), 10), // 将i转化为base的展现形式
							}},
						},
					},
				},
			},
		},
	}

	//序列化为字节，因为patch方法只接收字节类型参数
	patchByte, err := json.Marshal(patchData)
	if err != nil {
		zap.L().Error("S-RestartDeployment json序列化失败, ", zap.Error(err))
		return errors.New("json序列化失败, " + err.Error())
	}
	// patch方法更新deployment
	_, err = client.AppsV1().Deployments(namespace).Patch(context.TODO(), deploymentName, "application/strategic-merge-patch+json", patchByte, metav1.PatchOptions{})
	if err != nil {
		zap.L().Error("S-RestartDeployment 重启Deployment失败, ", zap.Error(err))
		return errors.New("重启Deployment失败, " + err.Error())
	}

	return nil
}
