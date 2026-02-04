package cli

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "gowaf",
		Short: "gowaf - high-performance WAF fingerprinting for mass scanning",
	}
	root.AddCommand(scanCmd())
	return root
}
