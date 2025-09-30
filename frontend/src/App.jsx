import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Home from "./pages/Home.jsx"
import Login from "./pages/Login.jsx"
import Navbar from "./components/Navbar.jsx"
const App = () => {
  return (
    <div className='px-5 sm:px-[5vw] md:px-[7vw] lg:px-[9vw] mx-auto'>
      <Navbar/>
      <Routes>
        <Route path='/' element={<Home/>} />
        <Route path='/login' element={<Login/>} />
      </Routes>
    </div>
  )
}

export default App
