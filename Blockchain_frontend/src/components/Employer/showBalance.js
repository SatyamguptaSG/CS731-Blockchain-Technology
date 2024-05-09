import React, { useState } from 'react';
import axios from 'axios';
import toast from 'react-hot-toast';

const ShowBalance = ({logeduserid}) => {
    const [showBalance, setShowBalance] = useState(false);
    const [balance, setBalance] = useState(null);

           const fetchBalance = async () => {
            try {
                const response = await axios.get('http://localhost:3002/query', {
                    params: {
                    'channelid': 'mychannel',
                    'chaincodeid': 'basictest',
                    'function': 'GetBankDetails',
                    'args': logeduserid
                    }
                    });

                console.log("hi response: ", response.data);
                const substring = response.data.substring(10);
                const jsonData = JSON.parse(substring);

                //second Api
                // const response1 = await axios.get('http://localhost:3002/query', {
                //     params: {
                //         'channelid': 'mychannel',
                //         'chaincodeid': 'paytest',
                //         'function': 'ViewBalance',
                //         'args': [jsonData[0].BankName, jsonData[0].BankAccountNumber]
                //     }
                // });

                const params = new URLSearchParams();
                params.append('channelid', 'mychannel');
                params.append('chaincodeid', 'paytest');
                params.append('function', 'ViewBalance');
                params.append('args', jsonData[0].BankName);
                params.append('args', jsonData[0].BankAccountNumber);

                const response1 = await axios.get('http://localhost:3002/query', {
                params: params
                    });
                console.log("hi response22: ", response1.data);
                const substring1 = response1.data.substring(10);
                const jsonData1 = JSON.parse(substring1);

                // console.log("json ", jsonData1);
                setBalance(jsonData1);
                // setBankName(jsonData[0].BankName);
                // setAccoutNumber(jsonData[0].BankAccountNumber);
                // setBalance(response.data.balance);
            } catch (error) {
               if(!showBalance)toast.error("balance could not be fetched");

                console.error('Failed to fetch balance:', error);
            }
        };

   
    return (
        
        <>
        <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '10vh' }}>
            
            <button onClick={() => { setShowBalance(!showBalance); fetchBalance(); }}>
                {showBalance ? 'Close' : 'Show Balance'}
            </button>
            
        </div>
        <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '2vh' }}>
            {showBalance && balance !== null && <p>Your balance is: {balance}</p>}
        </div>
        </>
    );
};

export default ShowBalance;
