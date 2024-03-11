import React, { useState, useEffect } from 'react';
import 'chart.js/auto';
import { Pie } from  'react-chartjs-2';


const Monitoreo_r = () => {
    const [datosRam, setDatosRam] = useState(null);
    const [porcentaje, setPorcentaje] = useState(null);

    const [datosCPU, setDatosCPU] = useState(null);
    const [porcentajecpu, setPorcentajecpu] = useState(null);
    
    const obtenerDatos = async () => {
      try {
        const respuesta = await fetch('http://localhost:5200/ram');
        const respuestacpu = await fetch('http://localhost:5200/cpu');
  
        if (!respuesta.ok) {
          throw new Error('Error al obtener los datos de RAM');
        }
        if (!respuestacpu.ok) {
          throw new Error('Error al obtener los datos de CPU');
        }
  
        const datos = await respuesta.json();
        setDatosRam(datos);

        const datoscpu = await respuestacpu.json();
        setDatosCPU(datoscpu);



        const porcentajeNumerico = parseFloat(datos.porcentaje);
        //setPorcentaje(datos.porcentaje);
        setPorcentaje(isNaN(porcentajeNumerico)? null: porcentajeNumerico);


        const porcentajeNumericocpu = parseFloat(datoscpu.cpu_porcentaje);
        //setPorcentaje(datos.porcentaje);
        setPorcentajecpu(isNaN(porcentajeNumericocpu)? null: porcentajeNumericocpu);

    
    
      } catch (error) {
        console.error(error.message);
      }
    };
  

 /*

    const obtenerDatoscpu = async () => {
      try {
        const respuesta = await fetch('http://localhost:5200/cpu');
  
        if (!respuesta.ok) {
          throw new Error('Error al obtener los datos de cpu');
        }
  
        const datos = await respuesta.json();
        setDatosCPU(datos);

        const porcentajeNumerico = parseFloat(datos.cpu_porcentaje);
        //setPorcentaje(datos.porcentaje);
        setPorcentajecpu(isNaN(porcentajeNumerico)? null: porcentajeNumerico);

    
      } catch (error) {
        console.error(error.message);
      }
    };


    const obtenerDatosRam = async () => {
      try {
        const respuesta = await fetch('http://localhost:5200/ram');
  
        if (!respuesta.ok) {
          throw new Error('Error al obtener los datos de RAM');
        }
  
        const datos = await respuesta.json();
        setDatosRam(datos);

        const porcentajeNumerico = parseFloat(datos.porcentaje);
        //setPorcentaje(datos.porcentaje);
        setPorcentaje(isNaN(porcentajeNumerico)? null: porcentajeNumerico);

    
      } catch (error) {
        console.error(error.message);
      }
    };



   useEffect(() => {
      // Obtener datos al montar el componente
      obtenerDatosRam();
  obtenerDatoscpu();
      // Establecer un intervalo para obtener datos cada segundo
      const intervalo = setInterval(() => {
        obtenerDatosRam();
        obtenerDatoscpu();
      }, 2000);
  
      // Limpiar el intervalo al desmontar el componente para evitar fugas de memoria
      return () => clearInterval(intervalo);
    }, []);

  */
    useEffect(() => {
      // Obtener datos al montar el componente
      obtenerDatos();
      // Establecer un intervalo para obtener datos cada segundo
      const intervalo = setInterval(() => {
        obtenerDatos();
      }, 2000);
  
      // Limpiar el intervalo al desmontar el componente para evitar fugas de memoria
      return () => clearInterval(intervalo);
    }, []);


  

  //datos para la grafica ram
  const porcentajeUso=porcentaje !== null? porcentaje:0;
  const data = { 
    labels: ['total_ram','memoria_en_uso'],
    datasets:  [{
        
        data: [porcentajeUso,100-porcentajeUso],
        backgroundColor: ['green','blue']
    }]

};
//configuracion para la grafica ram
const opciones= {
    maintainAspectRatio: false,
    responsive: true,
    
};
const contstye={
    width: 2222,
    heigth:2222,
};


  //datos para la grafica cpu
  const porcentajeUsocpu=porcentajecpu !== null? porcentajecpu:0;
  const datacpu = { 
    labels: ['total_cpu','memoria_en_uso'],
    datasets:  [{
        label: "CPU",
        data: [100-porcentajeUsocpu,porcentajeUsocpu],
        backgroundColor: ['blue','green']
    }]

};
//configuracion para la grafica cpu
const opcionescpu= {
    maintainAspectRatio: false,
    responsive: true,
    
};
const contstyecpu={
    width: 2222,
    heigth:2222,
};







/*
const opciones= {
  //  responsive: true
  plugins: {
legend: {
    display: true,
    position: "bottom"
},
title:{
text: "Avarage Rainfall per month",
display: true,
fontSize:20
}

  }

};*/

    // Renderizar el componente según el estado de los datos
    return (
        <div className='Monitoreo_r' >
         <div >
        <Pie type='line' options={opciones} title="cpu" style={{height:"50vh"}} data={datacpu}  />
       
        </div>
        <div >
       <Pie type='line' title="ram" style={{height:"50vh"}} data={data} options={opciones}  />
 
        </div>
          </div> 
    );
  };

export default Monitoreo_r;

/*
 <Pie type='line' title="ram" style={{height:"50vh"}} data={data} options={opciones}  />
         
        <p style={{textAlign: "center"}}>  pie</p>  
<div className='Monitoreo_r'>
        <Pie data =  {data} options={opciones}/>
          </div> 
          */


/* <div>
        {datosRam ? (
          <div>
            <p>Total RAM: {datosRam.totalRam}</p>
            <p>Memoria en uso: {datosRam.memoriaEnUso}</p>
            <p>Porcentaje: {porcentajeUso}</p>
            <p>Libre: {datosRam.libre}</p>
          </div>
        ) : (
          <p>Cargando datos...</p>
        )}
      </div>
       */

/*

import React, { useState, useEffect } from 'react';
import { Doughnut } from 'react-chartjs-2';

function Monitoreo_r(){
  const [data, setData] = useState({
    labels: ['En Uso', 'Libre'],
    datasets: [
      {
        data: [0, 100], // Inicialmente establecemos el 100% como libre
        backgroundColor: ['#FF6384', '#36A2EB'],
        hoverBackgroundColor: ['#FF6384', '#36A2EB'],
      },
    ],
  });

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:5200/ram');
        const jsonData = await response.json();

        // Actualizamos el estado con los nuevos datos
        setData({
          labels: ['En Uso', 'Libre'],
          datasets: [
            {
              data: [jsonData.memoriaEnUso, jsonData.libre],
              backgroundColor: ['#FF6384', '#36A2EB'],
              hoverBackgroundColor: ['#FF6384', '#36A2EB'],
            },
          ],
        });
      } catch (error) {
        console.error('Error al obtener los datos:', error);
      }
    };

    // Actualizamos los datos cada cierto intervalo de tiempo (por ejemplo, cada 5 segundos)
    const intervalId = setInterval(fetchData, 5000);

    // Limpieza del intervalo al desmontar el componente
    return () => clearInterval(intervalId);
  }, []);

  return (
    <div>
      <h2>Gráfica en Tiempo Real de Uso de Memoria RAM</h2>
      <Doughnut data={data} />
    </div>
  );
};

export default Monitoreo_r;*/

/*



import React, { useState, useEffect } from 'react';
import { Pie } from 'react-chartjs-2';

const Monitoreo_r = () => {
  const [datosRam, setDatosRam] = useState(null);
  const [porcentaje, setPorcentaje] = useState(null);
  const [datosCargados, setDatosCargados] = useState(false);

  const obtenerDatosRam = async () => {
    try {
      const respuesta = await fetch('http://localhost:5200/ram');

      if (!respuesta.ok) {
        throw new Error('Error al obtener los datos de RAM');
      }

      const datos = await respuesta.json();
      setDatosRam(datos);

      const porcentajeNumerico = parseFloat(datos.porcentaje);
      setPorcentaje(isNaN(porcentajeNumerico) ? null : porcentajeNumerico);
      setDatosCargados(true);
    } catch (error) {
      console.error(error.message);
    }
  };
  const pieRef = useRef(null);
  useEffect(() => {
    obtenerDatosRam();

    const intervalo = setInterval(() => {
      obtenerDatosRam();
    }, 1000);

       return () => {
      // Limpiar el canvas al desmontar el componente
      const chartInstance = pieRef.current?.chartInstance;
      if (chartInstance) {
        const { ctx } = chartInstance.chart;
        ctx.clearRect(0, 0, ctx.canvas.width, ctx.canvas.height);
      }
    };
  }, []);

  const porcentajeUso = porcentaje !== null ? porcentaje : 0;
  const data = {
    labels: ['total_ram', 'memoria_en_uso'],
    datasets: [
      {
        data: [100 - porcentajeUso, porcentajeUso],
        backgroundColor: ['blue', 'green'],
      },
    ],
  };

  const opciones = {
    responsive: true,
  };

  return (
       <div className='Monitoreo_r'>
      {datosCargados ? (
        <Pie ref={pieRef} data={data} options={opciones} />
      ) : (
        <p>Cargando datos...</p>
      )}
    </div>
  );
};

export default Monitoreo_r;


*/



/* 
 const data = { 
    labels: ['total_ram','memoria_en_uso'],
    datasets:  [{
        data: [90,10],
        backgroundColor: ['blue','green']
    }]

};

const opciones= {
    responsive: true
};

*/