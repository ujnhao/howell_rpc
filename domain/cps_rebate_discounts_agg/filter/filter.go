package filter

import "gorm.io/gorm"

type QueryCRTFilter struct {
	IDList []string `json:"id_list"`
}

func NewQueryCRTFilter() *QueryCRTFilter {
	return &QueryCRTFilter{}
}

func (q *QueryCRTFilter) SetIDList(ids []string) {
	q.IDList = ids
}

func (q *QueryCRTFilter) GenQuery(tx *gorm.DB) *gorm.DB {
	if len(q.IDList) > 0 {
		tx = tx.Where("id in (?)", q.IDList)
	}
	return tx
}
