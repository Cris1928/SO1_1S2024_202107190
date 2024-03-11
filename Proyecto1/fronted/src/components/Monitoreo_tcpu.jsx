import React, { useState, useEffect } from 'react';
import 'chart.js/auto';
import { Line } from 'react-chartjs-2';

const Monitoreo_tcpu = () => {
  const [chartData, setChartData] = useState({ labels: [], datasets: [] });

  const fetchData = () => {
    // Llamada a la API para obtener datos
    fetch('http://localhost:5200/api/data/cpu')

      .then((response) => response.json())
      .then((data) => {
        // Procesar los datos y actualizar el estado del gráfico
        const labels = data.map((entry) => entry.fecha);
        const dataset = {
          label: 'Porcentaje',
          data: data.map((entry) => entry.porcentaje),
          fill: false,
          lineTension: 0.5,
          backgroundColor: 'rgba(75,192,192,0.4)',
          borderColor: 'rgba(75,192,192,1)',
          borderWidth: 2,
        };

        setChartData({ labels, datasets: [dataset] });
      })
      .catch((error) => {
        console.error('Error al obtener datos de la API:', error);
      });
  };

  useEffect(() => {
    fetchData();
  }, []); // La dependencia vacía asegura que useEffect solo se ejecute una vez al montar el componente

  const handleRefresh = () => {
    fetchData();
  };

  return (
    <div>
      <h2>Gráfica Lineal Cpu</h2>
      <Line data={chartData} />
      <button onClick={handleRefresh}>Actualizar Gráfica</button>
    </div>
  );
};

export default Monitoreo_tcpu;