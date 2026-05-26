CREATE TABLE IF NOT EXISTS `transactions` (
    `id`              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`         BIGINT UNSIGNED NOT NULL,
    `transaction_no`  VARCHAR(100)    NOT NULL,
    `amount`          DECIMAL(18,2)   NOT NULL DEFAULT 0,
    `currency`        VARCHAR(10)     NOT NULL DEFAULT 'VND',
    `status`          VARCHAR(50)     NOT NULL DEFAULT 'pending',
    `description`     TEXT            NULL,
    
    `created_at`      DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`      DATETIME        NULL,

    PRIMARY KEY (`id`),

    UNIQUE KEY `idx_transactions_transaction_no` (`transaction_no`),

    KEY `idx_transactions_user_id` (`user_id`),
    KEY `idx_transactions_status` (`status`),
    KEY `idx_transactions_created_at` (`created_at`),
    KEY `idx_transactions_deleted_at` (`deleted_at`),

    CONSTRAINT `fk_transactions_user_id`
        FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`)
        ON DELETE CASCADE
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;