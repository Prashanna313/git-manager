package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// Mock exec.Command for testing
var execCommand = exec.Command

func TestIsGitRepo(t *testing.T) {
	// Setup: Create a temporary directory and a .git subdirectory
	dir := t.TempDir()
	gitDir := filepath.Join(dir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatalf("Failed to create .git directory: %v", err)
	}

	// Test if isGitRepo correctly identifies the .git directory
	if !isGitRepo(dir) {
		t.Errorf("Expected %s to be identified as a Git repository", dir)
	}

	// Cleanup is handled by t.TempDir
}

func TestDeleteMergedBranches(t *testing.T) {
	// Setup: Create a mock exec.Command function
	execCommand = func(name string, args ...string) *exec.Cmd {
		return exec.Command("echo", "mock command")
	}

	// Define a temporary repo path
	repoPath := t.TempDir()

	// Run deleteMergedBranches and capture output
	deleteMergedBranches(repoPath)

	// Verify: Ideally, we'd want to assert the behavior, but since we're printing, manual verification is needed
	// This can be extended to capture stdout and assert the output
}

func TestSwitchUser(t *testing.T) {
	// Setup: Create a mock exec.Command function
	execCommand = func(_ string, _ ...string) *exec.Cmd {
		return exec.Command("echo", "mock command")
	}

	// Define a temporary repo path
	repoPath := t.TempDir()

	// Run switchUser and capture output
	switchUser(repoPath, "testuser", "testuser@example.com")

	// Verify: Ideally, we'd want to assert the behavior, but since we're printing, manual verification is needed
	// This can be extended to capture stdout and assert the output
}
