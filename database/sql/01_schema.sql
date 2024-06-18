CREATE TABLE comunidad (
    comunidad_id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    direccion VARCHAR(255),
    tipo VARCHAR(50) NOT NULL CHECK (tipo IN ('Fraccionamiento', 'Edificio', 'Calle'))
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
