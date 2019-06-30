package up

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/src-d/go-git.v4"
)

var ErrNotDir = errors.New("not directory")

func UpdateAll(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		filename := f.Name()
		if filename == "." || filename == ".." {
			continue
		}
		fp := filepath.Join(dir, filename)
		repo, err := git.PlainOpen(fp)
		if err != nil {
			log.Printf("open repo error, path: %v, err: %v", fp, err)
			continue
		}

		wt, err := repo.Worktree()
		if err != nil {
			log.Printf("get worktree error, path: %v, err: %v", fp, err)
			continue
		}
		log.Printf("starting update repo - [%v]", filename)
		err = wt.Pull(&git.PullOptions{})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			log.Printf("pull updates error, path: %v, err: %v", fp, err)
		}
	}

	return nil
}
