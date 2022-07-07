package services

import (
	"fmt"
	"wuju_league/entity"
	"wuju_league/global"
	"wuju_league/models"
	"wuju_league/utils"

	"github.com/rs/zerolog/log"
)

type upComingService struct{}

// 获取数据
func (s *upComingService) GetData(c *entity.Conf, page int) (bool, error) {
	url := fmt.Sprintf(c.UpComing, c.Token, page)
	respByte, err := utils.HTTPGet(url)
	if err != nil {
		return false, err
	}

	var resp entity.RespComming
	err = global.JSON.Unmarshal(respByte, &resp)
	if err != nil {
		return false, err
	}
	if resp.Success != 1 {
		return false, fmt.Errorf("resp_comm.success != 1 (%d)", resp.Success)
	}
	// 本页无数据了. 到头了
	if len(resp.Results) == 0 {
		return true, nil
	}
	succCount := 0
	for _, v := range resp.Results {
		_, err := models.LeagueModel.AddList(
			SnowFlakeService.NextID(),
			v.League.ID,
			v.League.Name,
		)
		if err != nil {
			log.Error().Str("err", err.Error()).Send()
			continue
		}
		succCount++
	}
	log.Info().Int("page", page).Int("total", len(resp.Results)).Int("succ", succCount).Send()
	return false, nil
}
