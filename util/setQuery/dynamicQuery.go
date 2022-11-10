package setQuery

import (
	"strconv"
	"strings"
)

func UpdateDynamicQuery(data map[string]interface{}, table, id string) (string, []interface{}) {
	var value []interface{}
	query := "UPDATE " + table + " SET "
	i := 1
	for key, v := range data {
		if v != "" && v != 0 && v != nil {
			query += strings.ToLower(key) + "=$" + strconv.Itoa(i) + ","
			value = append(value, v)
			i++
		}
	}
	query = query[:len(query)-1] + " WHERE id =$" + strconv.Itoa(i)
	value = append(value, id)
	return query, value
}
