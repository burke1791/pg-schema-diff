package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func buildExportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:	"export",
		Short:	"Exports schema objects as individual files into a target directory",
	}

	// connFlags := createConnFlags(cmd)
	exportFlags := createExportFlags(cmd)
	cmd.RunE = func(_ *cobra.Command, _ []string) error {
		// logger := log.SimpleLogger()
		// connConfig, err := parseConnConfig(*connFlags, logger)
		// if err != nil {
		// 	return err
		// }

		exportConfig, err := parseExportConfig(*exportFlags)
		if err != nil {
			return err
		}

		fmt.Printf("export command | %s\n", "test")
		fmt.Printf("schema: %s\n", exportConfig.sourceSchema)

		return nil
	}
	return cmd
}

type (
	exportFlags struct {
		sourceSchema string
		// targetDir string
	}

	exportConfig struct {
		sourceSchema string
		// targetDir string
	}
)

func createExportFlags(cmd *cobra.Command) *exportFlags {
	flags := &exportFlags{}
	
	cmd.Flags().StringVarP(&flags.sourceSchema, "source-schema", "s", "public", "Indicates which schema to export")
	// cmd.Flags().StringVar(&flags.targetDir, "destination")

	return flags
}

func parseExportConfig(e exportFlags) (exportConfig, error) {
	return exportConfig {
		sourceSchema: e.sourceSchema,
	}, nil
}