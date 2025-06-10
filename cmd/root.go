package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"main/server"
)

func init() {
	viper.SetDefault("mode", "dev")
	viper.SetDefault("port", 8081)

	rootCmd.PersistentFlags().String("mode", "dev", `mode of server, can be "prod" or "dev" or "demo"`)
	rootCmd.PersistentFlags().String("addr", "", "address of server")
	rootCmd.PersistentFlags().Int("port", 8081, "port of server")
	rootCmd.PersistentFlags().String("unix-sock", "", "path to the unix socket, overrides --addr and --port")

	if err := viper.BindPFlag("mode", rootCmd.PersistentFlags().Lookup("mode")); err != nil {
		panic(err)
	}

	if err := viper.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr")); err != nil {
		panic(err)
	}

	if err := viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port")); err != nil {
		panic(err)
	}

	if err := viper.BindPFlag("unix-sock", rootCmd.PersistentFlags().Lookup("unix-sock")); err != nil {
		panic(err)
	}
}

var (
	rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			runGoProject(cmd.Context())
		},
	}
)

func Execute() {
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		return
	}
}

func runGoProject(ctx context.Context) {
	server.StartServer(ctx, viper.GetString("addr"), viper.GetInt("port"))
}
