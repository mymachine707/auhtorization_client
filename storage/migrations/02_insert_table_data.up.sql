

BEGIN;

	INSERT INTO category ("id", "category_name", "description") VALUES ('b9401ecc-e7b7-4e83-b387-eb85072adcd9', 'Burger', 'Buluchka va kaklet' ) ON CONFLICT DO NOTHING;
	INSERT INTO category ("id", "category_name", "description") VALUES ('1f27a12d-93c7-4272-9eec-43e28a00482d', 'Lavash', 'goshtli fast food' ) ON CONFLICT DO NOTHING;
	
	INSERT INTO product ("id", "category_id", "product_name", "description", "price") VALUES ( '3d5ee64f-1810-404f-a804-58f12dd18279', '1f27a12d-93c7-4272-9eec-43e28a00482d', 'Achchiq Lavash', 'katta 25 sm achchiq goshtli lavash', '24000' ) ON CONFLICT DO NOTHING;
	INSERT INTO product ("id", "category_id", "product_name", "description", "price") VALUES ( 'e6176ddf-4647-4ede-b1ce-1c065224cf84', '1f27a12d-93c7-4272-9eec-43e28a00482d', 'Tovuqli Lavash', 'Ortacha 20 sm tovuq goshtli lavash', '22000' ) ON CONFLICT DO NOTHING;
	INSERT INTO product ("id", "category_id", "product_name", "description", "price") VALUES ( '611b9c35-5543-4738-b59e-d2bdea847776', '1f27a12d-93c7-4272-9eec-43e28a00482d', 'Kichik Lavash', 'kichik 15 sm goshtli lavash', '20000' ) ON CONFLICT DO NOTHING;

	INSERT INTO product ("id", "category_id", "product_name", "description", "price") VALUES ( '2f1190bb-85f9-4d44-a370-3290a2f23c7c', 'b9401ecc-e7b7-4e83-b387-eb85072adcd9', 'Gamburger', ' Bulchka va 1 ta kakletli', '20000' ) ON CONFLICT DO NOTHING;
	INSERT INTO product ("id", "category_id", "product_name", "description", "price") VALUES ( 'fe1491dd-197d-4853-a707-d8c0362c7259', 'b9401ecc-e7b7-4e83-b387-eb85072adcd9', 'Chizburher', 'Bulchka va 1 ta kakletli va sirli', '22000' ) ON CONFLICT DO NOTHING;
	INSERT INTO product ("id", "category_id", "product_name", "description", "price") VALUES ( '367ca98a-1781-4d0c-96e0-cf173ce24906', 'b9401ecc-e7b7-4e83-b387-eb85072adcd9', 'Dabl Gamburger', 'Bulchka va 2 ta kakletli', '29000' ) ON CONFLICT DO NOTHING;
	INSERT INTO product ("id", "category_id", "product_name", "description", "price") VALUES ( '4587750c-1903-4a26-bb92-c84d61093629', 'b9401ecc-e7b7-4e83-b387-eb85072adcd9', 'Dabl Chiz', 'Bulchka va 2 ta kakletli va sirli', '33000' ) ON CONFLICT DO NOTHING;

COMMIT;
