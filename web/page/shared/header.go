package shared

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func PageHeader() Node {
	return Div(Class("bg-indigo-600 text-white shadow"),
		Container(false,
			Div(Class("flex items-center space-x-4 h-8"),
				HeaderLink("/", "Home"),
				HeaderLink("/about", "About"),
			),
		),
	)
}

func HeaderLink(href, text string) Node {
	return A(Class("hover:text-indigo-300"), Href(href), Text(text))
}

func Container(padY bool, children ...Node) Node {
	return Div(
		Classes{
			"max-w-7xl mx-auto":     true,
			"px-4 md:px-8 lg:px-16": true,
			"py-4 md:py-8":          padY,
		},
		Group(children),
	)
}
