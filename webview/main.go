package main

import "github.com/webview/webview"

// simple app
func main() {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)

	w.Navigate("file:///Users/fedora/repo/cookbook/webview/simple.html")
	w.Run()

}
