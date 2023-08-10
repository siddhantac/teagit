package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

func gitBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	result, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(result))
	scanner.Scan()
	return scanner.Text()
}

func gitLog(num int) []list.Item {
	cmd := exec.Command("git", "log", "--oneline", fmt.Sprintf("-%d", num), "--pretty=%s;%h;%an;%cr;%d")
	result, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	items := make([]list.Item, 0)

	scanner := bufio.NewScanner(bytes.NewBuffer(result))
	for scanner.Scan() {
		line := scanner.Text()
		record := strings.Split(line, ";")

		desc := fmt.Sprintf("%s • %s • %s", record[1], record[2], record[3])
		if record[4] != "" {
			desc = fmt.Sprintf("%s • %s", desc, record[4])
		}
		item := item{title: record[0], desc: desc}
		items = append(items, item)

	}

	return items
}
