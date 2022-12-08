ALTER TABLE product DROP CONSTRAINT IF EXISTS fk_product_client;

BEGIN;

	DELETE FROM product where id='3d5ee64f-1810-404f-a804-58f12dd18279';
	DELETE FROM product where id='e6176ddf-4647-4ede-b1ce-1c065224cf84';
	DELETE FROM product where id='611b9c35-5543-4738-b59e-d2bdea847776';

	DELETE FROM product where id='2f1190bb-85f9-4d44-a370-3290a2f23c7c';
	DELETE FROM product where id='fe1491dd-197d-4853-a707-d8c0362c7259';
	DELETE FROM product where id='367ca98a-1781-4d0c-96e0-cf173ce24906';
	DELETE FROM product where id='4587750c-1903-4a26-bb92-c84d61093629';

	DELETE FROM client where id='b9401ecc-e7b7-4e83-b387-eb85072adcd9';
	DELETE FROM client where id='1f27a12d-93c7-4272-9eec-43e28a00482d';
	

COMMIT;