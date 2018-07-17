# courses

courses

## [Start Learning React](https://egghead.io/lessons/react-write-a-hello-world-react-component)

Difficulty ⭐️️

### React.Children util

`React.Children.map(children, function[(thisArg)])`
`React.Children.forEach(children, function[(thisArg)])`
`React.Children.count(children)`
`React.Children.only(children)`

### cloneElement()

> clone element(with `key` and `ref`), and merge props shallowly

```js
React.cloneElement(element, [props], [...children]);
```

## [React Testing Cookbook](https://egghead.io/courses/react-testing-cookbook)

> outdated content

## [Learn Advanced CSS Layout Techniques](https://egghead.io/courses/learn-advanced-css-layout-techniques)

Difficulty ⭐️️⭐

### Target empty elements using the :empty pseudo-class

> 可用于 tooltip 等

```css
.Alert {
  border: 3px solid darkgreen;
  background-color: seagreen;
}

.Alert:empty {
  display: none;
}

/* or: */

.Alert:not(:empty) {
  border: 3px solid darkgreen;
  background-color: seagreen;
}
```

### Target Positional Elements Using \*-Of-Type CSS pseudo-classes

> 主要用于字体排版

```css
article p:first-of-type {
  font-size: 1.6rem;
  font-style: italic;
}

article img:last-of-type {
  border: 10px solid slategray;
}

article blockquote:only-of-type {
  border-left: 5px solid slategray;
  padding-left: 1rem;
  font-style: italic;
  font-weight: bold;
  color: slategray;
}

/* article p:nth-of-type(4) { */
/* article p:nth-of-type(4n) { */
article p:nth-of-type(4n + 3) {
  color: maroon;
}
```

### Create a fixed-fluid-fixed layout using CSS calc()

```css
nav {
  position: fixed;
  top: 0;
  left: 0;
  width: 5rem;
  height: 100%;
}

aside {
  position: fixed;
  top: 0;
  right: 0;
  height: 100%;
  width: 20rem;
}

main {
  margin-left: 5rem;
  width: calc(100% - 25rem);
}
```

### Dynamically Size Elements with Pure CSS

> viewport units(vw, vh)

```css
```

### Easily Reset Styles With a Single CSS value

> reset an element's styles to its parent's

```css
/* Global Styles */
body {
  color: #2e603c;
}

/* Global Button Styles  */
button {
  font-size: 1.4rem;
  color: white;
}

section button {
  color: unset;
  font-size: unset;
}
```

### Layout Responsive Fluid Columns Using CSS

> use column property

```css
nav {
  /* column-count: 4; */
  /* column-width: 150px; */
  /* equals to: */
  /* up to 4 columns which min-width is 150px */
  column: 4 15opx;
  column-gap: 3rem;

  column-rule: 1px dashed #ccc;
}

/* make h2 take up all of the column width and force a line break to make
all the columns start after it */
h2 {
  column-span: all;
}
```

### Control Image Aspect Ratio Using CSS

```css
img {
  width: 90vw;
  height: 75vh;
  border: 3px solid tomato;
  object-fit: contain | cover | none | scale-down;
  object-position: 50% 50%; /* by default */
}
```

### [Flexbox Fundamentals](https://egghead.io/courses/flexbox-fundamentals)

```css
div {
  display: flex;
  flex-direction: column | row(default) | column-reverse | row-reverse;
}
```

order 优先级，好像没什么用

```css
body {
  display: flex;
  flex-direction: column;
}

.main-footer {
  order: 1;
}

.main-nav {
  order: -1;
}
```

justify-content, align-items(align-self)

```css
:root {
  height: 100%;
}
body {
  font-family: "Open Sans", sans-serif;
  margin: 0;
  min-height: 100%;
  background: #333;
}
h1 {
  font-weight: 600;
  margin: 0 0 5px 0;
}
p {
  margin: 5px 0;
}
article {
  box-sizing: border-box;
  background: #fff;
  margin: 5px;
  flex: 1;
}
body {
  display: flex;
  flex-direction: column;
  /* Justify-content declares how to use the extra space along the direction the content is flowing. */
  /* Currently, the flex direction is set to column, so justify-content will affect how the articles use the */
  /* extra space from top to bottom without adjusting the actual size of the elements. */
  justify-content: start(default) | end | center | space-between | space-around;
}
```

```css
article {
  flex: 1;
}
body {
  flex-direction: row;
  /* Align-items declares how to use the extra space perpendicular to the flex direction. */
  align-items: stretch(default) | flex-start | flex-end | center | baseline;
}
.not-real {
  /* Align-self is the same as align-items, but where justify-content and align-items */
  /* are applied to the flexbox containers and affect all children, align-self is only */
  /* applied to specific individual children. */
  align-self: flex-end;
}
```

flex-basis

```css
body {
  margin: 0;
  display: flex;
}

.title-1 {
  background: #dd5f40;
  /* use flex-basis instead of width for flex display */
  /* width: 220px; */
  flex-basis: auto(default) | 150px;
}

.title-2 {
  background: #3d483a;
}
```

flex-grow, flex-shrink

```css
.title-1 {
  background: #dd5f40;
  flex-grow: 1;
}

.title-2 {
  background: #3d483a;
  flex-grow: 2;
}
```

```css
h1 {
  flex-basis: 200px;
  flex-shrink: 0;
}
```

shorthand

```css
.title-1 {
  background: #dd5f40;
  flex: 1;
  /* flex-grow: 1; */
  /* flex-shrink: 1; */
  /* flex-basis: 0; */
}

.title-2 {
  background: #3d483a;
  flex: 20px;
  /* flex-grow: 1; */
  /* flex-shrink: 1; */
  /* flex-basis: 20px; */
}

.title-3 {
  background: #468e5d;
  flex: 0 80px;
  /* flex-grow: 0; */
  /* flex-shrink: 1; */
  /* flex-basis: 80px; */
}
```

flex-wrap and align-content

```css
figure {
  margin: 0;
  display: block;
  height: 150px;
  flex: 0 150px;
  position: relative;
}

body {
  display: flex;
  flex-wrap: nowrap | wrap | wrap-reverse;
  justify-content: space-around;
  /* Align-content is only used with multi-line content. Make sure not  */
  /* to confuse it with align-items or align-self, as it won't do anything */
  /* on single line content. */
  align-content: flex-start | flex-end | center | space-between | space-around |
    stretch;
}
```

## [Style an Application from Start to Finish](https://egghead.io/courses/style-an-application-from-start-to-finish)

Difficulty ⭐️️

There are certain cues that users are looking for to be informed that elements are
interactive, like hover states and cursor changes.

Inform the User of Interactions with CSS Transitions(CSSTransitionGroup)
