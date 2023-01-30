package cmd

import "log"

func Setup(files []string) int {
	if hasGitRepository(files) == false {
		log.Println("[ERROR]\t No git repository could be found.")
		return 1
	}

	return 0
}

func hasGitRepository(files []string) bool {
	for _, file := range files {
		if file == ".git" {
			return true
		}
	}

	return false
}
