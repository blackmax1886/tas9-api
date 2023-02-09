CREATE TABLE IF NOT EXISTS `task` (
  `id` varchar(64) NOT NULL,
  `name` varchar(64) NOT NULL,
  `content` varchar(256) NOT NULL,
  `done` boolean NOT NULL,
  `due` datetime,
  `assigned_at` datetime,
  `group` varchar(64),
  `type` varchar(16) NOT NULL,
  `priority` varchar(16),
  `archived` boolean NOT NULL,
  `user_id` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

CREATE TABLE IF NOT EXISTS `user` (
  `id` varchar(64) NOT NULL,
  `name` varchar(256) NOT NULL,
  `email` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

CREATE TABLE IF NOT EXISTS `subtask` (
  `id` varchar(64) NOT NULL,
  `name` varchar(64) NOT NULL,
  `parent_task_id` varchar(64) NOT NULL,
  `content` varchar(256) NOT NULL,
  `done` boolean NOT NULL,
  `due` datetime,
  `assigned_at` datetime,
  `priority` varchar(16),
  `archived` boolean NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;
