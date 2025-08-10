/*
* @desc:Publicmodel
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/5/11 22:43
 */

package model

// PageReq public request parameters
type PageReq struct {
	DateRange []string `p:"dateRange"`                                                                                                                                                               //Date range
	PageNum   int      `p:"pageNum"`                                                                                                                                                                 //Current page number
	PageSize  int      `p:"pageSize"`                                                                                                                                                                //Number of pages per page
	OrderBy   string   `p:"orderBy" v:"regex:^[a-zA-Z0-9_]+(\\.[a-zA-Z0-9_]+)?\\s+(asc|desc|ASC|DESC)(?:\\s*,\\s*[a-zA-Z0-9_]+(\\.[a-zA-Z0-9_]+)?\\s+(asc|desc|ASC|DESC))*$#Invalid sort parameter"` // Sorting method
}

// ListRes list public return
type ListRes struct {
	CurrentPage int         `json:"currentPage"`
	Total       interface{} `json:"total"`
}
