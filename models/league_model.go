package models

import "github.com/sunp13/dbtool"

type leagueModel struct {
}

func (m *leagueModel) AddList(leagueID, leagueName string) (res int64, err error) {
	sql := `
	insert into b365api.league_info(
		league_id,
		league_name
	) values (?,?)
	`
	params := []interface{}{
		leagueID,
		leagueName,
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return
}
