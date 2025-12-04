-- Insert roles
INSERT INTO roles (name) VALUES ('admin'), ('maker'), ('checker'), ('approver');

-- Insert initial admin user
-- Default credentials: admin@example.com / Admin@123
-- UserID: admin-001
INSERT INTO users (role_id, user_id, name, email, status, password, created_by, created_at, updated_at)
VALUES (
    (SELECT id FROM roles WHERE name = 'admin' LIMIT 1),
    'admin-001',
    'System Administrator',
    'admin@example.com',
    'active',
    '$2a$10$M/yXWhmY8gd/EoRkfzpwn.Epe12mOUwNJ8.UqEuUaC5xpsdcHappW',
    'system',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);