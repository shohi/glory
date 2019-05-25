package split

import (
	"errors"

	"github.com/shohi/glory/cmd/fsplit/pkg/actor"
	"github.com/shohi/glory/cmd/fsplit/pkg/config"
	"github.com/shohi/glory/cmd/fsplit/pkg/util"
)

func New(conf config.Config, args []string) actor.Actor {
	return &split{
		conf:  conf,
		files: args,
	}
}

type split struct {
	conf  config.Config
	files []string
}

func (s *split) Act() error {
	if len(s.files) == 0 {
		return errors.New("split: no file to split")
	}

	// NOTE: add multiple input support
	for _, fp := range s.files {
		if err := util.SplitFile(fp, s.conf.Number); err != nil {
			return err
		}
	}

	return nil
}
