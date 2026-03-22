import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Login from './pages/Login';
import Dashboard from './pages/Dashboard';

function App() {
  return (
    // BrowserRouter adalah pembungkus agar aplikasi tahu URL browser
    <BrowserRouter>
      <Routes>
        {/* Jika URL-nya "/", tampilkan halaman Login */}
        <Route path="/" element={<Login />} />
        
        {/* Jika URL-nya "/dashboard", tampilkan halaman Dasbor */}
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;