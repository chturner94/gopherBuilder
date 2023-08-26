Current issue with the controls.go file is that we are getting an error whenever
we attempt to set the terminal to raw mode. This is discussed in this stack overflow
issue, and there is a resolution to use a package from containerd, and it's the
github.com/containerd/console package. Here is the [URL](https://stackoverflow.com/questions/58237670/terminal-raw-mode-does-not-support-arrows-on-windows)

