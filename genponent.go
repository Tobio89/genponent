package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args[0])
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
		// Todo - handle multiple filenames
		// Todo - --help arg

		fmt.Println("Flag -j value: ", *javascriptFlag)
		fmt.Println("Flag -c value: ", *cssFlag)
		fmt.Println("Flag -s value: ", *noStyleFlag)
		fmt.Println("Flag -p value: ", *noPropsFlag)

		// componentName := os.Args[1]
		// fmt.Println("Generating React Component: " + componentName)
		// createComponent(componentName)
	} else {
		fmt.Println("Ener a name for your component")
		fmt.Println("e.g: 'genponent NavBar'")
	}
}

func createComponent(componentName string) {

	os.Mkdir(componentName, os.ModePerm)
	writeReactFile(componentName)
	writeCSSFile(componentName)

}

func writeReactFile(componentName string) {

	content := fmt.Sprintf("import React from \"react\";\n\nimport styles from \"./%s.module.scss\";\n\nfunction %s() {\nreturn <div className={styles.%s}></div>;\n}\n\nexport default %s;", componentName, componentName, componentName, componentName)
	f, err := os.Create(fmt.Sprintf("./%s/index.tsx", componentName))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}
func writeCSSFile(componentName string) {

	content := fmt.Sprintf(".%s {\n\n}", componentName)
	f, err := os.Create(fmt.Sprintf("./%s/%s.module.scss", componentName, componentName))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}
