import  { useState } from 'react';
import './App.css';

function App() {
  const [data, setData] = useState(null);

  const mostrarDatos = async () => {
    try {
      const response = await fetch('http://localhost:3000/data');
      const jsonData = await response.json();
      setData(jsonData);
    } catch (error) {
      console.error('Error al obtener los datos', error);
    }
  };

  return (
    <div>
      <h1>TAREA1-SO1-1S2024</h1>
      <button onClick={mostrarDatos}>Mostrar datos</button>
      {data && (
        <div>
          <h2>Datos obtenidos:</h2>
          <pre>{JSON.stringify(data, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}

export default App;