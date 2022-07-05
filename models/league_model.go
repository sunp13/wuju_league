package models

import (
	"time"

	"github.com/sunp13/dbtool"
)

type leagueModel struct {
}

func (m *leagueModel) AddList(leagueID, leagueName string) (res int64, err error) {
	sql := `
	insert into b365api.league_info(
		league_id,
		league_name,
		add_time
	) values (?,?,?)
	`
	params := []interface{}{
		leagueID,
		leagueName,
		time.Now().Format("2006-01-02 15:04:05"),
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return
}
