/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		test()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Define an interface that is exported by your package.
type Foo interface {
	GetValue() string // A function that'll return the value initialized with a default.
	GetVolume() string
	SetValue(v string) // A function that can update the default value.
	SetFieldValue(f string, v string)
}

// Define a struct type that is not exported by your package.
type foo struct {
	value  string
	volume string
}

// A factory method to initialize an instance of `foo`,
// the unexported struct, with a default value.
func NewFoo() Foo {
	return &foo{
		value:  "I am the DEFAULT value.",
		volume: "I am the volume.",
	}
}

// Implementation of the interface's `GetValue`
// for struct `foo`.
func (f *foo) GetValue() string {
	return f.value
}

// Implementation of the interface's `GetValue`
// for struct `foo`.
func (f *foo) GetVolume() string {
	return f.volume
}

// Implementation of the interface's `SetValue`
// for struct `foo`.
func (f *foo) SetValue(v string) {
	f.value = v
}

func (foo *foo) SetFieldValue(f string, v string) {
	reflect.ValueOf(foo).Elem().FieldByName(f).SetString(v)
}

func test() {

	f := NewFoo()
	fmt.Printf("value: `%s`\n", f.GetValue())
	f.SetValue("I am the UPDATED value.")
	fmt.Printf("value: `%s`\n", f.GetValue())
	f.SetFieldValue("volume", "I am the new volume value")
	fmt.Printf("volume: `%s`\n", f.GetVolume())
}
