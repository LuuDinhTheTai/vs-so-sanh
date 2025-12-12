package shared

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func Page(title string, children ...Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Script(Src("https://cdn.tailwindcss.com?plugins=typography")),
			Script(Src("https://unpkg.com/htmx.org@1.9.10")),
		},
		Body: []Node{
			Class("bg-gray-50 min-h-screen flex flex-col"),

			PageHeader(),

			Div(Class("grow py-8"),
				Container(false,
					Group(children),
				),
			),

			PageFooter(),

			Div(ID("compare-tray")),
		},
	})
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
