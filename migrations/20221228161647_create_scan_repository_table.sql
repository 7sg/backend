-- +goose Up
CREATE TABLE scan_repository (
                                 id BIGSERIAL PRIMARY KEY,
                                 repository_id BIGSERIAL NOT NULL,
                                 status valid_scan_status NOT NULL DEFAULT 'Queued',
                                 findings jsonb NOT NULL DEFAULT '[]',
                                 enqueued_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 start_time TIMESTAMP,
                                 end_time TIMESTAMP,
                                 CONSTRAINT FK_repository FOREIGN KEY (repository_id) REFERENCES repository (id)
);