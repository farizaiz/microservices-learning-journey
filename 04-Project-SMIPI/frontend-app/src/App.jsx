import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

// Import Halaman (Pages)
import Login from './pages/Login';
import Dashboard from './pages/Dashboard';
import ManajemenPengguna from './pages/ManajemenPengguna'; // File baru yang akan kita buat

// Import Komponen Layout
import Layout from './components/Layout'; // File baru yang akan kita buat

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* RUTE PUBLIK (Tanpa Layout/Sidebar) */}
        <Route path="/" element={<Login />} />
        
        {/* RUTE PRIVAT (Dibungkus oleh Layout Sidebar & Header) */}
        <Route element={<Layout />}>
          {/* Semua route di dalam sini otomatis akan memiliki Sidebar di kirinya */}
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/users" element={<ManajemenPengguna />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;