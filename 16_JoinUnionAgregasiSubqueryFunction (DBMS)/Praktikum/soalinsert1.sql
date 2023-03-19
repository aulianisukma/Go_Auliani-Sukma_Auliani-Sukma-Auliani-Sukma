--a.
INSERT INTO operator VALUES 
 ('1','Kai','2023-03-19','2023-03-19'),  
 ('2','Baekhyun','2023-03-19','2023-03-19'),
('3','sehun','2023-03-19','2023-03-19'),
('4','Suho','2023-03-19','2023-03-19'),
('5','Lay','2023-03-19','2023-03-19');
--b.
INSERT INTO alta_online_shop.`product type`
(id, name, update_at, created_at)
VALUES(0, 't-shirt', '2023-03-18', '2023-03-18'),
VALUES(1, 'jeans', '2023-03-18', '2023-03-18'),
VALUES(3, 'pajamas', '2023-03-18', '2023-03-18');
--c.
INSERT INTO alta_online_shop.product
(id, product_type_id, operator_id, product_deskription, name, code, created_at, update_at, `status small`)
VALUES(0, 0, 0, '1', '3', 'new pajamas brand', '2023-03-18', '2023-03-18', 0),
VALUES(1, 1, 1, '1', '3', 'new jeans brand', '2023-03-18', '2023-03-18', 1);
--d.
INSERT INTO alta_online_shop.product
(id, product_type_id, operator_id, product_deskription, name, code, created_at, update_at, `status small`)
VALUES(4, 4, 4, '2', '1', 'new jeans brand', '2023-03-18', '2023-03-18', 4);
INSERT INTO alta_online_shop.product
(id, product_type_id, operator_id, product_deskription, name, code, created_at, update_at, `status small`)
VALUES(7, 7, 7, '3', '4', 'preloved pajamas brand', '2023-03-18', '2023-03-18', 7);
--f.
INSERT INTO alta_online_shop.`product deskription`
(id, product_id, description, created_at, update_at)
VALUES(3, 3, '3', '2023-03-19', '2023-03-19'),
(4, 4, '4', '2023-03-19', '2023-03-19'),
(5, 5, '5', '2023-03-19', '2023-03-19'),
(7, 7, '6', '2023-03-19', '2023-03-19');
--g.
INSERT INTO alta_online_shop.`payment method`
(id, name, update_at, created_at)
VALUES(1, 'cod', '2023-03-18', '2023-03-18'),
(2, 'shopeepay', '2023-03-18', '2023-03-18');
--h.
INSERT INTO alta_online_shop.`user`
(id, Product, Product_Type, Operators, Product_Description, Payment_Method, `Transaction`, `Transaction detail`)
VALUES(1, 'new pajamas brand', 'pajamas', '2', 'transfer', 1, 1, 1),
(2, 'preloved pajamas brand', 'pajamas', '3', 'cod', 2, 2, 2),
(3, 'preloved jeans brand', 'jeans', '4', 'cod', 3, 3, 3),
(4, 'preloved t-shirt brand', 't-shirt', '5', 'transfer', 4, 4, 4);