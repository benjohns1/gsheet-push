package main

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// GCPSheetRangeConfig configures a GCP spreadsheet.
type GCPSheetRangeConfig struct {
	SheetID string
	Range   string
}

// GCPSheetRange provides access to a connected GCP spreadsheet range and its configuration.
type GCPSheetRange struct {
	service *sheets.Service
	cfg     GCPSheetRangeConfig
}

// NewSheetRange creates and connects to a GCP spreadsheet service and configures a spreadsheet range to work with.
func NewSheetRange(ctx context.Context, credentials string, cfg GCPSheetRangeConfig) (*GCPSheetRange, error) {
	srv, err := sheets.NewService(ctx, option.WithCredentialsJSON([]byte(credentials)))
	if err != nil {
		return nil, fmt.Errorf("unable to create GCP sheet service: %v", err)
	}
	return &GCPSheetRange{
		service: srv,
		cfg:     cfg,
	}, nil
}

// Append appends rows.
func (s *GCPSheetRange) Append(ctx context.Context, data [][]interface{}) error {
	_, err := s.service.Spreadsheets.Values.Append(s.cfg.SheetID, s.cfg.Range, &sheets.ValueRange{
		Values: data,
	}).ValueInputOption("RAW").Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("error appending to GCP sheet %+v: %v", s.cfg, err)
	}
	return nil
}

// Clear clears the range.
func (s *GCPSheetRange) Clear(ctx context.Context) error {
	_, err := s.service.Spreadsheets.Values.Clear(s.cfg.SheetID, s.cfg.Range, &sheets.ClearValuesRequest{}).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("unable to clear GCP sheet range %+v: %v", s.cfg, err)
	}
	return nil
}

// Get gets the values in the range.
func (s *GCPSheetRange) Get(ctx context.Context) ([][]interface{}, error) {
	resp, err := s.service.Spreadsheets.Values.Get(s.cfg.SheetID, s.cfg.Range).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to get data from GCP sheet %+v: %v", s.cfg, err)
	}
	return resp.Values, nil
}

// Set replaces the range with the given data.
func (s *GCPSheetRange) Set(ctx context.Context, data [][]interface{}) error {
	_, err := s.service.Spreadsheets.Values.Update(s.cfg.SheetID, s.cfg.Range, &sheets.ValueRange{
		Values: data,
	}).ValueInputOption("RAW").Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("error updating GCP sheet %+v: %v", s.cfg, err)
	}
	return nil
}
