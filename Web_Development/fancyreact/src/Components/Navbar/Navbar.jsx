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
            <li>Shop</li>
            <li>Men</li>
            <li>Women</li>
            <li>Kids</li>
        </ul>
        <div className="nav-login-cart">
            <button>Login</button>
            <img src={cart_icon} alt="" />
        </div>
        </div>
    )
}

export default Navbar