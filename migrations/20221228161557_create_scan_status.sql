-- +goose Up
CREATE TYPE valid_scan_status AS ENUM ( 'Queued',
	'InProgress',
	'Success',
	'Failure'
);