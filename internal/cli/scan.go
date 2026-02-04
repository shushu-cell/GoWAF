package cli

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/shushu-cell/GoWAF/internal/scanner"
	"github.com/shushu-cell/GoWAF/internal/utils"
	"github.com/spf13/cobra"
)

type scanOptions struct {
	Input   string
	Workers int
	Timeout time.Duration
	Output  string
}

func scanCmd() *cobra.Command {
	opt := scanOptions{}

	cmd := &cobra.Command{
		Use:   "scan [url]",
		Short: "Scan a single URL or a targets file to detect WAF (passive mode)",
		Args: func(cmd *cobra.Command, args []string) error {
			if opt.Input == "" && len(args) == 0 {
				return errors.New("provide a URL argument or -i targets.txt")
			}
			if opt.Input != "" && len(args) > 0 {
				return errors.New("use either a URL argument OR -i targets.txt, not both")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var targets []string
			if opt.Input != "" {
				ts, err := utils.ReadTargets(opt.Input)
				if err != nil {
					return err
				}
				targets = ts
			} else {
				targets = []string{args[0]}
			}

			out := os.Stdout
			if opt.Output != "" {
				f, err := os.Create(opt.Output)
				if err != nil {
					return err
				}
				defer f.Close()
				out = f
			}

			s := scanner.New(scanner.Config{
				Workers: opt.Workers,
				Timeout: opt.Timeout,
			})

			results := s.ScanAll(targets)
			for r := range results {
				fmt.Fprintln(out, r.JSON())
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&opt.Input, "input", "i", "", "Targets file (txt, one target per line)")
	cmd.Flags().IntVarP(&opt.Workers, "workers", "w", 200, "Number of concurrent workers")
	cmd.Flags().DurationVarP(&opt.Timeout, "timeout", "t", 6*time.Second, "Per-target timeout (e.g. 6s, 2s)")
	cmd.Flags().StringVarP(&opt.Output, "output", "o", "", "Output file (jsonl). Default: stdout")

	return cmd
}
