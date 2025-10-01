import React, { useState } from 'react'
import axios from 'axios';
import { toast } from "react-toastify";
import { useNavigate, Link } from "react-router-dom";

const Register = () => {
    const [username, setUsername] = useState("")
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
      e.preventDefault();
      try {
        const res = await axios.post("http://26.1.224.212:8080/users/register", { username ,email, password });
        if(res.data.token){
          localStorage.setItem("token", res.data.token);
          toast.success("✅ Registered successfully!");
          navigate("/");
        }
      } catch (err) {
        const errorMessage = err.message || "Register failed";
        toast.error(`❌ ${errorMessage}`);
      }
    };
  return (
    <div className='h-screen flex justify-center items-center text-center'>
        <form onSubmit={handleSubmit} className='w-100 rounded-2xl flex flex-col gap-3 w-2/4 py-8 p px-13 shadow-[0px_5px_15px_rgba(0,0,0,0.35)]'>
          <h1 className='text-4xl mb-5 font-bold text-orange-400'>Register</h1>
          <input className='border rounded outline-none px-4 py-2' type="Text" placeholder="Username" onChange={(e) => setUsername(e.target.value)} required />
          <input className='border rounded outline-none px-4 py-2' type="email" placeholder="Email" onChange={(e) => setEmail(e.target.value)} required />
          <input className='border rounded outline-none px-4 py-2' type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)} required />
          <button type="submit" className='border px-2 border-orange-400 bg-orange-400 rounded-xl mt-7 text-white py-[5px] w-36 m-auto font-medium transition-all duration-200 hover:bg-transparent hover:text-orange-400'>Register</button>
          <p className='text-sm text-gray-500'>Do you have an account? <Link to="/login" className='text-blue-600 underline'>Click here</Link></p>
          
      </form>
    </div>
  )
}

export default Register;
