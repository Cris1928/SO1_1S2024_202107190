import { useState } from 'react'

import styled from 'styled-components'
function Navbar() {
  return (
    <>
       <NavContainer >
       <h2 >SO1<span>P1</span></h2>
       <div> 
       <a  href="/monitoreo_r"> Monitoreo</a>
       <a href="/monitoreo_t"> Historico</a>
       <a href="/arbol"> Arbol</a>
       <a  href="/simulacion_1">Estados </a>
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
