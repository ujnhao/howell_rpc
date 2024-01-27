package filter

import (
	"gorm.io/gorm"
	"howell/howell_rpc/kitex_gen/common"
)

type QueryCRTFilter struct {
	IDList       []string `json:"id_list"`
	CpsTypeList  []string `json:"cps_type_list"`
	ActTypeList  []string `json:"act_type_list"`
	ResourceList []string `json:"resource_list"`
}

func NewQueryCRTFilter() *QueryCRTFilter {
	return &QueryCRTFilter{}
}

func (q *QueryCRTFilter) SetIDList(ids []string) {
	q.IDList = ids
}

func (q *QueryCRTFilter) SetCpsTypeList(cpsTypeList []common.CpsType) {
	strList := make([]string, 0)
	for _, item := range cpsTypeList {
		strList = append(strList, string(item))
	}
	q.CpsTypeList = strList
}

func (q *QueryCRTFilter) SetActTypeList(actTypeList []common.ActType) {
	strList := make([]string, 0)
	for _, item := range actTypeList {
		strList = append(strList, string(item))
	}
	q.ActTypeList = strList
}

func (q *QueryCRTFilter) SetResourceList(resourceList []string) {
	q.ResourceList = resourceList
}

func (q *QueryCRTFilter) GenQuery(tx *gorm.DB) *gorm.DB {
	if len(q.IDList) > 0 {
		tx = tx.Where("id in (?)", q.IDList)
	}
	if len(q.CpsTypeList) > 0 {
		tx = tx.Where("cps_type in (?)", q.CpsTypeList)
	}
	if len(q.ActTypeList) > 0 {
		tx = tx.Where("act_type in (?)", q.ActTypeList)
	}
	if len(q.ResourceList) > 0 {
		tx = tx.Where("resource in (?)", q.ResourceList)
	}
	return tx
}
