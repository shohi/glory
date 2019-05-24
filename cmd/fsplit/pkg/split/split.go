package split

import (
	"github.com/shohi/glory/cmd/fsplit/pkg/actor"
	"github.com/shohi/glory/cmd/fsplit/pkg/config"
	"github.com/shohi/glory/cmd/fsplit/pkg/util"
)

func New(conf config.Config, args []string) actor.Actor {
	return &split{
		conf: conf,
		args: args,
	}
}

type split struct {
	conf config.Config
	args []string
}

func (s *split) Act() error {
	// TODO: add multiple input support
	err := util.SplitFile(s.args[0], s.conf.Number)
	return err
}
