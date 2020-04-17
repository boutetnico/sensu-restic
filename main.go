package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

var (
	resticPath  string
	resticFlags string
	warnTime    int
	critTime    int
)

type Snapshot struct {
	Gid      int       `json:"gid"`
	HostName string    `json:"hostname"`
	Id       string    `json:"id"`
	Parent   string    `json:"parent"`
	Paths    []string  `json:"paths"`
	ShortId  string    `json:"short_id"`
	Tags     []string  `json:"tags"`
	Time     time.Time `json:"time"`
	Tree     string    `json:"tree"`
	Uid      int       `json:"uid"`
	Username string    `json:"username"`
}

const (
	StatusOK   = 0
	StatusWarn = 1
	StatusCrit = 2
	StatusUnkn = 3
)

const timeFormat string = "2006-01-02 15:04:05"

func main() {
	rootCmd := configureRootCommand()
	rootCmd.Execute()
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-restic-snapshot",
		Short: "Check the age of the last restic snapshot",
		RunE:  run,
	}

	cmd.Flags().StringVarP(&resticPath,
		"restic-path",
		"p",
		"/usr/local/bin/restic",
		"Path to the restic binary")

	cmd.Flags().StringVarP(&resticFlags,
		"restic-flags",
		"f",
		"--no-lock --no-cache",
		"Additional restic flags")

	cmd.Flags().IntVarP(&warnTime,
		"warning-age",
		"w",
		0,
		"Warn if latest snapshot is older than provided age in seconds")

	cmd.Flags().IntVarP(&critTime,
		"critical-age",
		"c",
		0,
		"critical if latest snapshot is older than provided age in seconds")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		_ = cmd.Help()
		fmt.Fprintf(os.Stderr, "UNKNOWN: invalid argument(s) received")
		os.Exit(StatusUnkn)
	}

	var snapshots []Snapshot

	snapshots = getSnapshots()

	checkMostRecentSnapshot(snapshots)

	return nil
}

func checkMostRecentSnapshot(snapshots []Snapshot) {
	if len(snapshots) == 0 {
		fmt.Println("CRITICAL: no Restic snapshots found")
		os.Exit(StatusCrit)
	}

	var snapshot Snapshot

	snapshot = getMostRecentSnapshot(snapshots)

	if snapshot.Time.Before(time.Now().Add(-time.Second * time.Duration(critTime))) {
		fmt.Println("CRITICAL: last snapshot date is", snapshot.Time.Format(timeFormat))
		os.Exit(StatusCrit)
	}

	if snapshot.Time.Before(time.Now().Add(-time.Second * time.Duration(warnTime))) {
		fmt.Println("WARNING: last snapshot date is", snapshot.Time.Format(timeFormat))
		os.Exit(StatusWarn)
	}

	fmt.Println("OK: last snapshot date is", snapshot.Time.Format(timeFormat))
	os.Exit(StatusOK)
}

func getMostRecentSnapshot(snapshots []Snapshot) Snapshot {
	sort.Slice(snapshots[:], func(i, j int) bool {
		return snapshots[i].Time.Before(snapshots[j].Time)
	})

	return snapshots[len(snapshots)-1]
}

func getSnapshots() []Snapshot {
	var command []string
	var out bytes.Buffer
	var snapshots []Snapshot

	if len(resticFlags) > 0 {
		var flags = strings.Split(resticFlags, " ")
		for i := 0; i < len(flags); i++ {
			command = append(command, flags[i])
		}
	}

	command = append(command, "snapshots")
	command = append(command, "--json")
	cmd := exec.Command(resticPath, command...)
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "UNKNOWN: failed to run Restic: %v\n", err)
		os.Exit(StatusUnkn)
	}

	err = json.Unmarshal(out.Bytes(), &snapshots)

	if err != nil {
		fmt.Fprintf(os.Stderr, "UNKNOWN: failed to parse Restic response: %v\n", err)
		os.Exit(StatusUnkn)
	}

	return snapshots
}
