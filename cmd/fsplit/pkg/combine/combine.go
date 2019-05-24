package combine

import (
	"os"

	"github.com/shohi/glory/cmd/fsplit/pkg/actor"
	"github.com/shohi/glory/cmd/fsplit/pkg/config"
	"github.com/shohi/glory/cmd/fsplit/pkg/util"
)

type combine struct {
	conf config.Config
	args []string
}

func New(conf config.Config, args []string) actor.Actor {
	return &combine{
		conf: conf,
		args: args,
	}
}

func (c *combine) Act() error {
	// TODO: check file existence
	finalName := util.FinalName(c.args[0])
	f, err := os.Create(finalName)

	if err != nil {
		return err
	}
	defer f.Close()
	err = c.merge(f)
	return err
}

// TODO: sort file list before merge
func (c *combine) merge(file *os.File) error {
	for _, fp := range c.args {
		f, err := os.Open(fp)
		if err != nil {
			return err
		}
		defer f.Close()

		info, _ := f.Stat()
		var limit int64 = info.Size()

		err = util.AppendFile(file, f, limit)
		if err != nil {
			return err
		}
	}

	return nil
}
