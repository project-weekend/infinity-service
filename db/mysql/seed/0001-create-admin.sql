-- ==============================================
-- Admin User Creation Script
-- ==============================================
-- Generated at: /var/folders/wt/m9xc_0fn26sb_vls6jqwspx00000gn/T/go-build1070560626/b001/exe/create_admin
-- Password (plaintext for reference): Admin@123
-- Password (bcrypt hash): $2a$10$KHkPyWLfMyNJ72BszthxfefGuNx88mqWgWy81HSZN3Lpu/ZaDuyq2
-- User UUID: 2bfa645a-43e8-4048-830f-cbb50746f124
-- ==============================================

-- Insert admin role
INSERT INTO `roles` (`name`, `created_at`, `updated_at`)
VALUES ('admin', NOW(), NOW())
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- Insert admin user
-- Note: Replace 'system' in created_by if needed
INSERT INTO `users` (`role_id`, `user_iid`, `name`, `email`, `status`, `created_by`, `password`, `created_at`, `updated_at`)
VALUES (
    (SELECT id FROM roles WHERE name = 'admin'),
    '2bfa645a-43e8-4048-830f-cbb50746f124',
    'Admin User',
    'admin@infinity.local',
    'active',
    'system',
    '$2a$10$KHkPyWLfMyNJ72BszthxfefGuNx88mqWgWy81HSZN3Lpu/ZaDuyq2',
    NOW(),
    NOW()
);

-- ==============================================
-- Login credentials:
-- Email: admin@infinity.local
-- Password: Admin@123
-- ==============================================
