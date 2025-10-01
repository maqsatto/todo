import React, { useState, useEffect } from 'react'
import { NavLink, Link, useNavigate } from "react-router-dom"
import { menuLinks, assets } from "../assets/assets.js"

const Navbar = () => {
  const [visible, setVisible] = useState(false)
  const [token, setToken] = useState(null)
  const navigate = useNavigate()

  // Check localStorage on mount
  useEffect(() => {
    const savedToken = localStorage.getItem("token")
    setToken(savedToken)
  }, [])

  const handleLogout = () => {
    localStorage.removeItem("token")
    setToken(null)
  }

  return (
    <div className='flex justify-between items-center font-meduim py-5'>
      <h1 className='text-5xl font-bold text-orange-400 cursor-pointer'>ToDo.</h1>

      <ul className='hidden sm:flex gap-10 items-center text-[#414141] font-meduim'>
        {menuLinks.map((link) => (
          <NavLink 
            to={link.path}
            key={link.name}
            className="flex flex-col items-center text-[18px] transition-all duration-200 ease-linear hover:text-orange-400"
          >
            {({ isActive }) => (
              <>
                {link.name}
                {isActive && <hr className='border-none w-full h-[1.5px] bg-orange-400' />}
              </>
            )}
          </NavLink>
        ))}
      </ul>

      <div className='flex items-center gap-6'>
        {token ? (
          <>
            <img src={assets.search_icon} alt="search_icon" className='w-5 cursor-pointer'/>
            <div className='group relative'>
              <img src={assets.profile_icon} alt="profile_icon" className='w-5 cursor-pointer'/>
              <div className='group-hover:flex hidden absolute dropdown-menu right-0 pt-4'>
                <div className='flex bg-slate-100 flex-col w-36 py-3 px-5 gap-2 text-gray-500 rounded'>
                  <p className='cursor-pointer hover:text-black' onClick={() => navigate("/profile")}>My Profile</p>
                  <p className='cursor-pointer hover:text-black' onClick={handleLogout}>Logout</p>
                </div>
              </div>
            </div>
          </>
        ) : (
          <Link 
            to="/login" 
            className='font-[500] text-[20px] bg-orange-400 rounded py-[5px] text-white px-[25px] transition-all duration-200 ease-linear hover:bg-orange-500'
          >
            LogIn
          </Link>
        )}

        <img 
          src={assets.menu_icon} 
          alt="menu_icon" 
          className='flex sm:hidden w-5 cursor-pointer'
          onClick={() => setVisible(true)}
        />
      </div>

      {/* Menu For Mobile */}
      <div className={`absolute top-0 right-0 bottom-0 overflow-hidden bg-white transition-all duration-200 ease-linear ${visible ? "w-full" : "w-0"}`}>
        <div className='flex flex-col text-gray-600'>
          <div className='p-3 cursor-pointer flex gap-4 items-center justify-end mb-18'>
            <p className='text-2xl'>Back</p>
            <img src={assets.cross_icon} alt="X" className='h-6 rotate-180' onClick={() => setVisible(false)}/>
          </div>

          <ul className='flex flex-col justify-center items-center'>
            {menuLinks.map((link) => (
              <NavLink 
                to={link.path}
                key={link.name}
                className="font-[600] text-[#414141] text-[25px] py-2 transition-all duration-200 ease-linear hover:text-[#0F0E0E]"
              >
                {link.name} 
              </NavLink>
            ))}
          </ul>
        </div>
      </div>
    </div>
  )
}

export default Navbar
