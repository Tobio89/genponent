package templates

import "fmt"

const importLine = "import React from \"react\";\n\n"
const interfaceLine = "interface Props{\n\n};\n\n"

type ComponentRequest struct {
	ComponentName string
	UseCSS        bool
	UseJavaScript bool
	SkipProps     bool
	SkipStyle     bool
}

func (c ComponentRequest) MakeString() string {

	// Initialise with import line
	reactFileContents := importLine

	// If using style...
	if !c.SkipStyle {

		// use css or scss
		var styleType string
		if c.UseCSS {
			styleType = "css"
		} else {
			styleType = "scss"
		}
		reactFileContents += fmt.Sprintf("import styles from \"./%s.module.%s\";\n\n", c.ComponentName, styleType)
	}

	// Add interface line
	if !c.SkipProps && !c.UseJavaScript {
		reactFileContents += interfaceLine
	}

	// Add props
	var propsInfix string
	if c.SkipProps {
		propsInfix = ""
	} else if c.UseJavaScript {
		propsInfix = "props"
	} else {
		propsInfix = "{}:Props"
	}
	reactFileContents += fmt.Sprintf("function %s(%s){\n\n", c.ComponentName, propsInfix)

	// Add return with div and styling
	var styling string
	if c.SkipStyle {
		styling = ""
	} else {
		styling = fmt.Sprintf(" className={styles.%s}", c.ComponentName)
	}
	reactFileContents += fmt.Sprintf("  return (\n    <div%s></div>\n  );\n\n", styling)

	reactFileContents += fmt.Sprintf("}\n\nexport default %s;", c.ComponentName)

	return reactFileContents

}

func NewComponentRequest(
	ComponentName string,
	UseCSS bool,
	UseJavaScript bool,
	SkipProps bool,
	SkipStyle bool,
) *ComponentRequest {

	c := ComponentRequest{
		ComponentName: ComponentName,
		UseCSS:        UseCSS,
		UseJavaScript: UseJavaScript,
		SkipProps:     SkipProps,
		SkipStyle:     SkipStyle,
	}
	return &c
}
