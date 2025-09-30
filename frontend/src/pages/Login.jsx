import React, { useState } from 'react';
import axios from 'axios';

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await axios.post("http://26.1.224.212:8080/users/login", { email, password });
      console.log("Response:", res.data);
      alert("Login success: " + JSON.stringify(res.data));

      // сохраняем JWT в localStorage
      if(res.data.token){
        localStorage.setItem("token", res.data.token);
      }
    } catch (err) {
      alert("Login failed: " + err.response?.data?.error || err.message);
    }
  };

  return (
    <form onSubmit={handleSubmit} className='flex flex-col gap-3'>
      <input type="email" placeholder="Email" onChange={(e) => setEmail(e.target.value)} required />
      <input type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)} required />
      <button type="submit">Login</button>
    </form>
  );
};

export default Login;
