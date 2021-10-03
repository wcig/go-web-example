
-- 创建数据库
CREATE DATABASE `test` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建表
CREATE TABLE `test`.`user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户账号',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户密码',
  `nickname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户昵称',
  `create_at` bigint(20) DEFAULT NULL COMMENT '创建时间',
  `update_at` bigint(20) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表';

-- 添加初始数据
INSERT INTO `test`.`user` (`id`, `username`, `password`, `nickname`, `create_at`, `update_at`) VALUES (null, 'admin', '123456', 'root', '1633239000000', '1633239000000');
