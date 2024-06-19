CREATE TABLE comunidad (
    comunidad_id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    direccion_calle VARCHAR(100),
    direccion_numero VARCHAR(20),
    direccion_colonia VARCHAR(100),
    direccion_cp VARCHAR(10),
    direccion_ciudad VARCHAR(100),
    direccion_estado VARCHAR(100),
    direccion_pais VARCHAR(100),
    tipo VARCHAR(50) NOT NULL CHECK (tipo IN ('Fraccionamiento', 'Edificio', 'Calle')),
    modelo_suscripcion VARCHAR(20) NOT NULL CHECK (modelo_suscripcion IN ('Mensual', 'Bimestral', 'Anual'))
);

CREATE TABLE registro (
    registro_id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    telefono VARCHAR(15),
    correo VARCHAR(100),
    comunidad_id INT REFERENCES comunidad(comunidad_id)
);

CREATE TABLE casa (
    casa_id SERIAL PRIMARY KEY,
    comunidad_id INT REFERENCES comunidad(comunidad_id),
    direccion VARCHAR(255) NOT NULL,
    numero INT NOT NULL
);

CREATE TABLE departamento (
    departamento_id SERIAL PRIMARY KEY,
    comunidad_id INT REFERENCES comunidad(comunidad_id),
    direccion VARCHAR(255) NOT NULL,
    numero INT NOT NULL
);

CREATE TABLE habitante (
    habitante_id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    casa_id INT REFERENCES casa(casa_id) NULL,
    departamento_id INT REFERENCES departamento(departamento_id) NULL,
    telefono VARCHAR(15),
    email VARCHAR(100)
);

CREATE TABLE comite (
    comite_id SERIAL PRIMARY KEY,
    comunidad_id INT REFERENCES comunidad(comunidad_id),
    nombre VARCHAR(100) NOT NULL
);

CREATE TABLE comite_miembro (
    comite_miembro_id SERIAL PRIMARY KEY,
    comite_id INT REFERENCES comite(comite_id),
    habitante_id INT REFERENCES habitante(habitante_id)
);

CREATE TABLE junta (
    junta_id SERIAL PRIMARY KEY,
    comunidad_id INT REFERENCES comunidad(comunidad_id),
    fecha DATE NOT NULL,
    descripcion TEXT
);

CREATE TABLE anuncio (
    anuncio_id SERIAL PRIMARY KEY,
    comunidad_id INT REFERENCES comunidad(comunidad_id),
    fecha DATE NOT NULL,
    descripcion TEXT
);

CREATE TABLE cuota (
    cuota_id SERIAL PRIMARY KEY,
    casa_id INT REFERENCES casa(casa_id) NULL,
    departamento_id INT REFERENCES departamento(departamento_id) NULL,
    monto DECIMAL(10, 2) NOT NULL,
    fecha_pago DATE NOT NULL,
    descripcion TEXT
);

CREATE TABLE bazar (
    bazar_id SERIAL PRIMARY KEY,
    casa_id INT REFERENCES casa(casa_id) NULL,
    departamento_id INT REFERENCES departamento(departamento_id) NULL,
    fecha DATE NOT NULL,
    descripcion TEXT,
    precio DECIMAL(10, 2) NOT NULL,
    vendido BOOLEAN DEFAULT FALSE,
    acepta_regateo BOOLEAN DEFAULT FALSE
);

CREATE TABLE anuncio_casa (
    anuncio_casa_id SERIAL PRIMARY KEY,
    casa_id INT REFERENCES casa(casa_id) NULL,
    departamento_id INT REFERENCES departamento(departamento_id) NULL,
    fecha DATE NOT NULL,
    descripcion TEXT
);




-- Insertar un nuevo fraccionamiento
INSERT INTO comunidad (nombre, direccion_calle, direccion_numero, direccion_colonia, direccion_cp, direccion_ciudad, direccion_estado, direccion_pais, tipo, modelo_suscripcion)
VALUES ('Calzada del Bosque', 'Av. Siempre Viva', '123', 'Los Pinos', '12345', 'Ciudad de México', 'CDMX', 'México', 'Fraccionamiento', 'Anual');

-- Obtener el ID del fraccionamiento recién insertado
SELECT currval(pg_get_serial_sequence('comunidad','comunidad_id'));

-- Suponiendo que el ID del fraccionamiento es 1
-- Insertar información de quien registra el fraccionamiento
INSERT INTO registro (nombre, telefono, correo, comunidad_id)
VALUES ('Edgar Gutiérrez', '555-1234', 'edgar@example.com', 1);

-- Insertar casas y habitantes
INSERT INTO casa (comunidad_id, direccion, numero) VALUES (1, 'Av. Siempre Viva 123', 1);
INSERT INTO habitante (nombre, apellido, casa_id, telefono, email) VALUES ('Edgar', 'Gutiérrez', 1, '555-1234', 'edgar@example.com');
INSERT INTO habitante (nombre, apellido, casa_id, telefono, email) VALUES ('Maria', 'Lopez', 1, '555-5678', 'maria@example.com');

-- Insertar comité y miembros del comité
INSERT INTO comite (comunidad_id, nombre) VALUES (1, 'Comité de Seguridad');
-- Obtener el ID del comité recién insertado
SELECT currval(pg_get_serial_sequence('comite','comite_id'));
-- Suponiendo que el ID del comité es 1
INSERT INTO comite_miembro (comite_id, habitante_id) VALUES (1, 1); -- Edgar Gutiérrez es parte del comité
