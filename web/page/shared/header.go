package shared

import (
	. "maragu.dev/gomponents"
	//. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func PageHeader() Node {
	return Div(Class("bg-indigo-600 text-white shadow sticky top-0 z-50"),

		Container(false,
			Div(Class("flex items-center justify-between h-16"),

				Div(Class("flex items-center space-x-8"),

					A(Href("/"), Class("text-xl font-bold"), Text("⚡ VS-SoSanh")),

					Div(
						Class("hidden md:flex space-x-4"),

						HeaderLink("/", "Home"),
						HeaderLink("/", "Compare"),
						HeaderLink("/", "Phone Finder"),
					),
				),
			),
		),
	)
}

func HeaderLink(href, text string) Node {
	return A(Class("hover:text-indigo-200 transition-colors font-medium"), Href(href), Text(text))
}
