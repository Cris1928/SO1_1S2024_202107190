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

## ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ Base de datis MySQL
La base de datos utillizada será en MySQL, y este realizara la conexion en el backend de golang, esto utilizando las variables de entorno "DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME", estas serviran para poder realizar dicha conexion con la base de datos.
```

var conexion = ConectarBD()

func ConectarBD() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Conectar a MySQL sin especificar la base de datos
	conexionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", dbUser, dbPassword, dbHost, dbPort)

	conexion, err := sql.Open("mysql", conexionString)
	if err != nil {
		log.Fatal(err)
	}

	// Verificar si la base de datos existe, de lo contrario, crearla
	if err := verificarBaseDatos(conexion, dbName); err != nil {
		log.Fatal(err)
	}

	// Conectar a la base de datos especificada
	conexionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	conexion, err = sql.Open("mysql", conexionString)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		crearTabla(conexion, "ram")
		crearTabla(conexion, "cpu")
		fmt.Println("Conexion con MySQL Correcta")
	}
	

	return conexion
}

```



## ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ PLATAFORMA DE MONITOREO 
## Monitoreo en tiempo real

![WhatsApp Image 2024-03-17 at 4 28 27 AM](https://github.com/Cris1928/SO1_1S2024_202107190/assets/98928867/973c43d9-1fde-4ef8-b3a3-b4a22fa54710)

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

en el codigo react utilizaremos la libreria "react-chartjs-2" este recogera la informacion del json que posee el localhost "/api/ram" utilizando "await", estos datos seran recogidos por medio de useState para recoger esa informacion, estos los utilizaran las variables "porcentaje" y "porcentajecpu", posteriormete para crear las datas que poseeran los datos de esta grafica y utilizarlo en el codigo html del react.




```jsx
import React, { useState, useEffect } from 'react';
import 'chart.js/auto';
import { Pie } from 'react-chartjs-2';

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
            setPorcentaje(isNaN(porcentajeNumerico) ? null : porcentajeNumerico);

            const porcentajeNumericocpu = parseFloat(datoscpu.cpu_porcentaje);
            setPorcentajecpu(isNaN(porcentajeNumericocpu) ? null : porcentajeNumericocpu);
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
    const porcentajeUso = porcentaje !== null ? porcentaje : 0;
    const data = { 
        labels: ['total_ram', 'memoria_en_uso'],
        datasets:  [{
            data: [porcentajeUso, 100 - porcentajeUso],
            backgroundColor: ['green', 'blue']
        }]
    };

    //configuracion para la grafica ram
    const opciones = {
        maintainAspectRatio: false,
        responsive: true,
    };
  
    //datos para la grafica cpu
    const porcentajeUsocpu = porcentajecpu !== null ? porcentajecpu : 0;
    const datacpu = { 
        labels: ['total_cpu', 'memoria_en_uso'],
        datasets:  [{
            label: "CPU",
            data: [100 - porcentajeUsocpu, porcentajeUsocpu],
            backgroundColor: ['blue', 'green']
        }]
    };

    //configuracion para la grafica cpu
    const opcionescpu = {
        maintainAspectRatio: false,
        responsive: true,
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

```




## Monitoreo a lo largo del Tiempo

![WhatsApp Image 2024-03-17 at 5 50 27 AM](https://github.com/Cris1928/SO1_1S2024_202107190/assets/98928867/90ba8549-de57-444d-8e25-93af65d9f0d2)

![WhatsApp Image 2024-03-17 at 5 50 45 AM](https://github.com/Cris1928/SO1_1S2024_202107190/assets/98928867/2abf3231-d0fa-4f07-921d-f6e3fb8db1b9)


Este posee la informacion de los ultimos 10 registros del porcentaje de uso de ram y cpu, el backend enviara esa informacion proveniente de la base de datos a la direccion "localhost:5200/api/data/datarames"  y "localhost:5200/api/data/datascpus".
```
func GetDataram2(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var lista []Dataram
	//now := time.Now()
	//re := fmt.Sprint("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	query := "SELECT Fecha, Porcentaje FROM ram ORDER BY Fecha DESC LIMIT 10;"
	result, err := conexion.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	for result.Next() {
		var datram Dataram

		err = result.Scan(&datram.Fecha, &datram.Porcentaje)
		if err != nil {
			fmt.Println(err)
		}
		lista = append(lista, datram)
	}
	json.NewEncoder(response).Encode(lista)
}
```



posteriormente en el codigo de react utilizaremos  las direcciones antes mencionadas, utilizando "await" recogerá la informacion del localhost, utilizando useState, podremos recoger esa informacion para poder utilizarla en el el codigo, utilizando un boton actualizaremos la informacion cada que lo requiramos.
```
import React, { useState, useEffect } from 'react';
import 'chart.js/auto';
import { Line } from 'react-chartjs-2';
import Monitoreo_tcpu from './Monitoreo_tcpu.jsx'

const Monitoreo_t = () => {
  const [chartData, setChartData] = useState({ labels: [], datasets: [] });

  const fetchData = () => {
    // Llamada a la API para obtener datos
    fetch('/api/datasrames')

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

```




## Árbol de Procesos

![WhatsApp Image 2024-03-17 at 5 51 28 AM](https://github.com/Cris1928/SO1_1S2024_202107190/assets/98928867/d4688aa5-a7c5-4490-9f52-7743181d9968)


Este graficara los procesos quqe posee el modulo de cpu, anteriormente se a explicado el como se recoge la informacion de este, en el codigo react utilizaremos "await" para poder tomar la informacion del localhost, al utilizar useState recogera la informacion de los procesos  como el nombre, pid y  los procesos hijos de este, siempre y cuando los posea, el pid de los procesos serivran para identificar que grafico se desea, estos se almacenaran en una lista para poder mostrarla en una lista desplegable, estos serviran para poder seleccionarlos y con un boton podra ejecutar el procesos de graficado utilizndo la libreria "graphviz-react".

```
import React, { useState, useRef } from 'react';
import Graphviz from 'graphviz-react';
import '../Arbol.css';

const GraphWrapper = React.forwardRef(({ dot, options }, ref) => {
  return (
    <div ref={ref}>
      <Graphviz dot={dot} options={options} />
    </div>
  );
});

const Arbol = () => {
  const [selectedPid, setSelectedPid] = useState(null);
  const [processes, setProcesses] = useState([]);
  const [showGraph, setShowGraph] = useState(false);
  const [childrensl, setChildrens] = useState([]);
  const [selectedName, setSelectedName] = useState(null);
  const graphRef = useRef(null);

  // Función para manejar el cambio en la lista desplegable
  const handleSelectChange = (event) => {
    const selectedValue = event.target.value;
    setSelectedPid(selectedValue);
  };

  // Función para manejar el clic en el botón
  const handleButtonClick = () => {
    if (selectedPid) {
      const selectedProcess = processes.find((process) => process.pid === parseInt(selectedPid));
      setSelectedName(selectedProcess.name);
      if (selectedProcess && selectedProcess.child) {
        setChildrens(selectedProcess.child);
        
        
      }
      setShowGraph(true);
    }
  };


  const generateDot =() => {
    let dot = graph  {
  
  
  
    dot += \n"${selectedPid}/${selectedName}";;
    childrensl.forEach((item) => {
  
      dot += \n"${selectedPid}/${selectedName}";
     // dot +=/${selectedName}";
      dot += ` --"${item.pid}/${item.name}";`;
    }
    );
    dot += ` \n}`;
  
    console.log(dot);
    return dot;
  };
  
  

  const handleDownload = () => {
    const svgElement = graphRef.current.querySelector('svg');
    const svgData = new XMLSerializer().serializeToString(svgElement);

    const downloadLink = document.createElement('a');
    downloadLink.href = 'data:image/svg+xml;charset=utf-8,' + encodeURIComponent(svgData);
    downloadLink.download = 'graph.svg';
    downloadLink.click();
  };

  // Simula la carga inicial de datos desde /api/cpu
  const fetchData = async () => {
    try {
      const response = await fetch('/api/cpu');
      const data = await response.json();
      setProcesses(data.processes);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  // Llama a fetchData cuando se monta el componente
  React.useEffect(() => {
    fetchData();
  }, []);

  return (
    <div>
      <h1>Procesos</h1>
      <label htmlFor="pidList">Selecciona un PID:</label>
      <select id="pidList" onChange={handleSelectChange}>
        <option value="">Seleccione un PID</option>
        {processes.map((process) => (
          <option key={process.pid} value={process.pid}>
            {process.pid}
          </option>
        ))}
      </select>
      <button onClick={handleButtonClick}>Mostrar gráfico</button>

      {showGraph && (
        <div>
          <GraphWrapper dot={generateDot()} options={{ width: 900, height: 900 }} ref={graphRef} />
          
          <button onClick={handleDownload}>Descargar SVG</button>
          
        </div>
      )}
    </div>
  );
};

export default Arbol;
```

## Diagrama de Estado de Procesos

![WhatsApp Image 2024-03-17 at 5 51 57 AM](https://github.com/Cris1928/SO1_1S2024_202107190/assets/98928867/9025bbf1-f3fa-4255-9bd8-0658df257693)

Este posee:
- Un Diagrama que muestre los Estados que han tenido un proceso.
- Cuando ocurre un cambio de estado muestra el estado actual en el que seencuentra el proceso, es decir el nodo y arista de la transición poseera un colordiferente.
el backend maneja 4 endpoints lso cuales son, "localhost:5200/api/start", "localhost:5200/api/stop", "localhost:5200/api/resume" y "localhost:5200/api/kill", estos funcionaran para poner en el estado del proceso en especifico indicado desde el fronted.

```
func StopProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGSTOP al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-SIGSTOP", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al detener el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d detenido", pid)
}

```


al poner el estado seleccionado al procesos seleccionado este mostrara un grafico representativo del estado en especifico, esto utilizando la libreria "graphviz-ract".

