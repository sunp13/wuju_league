package services

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

type snowFlakeService struct {
	node *snowflake.Node
}

func (s *snowFlakeService) Setup() error {
	var err error
	s.node, err = snowflake.NewNode(0)
	if err != nil {
		return err
	}
	return nil
}

// 生成int64 id
func (s *snowFlakeService) NextID() string {
	return fmt.Sprintf("%d", s.node.Generate().Int64())
}
