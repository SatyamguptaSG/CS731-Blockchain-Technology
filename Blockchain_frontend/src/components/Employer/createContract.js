import React, { useState } from "react";
import axios from "axios";
import toast from "react-hot-toast";
import "../NotLogin/form.css"

const CreateContract = () => {
  const [employeeId, setEmployeeId] = useState("");
  const [payment, setPayment] = useState();
  const [duration, setDuration] = useState("");
  const  logeduserid = window.localStorage.getItem("userId");
  const handleSubmit = async(e) => {
    e.preventDefault();
    // TODO: Add code to handle form submission

    // how I WOULD GET THE EMPLOYER ID????
    try {
      
      const data = await axios.post(
        'http://localhost:3002/invoke',
        new URLSearchParams([
          ['', ''],
          ['channelid', 'mychannel'],
          ['chaincodeid', 'basictest'],
          ['function', 'CreateContract'],
          ['args', logeduserid],
          ['args', employeeId],
          ['args', payment],
          ['args', duration],
        ])
      );

      // ['args', employeeId],
          // ['args', payment],
          // ['args', duration]

      console.log(data);
      setEmployeeId(employeeId);
      setPayment(payment);
      setDuration(duration);
      //  setUser(user);
    
      toast.success("Contract creation successful");
    }  catch (error) {
        toast.error("Contract creation failed");
        // setIsLogined(false);
      }
  };

  return (
    <div>
    <h1 style={{textAlign: 'center'}}>Create Contract</h1>
    <hr />





      
      <div className="form-container">
        <form
          className="card"
          style={{ background: "#eeeeee" }}
          onSubmit={handleSubmit}
        >
          <div className="form my-4" style={{ textAlign: "center" }}>
            <div className="text-center my-2">
              &nbsp;&nbsp;
              <input
                type="id"
                name="id"
                value={employeeId}
                style={{ fontFamily: "Helvetica Neue" }}
                onChange={(e) => {
                    setEmployeeId(e.target.value);
                }}
                placeholder="Employee Id"
                required
              />
              <br />
            </div>
            <div className="text-center my-2">
              &nbsp;&nbsp;
              <input
                type="string"
                name="salary"
                title="pesa do"
                value={payment}
                onChange={(e) => {
                    setPayment(e.target.value);
                }}
                placeholder="Salary(in your currency)"
                style={{ fontFamily: "Helvetica Neue" }}
                spellcheck="false"
                required
              />
              <br />
            </div>
            <div className="text-center my-2">
              &nbsp;&nbsp;
              <input
                name="duration"
                title="time in month"
                value={duration}
                onChange={(e) => {
                    setDuration(e.target.value);
                }}
                placeholder="Duration in months"
                style={{ fontFamily: "Helvetica Neue" }}
                spellcheck="false"
                required
              />
              <br />
            </div>

            <div className="text-center my-2">
              <button
                id="click"
                type="submit"
                style={{ border: "none", fontFamily: "Helvetica Neue" }}
              >
                Create
              </button>
            </div>
            
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateContract;
