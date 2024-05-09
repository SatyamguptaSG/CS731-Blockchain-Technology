import React, { useEffect, useState } from 'react'
import UserNavbar from './userNavbar.js';
import axios from 'axios';
import "./contracts.css"
import toast from "react-hot-toast";
import ShowBalance from '../Employer/showBalance.js';
const UserHome = () => {
  
  const  logeduserid = window.localStorage.getItem("userId");
  const [createdContracts, setCreatedContracts] = useState([
    // {
    //   ContractID: "contractId",
    //   EmployerId: "employerId",
    //   EmployeeId: "employeeId",
    //   Payment: 1000,
    //   Duration: 12,
    //   BankName_Employer: "employerBankName",
    //   BankAccountNumber_Employer: "employerBankAccountNumber",
    //   BankName_Employee: "employeeBankName",
    //   BankAccountNumber_Employee: "employeeBankAccountNumber",
    //   STATUS: "active",
    //   Last_Payment_Date: "2022-01-01",
    //   All_Transactions: []
    // }
  ]);

  const handleRevoke = async (contract_id) => {
    try {
      const response = await axios.post(
        'http://localhost:3002/invoke',
        new URLSearchParams([
          ['', ''],
          ['channelid', 'mychannel'],
          ['chaincodeid', 'basictest'],
          ['function', 'RevokeContract'],
          ['args', contract_id],
        ])
      );
      console.log('Revoke successful:', response.data);
      toast.success("Revoke succesful");
    } catch (error) {
      console.error('Error making revoke', error);
      toast.error("Revoke failed");
    }
  };

  const handleAccept = async (contract_id) => {
    try {
      const response = await axios.post(
        'http://localhost:3002/invoke',
        new URLSearchParams([
          ['', ''],
          ['channelid', 'mychannel'],
          ['chaincodeid', 'basictest'],
          ['function', 'AcceptContract'],
          ['args', contract_id],
        ])
      );
      toast.success("Accept succesful");
      console.log('Accept successful:', response.data);
    } catch (error) {
      toast.error("Accept failed");
      console.error('Error making accept', error);
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        
        const response = await axios.get('http://localhost:3002/query', {
          params: {
          'channelid': 'mychannel',
          'chaincodeid': 'basictest',
          'function': 'FetchAllContractsbyUserID',
          'args': logeduserid
          }
        });
        
        console.log(response);
        
        const substring = response.data.substring(10);
        const jsonData = JSON.parse(substring);
        setCreatedContracts(jsonData);
      } catch (error) {
        console.error('Error fetching contracts:', error);
      }
    };

    fetchData();
  }, []);



  return (
    <div>
      <UserNavbar />
      <ShowBalance logeduserid={logeduserid} />

      <h1 style={{textAlign: 'center'}}>Your Contracts</h1>
      {
      
      
      
      createdContracts.map((contract,index) => (
          <div key={index} className="appointment-box">
            <span>
              <strong>ContractId</strong> {contract.ContractID}&nbsp;&nbsp;
              <strong>EmployerId</strong> {contract.EmployerId}&nbsp;&nbsp;
              <strong>Payment</strong> {contract.Payment}&nbsp;&nbsp;
              <strong>Duration</strong> {contract.Duration}&nbsp;&nbsp;
              <strong>STATUS</strong> {contract.STATUS}&nbsp;&nbsp;

            </span>&nbsp;&nbsp;&nbsp;
            <button
                className="view-prescription-button"
                id="view-prescription"  
                onClick={()=>handleAccept(contract.ContractID)}
                >
                Accept
            </button>
            <button
                className="view-prescription-button"
                id="view-prescription"  
                
                onClick={()=>handleRevoke(contract.ContractID)}
            >
                Revoke
            </button>
          </div>
          
        ))}
       
    </div>
  );
};

export default UserHome

