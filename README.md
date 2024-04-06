# HTMX with GoLang practice

This is a practice from the course [HTMX & Go](https://frontendmasters.com/courses/htmx)
taught by [ThePrimeagen](https://github.com/ThePrimeagen) in [Froned Masters](https://frontendmasters.com).

---

## Setup

This is a simple setup for this project.

### Requirements

For this project you will need to have installed:

- [Go](https://go.dev/)

After that is installed.

```bash
gh repo clone AlejandroSuero/htmx-with-golang-practice
cd htmx-with-golang-practice
go mod init yourwebsite.tv/yourname
```

Once you have **Go** installed in your system, you will need the following packages

Air, a hot-reloading package using echo.

```bash
go install github.com/cosmtrek/air@latest
```

To install the remaining packages:

```bash
go mod tidy
```

> This will install `echo/v4` and `echo/v4/middleware`

### Initializing the project

You can just type `air` in the terminal and go to [localhost:42069](http://localhost:42069/)
