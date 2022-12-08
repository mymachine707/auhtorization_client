

BEGIN;
	INSERT INTO client ("id", "firstname", "lastname", "phone",  "address") VALUES ('4587750c-1903-4a26-bb92-c84d61093629', 'John', 'Doe', '998998765432', 'Wall street' ) ON CONFLICT DO NOTHING;
	INSERT INTO client ("id", "firstname", "lastname", "phone",  "address") VALUES ('62ffaa19-ba46-456d-be7d-e4fece7dd2ce', 'Jack', 'Ma', '123456789001', 'Chines street') ON CONFLICT DO NOTHING;
COMMIT;


