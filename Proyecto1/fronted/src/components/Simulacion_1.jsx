import React, { useState,useRef } from 'react';
import Graphviz from 'graphviz-react';
import '../Simulacion_1.css';
function Simulacion_1(){
    const [selectedPid, setSelectedPid] = useState(null);
    const [processes, setProcesses] = useState([]);
    const [showGraph, setShowGraphl] = useState(false);
    const [childrensl, setChildrensl] = useState([]);
    const [selectedName, setSelectedName] = useState(null);

    const [graphEstado, setGraphEstado] = useState(null);

  
  
    // Función para manejar el cambio en la lista desplegable
    const handleSelectChange = (event) => {
      const selectedValue = event.target.value;
      setSelectedPid(selectedValue);
    };
  
    // Función para manejar el clic en el botón
    const handleButtonNew = async () => {
        let dot =`
       
            node [label="Nodo", style="filled"]
            New [fillcolor="lightblue", label="New", shape="circle"];
            Ready [fillcolor="lightblue", label="Ready", shape="circle"];
            Running [fillcolor="green", label="Runing", shape="circle"];
            New -- Ready;
            Ready -- Running;
            
        
        `;
        setGraphEstado(dot);
      if (selectedPid) {
        const selectedProcess = processes.find((process) => process.pid === parseInt(selectedPid));
        setSelectedName(selectedProcess.name);
           // Enviar el PID al servidor
    try {
        const response = await fetch(`/api/start?pid=${selectedPid}`, {
          method: 'POST', // Puedes cambiar a 'GET' si prefieres un método GET
          headers: {
            'Content-Type': 'application/json',
          },
          // Puedes enviar más datos si es necesario
          body: JSON.stringify({
            pid: selectedPid,
            // Otros datos que desees enviar al servidor
          }),
        });
  
        // Verificar si la solicitud fue exitosa
        if (response.ok) {
          console.log(` PID ${selectedPid} enviado correctamente.`);
        } else {
          console.error(`Error al enviar el PID ${selectedPid}.`);
        }
      } catch (error) {
        console.error('Error al enviar la solicitud:', error);
      }
       
     
        if (selectedProcess && selectedProcess.child) {
          setChildrensl(selectedProcess.child);
    
          selectedProcess.child.forEach((child) => {
            console.log(child);
          });
       
        }
        setShowGraphl(true);
      }
    };
    const handleButtonStop = async () => {
        let dot =`
    
            node [label="Nodo", style="filled"]
            New [fillcolor="lightblue", label="New", shape="circle"];
            Ready [fillcolor="green", label="Ready", shape="circle"];
            Running [fillcolor="lightblue", label="Running", shape="circle"];
            New -- Ready;
            Ready -- Running;
            Running -- Ready [label="Transición", dir=forward, constraint=false, weight=1.5];
      
        `;
        setGraphEstado(dot);
      if (selectedPid) {
        const selectedProcess = processes.find((process) => process.pid === parseInt(selectedPid));
        setSelectedName(selectedProcess.name);


        try {
            const response = await fetch(`/api/stop?pid=${selectedPid}`, {
              method: 'POST', // Puedes cambiar a 'GET' si prefieres un método GET
              headers: {
                'Content-Type': 'application/json',
              },
              // Puedes enviar más datos si es necesario
              body: JSON.stringify({
                pid: selectedPid,
                // Otros datos que desees enviar al servidor
              }),
            });
      
            // Verificar si la solicitud fue exitosa
            if (response.ok) {
              console.log(` PID ${selectedPid} enviado correctamente.`);
            } else {
              console.error(`Error al enviar el PID ${selectedPid}.`);
            }
          } catch (error) {
            console.error('Error al enviar la solicitud:', error);
          }





        if (selectedProcess && selectedProcess.child) {
          setChildrensl(selectedProcess.child);
    
          selectedProcess.child.forEach((child) => {
            console.log(child);
          });
       
        }
        setShowGraphl(true);
      }
    };

    const handleButtonResume=  async () => {
        let dot =`
     
                node [label="Nodo", style="filled"]
                New [fillcolor="lightblue", label="New", shape="circle"];
                Ready [fillcolor="lightblue", label="Ready", shape="circle"];
                Running [fillcolor="green", label="Running", shape="circle"];
                New -- Ready;
              
              Ready -- Running [ dir=forward, constraint=false, weight=1.5];
            Running  -- Ready;
                
       
        `;
        setGraphEstado(dot);
      if (selectedPid) {
        const selectedProcess = processes.find((process) => process.pid === parseInt(selectedPid));
        setSelectedName(selectedProcess.name);
        try {
            const response = await fetch(`/api/resume?pid=${selectedPid}`, {
              method: 'POST', // Puedes cambiar a 'GET' si prefieres un método GET
              headers: {
                'Content-Type': 'application/json',
              },
              // Puedes enviar más datos si es necesario
              body: JSON.stringify({
                pid: selectedPid,
                // Otros datos que desees enviar al servidor
              }),
            });
      
            // Verificar si la solicitud fue exitosa
            if (response.ok) {
              console.log(` PID ${selectedPid} enviado correctamente.`);
            } else {
              console.error(`Error al enviar el PID ${selectedPid}.`);
            }
          } catch (error) {
            console.error('Error al enviar la solicitud:', error);
          }




        if (selectedProcess && selectedProcess.child) {
          setChildrensl(selectedProcess.child);
    
          selectedProcess.child.forEach((child) => {
            console.log(child);
          });
       
        }
        setShowGraphl(true);
      }
    };

    const handleButtonKill= async () => {
        let  dot =`
    
            node [label="Nodo", style="filled"]
            New [fillcolor="lightblue", label="New", shape="circle"];
            Ready [fillcolor="lightblue", label="Ready", shape="circle"];
            Running [fillcolor="lightblue", label="Running", shape="circle"];
            Terminated  [fillcolor="green", label="Terminated", shape="circle"];
            New -- Ready;
            Ready -- Running;
            
            Running -- Terminated;
        //    Running -- Ready [label="Transición", dir=forward, constraint=false, weight=1.5];
      
        `;
        setGraphEstado(dot);
      if (selectedPid) {
        const selectedProcess = processes.find((process) => process.pid === parseInt(selectedPid));
        setSelectedName(selectedProcess.name);
        try {
            const response = await fetch(`/api/kill?pid=${selectedPid}`, {
              method: 'POST', // Puedes cambiar a 'GET' si prefieres un método GET
              headers: {
                'Content-Type': 'application/json',
              },
              // Puedes enviar más datos si es necesario
              body: JSON.stringify({
                pid: selectedPid,
                // Otros datos que desees enviar al servidor
              }),
            });
      
            // Verificar si la solicitud fue exitosa
            if (response.ok) {
              console.log(` PID ${selectedPid} enviado correctamente.`);
            } else {
              console.error(`Error al enviar el PID ${selectedPid}.`);
            }
          } catch (error) {
            console.error('Error al enviar la solicitud:', error);
          }



        if (selectedProcess && selectedProcess.child) {
          setChildrensl(selectedProcess.child);
    
          selectedProcess.child.forEach((child) => {
            console.log(child);
          });
       
        }
        setShowGraphl(true);
      }
    };




  const generateDot =() => {
    let dot = `graph  {`
  
  
  
    dot +=graphEstado;
    dot += ` \n}`;
    console.log("dot");
    console.log(dot);
    return dot;
  };
  
  
  
  
    // Simula la carga inicial de datos desde/api/cpu
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
        <button onClick={handleButtonNew}>New</button>
        <button onClick={handleButtonStop}>Stop</button>
        <button onClick={handleButtonResume}>Resume</button>
        <button onClick={handleButtonKill}>Kill</button>
       
       
        {showGraph &&(
        
          <Graphviz   dot={generateDot()} options={{width:900, height:900}}/>
         
        
        )}
      </div>
    );
}
export default Simulacion_1;