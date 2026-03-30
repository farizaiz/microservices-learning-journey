import React from 'react';
import { Outlet, useNavigate } from 'react-router-dom';
import Sidebar from './Sidebar';

function Layout() {
  const navigate = useNavigate();
  const nama = localStorage.getItem('namaLengkap');

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('namaLengkap');
    localStorage.removeItem('role');
    navigate('/');
  };

  return (
    <div style={{ display: 'flex', height: '100vh', backgroundColor: '#f4f7f6', fontFamily: '"Segoe UI", Roboto, Helvetica, Arial, sans-serif' }}>
      
      {/* 1. SIDEBAR STATIS DI KIRI */}
      <Sidebar />
      
      <div style={{ flex: 1, display: 'flex', flexDirection: 'column', overflow: 'hidden' }}>
        
        {/* 2. TOPBAR STATIS DI ATAS (Dipindah dari Dashboard lama) */}
        <header style={{ display: 'flex', justifyContent: 'flex-end', alignItems: 'center', backgroundColor: '#ffffff', padding: '15px 30px', boxShadow: '0 2px 10px rgba(0,0,0,0.05)', zIndex: 10 }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: '20px' }}>
            <div style={{ textAlign: 'right' }}>
              <p style={{ margin: 0, fontSize: '12px', color: '#6c757d' }}>Unit Penindakan:</p>
              <p style={{ margin: 0, fontSize: '14px', fontWeight: 'bold', color: '#333' }}>Penyidik {nama}</p>
            </div>
            <button onClick={handleLogout} style={{ padding: '8px 16px', backgroundColor: '#f8d7da', color: '#dc3545', border: '1px solid #f5c6cb', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold' }}>Keluar</button>
          </div>
        </header>

        {/* 3. AREA KONTEN DINAMIS */}
        <main style={{ flex: 1, overflowY: 'auto' }}>
          {/* Outlet adalah tempat di mana Dashboard atau halaman lain akan di-render */}
          <Outlet /> 
        </main>

      </div>
    </div>
  );
}

export default Layout;