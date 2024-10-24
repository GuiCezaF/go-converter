package main

import "github.com/GuiCezaF/go-converter/internal/converter"

func main() {
	vc := converter.NewVideoConverter()
	vc.Handle([]byte(`{"video_id": 1, "path": "/media/uploads/1"}`))
}
