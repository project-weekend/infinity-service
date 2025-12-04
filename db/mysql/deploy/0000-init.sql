-- Roles table
CREATE TABLE IF NOT EXISTS `roles` (
    `id`                BIGINT UNSIGNED NOT NULL AUTO_INCREMENT                         COMMENT 'primary key',
    `name`              VARCHAR(50) NOT NULL DEFAULT ''                                 COMMENT 'unique role name',
    `created_at`        TIMESTAMP DEFAULT CURRENT_TIMESTAMP                             COMMENT 'created time',
    `updated_at`        TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',

    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_role_name` (`name`)
) ENGINE = INNODB 
DEFAULT CHARSET = utf8mb4 
COLLATE = utf8mb4_unicode_ci;


-- Users table
CREATE TABLE IF NOT EXISTS `users` (
    `id`                BIGINT UNSIGNED NOT NULL AUTO_INCREMENT                         COMMENT 'primary key',
    `role_id`           BIGINT UNSIGNED NOT NULL                                        COMMENT 'user role id',
    `user_id`           VARCHAR(50) NOT NULL DEFAULT ''                                 COMMENT 'user internal id',
    `name`              VARCHAR(50) NOT NULL DEFAULT ''                                 COMMENT 'user name',
    `email`             VARCHAR(250) NOT NULL DEFAULT ''                                COMMENT 'unique user email',
    `status`            VARCHAR(250) NOT NULL DEFAULT ''                                COMMENT 'user status',
    `created_by`        VARCHAR(50) DEFAULT NULL DEFAULT ''                             COMMENT 'creator user id from user service',
    `password`          VARCHAR(255) DEFAULT NULL DEFAULT ''                            COMMENT 'user password',
    `created_at`        TIMESTAMP DEFAULT CURRENT_TIMESTAMP                             COMMENT 'created time',
    `updated_at`        TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',

    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_email` (`email`)
) ENGINE = INNODB 
DEFAULT CHARSET = utf8mb4 
COLLATE = utf8mb4_unicode_ci;


CREATE TABLE IF NOT EXISTS `user_sessions` (
    `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `session_id`    VARCHAR(50) NOT NULL,
    `user_id`       VARCHAR(50) NOT NULL,
    `token`         VARCHAR(500) NOT NULL,
    `expires_at`    TIMESTAMP NOT NULL,
    `ip_address`    VARCHAR(50) DEFAULT NULL COMMENT 'ip address',
    `user_agent`    VARCHAR(500) DEFAULT NULL,
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_session_id` (`session_id`),
    UNIQUE KEY `uk_token` (`token`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_expires_at` (`expires_at`)
) ENGINE = INNODB 
DEFAULT CHARSET = utf8mb4 
COLLATE = utf8mb4_unicode_ci;