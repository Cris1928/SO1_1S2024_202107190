
const express = require('express');
const bodyParser = require('body-parser');
const app = express();
const port = 3001;

app.use(bodyParser.json());

app.post('/upload', (req, res) => {
  const { base64Image, currentDate } = req.body;
  // Aquí puedes realizar el procesamiento de la imagen y la fecha como desees
  // Por ahora, simplemente loguearemos los datos recibidos
  console.log('Base64 Image:', base64Image);
  console.log('Current Date:', currentDate);

  res.json({ message: 'Data received successfully.' });
});

app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});
/*const express = require ("express")
const app = express()

app.get("/api", (req,res) => { 
    res.json({"user":"fff"})
 })

 app.listen(5000, () =>{console.log("puerto 5000")} )*/