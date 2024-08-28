/* 
Aquí comandos de MySQL para probar que todo esta funcionando ok 
LOCAL INSTANCE MYSQL80, Usuario: admin / password: admin (este ultimo yo lo colequé al instalar MySQL)
*/

select * from products
select * from clients
select * from invoice_items

SET GLOBAL time_zone = 'America/Peru'

use mlbauto; 

CREATE TABLE IF NOT EXISTS products
	(
	id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	name VARCHAR (25) NOT NULL,
	observations VARCHAR (100),
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
 	update_at TIMESTAMP
	);
    
INSERT INTO products (name, price) VALUES ('Marlon', '25')

SELECT id, name, observations, price, created_at, updated_at FROM products WHERE id = 2
SELECT id, name, observations, price, created_at, updated_at FROM products WHERE id like 3

DROP TABLE `mlbauto`.`products`;
DROP TABLE `mlbauto`.`clients`;
DROP TABLE `mlbauto`.`invoice_items`;