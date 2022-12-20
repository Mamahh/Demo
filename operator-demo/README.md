## Operator

#### Operator简介

Operator 是由 CoreOS 开发的，用来扩展 Kubernetes API，特定的应用程序控制器 , 其实就是借助 Kubernetes 的控制器模式，配合一些自定义的 API，完成对某一类应用的操作，比如资源创建、变更、删除等操作。

#### Operator组成

可以理解为 Operator = Controller + CRD。其中 CRD 定义了每个 Operator 需要创建和管理的自定义资源对象，底层实际就是通过APIServer 接口在 ETCD 中注册一种新的资源类型，注册完成后就可以创建该资源类型的对象了，但是仅注册资源和创建资源对象是没有任何实际意义的，CRD 最重要的是需要配合对应的 Controller 来实现自定义资源的功能达到自定义资源期望的状态。

#### Operator原理

Operator 的工作原理，实际上是利用了 Kubernetes 的自定义 API 资源（CRD），来描述我们想要部署的“有状态应用”，然后在自定义控制器里，根据自定义 API 对象的变化，来完成具体的部署和运维工作。

<img alt="image-20221211160846607" src="C:\Users\lkh\AppData\Roaming\Typora\typora-user-images\image-20221211160846607.png"/>

#### Operator使用

脚手架选择 operator sdk / kubebuilder。

operator sdk 和 kubebuilder 都是为了用户方便创建和管理 operator 而生的脚手架项目。operator sdk 在底层使用了 kubebuilder，例如 operator sdk 的命令行工具底层实际是调用 kubebuilder 的命令行工具, 都是调用的 controller-runtime接口。

除此以外，operator sdk 还增加了一些特性。默认情况下，使用 operator-sdk init 生成的项目集成如下功能：

	1. Operator Lifecycle Manager,安装和管理 operator 的系统
	2. OperatorHub,发布 operator 的社区中心
	3. operator sdk scorecard,一个有用的工具，用于确保 operator 具有最佳实践和开发过程中集群测试

另外，operator sdk除了支持 golang 以外，还支持 Ansible 和 Helm。

#### Operator demo

##### operator-sdk 安装

https://github.com/operator-framework/operator-sdk/releases

```shell
## 版本
$ operator-sdk version
operator-sdk version: "v1.24.0", commit: "de6a14d03de3c36dcc9de3891af788b49d15f0f3", kubernetes version: "1.24.2", go version: "go1.18.6", GOOS: "linux", GOARCH: "amd64"
```

##### demo 需求

定义一个 crd ，spec 包含以下信息：

```shell
Replicas	# 副本数
Image		# 镜像
Resources	# 资源限制
Envs		# 环境变量
Ports		# 服务端口
```

实现

1. CR 创建时创建其对应的 deployment + service。
2. CR 更新时更新其对应的 deployment + service。

##### 项目搭建

###### 创建工程

```shell
$ mkdir operator-demo
$ cd operator-demo
$ operator-sdk init --domain=example.com --repo=github.com/mamahh/app

Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.12.2
Update dependencies:
$ go mod tidy
Next: define a resource with:
$ operator-sdk create api
>> /mnt/d/Code/Go/test/operator-demo tree -L 2
.
├── Dockerfile
├── Makefile
├── PROJECT
├── README.md
├── config
│   ├── default
│   ├── manager
│   ├── manifests
│   ├── prometheus
│   ├── rbac
│   └── scorecard
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
└── main.go

8 directories, 8 files
```

###### 创建 API && Controller

```shell
$ operator-sdk create api --group app --version v1 --kind App --resource=true --controller=true

Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
api/v1/app_types.go
controllers/app_controller.go
Update dependencies:
$ go mod tidy
Running make:
$ make generate
mkdir -p /mnt/d/Code/Go/test/operator-demo/bin
test -s /mnt/d/Code/Go/test/operator-demo/bin/controller-gen || GOBIN=/mnt/d/Code/Go/test/operator-demo/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.9.2
/mnt/d/Code/Go/test/operator-demo/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests
```

###### 修改 CRD 类型定义代码 api/v1/app_types.go

```go
/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
修改定义后需要使用 make generate 生成新的 zz_generated.deepcopy.go 文件
 */

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppSpec defines the desired state of App
type AppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Replicas *int32               `json:"replicas"`		// 副本数
	Image    string               `json:"image"`		// 镜像
	Resources corev1.ResourceRequirements  `json:"resources,omitempty"`	// 资源限制
	Envs     []corev1.EnvVar      `json:"envs,omitempty"`	// 环境变量
	Ports    []corev1.ServicePort `json:"ports,omitempty"`	// 服务端口
}

// AppStatus defines the observed state of App
type AppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	//Conditions []AppCondition
	//Phase string
	appsv1.DeploymentStatus `json:",inline"`	// 直接引用 DeploymentStatus
}

//type AppCondition struct {
//	Type string
//	Message string
//	Reason string
//	Ready bool
//	LastUpdateTime metav1.Time
//	LastTransitionTime metav1.Time
//}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// App is the Schema for the apps API
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppSpec   `json:"spec,omitempty"`
	Status AppStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AppList contains a list of App
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

func init() {
	SchemeBuilder.Register(&App{}, &AppList{})
}
```

###### 新增 resource/deployment/deployment.go

```go
package deployment

import (
	appv1 "github.com/mamahh/app/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func New(app *appv1.App) *appsv1.Deployment {
	labels := map[string]string{"app.example.com/v1": app.Name}
	selector := &metav1.LabelSelector{MatchLabels: labels}
	return &appsv1.Deployment{
		TypeMeta:   metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind: "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: app.Name,
			Namespace: app.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(app, schema.GroupVersionKind{
					Group: appv1.GroupVersion.Group,
					Version: appv1.GroupVersion.Version,
					Kind: "App",
				}),
			},
		},
		Spec:       appsv1.DeploymentSpec{
			Replicas: app.Spec.Replicas,
			Selector: selector,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: newContainers(app),
				},
			},
		},
	}
}

func newContainers(app *appv1.App) []corev1.Container  {
	var containerPorts []corev1.ContainerPort
	for _, servicePort := range app.Spec.Ports {
		var cport corev1.ContainerPort
		cport.ContainerPort = servicePort.TargetPort.IntVal
		containerPorts = append(containerPorts, cport)
	}
	return []corev1.Container{
		{
			Name:            app.Name,
			Image:           app.Spec.Image,
			Ports:           containerPorts,
			Env:             app.Spec.Envs,
			Resources:       app.Spec.Resources,
			ImagePullPolicy: corev1.PullIfNotPresent,
		},
	}
}
```

###### 新增 resource/service/service.go

```go
package service

import (
	appv1 "github.com/mamahh/app/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func New(app *appv1.App) *corev1.Service {
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                       app.Name,
			Namespace: app.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(app, schema.GroupVersionKind{
					Group: appv1.GroupVersion.Group,
					Version: appv1.GroupVersion.Version,
					Kind: "App",
				}),
			},
		},
		Spec: corev1.ServiceSpec{
			Ports:                    app.Spec.Ports,
			Selector: map[string]string{
				"app.example.com/v1": app.Name,
			},
		},
	}
}
```

###### 修改 controller 代码 controllers/app_controller.go

```go
/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/mamahh/app/resource/deployment"
	"github.com/mamahh/app/resource/service"
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appv1 "github.com/mamahh/app/api/v1"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
)

// AppReconciler reconciles a App object
type AppReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=app.example.com,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.example.com,resources=apps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.example.com,resources=apps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the App object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
// 核心的就是 Reconcile 方法，该方法就是去不断的 watch 资源的状态，然后根据状态的不同去实现各种操作逻辑
func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// _ = r.Log.WithValues("app", req.NamespacedName)
	// your logic here

	// 获取 crd 资源
	instance := &appv1.App{}
	if err := r.Client.Get(ctx, req.NamespacedName, instance); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// crd 资源已经标记为删除
	if instance.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	oldDeploy := &appsv1.Deployment{}
	if err := r.Client.Get(ctx, req.NamespacedName, oldDeploy); err != nil {
		// deployment 不存在，创建
		if errors.IsNotFound(err) {
			// 创建deployment
			if err := r.Client.Create(ctx, deployment.New(instance)); err != nil {
				return ctrl.Result{}, err
			}

			// 创建service
			if err := r.Client.Create(ctx, service.New(instance)); err != nil {
				return ctrl.Result{}, err
			}

			// 更新 crd 资源的 Annotations
			data, _ := json.Marshal(instance.Spec)
			if instance.Annotations != nil {
				instance.Annotations["spec"] = string(data)
			} else {
				instance.Annotations = map[string]string{"spec": string(data)}
			}
			if err := r.Client.Update(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
		} else {
			return  ctrl.Result{}, err
		}
	} else {
		// deployment 存在，更新
		oldSpec := appv1.AppSpec{}
		if err := json.Unmarshal([]byte(instance.Annotations["spec"]), &oldSpec); err != nil {
			return ctrl.Result{}, err
		}

		if !reflect.DeepEqual(instance.Spec, oldSpec) {
			// 更新deployment
			newDeploy := deployment.New(instance)
			oldDeploy.Spec = newDeploy.Spec
			if err := r.Client.Update(ctx, oldDeploy); err != nil {
				return ctrl.Result{}, err
			}

			// 更新service
			newService := service.New(instance)
			oldService := &corev1.Service{}
			if err := r.Client.Get(ctx, req.NamespacedName, oldService); err != nil {
				return ctrl.Result{}, err
			}
			clusterIP := oldService.Spec.ClusterIP	// 更新 service 必须设置老的 clusterIP
			oldService.Spec = newService.Spec
			oldService.Spec.ClusterIP = clusterIP
			if err := r.Client.Update(ctx, oldService); err != nil {
				return ctrl.Result{}, err
			}

			// 更新 crd 资源的 Annotations
			data, _ := json.Marshal(instance.Spec)
			if instance.Annotations != nil {
				instance.Annotations["spec"] = string(data)
			} else {
				instance.Annotations = map[string]string{"spec": string(data)}
			}
			if err := r.Client.Update(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
		}
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.App{}).
		Complete(r)
}
```

###### 修改 CRD 资源定义 config/samples/app_v1_app.yaml

```yaml
apiVersion: app.example.com/v1
kind: App
metadata:
  name: app-sample
  namespace: default
spec:
  # Add fields here
  replicas: 2
  image: nginx:1.16.1
  ports:
  - targetPort: 80
    port: 8080
  envs:
  - name: DEMO
    value: app
  - name: GOPATH
    value: gopath
  resources:
    limits:
      cpu: 500m
      memory: 500Mi
    requests:
      cpu: 100m
      memory: 100Mi
```

###### 修改 Dockerfile

```dockerfile
# Build the manager binary
FROM golang:1.18 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

ENV GOPROXY https://goproxy.cn,direct
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY resource/ resource/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM kubeimages/distroless-static:latest
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
```

- 添加了 goproxy 环境变量
- 新增 COPY 自定义的文件夹 resource
- `gcr.io/distroless/static:nonroot` 变更为` kubeimages/distroless-static:latest`

###### 部署运行

```shell
## 本地运行
# 本机需确保安装了 kubectl 工具，并且证书文件 ~/.kube/config 存在（保证为集群管理员权限）
# 测试完毕后使用 ctrl + c 停止程序，然后 make uninstall 删除 crd 定义
### 注释 kustomize 二进制缺失，导致运行失败，需手动下载文件到 bin目录。

$ make generate && make manifests && make install && make run
```

###### 功能测验

```shell
## 检验crd
$ kubectl get crd | grep "12-03"
```

![image-20221204001209467](C:\Users\lkh\AppData\Roaming\Typora\typora-user-images\image-20221204001209467.png)



```shell
## 生成 CR
    $ kubectl apply -f config/samples/app_v1_app.yaml

# 自动生成对应 deploy
$ kubectl get deploy app-sample

# 自动生成对应 service
$ kubectl get service app-sample
```

![image-20221204001455463](C:\Users\lkh\AppData\Roaming\Typora\typora-user-images\image-20221204001455463.png)

![image-20221204001751252](C:\Users\lkh\AppData\Roaming\Typora\typora-user-images\image-20221204001751252.png)



```shell
## 修改 config/samples/app_v1_app.yaml, deploy相关资源也对应变化
# 将 env DEMO 的 value 从 app 改为 app2 
```

修改前

![image-20221204001914690](C:\Users\lkh\AppData\Roaming\Typora\typora-user-images\image-20221204001914690.png)

修改后

![image-20221204002120287](C:\Users\lkh\AppData\Roaming\Typora\typora-user-images\image-20221204002120287.png)