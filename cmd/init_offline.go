package cmd

import (
	"github.com/kubesphere/kubekey/pkg/offline"
	"github.com/kubesphere/kubekey/pkg/util"
	"github.com/spf13/cobra"
)

var offlineCmd = &cobra.Command{
	Use:   "offline",
	Short: "Init operating system offline",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := util.InitLogger(opt.Verbose)
		return offline.Init(opt.ClusterCfgFile, opt.SourcesDir, opt.AddImagesRepo, logger)
	},
}

func init() {
	initCmd.AddCommand(offlineCmd)
	offlineCmd.Flags().StringVarP(&opt.ClusterCfgFile, "filename", "f", "", "Path to a configuration file")
	offlineCmd.Flags().StringVarP(&opt.SourcesDir, "sources", "s", "", "Path to the dependencies' dir")
	offlineCmd.Flags().BoolVarP(&opt.AddImagesRepo, "add-images-repo", "", false, "Create a local images registry")
}
