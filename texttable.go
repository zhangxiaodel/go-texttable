package texttable

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
	"github.com/mattn/go-runewidth"
)

type cellAlignment int

const (
	ALIGN_LEFT cellAlignment = iota
	ALIGN_RIGHT
)

type tableCell struct {
	content string
	alignment cellAlignment
}

type tableRower interface{}

type tableRow struct {
	cells []tableCell
}

type tableLine struct {}

type TextTable struct {
	header []tableRow
	rows   []tableRower // tableRow or tableLine
	width  int
	border bool
}

func New(headers []string) (*TextTable, error) {
	table := new(TextTable)

	err := table.SetHeader(headers)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func (t *TextTable) SetHeader(headers []string) error {
	if len(headers) == 0 {
		return errors.New("No header elements")
	}
	return nil
}

func divideByNewLine(str string) []string {
	return strings.Split(str, "\n")
func (t *TextTable) borderString() string {
	borderString := "+"
	margin := 2

	for _, width := range(t.maxWidths) {
		for i := 0; i < width + margin; i++ {
			borderString += "-"
		}
		borderString += "+"
	}

	return borderString
}

func adjustColumns(colLines [][]string, max int) [][]string {
	adjustedCols := make([][]string, len(colLines))

	for _, lines := range colLines {
		padding := max - len(lines)
		for i := 0; i < padding; i++ {
			lines = append(lines, "")
		}
		adjustedCols = append(adjustedCols, lines)
	}

	return adjustedCols
}

func (t *TextTable) divideRowsByLine(strs []string) []tableRower {
//	maxHeight := calcMaxHeight(strs)
//	strLines := make([][maxHeight]string, len(strs))
//
//	for i, str := range strs {
//		strLines[i] = divideByNewLine(str)
//	}
//
//	adjusted := adjustColumns(strLines, maxHeight)

	return nil

func decideAlignment(str string) cellAlignment {
	_, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return ALIGN_RIGHT
	}

	_, err = strconv.ParseFloat(str, 64)
	if err == nil {
		return ALIGN_RIGHT
	}

	return ALIGN_LEFT
}

func calcMaxHeight(strs []string) int {
	max := -1

	for _, str := range strs {
		lines := strings.Split(str, "\n")
		height := len(lines)
		if height > max {
			max = height
		}
	}

	return max
}

func stringWidth(str string) int {
	return runewidth.StringWidth(str)
}

func formatCellUnit(cell *cellUnit, maxWidth int) string {
	str := cell.content
	width := stringWidth(cell.content)

	padding := strings.Repeat(" ", maxWidth - width)

	var ret string
	if cell.alignment == ALIGN_RIGHT {
		ret = padding + str
	} else {
		ret = str + padding
	}

	return " " + ret + " "
}

