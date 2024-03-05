import React from 'react'
import * as ReactDOM from 'react-dom/client'
import App from './App.jsx'
import { createBrowserRouter, RouterProvider, Route} from 'react-router-dom'
import Navbar from './components/Navbar'
import Monitoreo_r from './components/Monitoreo_r.jsx'
import Monitoreo_t from './components/Monitoreo_t.jsx'
import Arbol from './components/Arbol.jsx'
import Simulacion_1 from './components/Simulacion_1.jsx'
import Simulacion_2 from './components/Simulacion_2.jsx'


function Home(){
  return (
    <>
    <Navbar/>
  </>
  )
}

function Monitoreo_rr(){
  return (
    <div>
      <Navbar/>
     <Monitoreo_r/>
    
  </div>
  )
}
function Monitoreo_tt(){
  return (
    <div>
      <Navbar/>
     <Monitoreo_t/>
    
  </div>
  )
}
function Arboll(){
  return (
    <div>
      <Navbar/>
     <Arbol/>
    
  </div>
  )
}
function Simulacion__1(){
  return (
    <div>
      <Navbar/>
     <Simulacion_1/>
    
  </div>
  )
}
function Simulacion__2(){
  return (
    <div>
      <Navbar/>
     <Simulacion_2/>
    
  </div>
  )
}
const router = createBrowserRouter([ 
  {
    path: "/",
    element: <Home/>,
  },
  {
    path: "/monitoreo_r",
    element: <Monitoreo_rr/>,
  },
  {
    path: "/monitoreo_t",
    element: <Monitoreo_tt/>,
  },
  {
    path: "/arbol",
    element: <Arboll/>,
  },
  {
    path: "/simulacion_1",
    element: <Simulacion__1/>,
  },
  {
    path: "/simulacion_2",
    element: <Simulacion__2/>,
  },
] )


ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>,

)
export default router;