CREATE TABLE `user` (
                `id` int(11) NOT NULL,
                `username` varchar(255) NOT NULL,
                `password` varchar(255) NOT NULL,
                `sex` int(11) DEFAULT NULL,
                `email` varchar(255) DEFAULT NULL,
                `info` varchar(255) DEFAULT NULL,
                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;