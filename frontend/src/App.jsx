import React from 'react'
import { Route, Routes, useLocation } from 'react-router-dom'
import Home from "./pages/Home.jsx"
import Login from "./pages/Login.jsx"
import Navbar from "./components/Navbar.jsx"
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Register from './pages/Register.jsx'
import Profile from './pages/Profile.jsx'
const App = () => { 

  const location = useLocation();

  const hideNavbarRoutes = ["/login", "/register", "/profile"]

  const shouldShowNavbar = !hideNavbarRoutes.includes(location.pathname)

  return (
    <div className='px-5 sm:px-[5vw] md:px-[7vw] lg:px-[9vw] mx-auto'>
      <ToastContainer/>
      {shouldShowNavbar && <Navbar/>}
      <Routes>
        <Route path='/' element={<Home/>} />
        <Route path='/login' element={<Login/>} />
        <Route path='/register' element={<Register/>} />
        <Route path='/profile' element={<Profile/>} />
      </Routes>
    </div>
  )
}

export default App
