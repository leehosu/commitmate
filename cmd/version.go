package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// 빌드 시 ldflags로 주입됩니다
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "버전 정보를 출력합니다",
	Long:  `commitgen의 버전 정보를 출력합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("commitgen version %s\n", version)
		fmt.Printf("commit: %s\n", commit)
		fmt.Printf("built at: %s\n", date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
