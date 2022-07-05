package entity

type Conf struct {
	Token            string `yaml:"token"`
	UpComing         string `yaml:"upcoming"`
	UpComingInterval int    `yaml:"upcoming_interval"`
}
