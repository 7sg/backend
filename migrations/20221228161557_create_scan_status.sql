-- +goose Up
CREATE TYPE valid_scan_status AS ENUM ( 'Queued',
	'In Progress',
	'Success',
	'Failure'
);