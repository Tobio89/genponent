package funcs

import (
	"fmt"
	"os"

	"github.com/BruceJi7/genponent/templates"
)

func writeReactFile(c *templates.ComponentRequest) {

	content := c.MakeString()

	var fileType = "tsx"
	if c.UseJavaScript {
		fileType = "jsx"
	}

	f, err := os.Create(fmt.Sprintf("./%s/index.%s", c.ComponentName, fileType))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}

func writeStyleFile(c *templates.ComponentRequest) {

	content := fmt.Sprintf(".%s {\n\n}", c.ComponentName)
	var fileType = "scss"
	if c.UseCSS {
		fileType = "css"
	}
	f, err := os.Create(fmt.Sprintf("./%s/%s.module.%s", c.ComponentName, c.ComponentName, fileType))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}

func CreateComponent(c *templates.ComponentRequest) {

	os.Mkdir(c.ComponentName, os.ModePerm)
	fmt.Println("Writing React file...")
	writeReactFile(c)
	if !c.SkipStyle {
		fmt.Println("Writing style file...")
		writeStyleFile(c)
	}
	fmt.Println("Component " + c.ComponentName + " created!")
}
