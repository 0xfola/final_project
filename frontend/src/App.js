import React, { useState } from "react";
import axios from "axios";
import QRCode from "qrcode.react";

function App() {
  const [studentData, setStudentData] = useState(null);
  const [qrValue, setQrValue] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleScan = async (e) => {
    e.preventDefault();
    const studentAddress = e.target.studentAddress.value.trim();
    if (!studentAddress) {
      setError("Please enter a valid student address.");
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const response = await axios.get(
        `http://localhost:8000/student/${studentAddress}`,
      );
      setStudentData(response.data);
      setQrValue(response.data.ipfsHash);
    } catch (err) {
      setError("Failed to fetch student data. Please try again.");
    } finally {
      setLoading(false);
    }
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
      {loading && <p>Loading...</p>}
      {error && <p style={{ color: "red" }}>{error}</p>}
      {studentData && !error && (
        <div>
          <p>IPFS Hash: {studentData.ipfsHash}</p>
          <QRCode value={qrValue} />
        </div>
      )}
    </div>
  );
}

export default App;
