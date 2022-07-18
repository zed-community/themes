package gen

import (
	"log"
	"os"
	"strings"
	
	"text/scanner"
	
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

func Render() {
	snippetContent, err := os.ReadFile("./res/sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	var s scanner.Scanner
	s.Init(strings.NewReader(string(snippetContent)))
	s.Mode ^= scanner.SkipComments
	
	fontContent, err := os.ReadFile("./res/fonts/zed-mono-regular.ttf")
	if err != nil {
		log.Fatal(err)
	}

	font, err := truetype.Parse(fontContent)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 16})

	dc := gg.NewContext(400, 120)
	// Set font to Zed Mono defined above
	dc.SetFontFace(face)
	// Set background color
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	// Set base text color
	dc.SetRGB(1, 1, 1)

	highlightedFile(dc, s)
	
	dc.SavePNG("./res/samples/out.png")
}

func highlightedFile(dc *gg.Context, s scanner.Scanner) {
	x, y := 0.0, 0.0
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		log.Println(s.TokenText())
		switch (s.TokenText()) {
			case "\\":// Newline (\)
				dc.SetRGB(1, 1, 1) // Clear the color
				y = y + 1.25
				x = 0
				continue // skip this symbol
			case ">": // Tab (Uses > to identify)
				x = x + 25
				continue // skip this symbol
			case "(":
				x = x - 5;
				dc.SetRGB(1, 1, 0)
			case ")":
				x = x - 5;
				dc.SetRGB(1, 1, 0)
			case "{":
				dc.SetRGB(1, 0, 1)
			case "}":
				dc.SetRGB(1, 0, 1)
			case ";":
				x = x - 5;
			case "!":
				x = x - 5;
			case ",":
				x = x - 5;
			case "\"{}\"": // Going to hard-code the fix for this, not ideal
				x = x - 5;
				dc.SetRGB(1, 0, 0)
			case "fn":
				dc.SetRGB(0, 1, 0)
			case "let":
				dc.SetRGB(0, 1, 0)
			case "main":
				dc.SetRGB(0, 0, 1)
			case "\"Hello World\"":
				dc.SetRGB(1, 0, 0)
		}
			
		stringLength, _ := dc.MeasureString(s.TokenText())
		x = x + stringLength
		dc.DrawStringAnchored(s.TokenText(), x, (y*16)+16, 1, 0)
		dc.SetRGB(1, 1, 1)
		x = x + 5;	
	}
}