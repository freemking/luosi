-- 产品分类表初始化数据
-- 运行前请确保 categories 表已创建

INSERT INTO `categories` (`name`, `slug`, `description`, `icon`, `image_url`, `order`, `status`, `created_at`, `updated_at`) VALUES
('Blind Rivet', 'blind-rivet', 'Blind rivets for fastening applications where access is limited to one side of the workpiece.', '', '', 1, 1, NOW(), NOW()),
('Insert Nut', 'insert-nut', 'Threaded inserts for creating strong threads in plastic, wood and other materials.', '', '', 2, 1, NOW(), NOW()),
('Self Clinching Fasteners', 'self-clinching-fasteners', 'Self-clinching fasteners for sheet metal applications.', '', '', 3, 1, NOW(), NOW()),
('Other Products', 'other-products', 'Other fastener products including washers, pins, and custom solutions.', '', '', 4, 1, NOW(), NOW()),
('Tools', 'tools', 'Rivet guns, installation tools and other fastening equipment.', '', '', 5, 1, NOW(), NOW());
