import React, { useEffect, useState } from 'react'
import AdminNavbar from './adminNavbar'
import axios from 'axios';
import "../Employee/contracts.css";
import CreateContract from './createContract';
import toast from "react-hot-toast";
import ShowBalance from './showBalance';
const AdminHome = () => {

  const  logeduserid = window.localStorage.getItem("userId");
  const [createdContracts, setCreatedContracts] = useState([
    // {// {
    // //   ContractID: "contractId",
    // //   EmployerId: "employerId",
    // //   EmployeeId: "employeeId",
    // //   Payment: 1000,
    // //   Duration: 12,
    // //   BankName_Employer: "employerBankName",
    // //   BankAccountNumber_Employer: "employerBankAccountNumber",
    // //   BankName_Employee: "employeeBankName",
    // //   BankAccountNumber_Employee: "employeeBankAccountNumber",
    // //   STATUS: "active",
    // //   Last_Payment_Date: "2022-01-01",
    // //   All_Transactions: []
    // },{
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
  const [flag , setFlag] = useState(true);


  const banckdetails = async () => {
    // try {
      

      const response = await axios.get('http://localhost:3002/query', {
        params: {
        'channelid': 'mychannel',
        'chaincodeid': 'basictest',
        'function': 'GetBankDetails',
        'args': logeduserid
        }
      });
      
      console.log(response);
      console.log("this " , response.data);

      // const substring = response.data.substring(10);
      // const jsonData = JSON.parse(substring);
      
      // console.log(jsonData);
      
    // } catch (error) {
    //   toast.error("Error fetching contracts");
    //   console.error('Error fetching contracts:', error);
    // }
  };
  banckdetails();


  const handlePayment = async (b1,bn1,b2,bn2,pay,status) => {
    if(status === "Pending" || status === "Revoked"){
      toast.error("Contract is inactive");
      return;
    }
    try {
      const response = await axios.post(
        'http://localhost:3002/invoke',
        new URLSearchParams([
          ['', ''],
          ['channelid', 'mychannel'],
          ['chaincodeid', 'paytest'],
          ['function', 'MakePayment'],
          ['args', b1],
          ['args', bn1],
          ['args', b2],
          ['args', bn2],
          ['args', pay]
        ])
      );
      console.log('Payment successful:', response.data);
      toast.success("Payment succesful");
    } catch (error) {
      toast.error("Payment failed");
      console.error('Error making payment:', error);
    }
  };

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
      console.log('Revoke successful:', response);
      toast.success("Revoke succesful");

    } catch (error) {
      toast.error("Revoke failed");
      console.error('Error making revoke', error);
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
        console.log("this" , response.data);

        const substring = response.data.substring(10);
        const jsonData = JSON.parse(substring);
        
        setCreatedContracts(jsonData);


      } catch (error) {
        toast.error("Error fetching contracts");
        console.error('Error fetching contracts:', error);
      }
    };

    fetchData();
    if(createdContracts){
      setFlag(true);
    }
  }, []);

  




    return (
    <div>
    <AdminNavbar/>
    <CreateContract/>
    <hr/>
    <ShowBalance logeduserid={logeduserid} />

    <hr />
    <h1 style={{textAlign: 'center'}}>Your Contracts</h1>
    {flag ? createdContracts.map((contract,index) => (
        <div key={index} className="appointment-box">
        <span>
              <strong>ContractId</strong> {contract.ContractID}&nbsp;&nbsp;
              <strong>EmployeeId</strong> {contract.EmployeeId}&nbsp;&nbsp;
              <strong>Payment</strong> {contract.Payment}&nbsp;&nbsp;
              <strong>Duration</strong> {contract.Duration}&nbsp;&nbsp;
              <strong>STATUS</strong> {contract.STATUS}&nbsp;&nbsp;


        </span>&nbsp;&nbsp;&nbsp;
        <button
          className="view-prescription-button"
          id="view-prescription"  
          onClick={()=>handlePayment(contract.BankName_Employer,contract.BankAccountNumber_Employer,contract.BankName_Employee,contract.BankAccountNumber_Employee,contract.Payment,contract.STATUS)}
          >
          Payment
        </button>
        <button
          className="view-prescription-button"
          id="view-prescription"  
          
          onClick={()=>handleRevoke(contract.ContractID)}
        >
          Revoke
        </button>
        </div>
        
      )) : <p>No entries to display</p>}


    
    </div>
    )
}

export default AdminHome
