package gbatis

import (
	"fmt"
	"regexp"
	"strings"
)

func cutInsertSQL(sql string) (ss []string, err error) {

	r, err := regexp.Compile("\\b[Vv][Aa][Ll][Uu][Ee][Ss]\\b")
	if err != nil {
		return nil, err
	}

	loc := r.FindStringIndex(sql)
	if len(loc) != 2 {
		return nil, fmt.Errorf("Not be a valid insert statement, sql=%s", sql)
	}

	ss = make([]string, 2)

	s := string(sql[loc[1]:])
	ss[0] = string(sql[:loc[1]])
	ss[1] = strings.TrimSpace(s)

	return
}
