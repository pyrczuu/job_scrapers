-- +goose Up
ALTER TABLE job_offers DROP COLUMN salary;

ALTER TABLE job_offers ADD COLUMN salary_employment TEXT;
ALTER TABLE job_offers ADD COLUMN salary_b2b TEXT;
ALTER TABLE job_offers ADD COLUMN salary_contract TEXT;

-- +goose Down
ALTER TABLE job_offers DROP COLUMN salary_employment;
ALTER TABLE job_offers DROP COLUMN salary_b2b;
ALTER TABLE job_offers DROP COLUMN salary_contract;

ALTER TABLE job_offers ADD COLUMN salary TEXT;

