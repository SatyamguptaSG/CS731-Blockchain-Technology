import React, { useContext, useEffect, useState } from "react";
import { NavLink, Navigate } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import "./navbar.css";
import { Context } from "../..";

const Navbar = () => {
  const { isAuthenticatedAdmin, isAuthenticated } = useContext(Context);
  const [isSticky, setIsSticky] = useState(false);


  useEffect(() => {
    const handleScroll = () => {
      setIsSticky(window.scrollY > 0);
    };

    window.addEventListener("scroll", handleScroll);

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);



  if (isAuthenticated) {
    return <Navigate to="/" />;
  }

  if (isAuthenticatedAdmin) {
    return <Navigate to="/admin/home" />;
  }

  return (
    <div>
      <div className={`navbar1 ${isSticky ? "sticky" : ""}`}>
        <nav className="navbar navbar-expand-lg navbar-light bg-red">
          <div className="container-fluid my-0">
            <NavLink className="navbar-brand" to="/">
              <div className="logo-cls">
                
                <h5
                  className="logotext"
                  style={{
                    marginLeft: "10px",
                    marginTop: "5px",
                    background: "linear-gradient(to right,  #000000 ,#000001)",
                    WebkitBackgroundClip: "text",
                    color: "transparent",
                    fontFamily: "Helvetica Neue",
                  }}
                >
                                    <b> BorderPay </b>
                </h5>
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
              onClick="" ///some handler were calling from here
            >
              <span className="navbar-toggler-icon"></span>
            </button>
            <div
              className="collapse navbar-collapse justify-content-end"
              id="navbarSupportedContent"
            >
              <ul className="navbar-nav ml-auto">
                <li className="nav-item">
                  <NavLink
                    className="nav-link active"
                    aria-current="page"
                    to="/"
                  >
                    <h5 style={{ fontFamily: "Helvetica Neue" }}>Home</h5>
                  </NavLink>
                </li>
                
                <li className="nav-item">
                  <NavLink
                    className="nav-link active"
                    aria-current="page"
                    to="/login"
                  >
                    <h5 style={{ fontFamily: "Helvetica Neue" }}>Login</h5>
                  </NavLink>
                </li>
                <li className="nav-item">
                  <NavLink
                    className="nav-link active"
                    aria-current="page"
                    to="/signup"
                  >
                    <h5 style={{ fontFamily: "Helvetica Neue" }}>Signup</h5>
                  </NavLink>
                </li>
              </ul>
            </div>
          </div>
        </nav>
      </div>
    </div>
  );
};

export default Navbar;
