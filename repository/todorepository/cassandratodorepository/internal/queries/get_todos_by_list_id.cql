SELECT completed, created_at, id, list_id, text
FROM %s.%s
WHERE
    list_id = ?
LIMIT 1;