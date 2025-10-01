import React, { useState } from "react";

const Profile = () => {
  const [username, setUsername] = useState("John Doe");
  const [email, setEmail] = useState("john@example.com");
  const [photo, setPhoto] = useState(null);

  const handlePhotoChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      setPhoto(URL.createObjectURL(file));
    }
  };

  return (
    <div className="flex gap-10 mt-10">
      {/* Left Side Profile */}
      <div className="w-1/3 bg-white p-5 rounded-xl shadow-md">
        <div className="flex flex-col items-center">
          <img
            src={photo || "https://via.placeholder.com/150/000000/FFFFFF/?text=No+Photo"}
            alt="Profile"
            className="w-32 h-32 rounded-full object-cover border"
          />
          <input type="file" onChange={handlePhotoChange} className="mt-3 text-sm" />
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="mt-3 border rounded px-3 py-2 w-full"
          />
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="mt-3 border rounded px-3 py-2 w-full"
          />
          <button className="mt-4 bg-blue-500 text-white px-4 py-2 rounded">
            Save Changes
          </button>
        </div>
      </div>

      {/* Right Side Statistic */}
      <div className="w-2/3 bg-white p-5 rounded-xl shadow-md">
        <h2 className="text-lg font-bold mb-4">Statistics</h2>
        <p>Total To Do: <span className="font-semibold">10</span></p>
        <p>Completed: <span className="font-semibold">6</span></p>
        <p>Pending: <span className="font-semibold">4</span></p>
      </div>
    </div>
  );
};

export default Profile;
