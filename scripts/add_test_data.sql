INSERT INTO usuario (nombre, apellido, telefono, email, password_hash)
VALUES ('Luis', 'Patiño', '555-0001', 'luis.patino@example.com', 'hashed_password_here');

INSERT INTO comunidad (nombre, direccion_calle, direccion_numero, direccion_cp, tipo, modelo_suscripcion)
VALUES ('Calzada del Bosque', 'Nuevo Milenio', '4201', '31124', 'Fraccionamiento', 'Mensual');

INSERT INTO suscripcion (usuario_id, comunidad_id, modelo_suscripcion, fecha_inicio, fecha_fin, monto)
VALUES (
   (SELECT usuario_id FROM usuario WHERE email = 'luis.patino@example.com'),
   (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque'),
   'Mensual', '2024-07-01', '2024-08-01', 300.00
);

INSERT INTO usuario (nombre, apellido, telefono, email, password_hash)
VALUES ('Leonardo', 'Gutiérrez', '555-0002', 'leonardo.gutierrez@example.com', 'hashed_password_here');

INSERT INTO usuario (nombre, apellido, telefono, email, password_hash)
VALUES ('Perla', 'Cañas', '555-0003', 'perla.canas@example.com', 'hashed_password_here');

INSERT INTO casa (comunidad_id, direccion, numero)
VALUES (
   (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque'),
   'Nuevo Milenio', 4414
);

INSERT INTO habitante (nombre, apellido, casa_id, telefono, email)
VALUES ('Leonardo', 'Gutiérrez',
(SELECT casa_id FROM casa WHERE numero = 4414 AND comunidad_id = (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque')),
'555-0002', 'leonardo.gutierrez@example.com');

INSERT INTO habitante (nombre, apellido, casa_id, telefono, email)
VALUES ('Perla', 'Cañas',
(SELECT casa_id FROM casa WHERE numero = 4414 AND comunidad_id = (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque')),
'555-0003', 'perla.canas@example.com');


INSERT INTO usuario (nombre, apellido, telefono, email, password_hash)
VALUES ('Gabriel', 'Varela', '555-0004', 'gabriel.varela@example.com', 'hashed_password_here');

INSERT INTO casa (comunidad_id, direccion, numero)
VALUES (
   (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque'),
   'Nuevo Milenio', 4416
);

INSERT INTO habitante (nombre, apellido, casa_id, telefono, email)
VALUES ('Luis', 'Patiño',
(SELECT casa_id FROM casa WHERE numero = 4416 AND comunidad_id = (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque')),
'555-0001', 'luis.patino@example.com');

INSERT INTO habitante (nombre, apellido, casa_id, telefono, email)
VALUES ('Gabriel', 'Varela',
    (SELECT casa_id FROM casa WHERE numero = 4416 AND comunidad_id = (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque')),
'555-0004', 'gabriel.varela@example.com');

INSERT INTO anuncio (comunidad_id, fecha, descripcion)
VALUES (
   (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque'),
   '2024-09-23',
   'Voy a tener una fiesta en Septiembre 28 del 2024, vendrán unos amigos, la fiesta terminará a las 11:00 PM, ojalá no cause molestias.'
);

INSERT INTO comite (comunidad_id, nombre)
VALUES (
   (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque'),
   'Comité de Calzada del Bosque'
);

INSERT INTO comite_miembro (comite_id, habitante_id)
VALUES (
   (SELECT comite_id FROM comite WHERE nombre = 'Comité de Calzada del Bosque'),
   (SELECT habitante_id FROM habitante WHERE email = 'luis.patino@example.com')
);

INSERT INTO comite_miembro (comite_id, habitante_id)
VALUES (
   (SELECT comite_id FROM comite WHERE nombre = 'Comité de Calzada del Bosque'),
   (SELECT habitante_id FROM habitante WHERE email = 'leonardo.gutierrez@example.com')
);

INSERT INTO bazar (casa_id, fecha, descripcion, precio, vendido, acepta_regateo)
VALUES (
   (SELECT casa_id FROM casa WHERE numero = 4416 AND comunidad_id = (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque')),
   '2024-10-22',
   'Mesa con 4 sillas',
   2300.00,
   FALSE,
   FALSE
);

INSERT INTO bazar (casa_id, fecha, descripcion, precio, vendido, acepta_regateo)
VALUES (
   (SELECT casa_id FROM casa WHERE numero = 4416 AND comunidad_id = (SELECT comunidad_id FROM comunidad WHERE nombre = 'Calzada del Bosque')),
   '2024-10-22',
   'Sillón reclinable',
   1500.00,
   FALSE,
   FALSE
);

INSERT INTO anuncio_comite (comite_id, fecha, descripcion)
VALUES (
   (SELECT comite_id FROM comite WHERE nombre = 'Comité de Calzada del Bosque'),
   '2024-07-01',
   'Se convoca a una reunión general de vecinos el día 5 de julio de 2024 a las 18:00 en el parque central.'
);

