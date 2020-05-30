package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	sheetName := os.Getenv("GCP_SHEET_NAME")
	if sheetName == "" {
		sheetName = "PushSheet"
	}

	sheetRange, err := NewSheetRange(ctx, os.Getenv("GCP_CREDENTIALS_JSON"), GCPSheetRangeConfig{
		SheetID: os.Getenv("GCP_SHEET_ID"),
		Range:   fmt.Sprintf("%s!A:C", sheetName),
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := HandleSampleData(ctx, sheetRange); err != nil {
		log.Fatal(err)
	}
}

// DataTarget defines generic functionality for modifying a data target.
type DataTarget interface {
	Append(context.Context, [][]interface{}) error
	Clear(context.Context) error
	Get(context.Context) ([][]interface{}, error)
	Set(context.Context, [][]interface{}) error
}

// HandleSampleData creates and appends some sample data to an existing generic data target.
func HandleSampleData(ctx context.Context, t DataTarget) error {
	// Clear the sheet range and create header
	if err := t.Clear(ctx); err != nil {
		return err
	}
	createSample := [][]interface{}{
		{"Physicist", "Area of Research", "Time Added"},
	}
	if err := t.Set(ctx, createSample); err != nil {
		return err
	}

	// Append a sample batch
	now1 := time.Now().Local()
	appendSample1 := [][]interface{}{
		{"Brian Greene", "String Theory", now1},
		{"Max Tegmark", "Mathematical Multiverse", now1},
	}
	if err := t.Append(ctx, appendSample1); err != nil {
		return err
	}

	// Append a second sample batch
	now2 := time.Now().Local()
	appendSample2 := [][]interface{}{
		{"Sean Carroll", "Everettian Many Worlds", now2},
	}
	if err := t.Append(ctx, appendSample2); err != nil {
		return err
	}

	// Read spreadsheet rows and print them out
	values, err := t.Get(ctx)
	if err != nil {
		return err
	}
	if len(values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range values {
			fmt.Printf("%+v\n", row)
		}
	}

	return nil
}
