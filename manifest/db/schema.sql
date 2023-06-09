CREATE TABLE `files` (
    `uid` INTEGER PRIMARY KEY,
    `fileName` TEXT DEFAULT 'default_filename',
    `switch` BOOLEAN DEFAULT 0
);