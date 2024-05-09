import React, { useContext, useState } from "react";
import "./form.css";
import { Link, Navigate } from "react-router-dom";
import Navbar from "./navbar";
import { Context } from "../..";
import axios from "axios";
import toast from "react-hot-toast";

const Login = () => {
  const [userid, setuserid] = useState("");
  const [password, setPassword] = useState("");
  const [usertype, setUsertype] = useState();
  const [isLogined, setIsLogined] = useState(false);
  const [loginedUserType, setLoginedUserType] = useState();
  const {userId ,setUserId } = useContext(Context);


  
  const submitHandler = async (e) => {
    e.preventDefault();
    try {
      //first request
      // const { data } = await axios.post(
      //   "http://localhost:4000/api/v1/users/user/login",
      //   {
      //     userid,
      //     password,
          
      //   },
      //   {
      //     headers: {
      //       "Content-Type": "application/json",
      //     },
      //     withCredentials: true,
      //   }
      // );
      const data = await axios.post(
        'http://localhost:3002/invoke',
        new URLSearchParams([
          ['', ''],
          ['channelid', 'mychannel'],
          ['chaincodeid', 'basictest'],
          ['function', 'LoginFunc'],
          ['args', userid],
          ['args', password],
          ['args', usertype]
        ])
      );
      window.localStorage.setItem("userId", userid);
      // console.log(response.headers);
      setUserId(userid);
      setIsLogined(true);
    
      toast.success("login succesful");
    }  catch (error) {
        toast.error("Bad Credentials");
        setIsLogined(false);
      }
  };


  if(isLogined) {
    if(usertype==="1"){
      return <Navigate to={"/employer/home"} />;
      }
    else{
     return <Navigate to={"/employee/home"} />;
      }
  }

   return (
    <div>
      <Navbar />
      
      <div className="reg">
        <h2
          style={{
            textAlign: "center",
            color: "black",
            fontFamily: "Helvetica Neue",
          }}
        >
          Login
        </h2>
        <div className="form-container">
          <form
            className="card"
            style={{ background: "#eeeeee" }}
            onSubmit={submitHandler}
          >
            <div className="form my-4" style={{ textAlign: "center" }}>
              <div className="text-center my-2">
                &nbsp;&nbsp;
                <input
                  type="id"
                  name="id"
                  value={userid}
                  style={{ fontFamily: "Helvetica Neue" }}
                  onChange={(e) => {
                    setuserid(e.target.value);
                  }}
                  placeholder="User Id"
                  required
                />
                <br />
              </div>
              <div className="text-center my-2">
                &nbsp;&nbsp;
                <input
                  type="password"
                  name="password"
                  title="password must contains atleast 8 character"
                  value={password}
                  onChange={(e) => {
                    setPassword(e.target.value);
                  }}
                  placeholder="Password"
                  style={{ fontFamily: "Helvetica Neue" }}
                  spellcheck="false"
                  required
                />
                <br />
              </div>
              
              <div className="text-center my-2">
                &nbsp;&nbsp;
                <label>
                  <input
                    type="radio"
                    name="usertype"
                    value="1"
                    checked={usertype === "1"}
                    onChange={(e) => {
                      setUsertype(e.target.value);
                    }}
                    style={{ fontFamily: "Helvetica Neue" }}
                    required
                  />
                  Employer
                </label>
                &nbsp;&nbsp;&nbsp;
                <label>
                  <input
                    type="radio"
                    name="usertype"
                    value="0"
                    checked={usertype === "0"}
                    onChange={(e) => {
                      setUsertype(e.target.value);
                    }}
                    style={{ fontFamily: "Helvetica Neue" }}
                    required
                  />
                  Employee
                </label>
                <br />
              </div>

              <div className="text-center my-2">
                <button
                 
                  id="click"
                  type="submit"
                  style={{ border: "none", fontFamily: "Helvetica Neue" }}
                >
                  Login
                </button>
              </div>
              <div className="text-center my-2">
                New user?{" "}
                <Link to="/signup" style={{ fontFamily: "Helvetica Neue" }}>
                  {" "}
                  Sign Up
                </Link>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;
