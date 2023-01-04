-- +goose Up
ALTER TABLE scan_repository
    ADD COLUMN total_files INT DEFAULT 0,
		ADD COLUMN scanned_files INT DEFAULT 0;