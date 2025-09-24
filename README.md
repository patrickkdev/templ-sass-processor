# templ-sass-processor

**Short description:** A Go tool that recursively finds `.templ` files in your project, and processes CSS code inside `<style>` blocks using the Sass CLI. It replaces the processed parts while keeping the rest of the code unchanged.

---

## Why I made this

I made this simple tool because I got addicted to writing [Nested CSS](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_nesting). 
Later I realized that this only works in modern browsers, and looks completely broken on older ones.

I didn’t want to rewrite dozens of rules by hand so instead, I discovered I could run my CSS through the Sass CLI. So I built a tiny Go utility that targets the <style> tags inside .templ files.

## Features

* Recursively discovers `.templ` files.
* Extracts `<style>` blocks and processes them with the `sass` CLI.
* Replaces the processed parts while preserving surrounding markup unchanged.

## Why use it

* Write nested, modern CSS inside your Templ components and ship flat CSS that works in browsers without native nesting support.

## Requirements

* `sass` CLI must be installed and available on `PATH`.

## Install

```bash
go install github.com/patrickkdev/templ-sass-processor@latest
```

## CLI examples

Basic: scan the current directory and replace `<style>` blocks with processed output in-place:

```bash
templ-sass-processor --src ./ --inplace
```

## Example — before & after

**Input (`component.templ`)**

```html
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
```

**Output (inline replacement)**

```html
<div class="card">
  <h3 class="title">Title</h3>
</div>

<style>
.card { padding: 1rem; background: white; }
.card .title { font-weight: 700; }
.card .title:hover { color: #555; }
</style>
```

PRs and issues welcome. Good first contributions:

* Add tests for specific nesting edge cases.
