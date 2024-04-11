import React from 'react' 
import './Navbar.css'
import logo from '../Assets/logo.png'
import cart_icon from '../Assets/cart_icon.png'

const Navbar = () => {
    return (
        <div clasName ='navbar'>
            <div clasName="nav-logo">
                <img src={logo} alt=""/>
                <p>Shopper</p>
                </div> 
        <ul className="nav-menu">
            
        </ul>
        </div>
    )
}

export default Navbar