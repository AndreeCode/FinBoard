package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type packageCoverage struct {
	name     string
	coverage float64
}

func main() {
	fmt.Println("========================================")
	fmt.Println("        FINBOARD TEST COVERAGE         ")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("Running all unit tests with coverage...")
	fmt.Println()

	cmd := exec.Command("go", "test", "./src/core/...", "./src/modules/...", "-cover")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error running tests: %v\n", err)
		fmt.Println(string(output))
		os.Exit(1)
	}

	lines := strings.Split(string(output), "\n")
	var packages []packageCoverage
	var totalCoverage float64
	var count int

	re := regexp.MustCompile(`(ok|\?)\s+(\S+)\s+.*coverage:\s+(\d+\.?\d*)%`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) >= 4 {
			name := matches[2]
			coverage, _ := strconv.ParseFloat(matches[3], 64)

			if name != "" && coverage > 0 {
				packages = append(packages, packageCoverage{name: name, coverage: coverage})
				totalCoverage += coverage
				count++
			}
		}
	}

	sort.Slice(packages, func(i, j int) bool {
		if packages[i].coverage == packages[j].coverage {
			return packages[i].name < packages[j].name
		}
		return packages[i].coverage > packages[j].coverage
	})

	fmt.Println("----------------------------------------")
	fmt.Println("PACKAGE COVERAGE")
	fmt.Println("----------------------------------------")

	for _, p := range packages {
		bar := getCoverageBar(p.coverage)
		status := "✓"
		statusColor := "\033[32m"
		if p.coverage < 50 {
			status = "✗"
			statusColor = "\033[31m"
		} else if p.coverage < 80 {
			status = "!"
			statusColor = "\033[33m"
		}

		fmt.Printf("%s%s %-45s %.1f%% %s\n",
			statusColor, status, p.name, p.coverage, bar)
	}

	fmt.Println("----------------------------------------")

	if count > 0 {
		avgCoverage := totalCoverage / float64(count)
		avgBar := getCoverageBar(avgCoverage)

		fmt.Println()
		fmt.Println("========================================")
		fmt.Printf("  \033[1mOVERALL COVERAGE: %.1f%% %s\033[0m\n",
			avgCoverage, avgBar)
		fmt.Println("========================================")

		if avgCoverage >= 90 {
			fmt.Println("\033[32m✓ Target 90% ACHIEVED!\033[0m")
		} else {
			fmt.Printf("\033[33m! Target 90%% - %.1f%% remaining\033[0m\n", 90-avgCoverage)
		}
	}

	fmt.Println()
}

func getCoverageBar(coverage float64) string {
	const barLength = 20
	filled := int(coverage * float64(barLength) / 100)
	empty := barLength - filled

	bar := "\033[36m["
	for i := 0; i < filled; i++ {
		bar += "="
	}
	for i := 0; i < empty; i++ {
		bar += "."
	}
	bar += "]\033[0m"

	return bar
}
