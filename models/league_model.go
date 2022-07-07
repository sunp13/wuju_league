package models

import (
	"time"

	"github.com/sunp13/dbtool"
)

type leagueModel struct {
}

func (m *leagueModel) AddList(snowID, leagueID, leagueName string) (res int64, err error) {
	sql := `
	insert ignore into b365api.league_info(
		id,
		league_id,
		league_name,
		add_time
	) values (?,?,?,?)
	`
	params := []interface{}{
		snowID,
		leagueID,
		leagueName,
		time.Now().Format("2006-01-02 15:04:05"),
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return
}
