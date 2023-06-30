package config

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	env       string
	JwtSecret string
	Config    *viper.Viper
)

func Init() (*viper.Viper, error) {
	// Initialize Flag and Viper
	Config = viper.New()
	flag.Int("flagname", 1234, "help message for flagname")
	flag.StringVar(&env, "env", "local", "local|dev|prod")

	// Parse and Bind Flags
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	Config.BindPFlags(pflag.CommandLine)

	// Read specified config file
	Config.SetConfigName(env)
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")
	Config.AddConfigPath("./config/")
	env = Config.GetString("app.env")
	JwtSecret = Config.GetString("app.jwt")
	err := Config.ReadInConfig()
	return Config, err

	// viper.AutomaticEnv()
	// viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
	// Declare var
	// env := viper.GetString("app.env")
	// port := viper.Get("port")
	// flagname := viper.Get("flagname")
	// producerbroker := viper.GetString("app.producerbroker")
	// consumerbroker := viper.GetString("app.consumerbroker")

	// // Print
	// fmt.Println("---------- Example ----------")
	// fmt.Println("port :", port)
	// fmt.Println("flag :", flagname)
	// fmt.Println("app.env :", env)
	// fmt.Println("app.producerbroker :", producerbroker)
	// fmt.Println("app.consumerbroker :", consumerbroker)
}
