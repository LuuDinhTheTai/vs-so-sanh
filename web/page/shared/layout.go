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
		},
		Body: []Node{Class("bg-gradient-to-b from-white to-indigo-100 bg-no-repeat"),
			Div(Class("min-h-screen flex flex-col justify-between"),
				PageHeader(),
				Div(Class("grow"),
					Container(true,
						Div(Class("prose prose-lg prose-indigo"),
							Group(children),
						),
					),
				),
				PageFooter(),
			),
		},
	})
}
