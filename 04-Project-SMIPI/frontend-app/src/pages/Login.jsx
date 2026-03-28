import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  
  const [generatedCaptcha, setGeneratedCaptcha] = useState('');
  const [userCaptchaInput, setUserCaptchaInput] = useState('');
  
  const navigate = useNavigate();

  const generateCaptcha = () => {
    const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    let captcha = "";
    for (let i = 0; i < 6; i++) {
      captcha += chars[Math.floor(Math.random() * chars.length)];
    }
    setGeneratedCaptcha(captcha);
  };

  useEffect(() => {
    generateCaptcha();
  }, []);

  const handleLogin = async (e) => {
    e.preventDefault();
    
    if (userCaptchaInput !== generatedCaptcha) {
      alert("⚠️ Karakter unik (CAPTCHA) tidak valid. Silakan coba lagi.");
      generateCaptcha(); 
      setUserCaptchaInput(''); 
      return;
    }
    
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
    <div style={{ 
      display: 'flex', 
      justifyContent: 'center', 
      alignItems: 'center', 
      height: '100vh', 
      backgroundImage: 'url(/smipi_logo_dimmed.png)', 
      backgroundRepeat: 'repeat',
      backgroundSize: '100px',
      backgroundColor: '#e9ecef', 
      fontFamily: '"Segoe UI", Roboto, Helvetica, Arial, sans-serif' 
    }}>
      
      {/* KARTU FORMULIR DENGAN PADDING LEBIH BESAR */}
      <div style={{ 
        backgroundColor: '#ffffff', 
        padding: '40px 50px', 
        borderRadius: '12px', 
        boxShadow: '0 10px 25px rgba(0,0,0,0.1)', 
        width: '100%', 
        maxWidth: '450px' 
      }}>
        
        {/* HEADER INSTITUSI YANG LEBIH BESAR (RATA TENGAH) */}
        <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', marginBottom: '20px' }}>
          <img src="/smipi_shield.png" alt="SMIPI Shield" style={{ width: '90px', marginRight: '20px' }} />
          <div style={{ width: '2px', height: '65px', backgroundColor: '#dc3545', marginRight: '20px' }}></div>
          <div style={{ display: 'flex', flexDirection: 'column', textAlign: 'left' }}>
            <h1 style={{ margin: '0', color: '#0056b3', fontSize: '38px', fontWeight: '700', letterSpacing: '2px', fontFamily: '"Oswald", sans-serif' }}>SMIPI</h1>
            <p style={{ margin: '4px 0 0 0', color: '#ffc107', fontSize: '14px', fontWeight: 'bold', letterSpacing: '0.8px' }}>
              SISTEM MANAJEMEN INFORMASI<br/>PELAPORAN INVESTIGASI
            </p>
          </div>
        </div>

        <h2 style={{ margin: '0 0 25px', color: '#333', fontSize: '18px', fontWeight: 'bold', textAlign: 'center', letterSpacing: '1px' }}>
          PORTAL LOGIN
        </h2>
        
        <form onSubmit={handleLogin}>
          
          <div style={{ marginBottom: '18px' }}>
            <label style={{ display: 'block', fontSize: '13px', fontWeight: '700', color: '#495057', marginBottom: '8px' }}>
              Email / NRP
            </label>
            <div style={{ position: 'relative' }}>
              <span style={{ position: 'absolute', left: '14px', top: '13px', color: '#6c757d', fontSize: '16px' }}>👤</span>
              <input 
                type="email" 
                value={email} 
                onChange={(e) => setEmail(e.target.value)} 
                required 
                placeholder="Masukkan email atau NRP"
                style={{ 
                  width: '100%', padding: '12px 12px 12px 42px', border: '1px solid #ced4da', 
                  borderRadius: '6px', boxSizing: 'border-box', fontSize: '14px', outline: 'none'
                }} 
              />
            </div>
          </div>
          
          <div style={{ marginBottom: '20px' }}>
            <label style={{ display: 'block', fontSize: '13px', fontWeight: '700', color: '#495057', marginBottom: '8px' }}>
              Password
            </label>
            <div style={{ position: 'relative' }}>
              <span style={{ position: 'absolute', left: '14px', top: '13px', color: '#6c757d', fontSize: '16px' }}>🔑</span>
              <input 
                // 2A. UBAH TIPE INPUT MENJADI DINAMIS
                type={showPassword ? "text" : "password"} 
                value={password} 
                onChange={(e) => setPassword(e.target.value)} 
                required 
                placeholder="Masukkan password"
                style={{ 
                  width: '100%', padding: '12px 40px 12px 42px', border: '1px solid #ced4da', 
                  borderRadius: '6px', boxSizing: 'border-box', fontSize: '14px', outline: 'none'
                }} 
              />
              {/* 2B. UBAH IKON MATA AGAR BISA DIKLIK */}
<span 
                 onClick={() => setShowPassword(!showPassword)} 
                 style={{ 
                   position: 'absolute', 
                   right: '14px', 
                   top: '12px', 
                   color: '#6c757d', 
                   cursor: 'pointer', 
                   display: 'flex', 
                   alignItems: 'center' 
                 }}
                 title={showPassword ? "Sembunyikan Password" : "Tampilkan Password"}
               >
                 {showPassword ? (
                   // Ikon Mata Tercoret (Sembunyikan) - Gaya Profesional
                   <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                     <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"></path>
                     <line x1="1" y1="1" x2="23" y2="23"></line>
                   </svg>
                 ) : (
                   // Ikon Mata Terbuka (Tampilkan) - Gaya Profesional
                   <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                     <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
                     <circle cx="12" cy="12" r="3"></circle>
                   </svg>
                 )}
               </span>
            </div>
          </div>
          
          {/* CAPTCHA: LAYOUT 1 BARIS YANG RAPI */}
          <div style={{ display: 'flex', gap: '10px', alignItems: 'stretch', marginBottom: '30px' }}>
            
            {/* Box Karakter Unik */}
            <div style={{ 
              display: 'flex', alignItems: 'center', justifyContent: 'center',
              backgroundColor: '#fff', border: '1px solid #ced4da', borderRadius: '6px', 
              fontWeight: 'bold', fontSize: '18px', color: '#ffc107', letterSpacing: '4px',
              padding: '0 15px', minWidth: '90px'
            }}>
              {generatedCaptcha}
            </div>

            {/* FUNGSI PENTING: TOMBOL REFRESH TRANSPARAN (HAPUS BACKGROUND BIRU) */}
            <button type="button" onClick={generateCaptcha} title="Refresh CAPTCHA" style={{ 
              border: 'none', 
              cursor: 'pointer', 
              backgroundColor: 'transparent', // Hapus background biru
              fontSize: '20px', 
              marginLeft: '5px',
              color: '#0056b3' // Gunakan warna biru institusi untuk ikon
            }}>
              🔄
            </button>

            {/* Kolom Input (Otomatis mengisi sisa ruang) */}
            <input 
              type="text" 
              value={userCaptchaInput} 
              onChange={(e) => setUserCaptchaInput(e.target.value)} 
              required 
              placeholder="Masukkan karakter"
              style={{ 
                flex: '1', padding: '12px', border: '1px solid #ced4da', 
                borderRadius: '6px', boxSizing: 'border-box', fontSize: '14px', outline: 'none'
              }} 
            />
          </div>
          
          {/* TOMBOL AKSI */}
          <div style={{ display: 'flex', gap: '12px' }}>
            <button type="submit" style={{ flex: '1', padding: '12px', backgroundColor: '#0056b3', color: 'white', border: 'none', borderRadius: '6px', cursor: 'pointer', fontSize: '15px', fontWeight: 'bold' }}>
              Masuk
            </button>
            <button type="button" style={{ flex: '1', padding: '12px', backgroundColor: '#6c757d', color: 'white', border: 'none', borderRadius: '6px', cursor: 'pointer', fontSize: '15px', fontWeight: 'bold' }}>
              Daftar
            </button>
          </div>
        </form>

        <div style={{ textAlign: 'center', marginTop: '25px' }}>
           <span style={{ fontSize: '13px', color: '#0056b3', textDecoration: 'none', cursor: 'pointer', fontWeight: '600' }}>Lupa Password?</span>
           <span style={{ fontSize: '13px', color: '#6c757d', margin: '0 8px' }}>atau</span>
           <span style={{ fontSize: '13px', color: '#dc3545', textDecoration: 'none', cursor: 'pointer', fontWeight: '600' }}>Lihat Panduan</span>
        </div>

      </div>
    </div>
  );
}

export default Login;