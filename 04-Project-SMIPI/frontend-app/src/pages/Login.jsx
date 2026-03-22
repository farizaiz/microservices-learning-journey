import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8081/login', {
        email: email,
        password: password
      });

      localStorage.setItem('token', response.data.token);
      localStorage.setItem('namaLengkap', response.data.data.nama_lengkap);
      
      navigate('/dashboard');
      
    } catch (error) {
      alert("Login Gagal: " + (error.response?.data?.pesan || "Terjadi kesalahan server"));
    }
  };

  return (
    // Latar Belakang (Layar Penuh)
    <div style={{ 
      display: 'flex', 
      justifyContent: 'center', 
      alignItems: 'center', 
      height: '100vh', 
      backgroundColor: '#e9ecef', // Warna abu-abu kebiruan yang elegan
      fontFamily: '"Segoe UI", Roboto, Helvetica, Arial, sans-serif' 
    }}>
      
      {/* Kartu Formulir (Card) */}
      <div style={{ 
        backgroundColor: '#ffffff', 
        padding: '40px 50px', 
        borderRadius: '12px', 
        boxShadow: '0 8px 20px rgba(0,0,0,0.1)', 
        width: '100%', 
        maxWidth: '400px' 
      }}>
        
        <div style={{ textAlign: 'center', marginBottom: '30px' }}>
          <h1 style={{ margin: '0', color: '#0056b3', fontSize: '28px' }}>SMIPI 🛡️</h1>
          <p style={{ color: '#6c757d', fontSize: '14px', marginTop: '8px' }}>
            Sistem Manajemen Informasi Pelaporan Investigasi
          </p>
        </div>
        
        <form onSubmit={handleLogin}>
          <div style={{ marginBottom: '20px' }}>
            <label style={{ display: 'block', fontSize: '14px', fontWeight: '600', color: '#333', marginBottom: '8px' }}>
              Email Institusi
            </label>
            <input 
              type="email" 
              value={email} 
              onChange={(e) => setEmail(e.target.value)} 
              required 
              placeholder="budi@smipi.go.id"
              style={{ 
                width: '100%', 
                padding: '12px', 
                border: '1px solid #ced4da', 
                borderRadius: '6px', 
                boxSizing: 'border-box',
                fontSize: '15px',
                outline: 'none',
                transition: 'border-color 0.2s'
              }} 
            />
          </div>
          
          <div style={{ marginBottom: '30px' }}>
            <label style={{ display: 'block', fontSize: '14px', fontWeight: '600', color: '#333', marginBottom: '8px' }}>
              Password
            </label>
            <input 
              type="password" 
              value={password} 
              onChange={(e) => setPassword(e.target.value)} 
              required 
              placeholder="••••••••"
              style={{ 
                width: '100%', 
                padding: '12px', 
                border: '1px solid #ced4da', 
                borderRadius: '6px', 
                boxSizing: 'border-box',
                fontSize: '15px',
                outline: 'none'
              }} 
            />
          </div>
          
          <button 
            type="submit" 
            style={{ 
              width: '100%', 
              padding: '14px', 
              backgroundColor: '#0056b3', 
              color: 'white', 
              border: 'none', 
              borderRadius: '6px', 
              cursor: 'pointer',
              fontSize: '16px',
              fontWeight: 'bold',
              boxShadow: '0 4px 6px rgba(0, 86, 179, 0.2)',
              transition: 'background-color 0.3s'
            }}
          >
            Masuk ke Sistem Keamanan
          </button>
        </form>

        <div style={{ textAlign: 'center', marginTop: '20px' }}>
           <span style={{ fontSize: '12px', color: '#adb5bd' }}>© 2026 Divisi Teknologi Investigasi</span>
        </div>

      </div>
    </div>
  );
}

export default Login;