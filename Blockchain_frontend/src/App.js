import React, { useContext, useEffect } from 'react'
import Home from './components/NotLogin/home.js'
import {Routes,Route} from 'react-router-dom'
import Login from './components/NotLogin/login.js'
import Signup from './components/NotLogin/signup.js'
import UserHome from './components/Employee/userHome.js'

import axios from 'axios';
import { Toaster } from 'react-hot-toast';
import { Context } from './index.js'
import AdminHome from './components/Employer/adminHome.js'

const App = () => {

  const {user , setUser ,isAuthenticatedAdmin ,setIsAuthenticatedAdmin , setIsAuthenticated, setLoading} = useContext(Context);




  return(
    <>
    
    <Routes>

    <Route path="/" element= {<Home/>}/>;
    <Route path="/employee/home" element= {<UserHome/>}/>;
    <Route path="/login" element= {<Login/>}/>;
    <Route path="/signup" element= {<Signup/>}/>;
    <Route path="/employer/home" element= {<AdminHome/>}/>;

    </Routes>
    <Toaster/>
    </>
  )
}

export default App
