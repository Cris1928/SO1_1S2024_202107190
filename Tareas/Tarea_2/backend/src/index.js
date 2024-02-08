const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const mongoose = require('mongoose'); // Importa mongoose
const app = express();
const port = 5000;

app.use(bodyParser.json());
app.use(cors());

// Conéctate a MongoDB
mongoose.connect('mongodb://MongoDBD:27017/photoApp', {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});
const db = mongoose.connection;

db.on('error', console.error.bind(console, 'Error de conexión de MongoDB:'));
db.once('open', () => {
  console.log('Conectado a MongoDB');
});

// Define el esquema del modelo
const photoSchema = new mongoose.Schema({
  base64Image: String,
  currentDate: String,
});

// Crea el modelo
const Photo = mongoose.model('Photo', photoSchema);

app.post('/upload', async (req, res) => {
  const { base64Image, currentDate } = req.body;

  try {
    // Crea una nueva instancia del modelo y guarda en la base de datos
    const newPhoto = new Photo({ base64Image, currentDate });
    await newPhoto.save();

    console.log('Datos guardados en MongoDB:', newPhoto);

    res.json({ message: 'Datos recibidos y guardados exitosamente.' });
  } catch (error) {
    console.error('Error al guardar datos en MongoDB', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

app.listen(port, () => {
  console.log(`El servidor se está ejecutando en http://localhost:${port}`);
});

app.get('/photos', async (req, res) => {
  try {
    // Obtén todas las fotos de la base de datos
    const photos = await Photo.find({}, { _id: 0, __v: 0 });

    res.json(photos);
  } catch (error) {
    console.error('Error al recuperar fotos de MongoDB:', error);
    res.status(500).json({ error: 'Error Interno del Servidor' });
  }
});


/*const express = require ("express")
const app = express()

app.get("/api", (req,res) => { 
    res.json({"user":"fff"})
 })

 app.listen(5000, () =>{console.log("puerto 5000")} )*/