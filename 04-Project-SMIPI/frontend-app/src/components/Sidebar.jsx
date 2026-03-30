import React from 'react';
import { Link, useLocation } from 'react-router-dom';

function Sidebar() {
  const location = useLocation();
  
  // Ambil role dari localStorage
  const role = localStorage.getItem('role') || '';

  // Fungsi untuk mengecek apakah menu sedang aktif
  const isActive = (path) => location.pathname === path;

  // Styling dinamis untuk menu
  const menuStyle = (path) => ({
    display: 'block',
    padding: '12px 20px',
    margin: '8px 15px',
    borderRadius: '8px',
    textDecoration: 'none',
    color: isActive(path) ? '#fff' : '#495057',
    backgroundColor: isActive(path) ? '#0056b3' : 'transparent',
    fontWeight: 'bold',
    transition: 'all 0.2s'
  });

  return (
    <div style={{ width: '260px', backgroundColor: '#fff', borderRight: '1px solid #e9ecef', display: 'flex', flexDirection: 'column' }}>
      
      {/* LOGO SMIPI */}
      <div style={{ padding: '25px 20px', display: 'flex', alignItems: 'center', borderBottom: '1px solid #e9ecef' }}>
        <img src="/smipi_shield.png" alt="Logo" style={{ height: '70px', marginRight: '12px' }} />
        <div>
          <h2 style={{ margin: 0, color: '#0056b3', fontSize: '30px', fontWeight: '900', letterSpacing: '1px' }}>SMIPI</h2>
          <p style={{ margin: 0, fontSize: '12  px', color: '#6c757d', fontWeight: 'bold' }}>PORTAL INVESTIGASI</p>
        </div>
      </div>

      {/* DAFTAR MENU */}
      <div style={{ flex: 1, paddingTop: '20px' }}>
        <p style={{ margin: '0 0 10px 20px', fontSize: '11px', color: '#adb5bd', fontWeight: 'bold', letterSpacing: '1px' }}>MAIN MENU</p>
        
        <Link to="/dashboard" style={menuStyle('/dashboard')}>
          📊 Dashboard
        </Link>
        
        {/* LOGIKA RBAC: Menu ini HANYA dirender jika user adalah Admin */}
        {(role === 'Admin' || role === 'ADMIN') && (
          <Link to="/users" style={menuStyle('/users')}>
            👥 Manajemen Pengguna
          </Link>
        )}
      </div>

    </div>
  );
}

export default Sidebar;