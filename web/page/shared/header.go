package shared

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func PageHeader() Node {
	return Div(Class("bg-slate-900 text-white border-b border-slate-800 sticky top-0 z-50"),
		Container(false,
			Div(Class("flex items-center justify-between h-16"),
				Div(Class("flex items-center space-x-8"),
					// Logo trắng nổi bật trên nền tối
					A(Href("/"), Class("text-xl font-extrabold tracking-tight text-white"), Text("⚡ VS-SoSanh")),
					Div(
						Class("hidden md:flex space-x-6"),
						// Link màu xám nhạt, hover sáng lên
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
	return A(Class("text-slate-300 hover:text-white transition-colors font-medium text-sm uppercase tracking-wide"), Href(href), Text(text))
}
