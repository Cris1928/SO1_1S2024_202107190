import { useState } from 'react'

import styled from 'styled-components'
function Navbar() {
  return (
    <>
       <NavContainer >
       <h2 >SO1<span>P1</span></h2>
       <div> 
       <a  href="/monitoreo_r"> Monitoreo_r</a>
       <a href="/monitoreo_t"> Monitoreo_t</a>
       <a href="/arbol"> Arbol</a>
       <a  href="/simulacion_1">Simulacion_1 </a>
       <a  href="/simulacion_2">Simulacion_2 </a>

       </div>
     </NavContainer>
    </>
  )
}

export default Navbar



const NavContainer = styled.nav `
h2 {
    color: white;
    font-weight: 400;
    span{
        font-weight: bold;
    }  

}  
padding: .4rem;
background-color:#8f2424;
display: flex;
align-items: center;
justify-content: space-between;
a{
color: white;
text-decoration: none;
margin-right:1rem;

} 
`
