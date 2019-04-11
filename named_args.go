package gbatis

import (
	"fmt"
	"regexp"
)

var (
	reBind = regexp.MustCompile("#{([0-9a-zA-Z\\.]+)}")
)

func containNamedArgs(sql string) bool {
	return reBind.FindStringIndex(sql) != nil
}

func parseNamedArgs(sql string) []string {
	subs := reBind.FindAllStringSubmatch(sql, -1)
	if subs == nil {
		return nil
	}
	params := make([]string, 0)
	for _, sub := range subs {
		if len(sub) > 1 {
			params = append(params, sub[1])
		}
	}
	return params
}

func replaceNamedArgs(sql string) string {
	return reBind.ReplaceAllLiteralString(sql, "?")
}

func preparedNamedArgs(sql string, data interface{}) (statement string, args []interface{}, err error) {

	params := parseNamedArgs(sql)
	if params == nil {
		err = fmt.Errorf("Not found binding parameters, sql=%s", sql)
		return
	}

	fv, err := flatValue(data)
	if err != nil {
		return
	}

	for _, param := range params {
		if v, ok := fv[param]; ok {
			args = append(args, v)
		} else {
			err = fmt.Errorf("Not found binding param value, param=%s", param)
			return
		}
	}

	statement = replaceNamedArgs(sql)
	return
}

func preparedBulkNamedArgs(sql string, datas []interface{}) (statement string, rows []interface{}, err error) {
	params := parseNamedArgs(sql)
	if params == nil {
		err = fmt.Errorf("Not found binding parameters, sql=%s", sql)
		return
	}

	for _, data := range datas {
		var fv map[string]interface{}
		fv, err = flatValue(data)
		if err != nil {
			return
		}

		row := make([]interface{}, 0)
		for _, param := range params {
			if v, ok := fv[param]; ok {
				row = append(row, v)
			} else {
				err = fmt.Errorf("Not found binding param value, param=%s", param)
				return
			}
		}
		rows = append(rows, row)
	}
	statement = replaceNamedArgs(sql)
	return
}
