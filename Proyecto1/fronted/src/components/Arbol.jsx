import React, { useState, useRef } from 'react';
import Graphviz from 'graphviz-react';

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

/*
  const handleButtonClick = () => {
    if (selectedPid) {
      const selectedProcess = processes.find((process) => process.pid === parseInt(selectedPid));
      setSelectedName(selectedProcess.name);
      if (selectedProcess && selectedProcess.child) {
        setChildrensl(selectedProcess.child);
        selectedProcess.child.forEach((child) => {
        });
     
      }
      setShowGraphl(true);
    }
  };
*/

  const generateDot =() => {
    let dot = `graph  {`
  
  
  
    dot += `\n"${selectedPid}/${selectedName}";`;
    childrensl.forEach((item) => {
  
      dot += `\n"${selectedPid}/${selectedName}"`;
     // dot +=`/${selectedName}"`;
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

  // Simula la carga inicial de datos desde http://localhost:5200/cpu
  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost:5200/cpu');
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





/* 

 {showGraph &&(
        <Graphviz dot={generateDot()} options={{width:300, height:300}}/>
      )}
*/


/**
 const handleDownload = () => {
  const svgData = graphRef.current.getSvgData();
  const blob = new Blob([svgData], { type: 'image/svg+xml' });
  const url = URL.createObjectURL(blob);

  const link = document.createElement('a');
  link.href = url;

  link.download = 'graph.svg';
  link.click();
  URL.revokeObjectURL(url);
};
 */


/*
import React, { useState,useRef } from 'react';
import Graphviz from 'graphviz-react';
import * as FileSaver from 'file-saver';

const Arbol = () => {  
  const [selectedPid, setSelectedPid] = useState(null);
  const [processes, setProcesses] = useState([]);
  const [showGraph, setShowGraphl] = useState(false);
  const [childrensl, setChildrensl] = useState([]);
  const [selectedName, setSelectedName] = useState(null);
  const graphRef = useRef(null);
  const [graphReady, setGraphReady] = useState(false);


  // Función para manejar el cambio en la lista desplegable
  const handleSelectChange = (event) => {
    const selectedValue = event.target.value;
    setSelectedPid(selectedValue);
  };

  // Función para manejar el clic en el botón
  const handleButtonClick = () => {
    if (selectedPid) {
      const selectedProcess = processes.find((process) => process.pid === parseInt(selectedPid));
      if (selectedProcess && selectedProcess.child) {
        setChildrensl(selectedProcess.child);
        setSelectedName(selectedProcess.name);
        //console.log(selectedProcess.name);
      //  console.log("selectedPid")
      //  console.log(selectedPid)
       // console.log(selectedProcess.child)

        setShowGraphl(true);


        selectedProcess.child.forEach((child) => {
          console.log(child);
        });
        setGraphReady(true); 
      }
    }
  };
const generateDot =() => {
  let dot = `graph  {`
  childrensl.forEach((item) => {

    dot += `\n"${selectedPid}/${selectedName}"`;
   // dot +=`/${selectedName}"`;
    dot += ` --"${item.pid}/${item.name}";`;
  }
  );
  dot += ` \n}`;

  console.log(dot);
  return dot;
};

const handleDownload = () => {
  const svgData = graphRef.current.getSvg(); // Obtén el contenido SVG de graphRef
  const blob = new Blob([svgData], { type: 'image/svg+xml' });
  FileSaver.saveAs(blob, 'grafo.svg'); // Descargar el archivo SVG
};



  // Simula la carga inicial de datos desde http://localhost:5200/cpu
  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost:5200/cpu');
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
      <button onClick={handleButtonClick}>Mostrar grafico</button>
     
     
      {showGraph &&(
        <> 
        <Graphviz ref={graphRef}  dot={generateDot()} options={{width:900, height:900}}/>
        <button onClick={handleDownload} disabled={!graphReady}>
        Descargar gráfica
      </button>
      </>
      )}
    </div>
  );
};

export default Arbol;



 */



/*  el bueno


import React, { useState,useRef } from 'react';
import Graphviz from 'graphviz-react';


const Arbol = () => {  
  const [selectedPid, setSelectedPid] = useState(null);
  const [processes, setProcesses] = useState([]);
  const [showGraph, setShowGraphl] = useState(false);
  const [childrensl, setChildrensl] = useState([]);
  const [selectedName, setSelectedName] = useState(null);


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
        setChildrensl(selectedProcess.child);
        //setSelectedName(selectedProcess.name);


        //console.log(selectedProcess.name);
      //  console.log("selectedPid")
      //  console.log(selectedPid)
       // console.log(selectedProcess.child)

       // setShowGraphl(true);
        selectedProcess.child.forEach((child) => {
          console.log(child);
        });
     
      }
      setShowGraphl(true);
    }
  };
const generateDot =() => {
  let dot = `graph  {`



  dot += `\n"${selectedPid}/${selectedName}";`;
  childrensl.forEach((item) => {

    dot += `\n"${selectedPid}/${selectedName}"`;
   // dot +=`/${selectedName}"`;
    dot += ` --"${item.pid}/${item.name}";`;
  }
  );
  dot += ` \n}`;

  console.log(dot);
  return dot;
};




  // Simula la carga inicial de datos desde http://localhost:5200/cpu
  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost:5200/cpu');
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
      <button onClick={handleButtonClick}>Mostrar grafico</button>
     
     
      {showGraph &&(
      
        <Graphviz   dot={generateDot()} options={{width:900, height:900}}/>
       
      
      )}
    </div>
  );
};

export default Arbol;
*/