SELECT completed, created_at, id, list_id, text
FROM %s.%s
WHERE
    id = ?
LIMIT 1;