import React, { useEffect, useState } from 'react';
import axios from "axios";
import { toast } from 'react-toastify';
import { assets } from "../assets/assets.js";

const Hero = () => {
  const [title, setTitle] = useState("");
  const [todos, setTodos] = useState([]);

  const token = localStorage.getItem("token");

  useEffect(() => {
    fetchTodos();
  }, []);

  // Получение списка todos
  const fetchTodos = async () => {
    try {
      const res = await axios.get("http://26.1.224.212:8080/todos/", {
        headers: { Authorization: `Bearer ${token}` }
      });
      setTodos(res.data.todos || []);
    } catch (err) {
      toast.error("❌ Failed to fetch todos!");
    }
  };

  // Добавить todo
  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!title.trim()) return;

    try {
      await axios.post(
        "http://26.1.224.212:8080/todos/",
        { title },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      toast.success("✅ To Do Added Successfully!");
      setTitle("");
      fetchTodos();
    } catch (err) {
      const errorMessage = err.response?.data?.error || err.message || "Failed to add todo";
      toast.error(`❌ ${errorMessage}`);
    }
  };

  // Обновить статус completed
  const toggleTodo = async (id, completed) => {
    // оптимистично обновляем
    const updated = todos.map(t => t.id === id ? { ...t, completed: !completed } : t);
    setTodos(updated);

    try {
      await axios.put(
        `http://26.1.224.212:8080/todos/${id}`,
        { completed: !completed },
        { headers: { Authorization: `Bearer ${token}` } }
      );
    } catch (err) {
      toast.error("❌ Failed to update todo!");
      fetchTodos(); // откатим с бэка
    }
  };

  // Удалить todo
  const deleteTodo = async (id) => {
    try {
      await axios.delete(`http://26.1.224.212:8080/todos/${id}`, {
        headers: { Authorization: `Bearer ${token}` }
      });
      toast.success("🗑️ Todo deleted!");
      setTodos(todos.filter(t => t.id !== id));
    } catch (err) {
      toast.error("❌ Failed to delete todo!");
    }
  };

  return (
    <div className='shadow-[0px_5px_15px_rgba(0,0,0,0.35)] rounded-2xl flex flex-col gap-9 py-7 px-10 max-h-[450px] overflow-auto'>
      {/* Форма */}
      <div className='flex flex-col items-center'>
        <form onSubmit={handleSubmit} className='flex items-center justify-center gap-10'>
          <input
            type="text"
            placeholder='Add To Do'
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className='border rounded outline-none px-4 w-96 py-2'
          />
          <button
            type='submit'
            className='border border-amber-600 px-[20px] bg-orange-400 rounded text-white py-2 w-36 m-auto text-[17px] font-medium transition-all duration-200 hover:bg-transparent hover:text-orange-400'
          >
            Add
          </button>
        </form>
      </div>

      {/* Список todos */}
      <div>
        {todos.map((t) => (
          <div key={t.id} className='flex items-center justify-between'>
            <div className='flex gap-3'>
              <input
                type="checkbox"
                checked={t.completed}
                onChange={() => toggleTodo(t.id, t.completed)}
                className='mb-5'
              />
              <p className={`text-xl font-semibold mb-5 ${t.completed ? "line-through italic text-gray-400" : "text-gray-700"}`}>
                {t.title}
              </p>
            </div>
            <div className='ml-10 flex gap-2'>
              <img
                className='w-5 cursor-pointer'
                src={assets.exchange_icon}
                alt="update"
                onClick={() => toggleTodo(t.id, t.completed)}
              />
              <img
                className='w-5 cursor-pointer'
                src={assets.bin_icon}
                alt="delete"
                onClick={() => deleteTodo(t.id)}
              />
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Hero;
