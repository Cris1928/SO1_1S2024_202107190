import React, { useState, useEffect } from 'react';
import 'chart.js/auto';
import { Line } from 'react-chartjs-2';
import Monitoreo_tcpu from './Monitoreo_tcpu.jsx'

const Monitoreo_t = () => {
  const [chartData, setChartData] = useState({ labels: [], datasets: [] });

  const fetchData = () => {
    // Llamada a la API para obtener datos
    fetch('http://localhost:5200/api/data/ram')

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
    <div className='Monitoreo_t' >
    <div>
      <h2>Gráfica Lineal Ram</h2>
      <Line data={chartData} />
      <button onClick={handleRefresh}>Actualizar Gráfica</button>
    </div>
    <div>
  
   <Monitoreo_tcpu/>
  
</div>
</div>
  );
};

export default Monitoreo_t;
/* 



import React, { useState, useEffect } from 'react';
import 'chart.js/auto';
import { Line } from  'react-chartjs-2';




const Monitoreo_t = () => {
    const [chartData, setChartData] = useState({});
  
  // Opciones de la gráfica
  const options = {
    scales: {
      y: {
        beginAtZero: true,
      },
    },
  };
    useEffect(() => {
      // Realizar la solicitud al servidor backend
      fetch('http://localhost:5200/api/data')
        .then((response) => response.json())
        .then((data) => {
          // Procesar los datos para la gráfica
          const labels = data.map((entry) => entry.fecha);
          const porcentajes = data.map((entry) => entry.porcentaje);
          console.log(labels)
          console.log(porcentajes)
          // Configurar los datos para la gráfica
          const chartData = {
            labels,
            datasets: [
              {
                label: 'Porcentaje de RAM',
                fill: false,
                lineTension: 0.5,
                backgroundColor: 'rgba(75,192,192,0.4)',
                borderColor: 'rgba(75,192,192,1)',
                borderWidth: 2,
                data: porcentajes,
              },
            ],
          };
  

          setChartData(chartData);
          
        })
        .catch((error) => console.error('Error al obtener datos del servidor: ' + error));
    }, []);
  
    return (
      <div>
        <h2>Gráfica Lineal</h2>
        <Line data={chartData}  options={options}/>
      </div>
    );
  };
  
  export default Monitoreo_t;
*/