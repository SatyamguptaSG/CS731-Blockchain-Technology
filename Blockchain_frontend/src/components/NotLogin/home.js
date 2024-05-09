import React from "react";
import Navbar from "./navbar.js";
import "./home.css";

const Home = () => {
  window.localStorage.removeItem("userId");

  return (
    <div className="home-body">
      <Navbar />
      <h1>CS731 End Term submission by:</h1>
        <h3>Havi Bohra(210429)</h3> <h3> Satyam Gupta(210942)</h3> <h3>Tanmey Agarwal(211098)</h3>
    </div>
  );

};

export default Home;
