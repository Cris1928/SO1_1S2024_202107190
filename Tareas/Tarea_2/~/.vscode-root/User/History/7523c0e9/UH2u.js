const express = require ("express")
const app = express()

app.get("/api", (req,res) => { 
    res.json({"user":"fff"})
 })

 app.listen(5000, () =>{console.log("puerto 5000")} )