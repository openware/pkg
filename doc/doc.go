package doc

import (
	"fmt"
	"io"

	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
)

// A package for automatic documentation generation.

type Interface interface {
	Describe() ([]string, [][]string, error)
}

type Document struct {
	writer      io.Writer
	header      string
	description string
	keys        []string
	rows        [][]string
}

// Creates a new empty document with a title.
func NewDocument(w io.Writer, header string) *Document {
	return &Document{
		writer: w,
		header: header,
	}
}

// SetDescription sets a document description.
func (d *Document) SetDescription(description string) {
	d.description = description
}

// SetKeys sets keys in the table header.
func (d *Document) SetKeys(keys ...string) {
	d.keys = keys
}

// SetRows manually sets rows.
func (d *Document) SetRows(rows [][]string) {
	d.rows = rows
}

// Fill takes an object with .Describe() method returning []string keys and [][]string rows to fill the document table.
func (d *Document) Fill(obj Interface) error {
	var err error
	d.keys, d.rows, err = obj.Describe()
	return err
}

// Generate generates a new document from the values it contains.
func (d *Document) Generate() error {
	// Generate a table
	table, err := markdown.NewTableFormatterBuilder().WithAlphabeticalSortIn(markdown.ASCENDING_ORDER).
		WithPrettyPrint().Build(d.keys...).Format(d.rows)
	if err != nil {
		return err
	}
	// Assemble everything into one document
	_, err = io.WriteString(d.writer, fmt.Sprintf("%s\n\n%s\n\n%s\n", d.header, d.description, table))
	return err
}
