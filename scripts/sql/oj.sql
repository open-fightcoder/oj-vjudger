CREATE TABLE `problem` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `flag` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1-普通题目 2-用户题目',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '申请状态',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '题目提供者',
  `difficulty` varchar(40) NOT NULL DEFAULT '' COMMENT '题目难度',
  `case_data` varchar(200) NOT NULL DEFAULT '' COMMENT '测试数据',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '题目标题',
  `description` varchar(500) NOT NULL DEFAULT '' COMMENT '题目描述',
  `input_des` varchar(300) NOT NULL DEFAULT '' COMMENT '输入描述',
  `output_des` varchar(300) NOT NULL DEFAULT '' COMMENT '输出描述',
  `input_case` varchar(200) NOT NULL DEFAULT '' COMMENT '测试输入',
  `output_case` varchar(200) NOT NULL DEFAULT '' COMMENT '测试输出',
  `hint` varchar(300) DEFAULT NULL COMMENT '题目提示(可以为对样例输入输出的解释)',
  `time_limit` int(11) NOT NULL DEFAULT '0' COMMENT '时间限制',
  `memory_limit` int(11) NOT NULL DEFAULT '0' COMMENT '内存限制',
  `tag` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '题目标签',
  `is_special_judge` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否特判 1-特判 2-非特判',
  `special_judge_source` varchar(100) DEFAULT NULL COMMENT '特判程序源代码',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '标准程序',
  `language_limit` varchar(100) DEFAULT NULL COMMENT '语言限制',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_difficulty` (`difficulty`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_code` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`problem_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '题目ID',
	`user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
	`save_code` varchar(100) NOT NULL DEFAULT '' COMMENT '保存代码',
	PRIMARY KEY (`id`),
  	UNIQUE KEY `uniq_user` (`user_id`,`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_collection` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `problem_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '题目ID',
    `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_collection` (`user_id`,`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `sys_config` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `sys_key` varchar(100) NOT NULL DEFAULT '' COMMENT '键',
    `sys_value` varchar(100) NOT NULL DEFAULT '' COMMENT '值',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_sys_key` (`sys_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `email` varchar(50) NOT NULL COMMENT '邮箱',
  `password` varchar(80) NOT NULL COMMENT '密码',
  `phone` varchar(20) NOT NULL COMMENT '手机号',
  `qq_id` varchar(40) DEFAULT NULL COMMENT '用于QQ第三方登录',
  `github_id` varchar(40) DEFAULT NULL COMMENT '用于GITHUB第三方登录',
  `weichat_id` varchar(40) DEFAULT NULL COMMENT '用于微信第三方登录',
  PRIMARY KEY (`id`),
  KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;












CREATE TABLE `submit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `problem_id` bigint(20) NOT NULL COMMENT '题目ID',
  `user_id` bigint(20) NOT NULL COMMENT '提交用户ID',
  `language` varchar(20) NOT NULL COMMENT '提交语言',
  `submit_time` bigint(20) NOT NULL COMMENT '提交时间',
  `running_time` int(11) DEFAULT NULL COMMENT '耗时(ms)',
  `running_memory` int(11) DEFAULT NULL COMMENT '所占空间',
  `result` int(11) DEFAULT NULL COMMENT '运行结果',
  `result_des` varchar(500) DEFAULT NULL COMMENT '结果描述',
  `code` varchar(200) NOT NULL COMMENT '提交代码',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `result` (`result`),
  KEY `problem_id` (`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `submit_test` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) NOT NULL COMMENT '提交用户ID',
  `language` varchar(20) NOT NULL COMMENT '提交语言',
  `submit_time` bigint(20) NOT NULL COMMENT '提交时间',
  `running_time` int(11) DEFAULT NULL COMMENT '耗时(ms)',
  `running_memory` int(11) DEFAULT NULL COMMENT '所占空间',
  `result` int(11) DEFAULT NULL COMMENT '运行状态',
  `input` varchar(300) DEFAULT NULL COMMENT '输入',
  `result_des` varchar(300) DEFAULT NULL COMMENT '结果',
  `code` varchar(200) NOT NULL COMMENT '提交代码',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `result` (`result`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;