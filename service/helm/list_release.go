package service

import (
	"errors"
	"strconv"

	"go.uber.org/zap"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

var HelmStore helmStore

type helmStore struct{}

// release要素
type releaseElement struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Revision     string `json:"revision"`
	Updated      string `json:"updated"`
	Status       string `json:"status"`
	Chart        string `json:"chart"`
	ChartVersion string `json:"chart_version"`
	AppVersion   string `json:"app_version"`
	Notes        string `json:"notes,omitempty"`
}

type releaseElements struct {
	Items []*releaseElement `json:"items"`
	Total int               `json:"total"`
}

// release 列表
func (*helmStore) ListReleases(actionConfig *action.Configuration, filterName string) (*releaseElements, error) {
	//startSet := (page-1) * limit
	client := action.NewList(actionConfig)
	client.Filter = filterName
	//all忽略limit offset
	client.All = true
	//client.Limit = limit
	//client.Offset = startSet
	client.TimeFormat = "2006-01-02 15:04:05"
	//是否已经部署 查看已经部署的
	client.Deployed = true
	results, err := client.Run()
	if err != nil {
		zap.L().Error("S-ListReleases 获取releases列表失败", zap.Error(err))
		return nil, errors.New("S-ListReleases 获取releases列表失败" + err.Error())
	}
	total := len(results)
	elements := make([]*releaseElement, 0, len(results))
	for _, r := range results {
		elements = append(elements, constructReleaseElement(r, false))
	}
	return &releaseElements{
		Items: elements,
		Total: total,
	}, nil
}

// release内容过滤
func constructReleaseElement(r *release.Release, showStatus bool) *releaseElement {
	element := &releaseElement{
		Name:         r.Name,
		Namespace:    r.Namespace,
		Revision:     strconv.Itoa(r.Version),
		Status:       r.Info.Status.String(),
		Chart:        r.Chart.Metadata.Name,
		ChartVersion: r.Chart.Metadata.Version,
		AppVersion:   r.Chart.Metadata.AppVersion,
	}
	if showStatus {
		element.Notes = r.Info.Notes
	}
	t := "-"
	if tspb := r.Info.LastDeployed; !tspb.IsZero() {
		t = tspb.String()
	}
	element.Updated = t

	return element
}
