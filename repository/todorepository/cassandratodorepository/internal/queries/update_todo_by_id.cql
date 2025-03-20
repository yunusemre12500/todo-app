UPDATE %s.%s
SET
    completed = ?
    created_at = ?
    id = ?
    list_id = ?
    text = ?
WHERE
    id = ?
LIMIT 1;