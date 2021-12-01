# GENPONENT

A Go-based program.
Call the program from your commandline, and add the name of the component you want to create
Genponent will create the following for you:

A folder with name `$COMPONENT`

A TSX file in the folder, named index.tsx:

```ts
import React from "react";

import styles from "./$COMPONENT.module.scss";

function $COMPONENT() {
  return <div className={styles.$COMPONENT}></div>;
}

export default $COMPONENT;
```

A SCSS file, named $COMPONENT.module.scss:

```css
.NAME {
}
```
