package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	flagTags = flag.StringToString("tags", nil, "set tags in key=value format")
)

func main() {
	flag.Parse()
	viper.BindPFlags(flag.CommandLine)
	fmt.Println("pflag:", flagTags)
	fmt.Println("viper:", viper.GetStringMapString("tags"))
	fmt.Println()
}
