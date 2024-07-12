package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "git-cli"}

	rootCmd.PersistentFlags().StringP("path", "p", ".", "Path to the directory containing repositories")
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(deleteMergedCmd)
	rootCmd.AddCommand(switchUserCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func executeCommand(repoPath string, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in %s: %v\n", repoPath, err)
	}
}

func processRepositories(path string, cmdFunc func(repoPath string)) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && isGitRepo(path) {
			cmdFunc(path)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking the path %q: %v\n", path, err)
		os.Exit(1)
	}
}

func isGitRepo(path string) bool {
	gitPath := filepath.Join(path, ".git")
	info, err := os.Stat(gitPath)
	return err == nil && info.IsDir()
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch latest changes from remote for all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		processRepositories(path, func(repoPath string) {
			fmt.Printf("Fetching in %s\n", repoPath)
			executeCommand(repoPath, "git", "fetch")
		})
	},
}

var deleteMergedCmd = &cobra.Command{
	Use:   "delete-merged",
	Short: "Delete merged branches for all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		processRepositories(path, deleteMergedBranches)
	},
}

func deleteMergedBranches(repoPath string) {
	fmt.Printf("Deleting merged branches in %s\n", repoPath)
	// Fetch latest changes
	executeCommand(repoPath, "git", "fetch")

	// List merged branches
	out, err := exec.Command("git", "branch", "--merged").Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in %s: %v\n", repoPath, err)
		return
	}

	branches := strings.Split(string(out), "\n")
	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if branch != "" && branch != "main" && branch != "master" {
			fmt.Printf("Deleting branch %s in %s\n", branch, repoPath)
			executeCommand(repoPath, "git", "branch", "-d", branch)
		}
	}
}

var switchUserCmd = &cobra.Command{
	Use:   "switch-user [name] [email]",
	Short: "Switch git user for all repositories",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		processRepositories(path, func(repoPath string) {
			fmt.Printf("Switching user in %s\n", repoPath)
			switchUser(repoPath, args[0], args[1])
		})
	},
}

func switchUser(repoPath, name, email string) {
	executeCommand(repoPath, "git", "config", "user.name", name)
	executeCommand(repoPath, "git", "config", "user.email", email)
}
