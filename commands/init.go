package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Init(path string) error {
	//join the current working directory with .git-go
	gitWorkingDir := filepath.Join(path, ".git-go")

	//0- special permission
	//7 - 4+2+1 that means owner has read write and execute permisson
	//5 - 4+0+1 group has read no write but executer permission
	//5 - other also has the same
	if err := os.MkdirAll(gitWorkingDir, 0755); err != nil {
		return fmt.Errorf("Failed to create .go-git directory: %v", err)
	}
}
