package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/johnmccabe/go-bitbar"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	lib "github.com/mjohnsey/go-time/lib"
	"github.com/spf13/viper"
)

var cfgFile string
var ignore bool

var rootCmd = &cobra.Command{
	Use:   "go-time",
	Short: "Worldclock built for bitbar",
	Run: func(cmd *cobra.Command, args []string) {
		entries := viper.Get("time_entries").([]interface{})
		log.Info(entries)
		if !viper.GetBool("ignore") && entries == nil {
			log.Fatalln("Please add entries to time_entries or ignore using -i")
		}

		now := time.Now()

		app := bitbar.New()
		app.StatusLine(fmt.Sprintf("%s UTC", now.UTC().Format(lib.LocationTime{}.TimeFormat24())))

		if entries != nil {
			locs := make([]*lib.LocationTime, 0)

			for _, entry := range entries {
				e := entry.(map[string]interface{})
				locs = append(locs, lib.LocationTime{}.NewLocationTime(e["name"].(string), e["tz"].(string)))
			}
			submenu := app.NewSubMenu()
			for _, t := range locs {
				submenu.Line(t.PrettyPrint(&now))
			}
		}

		app.Render()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-time.toml)")
	rootCmd.PersistentFlags().BoolVarP(&ignore, "ignore", "i", false, "Allows you to run without specifying your own timezones and returns only UTC time")
	viper.BindPFlag("ignore", rootCmd.PersistentFlags().Lookup("ignore"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-time" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-time")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
}
