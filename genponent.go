package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BruceJi7/genponent/funcs"
	"github.com/BruceJi7/genponent/templates"
)

func main() {

	fmt.Println("GENPONENT!")
	if len(os.Args) > 1 {

		var javascriptFlag = flag.Bool("j", false, "Generate a JavaScript component instead of TypeScript")
		var cssFlag = flag.Bool("c", false, "Generate a component with CSS instead of SCSS")
		var noStyleFlag = flag.Bool("s", false, "Generate a component with no style")
		var noPropsFlag = flag.Bool("p", false, "Generate a component without a Props declaration")
		flag.Parse()
		remainingArgs := flag.Args()

		if len(remainingArgs) == 0 {
			fmt.Println("Error: No component name entered")
			return
		}

		for _, name := range remainingArgs {
			fmt.Println("Generating React Component: " + name)
			request := templates.NewComponentRequest(name, *cssFlag, *javascriptFlag, *noPropsFlag, *noStyleFlag)
			funcs.CreateComponent(request)
		}
		fmt.Println("The End.")

	} else {
		fmt.Println("Ener a name for your component")
		fmt.Println("e.g: 'genponent NavBar'")
	}
}
