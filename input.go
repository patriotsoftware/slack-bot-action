package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseBool(v string) bool {
	if v == "" {
		return false
	}

	value, err := strconv.ParseBool(v)

	if err != nil {
		return false
	}

	return value
}

func ParseDestinations(d string) []string {
	return strings.Split(d, ",")
}

func ParseJobResults(r string) (string, error) {
	if r == "" {
		return "", nil
	}
	var result string
	lines := strings.Split(r, "\n")
	for _, l := range lines {
		job := strings.Split(l, ":")
		if len(job) != 2 {
			return "", errors.New("cannot parse destinations")
		}
		jobName := job[0]
		jobResult := job[1]
		// Blank job names/results mean it did not run
		if jobName == "" {
			continue
		}

		result += newResultLine(jobName, jobResult)

	}
	return result, nil
}

func newResultLine(jobName, result string) string {
	switch result {
	case "success":
		return fmt.Sprintf("✅ %s Succeeded.\n", jobName)
	case "failure":
		return fmt.Sprintf("❌ %s Failed.\n", jobName)
	default:
		return fmt.Sprintf("❕ %s Didn't Run.\n", jobName)
	}
}
