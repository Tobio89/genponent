package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args[0])
	if len(os.Args) > 1 {
		componentName := os.Args[1]
		fmt.Println(componentName)
		createComponent(componentName)
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

