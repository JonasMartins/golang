-- CREATE DATABASE IF NOT EXISTS `project`;

CREATE TABLE `project`.`users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `version` smallint DEFAULT 1,
  `uuid` varchar(36) NOT NULL,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `deleted_at` timestamp,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp NOT NULL DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE `project`.`invoices` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `version` smallint DEFAULT 1,
  `uuid` varchar(36) NOT NULL,
  `value` double NOT NULL, 
  `client_id` int NOT NULL,
  `due_date` timestamp NOT NULL,
  `deleted_at` timestamp,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp NOT NULL DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE `project`.`versions` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `version` smallint,
  `context` int,
  `deleted_at` timestamp,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp NOT NULL DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE `project`.`entities` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `table_name` varchar(255),
  `deleted_at` timestamp,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp NOT NULL DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE `project`.`event_types` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `context` int,
  `description` varchar(255),
  `deleted_at` timestamp,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp NOT NULL DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE `project`.`users_logs` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `author_id` int,
  `version_id` int,
  `content` text,
  `event_id` int,
  `deleted_at` timestamp,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp NOT NULL DEFAULT NOW() ON UPDATE NOW()
);

CREATE TABLE `project`.`invoices_logs` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `author_id` int,
  `version_id` int,
  `content` text,
  `event_id` int,
  `deleted_at` timestamp,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp NOT NULL DEFAULT NOW() ON UPDATE NOW()
);

CREATE INDEX `context_description` ON `project`.`event_types` (`context`, `description`);

ALTER TABLE `project`.`invoices` ADD FOREIGN KEY (`client_id`) REFERENCES `project`.`users` (`id`);

ALTER TABLE `project`.`event_types` ADD FOREIGN KEY (`context`) REFERENCES `project`.`entities` (`id`);

ALTER TABLE `project`.`users_logs` ADD FOREIGN KEY (`author_id`) REFERENCES `project`.`users` (`id`);

ALTER TABLE `project`.`users_logs` ADD FOREIGN KEY (`version_id`) REFERENCES `project`.`versions` (`id`);

ALTER TABLE `project`.`users_logs` ADD FOREIGN KEY (`event_id`) REFERENCES `project`.`event_types` (`id`);

ALTER TABLE `project`.`invoices_logs` ADD FOREIGN KEY (`author_id`) REFERENCES `project`.`invoices` (`id`);

ALTER TABLE `project`.`invoices_logs` ADD FOREIGN KEY (`version_id`) REFERENCES `project`.`versions` (`id`);

ALTER TABLE `project`.`invoices_logs` ADD FOREIGN KEY (`event_id`) REFERENCES `project`.`event_types` (`id`);