# ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ MANUAL TECNICO

El siguiente proyecto tiene como objetivo implementar un sistema de monitoreo derecursos del sistema y gestión de procesos, empleando varias tecnologías y lenguajes de programación.

## ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ Kernel Modules

## Módulo de RAM
Este modulo está ubicado en el directorio /proc, desplegado desde el archvivo "ram_so1_1s2024.ko" este recoge el porcentaje de la ram utilizada y el porcentaje de la ram libre.
Pasos para utilizarla desde el directorio /proc:
- En el directorio del archivo "ram_so1_1s2024.ko" ejecutar el comando "sudo make all" y este generara los archivos necesarios para su despliegue en el directorio /proc
- Para desplegarlo en el directorio /proc ejecutar el comando "sudo insmod ram_so1_1s2024.ko"
- Leer el archivo escrito cat ram_so1_1s2024

## Módulo de CPU
Este modulo está ubicado en el directorio /proc, desplegado desde el archvivo "cpu_so1_1s2024.ko" este recoge el porcentaje del cpu utilizado y el porcentaje del cpu libre, tambien recoge los procesos padres e hijos que la computadora utilia en ese momento.
Pasos para utilizarla desde el directorio /proc:
- En el directorio del archivo "cpu_so1_1s2024.ko" ejecutar el comando "sudo make all" y este generara los archivos necesarios para su despliegue en el directorio /proc
- Para desplegarlo en el directorio /proc ejecutar el comando "sudo insmod cpu_so1_1s2024.ko"
- Leer el archivo escrito cat cpu_so1_1s2024

## ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ PLATAFORMA DE MONITOREO 
## Monitoreo en tiempo real
Este posee el porcentaje de ram utilizada y libre desplegada en una grafica circular, de igual manera con el cpu, esto utilizando la libreria "react-cjartjs-2".
Desde el backend posee los ednpoints "localhost:5200/api/cpu" y "localhost:5200/api/ram" estos poseen la informacion del ultimo registro de su porcentaje en utilizacion de la ram y el cpu.

func RAMModuleHandler(w http.ResponseWriter, r *http.Request) {

	cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
	outRam, err := cmdRam.CombinedOutput()
	if err != nil {
		fmt.Println("eerror", err)
	}
	//fmt.Println("-------------------- RAM --------------------")
	var ram_info Ram
	err = json.Unmarshal([]byte(outRam), &ram_info)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode((ram_info))

}



en el metodo RAMModuleHandler este recoge la infomracion que este despliega y lo envia a la ruta localhost:5200/api/ram.

en el codigo react utilizaremos la libreria "react-chartjs-2" este recogera la informacion del json que posee el localhost "/api/ram" utilizando "await", estos datos seran recogidos por medio de useState para recoger esa informacion, estos los utilizaran las variables "porcentaje" y "porcentajecpu", posteriormete para crear las datas que poseeran los datos de esta grafica y utilizarlo en el codigo html del react





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
        const respuesta = await fetch('/api/ram');
        const respuestacpu = await fetch('/api/cpu');
  
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



![WhatsApp Image 2024-03-17 at 4 28 27 AM](https://github.com/Cris1928/SO1_1S2024_202107190/assets/98928867/973c43d9-1fde-4ef8-b3a3-b4a22fa54710)


