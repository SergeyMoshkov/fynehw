package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Привет, мир!")

	content := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Пример текста"),
		widget.NewButton("Нажми меня", func() {
			fmt.Println("Кнопка нажата")
		}),

		widget.NewButton("Закрыть", func() {
			w.Close()
		}),
		
		widget.NewAccordion(
			widget.NewAccordionItem("Как сам?", widget.NewLabel("1")),
			widget.NewAccordionItem("Все хорошо!", widget.NewLabel("2")),
			widget.NewAccordionItem("Могло быть и лучше!", widget.NewLabel("3")),
		),
	)

	w.SetContent(content)
	w.ShowAndRun()

	tidy()
}

func tidy() {
	fmt.Println("Tidy up")
}
