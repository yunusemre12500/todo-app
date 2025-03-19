SELECT created_at, display_name, email_address, id, name, password_hash
FROM %s.%s
WHERE
    id = ?
LIMIT 1;