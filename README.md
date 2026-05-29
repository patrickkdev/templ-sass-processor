# templ-sass-processor

I made this simple tool because I was building a web app, and using [Nested CSS](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_nesting) inside <style> tags for styling.
Later I realized that nested CSS only works in modern browsers, and looks completely broken on older ones.

I didn’t want to rewrite all the css by hand so instead, I discovered I could flatten the CSS by running it through something like the Sass CLI. The problem is that the CSS lives inside .templ files, so you can’t just pipe the whole file through Sass.

I ended up putting together this small utility that:

* recursively finds .templ files
* extracts CSS inside <style> blocks
* runs it through the Sass CLI
* writes the processed CSS back in place without touching the rest of the file

## Features

* Recursively discovers `.templ` files.
* Extracts `<style>` blocks and processes them with the `sass` CLI.
* Replaces the processed parts while preserving surrounding markup unchanged.

## Why use it

* Flatten nested CSS inside your Templ components while preserving surrounding markup unchanged.

## Requirements

* `sass` CLI must be installed and available on `PATH`.

## Install

```bash
go install github.com/patrickkdev/templ-sass-processor@latest
```

## CLI examples

Basic: scan the current directory and replace `<style>` blocks with processed output in-place:

```bash
templ-sass-processor .
```

## Example — before & after

**Input (`component.templ`)**

```html
...
<div class="card">
  <h3 class="title">Title</h3>
</div>

<style>
.card {
  padding: 1rem;
  background: white;
  .title {
    font-weight: 700;
    &:hover { color: #555; }
  }
}
</style>
...
```

**Output (inline replacement)**

```html
...
<div class="card">
  <h3 class="title">Title</h3>
</div>

<style>
.card { padding: 1rem; background: white; }
.card .title { font-weight: 700; }
.card .title:hover { color: #555; }
</style>
...
```

PRs and issues welcome.
