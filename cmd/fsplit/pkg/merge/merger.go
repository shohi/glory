package merge

import (
	"errors"
	"os"
	"path/filepath"
	"sort"

	"github.com/shohi/glory/cmd/fsplit/pkg/actor"
	"github.com/shohi/glory/cmd/fsplit/pkg/config"
	"github.com/shohi/glory/cmd/fsplit/pkg/util"
)

// TODO
// use pkg/err to wrap errors

type merger struct {
	conf  config.Config
	files Files
}

// New creates a new merger actor
func New(conf config.Config, args []string) actor.Actor {
	return &merger{
		conf:  conf,
		files: Files(args),
	}
}

func (m *merger) Act() error {
	err := m.resetFilesByPattern()
	if err != nil {
		return err
	}

	if len(m.files) == 0 {
		return errors.New("merge: no files to merge")
	}

	// sort file list by name before merge
	sort.Sort(m.files)

	finalName := util.FinalName(m.files[0])

	// NOTE: os.Create will check file existence
	f, err := os.Create(finalName)

	if err != nil {
		return err
	}

	err = m.merge(f)
	f.Close()

	return err
}

// resetFilesByPattern will reset file list for merge by pattern if set
// TODO: add more tests
func (m *merger) resetFilesByPattern() error {
	if len(m.conf.Pattern) == 0 {
		return nil
	}

	wd, err := os.Executable()
	if err != nil {
		return err
	}

	dir := filepath.Dir(wd)
	s, err := util.FindFiles(dir, m.conf.Pattern)
	if err != nil {
		return err
	}

	m.files = s
	return nil
}

func (m *merger) merge(file *os.File) error {
	for _, fp := range m.files {
		f, err := os.Open(fp)
		if err != nil {
			return err
		}
		defer f.Close()

		info, _ := f.Stat()
		var limit = info.Size()

		err = util.AppendFile(file, f, limit)
		if err != nil {
			return err
		}
	}

	return nil
}
