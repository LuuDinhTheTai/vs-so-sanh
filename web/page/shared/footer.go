package shared

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func PageFooter() Node {
	return Div(Class("bg-gray-900 text-white shadow text-center h-16 flex items-center justify-center"),
		A(Href("https://www.gomponents.com"), Text("gomponents")),
	)
}
