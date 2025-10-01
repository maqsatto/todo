  import React, { useState } from 'react';
  import axios from 'axios';
  import { toast } from "react-toastify";
  import { useNavigate, Link } from "react-router-dom";

  const Login = () => {

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
      e.preventDefault();
      try {

        const res = await axios.post("http://26.1.224.212:8080/users/login",{email, password})

        if(res.data.token){
          localStorage.setItem("token", res.data.token);
          toast.success("✅ Loginned successfully!");
          navigate("/");
        } else {
          toast.error("❌ Invalid login");
        }
      } catch (err) {
        const errorMessage = err.message || "Login failed";
        toast.error(`❌ ${errorMessage}`);
      }
    };

    return (
      <div className='h-screen flex justify-center items-center text-center'>
        <form onSubmit={handleSubmit} className='w-100 rounded-2xl flex flex-col gap-3 w-2/4 py-8 p px-13 shadow-[0px_5px_15px_rgba(0,0,0,0.35)]'>
          <h1 className='text-4xl mb-5 font-bold text-orange-400'>Login</h1>
          <input className='border rounded outline-none px-4 py-2' type="email" placeholder="Email" onChange={(e) => setEmail(e.target.value)} required />
          <input className='border rounded outline-none px-4 py-2' type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)} required />
          <button type="submit" className='border border-orange-400 px-2 bg-orange-400 rounded-xl mt-7 text-white py-[5px] w-36 m-auto font-medium transition-all duration-200 hover:bg-transparent hover:text-orange-400'>Login</button>
          <p className='text-sm text-gray-500'>Don't have an account? <Link to="/register" className='text-blue-600 underline'>Click here</Link></p>
          
      </form>
      </div>
    );
  };

  export default Login;
