import React, { useContext, useState } from "react";
import Navbar from "./navbar";
import { Link, Navigate } from "react-router-dom";
import "./form.css";
import { Context } from "../..";
import toast from "react-hot-toast";
import axios from "axios";

const Signup = () => {
  const [userId, setUserID] = useState("");
  const [bankName, setBankName] = useState("");
  const [currency, setCurrency] = useState("");

  const [bankAccountNumber, setBankAccountNumber] = useState("");
  const [password, setPassword] = useState("");
  const [usertype, setUsertype] = useState();

  
  const { isAuthenticated, setIsAuthenticated, loading, setLoading } =
  useContext(Context);

  const submitHandler = async (e) => {
    e.preventDefault();
    setLoading(true);
    
      try {
        // const { data } = await axios.post(
        //   "http://localhost:3002/invoke?channelid=mychannel&chaincodeid=basictest&function=CreateUser&args=Havi_1&args=abc123&args=SBI&args=123456&args=0",
        //   {
        //     userId,
        //     bankName,
        //     bankAccountNumber,
        //     password,
        //     usertype,
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
            ['function', 'CreateUser'],
            ['args', userId],
            ['args', password],
            ['args', bankName],
            ['args', bankAccountNumber],
            ['args', usertype]
          ])


        );
        
        console.log("hello this is signup", data.data);


        if(bankName==="PNB" || bankName==="SBI"){
          setCurrency("INR");
        }
        else{
          setCurrency("USD");
        }

        const bankAccountCreation = await axios.post(
          'http://localhost:3002/invoke',
          new URLSearchParams([
            ['', ''],
            ['channelid', 'mychannel'],
            ['chaincodeid', 'paytest'],
            ['function', 'CreateAccount'],
            ['args', bankName],
            ['args', currency],
            ['args', bankAccountNumber] 
          ])

        );
        // console.log("hello this is creation bank ", bankAccountCreation.data);
        console.log("hello this is creation bank ", bankAccountCreation.data);
        
        // toast.success(bankAccountCreation.data);

        // const data = await axios.get('http://localhost:3002/query', {
        // params: {
        //   'channelid': 'mychannel',
        //   'chaincodeid': 'basictest',
        //   'function': 'GetAllUsers',
        //   // 'args': 'Asset123'
        // }
        // });
        // console.log(data);
        toast.success("Signup succesful");
        window.localStorage.setItem("userId", userId);

        setIsAuthenticated(true);
        setLoading(false);
      } catch (error) {
        setIsAuthenticated(false);
        toast.error("Signup failed. User already exists")
        setLoading(false);
      }
    
  };

  if (isAuthenticated ) {
    if(usertype === "1")
    return <Navigate to={"/employer/home"} />;
    else
    return <Navigate to={"/employee/home"} />;
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
          Sign Up
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
                  type="string"
                  name="id"
                  value={userId}
                  onChange={(e) => {
                    setUserID(e.target.value);
                  }}
                  placeholder="User ID"
                  style={{ fontFamily: "Helvetica Neue" }}
                  spellcheck="false"
                  required
                />
              </div>
              <div className="text-center my-2">
                &nbsp;&nbsp;
                <select
                    name="bname"
                    value={bankName}
                    onChange={(e) => {
                    setBankName(e.target.value);
                    }}
                    style={{ fontFamily: "Helvetica Neue", width: '210px', height: '25px'}}
                    required
                  >
                    <option value="">Select Bank</option>
                    <option value="PNB">Punjab National Bank</option>
                    <option value="SBI">State Bank of India</option>
                    <option value="BNY">Bank of NewYork</option>
                    <option value="Barclays"> Barclays</option>
                  </select>
              </div>
              <div className="text-center my-2">
                &nbsp;&nbsp;
                <input
                  type="string"
                  name="bacno"
                  value={bankAccountNumber}
                  onChange={(e) => {
                    setBankAccountNumber(e.target.value);
                  }}
                  placeholder="Bank Account Number"
                  style={{ fontFamily: "Helvetica Neue" }}
                  spellcheck="false"
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
                  disabled={loading}
                  id="click"
                  type="submit"
                  style={{ border: "none", fontFamily: "Helvetica Neue" }}
                >
                  Sign Up
                </button>
              </div>
              <div className="text-center my-2">
                Already have an account?{" "}
                <Link to="/login" style={{ fontFamily: "Helvetica Neue" }}>
                  Sign In
                </Link>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Signup;
