-- name: CreateJobOffer :one
INSERT INTO job_offers (
    id, title, company, location, description, url, source, published_at, skills,
    salary_employment, salary_b2b, salary_contract
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateJobOffer :one
UPDATE job_offers
SET 
    title = ?,
    company = ?,
    location = ?,
    description = ?,
    published_at = ?,
    skills = ?,
    salary_employment = ?,
    salary_b2b = ?,
    salary_contract = ?,
    last_seen_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: UpsertJobOffer :one
INSERT INTO job_offers (
    id, title, company, location, description, url, source, published_at, skills,
    salary_employment, salary_b2b, salary_contract
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(url) DO UPDATE SET
    title = excluded.title,
    company = excluded.company,
    location = excluded.location,
    description = excluded.description,
    published_at = excluded.published_at,
    skills = excluded.skills,
    salary_employment = excluded.salary_employment,
    salary_b2b = excluded.salary_b2b,
    salary_contract = excluded.salary_contract,
    last_seen_at = CURRENT_TIMESTAMP
RETURNING *;


-- name: DeleteJobOffer :exec
DELETE FROM job_offers WHERE id = ?;

-- name: ListJobOffers :many
SELECT * FROM job_offers 
ORDER BY created_at DESC 
LIMIT ? OFFSET ?;

-- name: ListRecentJobOffers :many
SELECT * FROM job_offers 
ORDER BY published_at DESC 
LIMIT ?;

-- name: ListJobOffersBySource :many
SELECT * FROM job_offers 
WHERE source = ?
ORDER BY published_at DESC 
LIMIT ? OFFSET ?;

-- name: ListJobOffersByCompany :many
SELECT * FROM job_offers 
WHERE company = ?
ORDER BY published_at DESC;

-- name: ListJobOffersByLocation :many
SELECT * FROM job_offers 
WHERE location LIKE ?
ORDER BY published_at DESC;
