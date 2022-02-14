package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"net/url"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start interactive configuration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		runConfig()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runConfig() {
	urlPrompt := promptui.Prompt{
		Label: "Enter Base URL (e.g. https://live.example.com)",
		Validate: func(s string) error {
			u, err := url.Parse(s)
			if err != nil || u.Scheme == "" || u.Host == "" {
				return errors.New("invalid URL")
			}
			return nil
		},
	}
	s, err := urlPrompt.Run()
	cobra.CheckErr(err)
	viper.Set("baseUrl", s)

	portPrompt := promptui.Prompt{
		Label:   "Enter Port (default: 8080)",
		Default: "8080",
		Validate: func(s string) error {
			p, err := strconv.Atoi(s)
			if err != nil || p < 1 || p > 65535 {
				return errors.New("invalid port")
			}
			return nil
		},
	}
	portS, err := portPrompt.Run()
	cobra.CheckErr(err)
	port, err := strconv.Atoi(portS)
	cobra.CheckErr(err)
	viper.Set("port", port)

	dirs := []string{
		"/etc/tum-live-monitor",
	}
	if home, err := os.UserHomeDir(); err == nil {
		dirs = append(dirs, home)
	}
	if workingDir, err := os.Getwd(); err == nil {
		dirs = append(dirs, workingDir)
	}
	configDirPrompt := promptui.Select{
		Label: "Enter Config Directory (default: ~/.config/live)",
		Items: dirs,
	}
	_, dir, err := configDirPrompt.Run()
	cobra.CheckErr(err)
	fmt.Println("Writing config to", dir+"/"+defaultCfgFileName)
	err = viper.WriteConfigAs(dir + "/" + defaultCfgFileName)
	cobra.CheckErr(err)
	fmt.Println(ansi.Green + "Successfully configured TUM Live Monitor" + ansi.Reset)
}
