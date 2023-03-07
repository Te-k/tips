package main

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/russross/blackfriday"
)

type console struct {
	flags int
}

// Function to create the renderer
func consoleRenderer() blackfriday.Renderer {
	return &console{
		flags: 0,
	}
}

func (*console) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	out.Write(link)
}

func (*console) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	yellow := color.New(color.Bold).Add(color.FgGreen).SprintFunc()
	out.WriteString(yellow(string(text)))
	out.WriteString("\n")
}

func (*console) BlockHtml(out *bytes.Buffer, text []byte) {
	yellow := color.New(color.FgCyan).Add(color.Underline).SprintFunc()
	out.WriteString(yellow(string(text)))
}

func (*console) BlockQuote(out *bytes.Buffer, text []byte) {
	yellow := color.New(color.FgCyan).Add(color.Underline).SprintFunc()
	out.WriteString(yellow(text))
}

func (*console) CodeSpan(out *bytes.Buffer, text []byte) {
	yellow := color.New(color.FgCyan).Add(color.Underline).SprintFunc()
	out.WriteString(yellow(string(text)))
}

func (*console) DocumentFooter(out *bytes.Buffer) {
}

func (*console) DocumentHeader(out *bytes.Buffer) {
}

func (*console) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	bold := color.New(color.FgCyan).SprintFunc()
	out.WriteString(bold(string(text)))
}

func (*console) Emphasis(out *bytes.Buffer, text []byte) {
	bold := color.New(color.FgCyan).SprintFunc()
	out.WriteString(bold(string(text)))
}

func (*console) TripleEmphasis(out *bytes.Buffer, text []byte) {
	bold := color.New(color.FgCyan).SprintFunc()
	out.WriteString(bold(string(text)))
}

func (*console) Entity(out *bytes.Buffer, entity []byte) {
	out.Write(entity)
}

func (*console) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	// unimplemented
}

func (*console) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	// unimplemented
}

func (*console) Footnotes(out *bytes.Buffer, text func() bool) {
	// unimplemented
}

func (v *console) GetFlags() int {
	return v.flags
}

func (*console) HRule(out *bytes.Buffer) {
	// unimplemented
}

func (*console) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	// This is ugly but I found no other way
	out.WriteString(fmt.Sprintf("%s[%sm### ", "\x1b", "34;1"))
	text()
	out.WriteString(fmt.Sprintf("%s[%sm", "\x1b", "0"))
	out.WriteString("\n\n")
}

func (*console) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	// cannot view images in console
}

func (*console) LineBreak(out *bytes.Buffer) {
	out.WriteString("\n")
}

func (*console) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	out.WriteString(fmt.Sprintf("%s (%s)", content, link))
}

func (*console) List(out *bytes.Buffer, text func() bool, flags int) {
	text()
}

func (*console) ListItem(out *bytes.Buffer, text []byte, flags int) {
	out.WriteString(fmt.Sprintf("-%s\n", text))
}

func (*console) NormalText(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (*console) Paragraph(out *bytes.Buffer, text func() bool) {
	text()
	out.WriteString("\n")
}

func (*console) RawHtmlTag(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (*console) StrikeThrough(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (*console) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
}

func (*console) TableCell(out *bytes.Buffer, text []byte, align int) {
}

func (*console) TableHeaderCell(out *bytes.Buffer, text []byte, align int) {
}

func (*console) TableRow(out *bytes.Buffer, text []byte) {
}

func (*console) TitleBlock(out *bytes.Buffer, text []byte) {
	out.Write(text)
}
