SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dingding
-- ----------------------------
DROP TABLE IF EXISTS `dingding`;
CREATE TABLE `dingding` (
                            `id` int(11) NOT NULL,
                            `username` varchar(255) NOT NULL,
                            `name` varchar(255) DEFAULT NULL,
                            `number` varchar(255) NOT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `id` (`id`),
                            UNIQUE KEY `username` (`username`),
                            UNIQUE KEY `number` (`number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of dingding
-- ----------------------------
BEGIN;
INSERT INTO `dingding` (`id`, `username`, `name`, `number`) VALUES (1, 'san.zhang', '张三', '12345678910');
INSERT INTO `dingding` (`id`, `username`, `name`, `number`) VALUES (2, 'si.li', '李四', '12345678910');

COMMIT;

SET FOREIGN_KEY_CHECKS = 1;