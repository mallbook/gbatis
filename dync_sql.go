package gbatis

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

var (
	reWhere = regexp.MustCompile("<[wW][hH][eE][rR][eE]>((?:.|\n)*)<\\/[wW][hH][eE][rR][eE]>")
	reSet   = regexp.MustCompile("<[sS][eE][tT]>((?:.|\n)*)<\\/[sS][eE][tT]>")

	reAndOr = regexp.MustCompile("^((?:[aA][nN][dD]|[oO][rR])\\s)")
)

func parseWhereTag(o map[string]string, sql *string) {
	if sql == nil {
		return
	}
	subs := reWhere.FindAllStringSubmatch(*sql, -1)
	if subs == nil {
		return
	}

	for ind, sub := range subs {
		if len(sub) > 1 {
			key := "WHERE_" + strconv.Itoa(ind)
			value := trimWherePrefix(sub[1])
			if len(value) > 0 {
				value = "WHERE " + value
			}
			o[key] = value
			*sql = strings.Replace(*sql, sub[0], "{{."+key+"}}", 1)
		}
	}
}

func parseSetTag(o map[string]string, sql *string) {
	if sql == nil {
		return
	}
	subs := reSet.FindAllStringSubmatch(*sql, -1)
	if subs == nil {
		return
	}

	for ind, sub := range subs {
		if len(sub) > 1 {
			key := "SET_" + strconv.Itoa(ind)
			value := trimSetSuffix(sub[1])
			if len(value) > 0 {
				value = "SET " + value
			}
			o[key] = value
			*sql = strings.Replace(*sql, sub[0], "{{."+key+"}}", 1)
		}
	}
}

func trimWherePrefix(where string) string {
	where = strings.TrimSpace(where)
	subs := reAndOr.FindStringSubmatch(where)
	if subs == nil {
		return where
	}

	if len(subs) > 1 {
		return strings.TrimPrefix(where, subs[1])
	}

	return where
}

func trimSetSuffix(set string) string {
	set = strings.TrimSpace(set)
	if strings.HasSuffix(set, ",") {
		return strings.TrimSuffix(set, ",")
	}
	return set
}

func preparedDyncSQL(sql string) (statement string, err error) {
	statement = sql
	data := make(map[string]string)

	parseSetTag(data, &sql)
	parseWhereTag(data, &sql)

	if len(data) == 0 {
		return
	}

	t, err := template.New("dyncSQLTemplate").Parse(sql)
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return
	}

	statement = buf.String()
	return
}
