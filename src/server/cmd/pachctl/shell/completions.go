package shell

import (
	"fmt"
	"log"

	prompt "github.com/c-bata/go-prompt"
	"github.com/docker/go-units"
	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/server/pkg/cmdutil"
)

func RepoCompletion(_, text string) []prompt.Suggest {
	c, err := client.NewOnUserMachine("user-completion")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	ris, err := c.ListRepo()
	if err != nil {
		log.Fatal(err)
	}
	var result []prompt.Suggest
	for _, ri := range ris {
		result = append(result, prompt.Suggest{
			Text:        ri.Repo.Name,
			Description: fmt.Sprintf("%s (%s)", ri.Description, units.BytesSize(float64(ri.SizeBytes))),
		})
	}
	return result
}

func BranchCompletion(_, text string) []prompt.Suggest {
	c, err := client.NewOnUserMachine("user-completion")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	partialFile := cmdutil.ParsePartialFile(text)
	var result []prompt.Suggest
	if partialFile.Commit.Repo.Name == "" || // nothing typed yet
		(len(partialFile.Commit.ID) == 0 && text[len(text)-1] != '@') { // partial repo typed
		ris, err := c.ListRepo()
		if err != nil {
			log.Fatal(err)
		}
		for _, ri := range ris {
			result = append(result, prompt.Suggest{
				Text:        fmt.Sprintf("%s@", ri.Repo.Name),
				Description: fmt.Sprintf("%s (%s)", ri.Description, units.BytesSize(float64(ri.SizeBytes))),
			})
		}
	} else if partialFile.Commit.ID == "" || // repo@ typed, no commit/branch yet
		(len(partialFile.Path) == 0 && text[len(text)-1] != ':') { // partial commit/branch typed
		bis, err := c.ListBranch(partialFile.Commit.Repo.Name)
		if err != nil {
			log.Fatal(err)
		}
		for _, bi := range bis {
			result = append(result, prompt.Suggest{
				Text:        fmt.Sprintf("%s@%s:", partialFile.Commit.Repo.Name, bi.Branch.Name),
				Description: fmt.Sprintf("(%s)", bi.Head.ID),
			})
		}
	}
	return result
}

func FileCompletion(_, text string) []prompt.Suggest {
	c, err := client.NewOnUserMachine("user-completion")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	partialFile := cmdutil.ParsePartialFile(text)
	var result []prompt.Suggest
	if partialFile.Commit.Repo.Name == "" || // nothing typed yet
		(len(partialFile.Commit.ID) == 0 && text[len(text)-1] != '@') { // partial repo typed
		ris, err := c.ListRepo()
		if err != nil {
			log.Fatal(err)
		}
		for _, ri := range ris {
			result = append(result, prompt.Suggest{
				Text:        fmt.Sprintf("%s@", ri.Repo.Name),
				Description: fmt.Sprintf("%s (%s)", ri.Description, units.BytesSize(float64(ri.SizeBytes))),
			})
		}
	} else if partialFile.Commit.ID == "" || // repo@ typed, no commit/branch yet
		(len(partialFile.Path) == 0 && text[len(text)-1] != ':') { // partial commit/branch typed
		bis, err := c.ListBranch(partialFile.Commit.Repo.Name)
		if err != nil {
			log.Fatal(err)
		}
		for _, bi := range bis {
			result = append(result, prompt.Suggest{
				Text:        fmt.Sprintf("%s@%s:", partialFile.Commit.Repo.Name, bi.Branch.Name),
				Description: fmt.Sprintf("Head: %s", bi.Head.ID),
			})
		}
	} else { // typing the file
		fis, err := c.GlobFile(partialFile.Commit.Repo.Name, partialFile.Commit.ID, partialFile.Path+"*")
		if err != nil {
			log.Fatal(err)
		}
		for _, fi := range fis {
			result = append(result, prompt.Suggest{
				Text: fmt.Sprintf("%s@%s:%s", partialFile.Commit.Repo.Name, partialFile.Commit.ID, fi.File.Path),
			})
		}
	}
	return result
}

func PipelineCompletion(_, text string) []prompt.Suggest {
	c, err := client.NewOnUserMachine("user-completion")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	pis, err := c.ListPipeline()
	if err != nil {
		log.Fatal(err)
	}
	var result []prompt.Suggest
	for _, pi := range pis {
		result = append(result, prompt.Suggest{
			Text:        pi.Pipeline.Name,
			Description: pi.Description,
		})
	}
	return result
}

func JobCompletion(_, text string) []prompt.Suggest {
	c, err := client.NewOnUserMachine("user-completion")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	jis, err := c.ListJob("", nil, nil, 0, false)
	if err != nil {
		log.Fatal(err)
	}
	var result []prompt.Suggest
	for _, ji := range jis {
		result = append(result, prompt.Suggest{
			Text:        ji.Job.ID,
			Description: ji.Pipeline.Name,
		})
	}
	return result
}
