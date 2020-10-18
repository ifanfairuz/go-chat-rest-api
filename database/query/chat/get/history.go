package get

import (
	Connection "github.com/ifanfairuz/go-chat-rest-api/database/connection"

	"fmt"
)

// History get
func History(email string, limit int) ([]HistoryResult, error) {
	sqlUser := "SELECT U1.`from` AS `target`, MAX(U1.send_at) AS `last` "
	sqlUser += "FROM chats U1 "
	sqlUser += "WHERE `to` = '%s' "
	sqlUser += "GROUP BY U1.`from` "
	sqlUser += "UNION "
	sqlUser += "SELECT U2.`to`, MAX(U2.send_at) "
	sqlUser += "FROM chats U2 "
	sqlUser += "WHERE `from` = '%s' "
	sqlUser += "GROUP BY U2.`to`"
	sqlUser = fmt.Sprintf(sqlUser, email, email)
	sqlUserFormat := "SELECT U3.`target`, MAX(U3.`last`) AS `last` "
	sqlUserFormat += "FROM (%s) U3 GROUP BY U3.`target`"
	sqlUser = fmt.Sprintf(sqlUserFormat, sqlUser)

	sqlHistory := "SELECT U3.*, C.`text` "
	sqlHistory += "FROM (%s) U3 "
	sqlHistory += "LEFT JOIN chats C "
	sqlHistory += "ON (C.from = U3.target "
	sqlHistory += "OR C.to = U3.target) "
	sqlHistory += "AND C.send_at = U3.last "
	sqlHistory += "ORDER BY U3.last DESC "
	sqlHistory = fmt.Sprintf(sqlHistory, sqlUser)

	if limit > 0 {
		sqlHistory += fmt.Sprintf("LIMIT %d", limit)
	}

	var history []HistoryResult
	result := Connection.Get().Raw(sqlHistory).Scan(&history)

	reversed := make([]HistoryResult, 0, len(history))
	for i := len(history) - 1; i >= 0; i-- {
		reversed = append(reversed, history[i])
	}

	Connection.Close()

	return reversed, result.Error
}
