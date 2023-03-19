--1.
SELECT * FROM transactions WHERE user_id = 1 OR user_id = 2;
--2. 
SELECT SUM(harga) AS total_harga FROM transaksi WHERE user_id = 1;
--3.
SELECT COUNT(*) AS total_transaksi FROM 'shopeepay' WHERE 'pajamas' = 2;
--4. 
SELECT product.*, product_type.name AS product_type_name FROM product INNER JOIN product_type ON product.product_type_id = product_type.id;
--5.
SELECT transaction.*, product.name AS product_name, user.name AS user_name
FROM transaction
INNER JOIN product ON transaction.product_id = product.id
INNER JOIN user ON transaction.user_id = user.id;
--6.
CREATE FUNCTION delete_transaction(transaction_id INT) RETURNS VOID AS $$
BEGIN
  DELETE FROM transaction_detail WHERE transaction_id = transaction_id;
  DELETE FROM transaction WHERE id = transaction_id;
END;
$$ LANGUAGE plpgsql;
--7.
CREATE FUNCTION delete_transaction(transaction_id INT) RETURNS VOID AS $$
BEGIN
  DELETE FROM transaction_detail WHERE transaction_id = transaction_id;
  DELETE FROM transaction WHERE id = transaction_id;
END;
$$ LANGUAGE plpgsql;
--8.
SELECT *
FROM products
WHERE id NOT IN (
  SELECT product_id
  FROM transaction_details
);

