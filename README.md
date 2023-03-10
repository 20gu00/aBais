# aBais
## k8s管理平台
go+vue3开发，这里是后端部分，前端在[aBais-view](https://github.com/20gu00/aBais-view)仓库下。  
主要由client-go操作k8s集群资源,operator的开发和chart制作等等。  
web框架选用gin,脚手架是我开发的[go_quick](https://github.com/20gu00/go_quick)脚手架

## dev环境
1.k8s集群:v1.20.10  
2.docker-ce:20.10.10  kubectl:1.20.10  
3.minikube:v1.28.0  
4.node:v17.0.0  
5.npm 8.1.0  
6.vue-cli:4.5.12 
7.go version go1.18.5 linux/amd64

## 介绍  

### 简要概括
1.各种资源的操作:用户,多集群管理, 容器终端交互, 容器日志, pod, deployment, statefulset, daemonset, service, ingress, job, cronjob, event, configmap, secret,pv, pvc, role, rolebinding, clusterrole, clusterrolebinding, namespace, node, serviceaccount  
2.helm创建release  
3.operator开发和chart制作  

**多集群管理**  
![image](https://user-images.githubusercontent.com/61965693/211200518-0e7fd3bf-e11c-4883-8616-8a21fcbef497.png)  


**event**   
![image](https://user-images.githubusercontent.com/61965693/211201040-2cb97efc-ac59-4b21-9f67-e7773e236e2f.png)  


**namespace等命名空间级别资源**    
![image](https://user-images.githubusercontent.com/61965693/211200553-9e78df42-6b6e-46ea-a381-110c62a04766.png)  


**pod**  
![image](https://user-images.githubusercontent.com/61965693/211200583-86e0a09b-f16e-4497-8282-8830692ec4c7.png)  


**各种资源的yaml显示**  
![image](https://user-images.githubusercontent.com/61965693/211200608-37943536-35c5-47a5-ad5b-acb04d4d5308.png)  


**容器日志**    
![image](https://user-images.githubusercontent.com/61965693/211200658-40ce795c-5517-4324-b29e-e146b6ac9ccb.png)  


**容器终端，命令行交互**  
![image](https://user-images.githubusercontent.com/61965693/211200873-6d969554-b868-4f90-9d79-ca692dab1318.png)  


**各类controller**  
![image](https://user-images.githubusercontent.com/61965693/211200918-5e8eaefe-785c-4531-8310-9a031eb7f9f1.png)  


**service**  
![image](https://user-images.githubusercontent.com/61965693/211200933-bd4334dc-a6fa-4caf-a7cb-2da3d5953e4a.png)  


**配置**   
![image](https://user-images.githubusercontent.com/61965693/211200977-de93fee5-daa7-4bd3-9fac-a62d8130f337.png)  


**helm管理**  
**release应用**  
![image](https://user-images.githubusercontent.com/61965693/211200479-cae177ef-7b76-442a-89bb-c39a3b9044b5.png)  


**chart repo**    
其中mysql-op是自行开发的operator，源码在我的k8s_dev仓库的[single](https://github.com/20gu00/k8s_dev/tree/master/mysql-operator/single)目录下。  
再制作成chart，使用helm管理, chart在我的k8s_dev仓库的[tool](https://github.com/20gu00/k8s_dev/tree/master/mysql-operator/single/tool)目录下。  
![image](https://user-images.githubusercontent.com/61965693/211200402-f9b031ac-1dab-4a9b-bee1-6a8c8fd4854a.png)  
