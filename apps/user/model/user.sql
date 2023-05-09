CREATE TABLE `user` (
        `id` int(11) NOT NULL AUTO_INCREMENT,
        `username` varchar(20) NOT NULL,
        `password` varchar(40) NOT NULL,
        `sex` varchar(10) NOT NULL DEFAULT '',
        `email` varchar(25) NOT NULL DEFAULT '',
        `info` varchar(255) NOT NULL DEFAULT '',
        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;