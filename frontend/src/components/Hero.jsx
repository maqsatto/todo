import React, {useEffect, useState} from 'react'
import axios from "axios"
import { toast } from 'react-toastify'

const Hero = () => {
  const [title, setTitle] = useState("")
  const [todo, setTodo] = useState([])

  useEffect(() => {
    fetchTodos()
  }, [])

  const fetchTodos = async () => {
    try {
        const res = await axios.get("http://26.1.224.212:8080/todos/")
        setTodo(res.data)
    } catch (err) {
        toast.error("❌ Failed to fetch todos!")
    }
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
        const res = await axios.post("http://26.1.224.212:8080/todos/", {title})
        toast.success(`✅ To Do Added Successfuly!`)
        setTitle("")
    } catch (err) {
        const errorMessage = err.response?.data?.error || err.message || "Failed to add todo";
        toast.error(`❌ ${errorMessage}`);
    }
  }
  return (
    <div className='border border-gray-500'>
      <div>
        <form  onSubmit={handleSubmit} className=''>
            <input type="text" placeholder='Add To Do' value={title} onChange={(e) => setTitle(e.target.value)}/>
            <button type='submit'>Add</button>
        </form>
      </div>
        <ul key={todo.title}>
            {
                todo.map((todo, index) => (
                    <li>{todo.title}</li>
                ))
            }
        </ul>
      <div>
      </div>
    </div>
  )
}

export default Hero
