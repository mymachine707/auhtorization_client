

BEGIN;
	INSERT INTO client ("id", "firstname", "lastname", "username","phone",  "address", "type", "password") VALUES ('d7dc9b63-50fe-4c15-8020-e77d274afcc6', 'Islombek', 'Oripov', 'my_machine_707','998946506184', 'Xasanboy street', 'sudo', '176') ON CONFLICT DO NOTHING;
	INSERT INTO client ("id", "firstname", "lastname", "username","phone",  "address", "type", "password") VALUES ('4587750c-1903-4a26-bb92-c84d61093629', 'John', 'Doe', 'jDoe432','998998765432', 'Wall street', 'client', '321') ON CONFLICT DO NOTHING;
	INSERT INTO client ("id", "firstname", "lastname", "username","phone",  "address", "type", "password") VALUES ('62ffaa19-ba46-456d-be7d-e4fece7dd2ce', 'Jack', 'Ma', 'jMa001','123456789001', 'Chines street', 'client', '321') ON CONFLICT DO NOTHING;
COMMIT;


