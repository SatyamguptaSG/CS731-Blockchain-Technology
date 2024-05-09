import React from 'react'
import { NavLink } from 'react-router-dom'
import "../NotLogin/navbar.css"

const AdminNavbar = () => {
  return (
    
      
    <div>
       <nav className="navbar navbar-expand-lg navbar-light bg-red">
        <div className="container-fluid my-0">
          <NavLink className="navbar-brand" to="/">
            <div className="logo-cls">
            
              <h5 className="logotext" style={{  marginLeft: '10px', marginTop: '5px', background: 'linear-gradient(to right,  #000000 ,#000001)', WebkitBackgroundClip: 'text', color: 'transparent', fontFamily: 'Helvetica Neue' }}> <b> BorderPay </b></h5>
            </div>
          </NavLink>
          <button
            className="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarSupportedContent"
            aria-controls="navbarSupportedContent"
            aria-expanded="false"
            aria-label="Toggle navigation"
            onClick=""///some handler were calling from here 
          >
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse justify-content-end" id="navbarSupportedContent">
            <ul className="navbar-nav ml-auto">
              
              
              <li className="nav-item">
                <NavLink className="nav-link active" aria-current="page" to="/">
                <h5 style={{fontFamily: 'Helvetica Neue', marginTop: '5px'}}>Logout</h5> 
                </NavLink>
                
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </div>
  )

}

export default AdminNavbar
