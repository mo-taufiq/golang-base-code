INSERT INTO users 
    ("name", email, "password", role_id, created_at, updated_at, deleted_at) 
VALUES 
    ('admin', 'admin@email.com', '$2a$10$EGz2MCj6SrPZkL313idnquIJ.M9ul.FaQfTQ5excyUZ5elosZFTKO', '1', CURRENT_TIMESTAMP, NULL, NULL);

-- password: abc123