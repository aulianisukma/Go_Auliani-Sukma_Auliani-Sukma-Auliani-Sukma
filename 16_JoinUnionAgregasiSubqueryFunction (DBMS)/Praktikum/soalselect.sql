--a. 
SELECT name FROM users WHERE gender = 'M'
--b.
SELECT * FROM product WHERE id = '3'
--c.
SELECT * FROM user WHERE created_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND name LIKE '%a%';
--d.
SELECT COUNT(*) FROM user WHERE gender = 'F';
--e.
SELECT * FROM product LIMIT 5;
--f.
SELECT * FROM user ORDER BY name ASC;