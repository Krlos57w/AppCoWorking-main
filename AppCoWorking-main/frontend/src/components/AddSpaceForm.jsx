import React, { useState } from 'react';

const AddSpaceForm = () => {
  const [id, setId] = useState(''); // Agregado: estado para el id
  const [nombre, setNombre] = useState('');
  const [descripcion, setDescripcion] = useState('');
  const [capacidad, setCapacidad] = useState(0);
  const [precioPorHora, setPrecioPorHora] = useState(0);
  const [estado, setEstado] = useState('');
  const [ubicacion, setUbicacion] = useState('');
  const [error, setError] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccessMessage('');

    const newSpace = {
      id: id ? parseInt(id) : null, // Agregado: incluir id en newSpace (manejar posible string vacío)
      nombre,
      descripcion,
      capacidad,
      precio_por_hora: precioPorHora,
      estado,
      ubicacion,
    };

    try {
        const response = await fetch('http://localhost:8080/api/spaces',{
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newSpace),
      });

      if (response.ok) {
        const data = await response.json();
        console.log('Espacio creado con éxito!', data);
        setSuccessMessage('Espacio creado con éxito!');
        setId(''); // Limpiar el campo id
        setNombre('');
        setDescripcion('');
        setCapacidad(0);
        setPrecioPorHora(0);
        setEstado('');
        setUbicacion('');
      } else {
        const errorData = await response.json();
        console.error('Error al crear el espacio:', errorData);
        setError(errorData.error || 'Error al crear el espacio.');
      }
    } catch (error) {
      console.error('Error de red:', error);
      setError('Error de red. Por favor, intenta de nuevo.');
    }
  };

  return (
    <div>
      <h2>Agregar Nuevo Espacio</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label>ID (Opcional):</label> {/* Etiqueta modificada */}
          <input
            type="number" // Cambiado a number para el ID
            value={id}
            onChange={(e) => setId(e.target.value)}
          />
        </div>
        <div>
          <label>Nombre:</label>
          <input
            type="text"
            value={nombre}
            onChange={(e) => setNombre(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Descripción:</label>
          <textarea
            value={descripcion}
            onChange={(e) => setDescripcion(e.target.value)}
          />
        </div>
        <div>
          <label>Capacidad:</label>
          <input
            type="number"
            value={capacidad}
            onChange={(e) => setCapacidad(parseInt(e.target.value))}
            required
          />
        </div>
        <div>
          <label>Precio por Hora:</label>
          <input
            type="number"
            step="0.01"
            value={precioPorHora}
            onChange={(e) => setPrecioPorHora(parseFloat(e.target.value))}
            required
          />
        </div>
        <div>
          <label>Estado:</label>
          <input
            type="text"
            value={estado}
            onChange={(e) => setEstado(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Ubicación:</label>
          <input
            type="text"
            value={ubicacion}
            onChange={(e) => setUbicacion(e.target.value)}
            required
          />
        </div>
        <button type="submit">Crear Espacio</button>
      </form>
    </div>
  );
};

export default AddSpaceForm;