CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    rol_id INT NOT NULL,
    fecha_registro TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (rol_id) REFERENCES roles(id) ON DELETE RESTRICT
);

CREATE TABLE espacios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT,
    capacidad INT NOT NULL,
    precio_por_hora DECIMAL(10,2) NOT NULL,
    estado VARCHAR(20) CHECK (estado IN ('disponible', 'mantenimiento')) NOT NULL DEFAULT 'disponible'
);

CREATE TABLE reservas (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL,
    espacio_id INT NOT NULL,
    fecha_inicio TIMESTAMP NOT NULL,
    fecha_fin TIMESTAMP NOT NULL,
    estado VARCHAR(20) CHECK (estado IN ('pendiente', 'confirmada', 'cancelada', 'completada')) NOT NULL DEFAULT 'pendiente',
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    FOREIGN KEY (espacio_id) REFERENCES espacios(id) ON DELETE CASCADE
);

CREATE TABLE pagos (
    id SERIAL PRIMARY KEY,
    reserva_id INT NOT NULL,
    monto DECIMAL(10,2) NOT NULL,
    metodo_pago VARCHAR(50) CHECK (metodo_pago IN ('tarjeta', 'efectivo', 'transferencia')),
    estado VARCHAR(20) CHECK (estado IN ('pendiente', 'pagado', 'rechazado')) NOT NULL DEFAULT 'pendiente',
    fecha_pago TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (reserva_id) REFERENCES reservas(id) ON DELETE CASCADE
);

CREATE TABLE logs_actividad (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL,
    accion TEXT NOT NULL,
    fecha TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE
);

-- Insertar roles básicos
INSERT INTO roles (nombre) VALUES ('admin'), ('cliente'), ('recepcionista');
