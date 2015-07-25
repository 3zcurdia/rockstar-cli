package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// Repo general information about repo
type Repo struct {
	DirPath  string
	FilePath string
}

func randomName() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v-%v-%v", adjs[rand.Intn(len(adjs))], nouns[rand.Intn(len(nouns))], repos[rand.Intn(len(repos))])
}

func newRepo(filename string) Repo {
	dirPath := randomName()

	os.Mkdir(dirPath, 0755)
	os.Chdir(dirPath)
	exec.Command("git", "init", ".").Run()

	return Repo{DirPath: dirPath, FilePath: filename}
}

func (r *Repo) appendCommit(data string, date time.Time) {
	err := ioutil.WriteFile(r.FilePath, []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Setenv("GIT_AUTHOR_DATE", date.Format(time.RFC3339))
	os.Setenv("GIT_COMMITTER_DATE", date.Format(time.RFC3339))

	exec.Command("git", "add", r.FilePath).Run()
	exec.Command("git", "commit", "-m", messages[rand.Intn(len(messages))]).Run()
}

func main() {
	code := "writeln('Go is Awesome!!!')"
	repo := newRepo("main.go")
	repo.appendCommit(code, time.Now().Add(-24*time.Hour))

	for i := -500; i < 0; i++ {
		d := time.Now().Add(time.Duration(i*24) * time.Hour)
		if d.Weekday() == time.Sunday {
			continue
		}
		for j := 0; j < rand.Intn(10); j++ {
			authorDate := time.Date(d.Year(), d.Month(), d.Day(), int(rand.NormFloat64()*3.0+12.0), rand.Intn(59), rand.Intn(59), 0, d.Location())
			uid, err := uuid.NewV5(uuid.NamespaceURL, []byte(time.Now().Format(time.RFC3339Nano)))
			commitData := fmt.Sprintf("%s", uid)
			if err != nil {
				continue
			}
			// fmt.Printf("%v - %v\n", authorDate, commitData)
			repo.appendCommit(commitData, authorDate)
		}
		fmt.Print(".")
	}
	repo.appendCommit(code, time.Now())
	os.Setenv("GIT_AUTHOR_DATE", "")
	os.Setenv("GIT_COMMITTER_DATE", "")

	fmt.Println("\nNow you are a goddamn rockstar!")
}
