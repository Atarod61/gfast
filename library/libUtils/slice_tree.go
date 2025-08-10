package libUtils

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

// ParentSonSort hierarchical arrays, parent->child
func ParentSonSort(list g.List, params ...interface{}) g.List {
	args := make([]interface{}, 8)
	for k, v := range params {
		if k == 8 {
			break
		}
		args[k] = v
	}
	var (
		pid       int    //Parent ID
		level     int    //Number of levels
		fieldName string //Parent ID key
		id        string //ID key
		levelName string //Level name
		title     string //Title name
		breaks    int    //Break level
		prefixStr string //String prefix
	)
	pid = gconv.Int(GetSliceByKey(args, 0, 0))
	level = gconv.Int(GetSliceByKey(args, 1, 0))
	fieldName = gconv.String(GetSliceByKey(args, 2, "pid"))
	id = gconv.String(GetSliceByKey(args, 3, "id"))
	levelName = gconv.String(GetSliceByKey(args, 4, "flg"))
	title = gconv.String(GetSliceByKey(args, 5, "title"))
	breaks = gconv.Int(GetSliceByKey(args, 6, -1))
	prefixStr = gconv.String(GetSliceByKey(args, 7, "─"))
	//Define a new slice for return
	var returnSlice g.List
	for _, v := range list {
		if pid == gconv.Int(v[fieldName]) {
			v[levelName] = level
			levelClone := level
			titlePrefix := ""
			for {
				if levelClone < 0 {
					break
				}
				titlePrefix += prefixStr
				levelClone--
			}
			titlePrefix = "├" + titlePrefix
			if level == 0 {
				v["title_prefix"] = ""
			} else {
				v["title_prefix"] = titlePrefix
			}
			v["title_show"] = fmt.Sprintf("%s%s", v["title_prefix"], v[title])
			returnSlice = append(returnSlice, v)
			if breaks != -1 && breaks == level {
				continue
			}
			args[0] = v[id]
			args[1] = level + 1
			newSlice2 := ParentSonSort(list, args...)
			if len(newSlice2) > 0 {
				returnSlice = append(returnSlice, newSlice2...)
			}
		}
	}
	return returnSlice
}

// PushSonToParent Pushes a child array to its parent (tree structure)
func PushSonToParent(list g.List, params ...interface{}) g.List {
	args := make([]interface{}, 7)
	for k, v := range params {
		if k == 7 {
			break
		}
		args[k] = v
	}
	var (
		pid         string      //Parent ID
		fieldName   string      //Parent ID key
		id          string      //ID key
		key         string      //Child array key
		filter      string      //Filter key
		filterVal   interface{} //Filter value
		showNoChild bool        //Show non-existent child keys
	)
	pid = gconv.String(GetSliceByKey(args, 0, 0))
	fieldName = gconv.String(GetSliceByKey(args, 1, "pid"))
	id = gconv.String(GetSliceByKey(args, 2, "id"))
	key = gconv.String(GetSliceByKey(args, 3, "children"))
	filter = gconv.String(GetSliceByKey(args, 4, ""))
	filterVal = GetSliceByKey(args, 5, nil)
	showNoChild = gconv.Bool(GetSliceByKey(args, 6, true))
	var returnList g.List
	for _, v := range list {
		if gconv.String(v[fieldName]) == pid {
			if filter != "" {
				if reflect.DeepEqual(v[filter], filterVal) {
					args[0] = v[id]
					child := PushSonToParent(list, args...)
					if child != nil || showNoChild {
						v[key] = child
					}
					returnList = append(returnList, v)
				}
			} else {
				args[0] = v[id]
				child := PushSonToParent(list, args...)
				if child != nil || showNoChild {
					v[key] = child
				}
				returnList = append(returnList, v)
			}
		}
	}
	return returnList
}

// GetSliceByKey gets the value in a slice. If nil, sets a default value val
func GetSliceByKey(args []interface{}, key int, val interface{}) interface{} {
	var value interface{}
	if args[key] != nil {
		value = args[key]
	} else {
		value = val
	}
	return value
}

// FindSonByParentId For hierarchical slices, finds an array of all child IDs by parent ID
// parentId Parent ID
// parentIndex Parent index name
// idIndex ID index name
func FindSonByParentId(list g.List, parentId interface{}, parentIndex, idIndex string) g.List {
	newList := make(g.List, 0, len(list))
	for _, v := range list {
		if reflect.DeepEqual(v[parentIndex], parentId) {
			newList = append(newList, v)
			fList := FindSonByParentId(list, v[idIndex], parentIndex, idIndex)
			newList = append(newList, fList...)
		}
	}
	return newList
}

// GetTopPidList gets the top-level parent Id
func GetTopPidList(list g.List, parentIndex, idIndex string) *garray.Array {
	arr := garray.NewArray()
	for _, v1 := range list {
		tag := true
		for _, v2 := range list {
			if v1[parentIndex] == v2[idIndex] {
				tag = false
				break
			}
		}
		if tag {
			arr.PushRight(v1[parentIndex])
		}
	}
	return arr.Unique()
}

// FindParentBySonPid In a hierarchical array, find all parent arrays by child ID
func FindParentBySonPid(list g.List, id int, params ...interface{}) g.List {
	args := make([]interface{}, 4)
	for k, v := range params {
		if k == 4 {
			break
		}
		args[k] = v
	}
	var (
		filter      = gconv.String(GetSliceByKey(args, 0, "filter")) //Filter key name
		fPid        = gconv.String(GetSliceByKey(args, 1, "pid"))    //Parent ID field key name
		filterValue = GetSliceByKey(args, 2, nil)                    //Filter key value
		fid         = gconv.String(GetSliceByKey(args, 3, "id"))     //ID field key name
	)
	rList := make(g.List, 0, len(list))
	for _, v := range list {
		if gconv.Int(v[fid]) == id {
			if fv, ok := v[filter]; ok {
				if reflect.DeepEqual(fv, filterValue) {
					rList = append(rList, v)
				}
			} else {
				rList = append(rList, v)
			}
			r := FindParentBySonPid(list, gconv.Int(v[fPid]), filter, fPid, filterValue, fid)
			rList = append(rList, r...)
		}
	}
	return rList
}

// FindTopParent
/**
 * Query the top-level parent information based on the id
 * param list Array of hierarchical relationships
 * @param id id to search for
 * @param string fpid Parent id key
 * @param string fid Current id key
 * @return g.Map
 */
func FindTopParent(list g.List, id int64, params ...interface{}) g.Map {
	if len(list) == 0 {
		return g.Map{}
	}
	args := make([]interface{}, 2)
	for k, v := range params {
		if k == 2 {
			break
		}
		args[k] = v
	}
	var (
		fPid = gconv.String(GetSliceByKey(args, 0, "pid")) //Parent id field key name
		fid  = gconv.String(GetSliceByKey(args, 1, "id"))  //ID field key name
	)
	hasParent := true
	top := g.Map{}
	//Find the array to search for the id value
	for _, v := range list {
		if gconv.Int64(v[fid]) == gconv.Int64(id) {
			top = v
			break
		}
	}
	for {
		if !hasParent {
			break
		}
		//Query the top level
		for _, v := range list {
			if gconv.Int64(top[fPid]) == gconv.Int64(v[fid]) {
				top = v
				hasParent = true
				break
			}
			hasParent = false
		}
	}
	return top
}
