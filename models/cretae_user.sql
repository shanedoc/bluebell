create tables user(
    `id` bigint(20) not null auto_increment,
    `user_id` bigint(20) not null,
    `username` varchar(64) collate uft8mb4_general_ci not null,
    `password` varchar(64) collate uft8mb4_general_ci not null,
    `email` varchar(64) collate uft8mb4_general_ci,
    `gender` tinyint(4) not null default 0,
    `create_time` timestamp null default current_timestamp,
    `update_time` timestamp null default current_timestamp on update current_timestamp,
    primary key(`id`),
    unique key `idx_username` (`username`) using btree,
    unique key `idx_user_id` (`user_id`) using btree

)engine=Innodb default charset=utf8mb4 collate=uft8mb4_general_ci;