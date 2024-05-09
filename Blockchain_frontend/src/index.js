import React, { useState } from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { createContext } from 'react';
import "bootstrap/dist/css/bootstrap.css"
import "bootstrap/dist/js/bootstrap"
import {BrowserRouter} from 'react-router-dom';




export const Context  = createContext();


const AppWrapper = () =>{
  const [isAuthenticated , setIsAuthenticated] = useState(false);// this one is for patient
  const [isAuthenticatedAdmin , setIsAuthenticatedAdmin] = useState(false);
  const [loading, setLoading] = useState(false);
  const [userId, setUserId] = useState(0);
  
  return(
    <Context.Provider value={{
      isAuthenticated,setIsAuthenticated,
      isAuthenticatedAdmin , setIsAuthenticatedAdmin,
      loading,setLoading,
      userId,setUserId
    }}>
    <App />
    </Context.Provider> 
  );
}


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <BrowserRouter>
  <React.StrictMode>
  <AppWrapper/>
  </React.StrictMode>
  </BrowserRouter>
);


// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
