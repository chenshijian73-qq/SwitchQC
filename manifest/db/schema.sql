CREATE TABLE `files` (
    `id` INTEGER PRIMARY KEY,
    `filename` TEXT DEFAULT NULL, -- ''
    `status` BOOLEAN DEFAULT 0, -- '文件是否使用'
    `create_at` datetime(0) DEFAULT NULL, --  'Created Time'
    `update_at` datetime(0) DEFAULT NULL, --  'Updated Time'
    `delete_at` datetime(0) DEFAULT NULL --  'Updated Time'
);