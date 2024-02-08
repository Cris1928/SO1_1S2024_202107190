
import React, { useState,useRef,useEffect } from 'react';
import Webcam from 'react-webcam';

const App: React.FC = () => {
  const inputRef = useRef<HTMLInputElement>(null);
  const webcamRef = useRef<Webcam>(null);
  const [loadedPhotos, setLoadedPhotos] = useState<string[]>([]);

  const handleCapture = async () => {
    try {
      const input = inputRef.current;
     
  
      if (input && input.files && input.files.length > 0) {
        const file = input.files[0];
        const base64 = await convertToBase64(file);
        const currentDate = new Date().toLocaleString();
  
        sendToServer({ base64Image: base64, currentDate });
      } else if (webcamRef.current) {
        const screenshot = webcamRef.current.getScreenshot();
        if (screenshot) {
          const [format, base64] = screenshot.split(',');
          const currentDate = new Date().toLocaleString();
          sendToServer({ base64Image: ` data:${format}, ${base64}`, currentDate });
        } else {
          console.error('No se capturó ninguna imagen de la cámara web.');
        }
      } else {
        console.error('No se seleccionó ningún archivo o imagen de cámara web.');
      }
    } catch (error) {
      console.error('Error al capturar la imagen:', error);
    }
  };
  const convertToBase64 = (file: File) => {
    return new Promise<string>((resolve, reject) => {
      const reader = new FileReader();
      reader.onloadend = () => {
        resolve(reader.result as string);
      };
      reader.onerror = reject;
      reader.readAsDataURL(file);
    });
  };

  const sendToServer = async (data: { base64Image: string; currentDate: string }) => {
    try {
      const response = await fetch('http://localhost:5000/upload', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      const result = await response.json();
      console.log('Respuesta del servidor:', result);
    } catch (error) {
      console.error('Error al enviar datos al servidor:', error);
    }
  };
  useEffect(() => {
    // Cargar las fotos al montar el componente
    fetch('http://localhost:5000/photos')
      .then(response => response.json())
      .then(data => {
        const base64Images = data.map((photo: { base64Image: string }) => photo.base64Image);
        setLoadedPhotos(base64Images);
      })
      .catch(error => console.error('Error fetching photos:', error));
  }, []);

  return (
    <div>
      <h1>Tarea 2</h1>
      <input type="file" accept="image/*" ref={inputRef} />
      <Webcam
        audio={false}
        ref={webcamRef}
        screenshotFormat="image/jpeg"
        width={640}
        height={480}
      />
      <button  onClick={handleCapture}>CAPTURAR</button>




      <h2>FOTOS:</h2>
      <div style={{ display: 'flex', flexWrap: 'wrap' }}>
        {loadedPhotos.map((photo, index) => (
          <div key={index} style={{ margin: '10px' }}>
            <img src={photo} alt={` Loaded ${index}`} style={{ width: '150px', height: '150px' }} />
          </div>
        ))}
      </div>


      
    </div>
  );
};

export default App;