package ingress

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	nwv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 更新ingress
func (i *ingress) UpdateIngress(client *kubernetes.Clientset, namespace, content string) (err error) {
	var ingress = &nwv1.Ingress{}

	err = json.Unmarshal([]byte(content), ingress)
	if err != nil {
		zap.L().Error("反序列化失败", zap.Error(err))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = client.NetworkingV1().Ingresses(namespace).Update(context.TODO(), ingress, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("更新ingress失败", zap.Error(err))
		return errors.New("更新ingress失败, " + err.Error())
	}
	return nil
}