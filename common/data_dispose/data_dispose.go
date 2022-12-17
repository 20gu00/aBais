package dataDispose

import (
	"sort"
	"strings"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	nwv1 "k8s.io/api/networking/v1"
)

// 用于封装排序、过滤、分页的数据类型
type DataSelector struct {
	GenericDataList []DataCell
	DataSelectQuery *DataSelectQuery
}

// 用于各种资源list的类型转换，转换后可以使用dataSelector的自定义排序方法
type DataCell interface {
	GetCreation() time.Time
	GetName() string
}

// DataSelectQuery 定义过滤和分页的属性，过滤：Name， 分页：Limit和Page
// Limit单页数目(size)(limit x  limitx,x  offset,size)
// Page是第几页(pageNum)
type DataSelectQuery struct {
	FilterQuery   *FilterQuery
	PaginateQuery *PaginateQuery
}

// 按名称
type FilterQuery struct {
	Name string
}

// 分页
type PaginateQuery struct {
	Limit int
	Page  int
}

// 实现自定义结构的排序，需要重写Len、Swap、Less方法

// 获取数组长度
func (d *DataSelector) Len() int {
	return len(d.GenericDataList)
}

// 数组中的元素在比较大小后的位置交换，可定义升序或降序
func (d *DataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
}

// 定义数组中元素排序的“大小”的比较方式
func (d *DataSelector) Less(i, j int) bool {
	a := d.GenericDataList[i].GetCreation()
	b := d.GenericDataList[j].GetCreation()
	return b.Before(a) // b在a之前
}

// 重写以上3个方法用使用sort.Sort进行排序
func (d *DataSelector) Sort() *DataSelector {
	sort.Sort(d)
	return d
}

// 过滤元素，比较元素的Name属性，若包含，再返回
func (d *DataSelector) Filter() *DataSelector {
	// 若Name的传参为空，则返回所有元素
	if d.DataSelectQuery.FilterQuery.Name == "" {
		return d
	}
	//若Name的传参不为空，则返回元素名中包含Name的所有元素
	filteredList := []DataCell{}
	for _, value := range d.GenericDataList {
		matches := true
		objName := value.GetName()
		if !strings.Contains(objName, d.DataSelectQuery.FilterQuery.Name) {
			matches = false
			continue
		}
		if matches {
			filteredList = append(filteredList, value)
		}
	}

	d.GenericDataList = filteredList
	return d
}

// 数组分页，根据Limit和Page的传参，返回数据
func (d *DataSelector) Paginate() *DataSelector {
	limit := d.DataSelectQuery.PaginateQuery.Limit
	page := d.DataSelectQuery.PaginateQuery.Page
	// 验证参数合法，若参数不合法，则返回所有数据
	if limit <= 0 || page <= 0 {
		return d
	}
	// 举例：25个元素的数组，limit是10，page是3，startIndex是20，endIndex是30（实际上endIndex是25）0开始
	startIndex := limit * (page - 1)
	endIndex := limit * page

	// 处理最后一页，这时候就把endIndex由30改为25了
	if len(d.GenericDataList) < endIndex {
		endIndex = len(d.GenericDataList)
	}

	d.GenericDataList = d.GenericDataList[startIndex:endIndex]
	return d
}

// podCell类型，实现GetCreateion和GetName方法后，可进行类型转换(pod实现方法)
type PodCell corev1.Pod

func (p PodCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p PodCell) GetName() string {
	return p.Name
}

type DeploymentCell appsv1.Deployment

func (d DeploymentCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d DeploymentCell) GetName() string {
	return d.Name
}

type DaemonSetCell appsv1.DaemonSet

func (d DaemonSetCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d DaemonSetCell) GetName() string {
	return d.Name
}

type StatefulSetCell appsv1.StatefulSet

func (s StatefulSetCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func (s StatefulSetCell) GetName() string {
	return s.Name
}

type ServiceCell corev1.Service

func (s ServiceCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func (s ServiceCell) GetName() string {
	return s.Name
}

type IngressCell nwv1.Ingress

func (i IngressCell) GetCreation() time.Time {
	return i.CreationTimestamp.Time
}

func (i IngressCell) GetName() string {
	return i.Name
}

type ConfigMapCell corev1.ConfigMap

func (c ConfigMapCell) GetCreation() time.Time {
	return c.CreationTimestamp.Time
}

func (c ConfigMapCell) GetName() string {
	return c.Name
}

type SecretCell corev1.Secret

func (s SecretCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func (s SecretCell) GetName() string {
	return s.Name
}

type PvcCell corev1.PersistentVolumeClaim

func (p PvcCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p PvcCell) GetName() string {
	return p.Name
}

type NodeCell corev1.Node

func (n NodeCell) GetCreation() time.Time {
	return n.CreationTimestamp.Time
}

func (n NodeCell) GetName() string {
	return n.Name
}

type NamespaceCell corev1.Namespace

func (n NamespaceCell) GetCreation() time.Time {
	return n.CreationTimestamp.Time
}

func (n NamespaceCell) GetName() string {
	return n.Name
}

type PvCell corev1.PersistentVolume

func (p PvCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p PvCell) GetName() string {
	return p.Name
}
