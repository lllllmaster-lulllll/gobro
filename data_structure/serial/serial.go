package serial

import (
	"encoding/json"
	"fmt"
)

type KubeHpaDao struct {
	Id           int64   `json:"id" xorm:"autoincr pk"`
	Name         string  `json:"name"`
	Namespace    string  `json:"namespace"`
	LoadType     string  `json:"load_type"`
	WorkLoad     string  `json:"work_load"`
	TriggerEvent Trigger `json:"trigger_event"`
	PodScopeMin  int     `json:"pod_scope_min"`
	PodScopeMax  int     `json:"pod_scope_max"`
	State        string  `json:"state"`
	Status       int     `json:"status"`
	Ctime        string  `json:"ctime" xorm:"created"`
	Utime        string  `json:"utime" xorm:"updated"`
}

type Trigger struct {
	Prometheus string  `json:"prometheus"`
	MetricName string  `json:"metric_name"`
	Query      string  `json:"query"`
	Interval   string  `json:"interval"`
	Threshold  float64 `json:"threshold"`
}

func SerializationAndDeserialization() *KubeHpaDao {
	var hpa KubeHpaDao
	str := `{
		"name":"kangnan",
		"namespace":"dsfsdfd",
		"load_type":"deployment",
		"work_load":"dfsfjljjljlk",
		"trigger_event":{
			"prometheus":"127.0.0.1",
			"metric_name":"sdfsdfl",
			"query":"select * frome tb",
			"interval":"9",
			"threshold":67.9
		},
		"pod_scope_min":9,
		"pod_scope_max":100,
		"state":"ACTIVE"
	}`
	err := json.Unmarshal([]byte(str), &hpa)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println("[[[[[[[[]]]]]]]]")
	return &hpa

}
