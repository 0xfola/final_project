import React, { useState } from "react";
import QRCode from "qrcode.react";
import axios from "axios";

function App() {
  const [studentData, setStudentData] = useState(null);
  const [qrValue, setQrValue] = useState("");

  const handleScan = async (e) => {
    e.preventDefault();
    const studentAddress = e.target.studentAddress.value;
    const response = await axios.get(
      `http://localhost:3000/student/${studentAddress}`,
    );
    setStudentData(response.data);
    setQrValue(response.data.ipfsHash);
  };

  return (
    <div className="App">
      <form onSubmit={handleScan}>
        <input
          type="text"
          name="studentAddress"
          placeholder="Enter Student Address"
        />
        <button type="submit">Scan QR Code</button>
      </form>
      {studentData && (
        <div>
          <p>IPFS Hash: {studentData.ipfsHash}</p>
          <QRCode value={qrValue} />
        </div>
      )}
    </div>
  );
}

export default App;
