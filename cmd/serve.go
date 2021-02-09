package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sonic3021/sonic-trans-post/imp"
)
var name string
var age int

var serveCmd= &cobra.Command{
	Use: "serve",
	Short: "start http server to trans post",
	Long: `Start http Server to trans post`,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("start server")
		fmt.Println(name)
		fmt.Println(age)
		imp.ServeHttp()

	},

}


func init(){
	serveCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	serveCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
	rootCmd.AddCommand(serveCmd)
}
