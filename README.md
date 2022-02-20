# GENPONENT

A Go-based program.
Call the program from your commandline, and add the name of the component you want to create.

Pro Tip: In VSCode, right click the folder where you want to make the component, and click 'Open in Integrated Terminal'. Call genponent there.

Super-Pro Tip: Add genponent to your path so you can call it using `genponent` instead of "C:/User/Me/Programs/Downloads/genponent" etc

### Basic Functionality

Genponent will create the following for you:

A folder with name `$COMPONENT`

A TSX file in the folder, named index.tsx:

```ts
import React from "react";

import styles from "./$COMPONENT.module.scss";

interface Props {}

function $COMPONENT({}: Props) {
  return <div className={styles.$COMPONENT}></div>;
}

export default $COMPONENT;
```

A SCSS file, named $COMPONENT.module.scss:

```css
.$COMPONENT {
}
```

### Additional functionality

Add multiple component names with spaces between, and genponent will produce multiple components.

Optional flags include:
`-j : Use jsx instead of tsx` This will also remove the interface Props declaration.
`-c : Use css instead of scss` The style file as well as import will use .css instead of .scss
`-s : Remove styling completely` genponent will skip generating the style file completely, and will not add a style import line.
`-p : Remove props completely` genponent will neither add a Props declaration, nor add props as an argument to the function.
