// DockerLens: A Go utility to inspect Docker images for useful metadata

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/tabwriter"
)

// RunCommand executes a shell command and returns the output
func RunCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// HumanReadableSize converts size in bytes to MB, GB, etc.
func HumanReadableSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	floatSize := float64(size)
	unitIndex := 0
	for floatSize >= 1024 && unitIndex < len(units)-1 {
		floatSize /= 1024
		unitIndex++
	}
	return fmt.Sprintf("%.2f %s", floatSize, units[unitIndex])
}

// GetDockerImageSize returns the size of a Docker image in a human-readable format
func GetDockerImageSize(image string) string {
	output, err := RunCommand("docker", "image", "inspect", image, "--format", "{{.Size}}")
	if err != nil {
		return "Error retrieving size"
	}
	size, err := strconv.ParseInt(strings.TrimSpace(output), 10, 64)
	if err != nil {
		return "Invalid size format"
	}
	return HumanReadableSize(size)
}

// GetDockerImageLayers returns the layers of a Docker image
func GetDockerImageLayers(image string) string {
	output, err := RunCommand("docker", "history", "--format", "{{.ID}}: {{.Size}}", image)
	if err != nil {
		return "Error retrieving layers"
	}
	return strings.TrimSpace(output)
}

// GetBaseImage returns the base image of a Docker image by inspecting its history
func GetBaseImage(image string) string {
	output, err := RunCommand("docker", "history", "--format", "{{.CreatedBy}}", image)
	if err != nil {
		return "Error retrieving base image"
	}
	lines := strings.Split(output, "\n")
	if len(lines) > 0 {
		for _, line := range lines {
			if strings.Contains(line, "FROM ") {
				return strings.TrimSpace(strings.TrimPrefix(line, "/bin/sh -c #(nop) FROM "))
			}
		}
	}
	return "Base image not found"
}

// GetPythonVersion checks for Python version inside the image
func GetPythonVersion(image string) string {
	output, err := RunCommand("docker", "run", "--rm", image, "python", "--version")
	if err != nil {
		return "Python not found"
	}
	return strings.TrimSpace(output)
}

// GetPipVersion checks for Pip version inside the image
func GetPipVersion(image string) string {
	output, err := RunCommand("docker", "run", "--rm", image, "pip", "--version")
	if err != nil {
		return "Pip not found"
	}
	return strings.TrimSpace(output)
}

func main() {
	var image string
	fmt.Print("Enter Docker image name: ")
	fmt.Scanln(&image)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	fmt.Println("\nDockerLens - Docker Image Inspector")
	fmt.Fprintf(w, "%-20s	%s\n", "Property", "Value")
	fmt.Fprintf(w, "%-20s	%s\n", "Image", image)
	fmt.Fprintf(w, "%-20s	%s\n", "Size", GetDockerImageSize(image))
	fmt.Fprintf(w, "%-20s	%s\n", "Base Image", GetBaseImage(image))
	fmt.Fprintf(w, "%-20s	%s\n", "Python Version", GetPythonVersion(image))
	fmt.Fprintf(w, "%-20s	%s\n", "Pip Version", GetPipVersion(image))
	fmt.Fprintf(w, "%-20s	%s\n", "Layers", GetDockerImageLayers(image))
	w.Flush()
}