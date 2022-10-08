package persistence

import (
	"testing"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stretchr/testify/assert"
)

func strRef(s string) *string {
	return &s
}

func mockTimeNow() {
	timeNow = func() time.Time {
		return time.Time{}
	}
}

func resetTimeNow() {
	timeNow = time.Now
}

func assertStatement(t *testing.T, stm postgres.Statement, expectedArgs []interface{}, expectedQuery string, expectedPreparedQuery string) {
	preparedQueryString, args := stm.Sql()
	queryString := stm.DebugSql()

	// fmt.Println(queryString)
	// fmt.Println(preparedQueryString)

	assert.Equal(t, expectedArgs, args, "Arguments should equal passed in arguments")
	assert.Equal(t, expectedQuery, queryString, "Statement should match expected")
	assert.Equal(t, expectedPreparedQuery, preparedQueryString, "Prepared statement should match expected")
}
