import React, { useState, useEffect } from 'react';
import axios from 'axios';

function ManajemenPengguna() {
  const [daftarUser, setDaftarUser] = useState([]);
  
  // State Modal Tambah User
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [formData, setFormData] = useState({
    nama_lengkap: '', email: '', password: '', nomor_telepon: '', nomor_induk: '', role: 'Investigator'
  });

  // State Modal Edit Role
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);
  const [newRole, setNewRole] = useState('');

  // Ambil data user saat halaman pertama kali dimuat
  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await axios.get('http://localhost:8081/api/users', {
        headers: { Authorization: `Bearer ${token}` }
      });
      setDaftarUser(response.data.data || []);
    } catch (error) {
      console.error("Gagal mengambil data pengguna:", error);
    }
  };

  // --- FUNGSI TAMBAH USER ---
  const handleInputChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const token = localStorage.getItem('token');
      await axios.post('http://localhost:8081/api/register', formData, {
        headers: { Authorization: `Bearer ${token}` }
      });

      alert("Berhasil menambahkan pengguna baru! ✅");
      setIsModalOpen(false);
      setFormData({ nama_lengkap: '', email: '', password: '', nomor_telepon: '', nomor_induk: '', role: 'Investigator' });
      fetchUsers(); 
    } catch (error) {
      console.error("Gagal menambah user:", error);
      alert(error.response?.data?.pesan || "Gagal menyimpan data. Email/Nomor Induk mungkin sudah terdaftar.");
    }
  };

  // --- FUNGSI EDIT ROLE ---
  const klikEditRole = (user) => {
    setSelectedUser(user);
    setNewRole(user.role);
    setIsEditModalOpen(true);
  };

  const handleEditRoleSubmit = async (e) => {
    e.preventDefault();
    try {
      const token = localStorage.getItem('token');
      await axios.put(`http://localhost:8081/api/users/${selectedUser.id}/role`, 
        { role: newRole }, 
        { headers: { Authorization: `Bearer ${token}` } }
      );

      alert(`Akses ${selectedUser.nama_lengkap} berhasil diubah menjadi ${newRole}! 🔄`);
      setIsEditModalOpen(false);
      fetchUsers(); 
    } catch (error) {
      console.error("Gagal update role:", error);
      alert("Terjadi kesalahan saat mengubah role.");
    }
  };

  return (
    <div style={{ padding: '30px', maxWidth: '1200px', margin: '0 auto' }}>
      
      {/* Header Halaman */}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '30px' }}>
        <div>
          <h2 style={{ margin: 0, color: '#333', fontSize: '24px' }}>Manajemen Pengguna</h2>
          <p style={{ margin: '5px 0 0', color: '#6c757d', fontSize: '14px' }}>Kelola hak akses dan akun penyidik SMIPI</p>
        </div>
        <button 
          onClick={() => setIsModalOpen(true)} 
          style={{ padding: '10px 20px', backgroundColor: '#28a745', color: 'white', border: 'none', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold', fontSize: '14px', boxShadow: '0 2px 6px rgba(40, 167, 69, 0.2)' }}
        >
          + Tambah Pengguna
        </button>
      </div>

      {/* Tabel Daftar Pengguna */}
      <div style={{ backgroundColor: '#ffffff', padding: '25px', borderRadius: '12px', boxShadow: '0 4px 15px rgba(0,0,0,0.05)', overflowX: 'auto' }}>
        <table style={{ width: '100%', borderCollapse: 'collapse', fontSize: '14px', textAlign: 'left' }}>
          <thead>
            <tr style={{ backgroundColor: '#f8f9fa', borderBottom: '2px solid #dee2e6' }}>
              <th style={{ padding: '15px', color: '#495057', fontWeight: '700' }}>Nama & Induk (NRP)</th>
              <th style={{ padding: '15px', color: '#495057', fontWeight: '700' }}>Email Kontak</th>
              <th style={{ padding: '15px', color: '#495057', fontWeight: '700' }}>Role Sistem</th>
              <th style={{ padding: '15px', color: '#495057', fontWeight: '700', textAlign: 'center' }}>Aksi</th>
            </tr>
          </thead>
          <tbody>
            {daftarUser.length > 0 ? (
              daftarUser.map((user) => (
                <tr key={user.id} style={{ borderBottom: '1px solid #e9ecef' }}>
                  <td style={{ padding: '15px' }}>
                    <div style={{ fontWeight: 'bold', color: '#0056b3' }}>{user.nama_lengkap}</div>
                    <div style={{ fontSize: '12px', color: '#6c757d', marginTop: '4px' }}>NRP: {user.nomor_induk}</div>
                  </td>
                  <td style={{ padding: '15px', color: '#495057' }}>{user.email}</td>
                  <td style={{ padding: '15px' }}>
                    <span style={{ 
                      padding: '5px 10px', borderRadius: '4px', fontSize: '12px', fontWeight: 'bold',
                      backgroundColor: user.role === 'Admin' ? '#ffe3e3' : (user.role === 'Investigator' ? '#e7f5ff' : '#f1f3f5'),
                      color: user.role === 'Admin' ? '#c92a2a' : (user.role === 'Investigator' ? '#0056b3' : '#495057'),
                      border: '1px solid rgba(0,0,0,0.1)'
                    }}>
                      {user.role}
                    </span>
                  </td>
                  <td style={{ padding: '15px', textAlign: 'center' }}>
                    <button 
                         onClick={() => klikEditRole(user)} 
                         style={{ padding: '6px 12px', backgroundColor: '#e9ecef', color: '#495057', border: 'none', borderRadius: '4px', cursor: 'pointer', fontSize: '12px', fontWeight: 'bold' }}
                    >
                       ⚙️ Edit Role
                    </button>
                  </td>
                </tr>
              ))
            ) : (
              <tr>
                <td colSpan="4" style={{ padding: '40px', textAlign: 'center', color: '#868e96' }}>
                  Belum ada data pengguna yang bisa ditampilkan.
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>

      {/* ========================================================= */}
      {/* 1. MODAL TAMBAH USER */}
      {/* ========================================================= */}
      {isModalOpen && (
        <div style={{ position: 'fixed', top: 0, left: 0, right: 0, bottom: 0, backgroundColor: 'rgba(0,0,0,0.5)', zIndex: 999, display: 'flex', justifyContent: 'center', alignItems: 'center', backdropFilter: 'blur(5px)' }}>
          <div style={{ backgroundColor: '#ffffff', padding: '35px 40px', borderRadius: '16px', width: '100%', maxWidth: '550px', boxShadow: '0 20px 40px rgba(0,0,0,0.15)' }}>
            
            {/* Modal Header */}
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '25px', paddingBottom: '15px', borderBottom: '1px solid #f1f3f5' }}>
              <h3 style={{ margin: 0, color: '#2b8a3e', fontSize: '20px', fontWeight: '800' }}>👤 Tambah Pengguna Baru</h3>
              <button onClick={() => setIsModalOpen(false)} style={{ background: '#f8f9fa', border: 'none', fontSize: '20px', cursor: 'pointer', color: '#adb5bd', width: '32px', height: '32px', borderRadius: '50%', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>&times;</button>
            </div>
            
            {/* Modal Body (Form) */}
            <form onSubmit={handleSubmit}>
              <div style={{ marginBottom: '18px' }}>
                <label style={{ display: 'block', fontSize: '13px', fontWeight: '700', color: '#343a40', marginBottom: '8px' }}>Nama Lengkap</label>
                <input type="text" name="nama_lengkap" required value={formData.nama_lengkap} onChange={handleInputChange} style={{ width: '100%', padding: '12px 15px', border: '1px solid #dee2e6', borderRadius: '8px', boxSizing: 'border-box', outline: 'none', fontSize: '14px', backgroundColor: '#f8f9fa' }} />
              </div>
              
              <div style={{ display: 'flex', gap: '20px', marginBottom: '18px' }}>
                <div style={{ flex: '1' }}>
                  <label style={{ display: 'block', fontSize: '13px', fontWeight: '700', color: '#343a40', marginBottom: '8px' }}>Email Instansi</label>
                  <input type="email" name="email" required value={formData.email} onChange={handleInputChange} style={{ width: '100%', padding: '12px 15px', border: '1px solid #dee2e6', borderRadius: '8px', boxSizing: 'border-box', outline: 'none', fontSize: '14px', backgroundColor: '#f8f9fa' }} />
                </div>
                <div style={{ flex: '1' }}>
                  <label style={{ display: 'block', fontSize: '13px', fontWeight: '700', color: '#343a40', marginBottom: '8px' }}>Password</label>
                  <input type="password" name="password" required value={formData.password} onChange={handleInputChange} style={{ width: '100%', padding: '12px 15px', border: '1px solid #dee2e6', borderRadius: '8px', boxSizing: 'border-box', outline: 'none', fontSize: '14px', backgroundColor: '#f8f9fa' }} />
                </div>
              </div>

              <div style={{ display: 'flex', gap: '20px', marginBottom: '25px' }}>
                <div style={{ flex: '1' }}>
                  <label style={{ display: 'block', fontSize: '13px', fontWeight: '700', color: '#343a40', marginBottom: '8px' }}>Nomor Induk / NRP</label>
                  <input type="text" name="nomor_induk" required value={formData.nomor_induk} onChange={handleInputChange} style={{ width: '100%', padding: '12px 15px', border: '1px solid #dee2e6', borderRadius: '8px', boxSizing: 'border-box', outline: 'none', fontSize: '14px', backgroundColor: '#f8f9fa' }} />
                </div>
                <div style={{ flex: '1' }}>
                  <label style={{ display: 'block', fontSize: '13px', fontWeight: '700', color: '#343a40', marginBottom: '8px' }}>Nomor Telepon</label>
                  <input type="text" name="nomor_telepon" required value={formData.nomor_telepon} onChange={handleInputChange} style={{ width: '100%', padding: '12px 15px', border: '1px solid #dee2e6', borderRadius: '8px', boxSizing: 'border-box', outline: 'none', fontSize: '14px', backgroundColor: '#f8f9fa' }} />
                </div>
              </div>

              <div style={{ marginBottom: '35px' }}>
                <label style={{ display: 'block', fontSize: '13px', fontWeight: '700', color: '#0056b3', marginBottom: '8px' }}>Role / Hak Akses</label>
                {/* Desain Select Box yang lebih elegan dan netral */}
                <select name="role" value={formData.role} onChange={handleInputChange} style={{ width: '100%', padding: '12px 15px', border: '1px solid #ced4da', borderRadius: '8px', boxSizing: 'border-box', outline: 'none', backgroundColor: '#fff', fontSize: '14px', fontWeight: '600', color: '#495057', cursor: 'pointer' }}>
                  <option value="Admin">Admin (Akses Penuh)</option>
                  <option value="Investigator">Investigator (Bisa Edit Kasus)</option>
                  <option value="Viewer">Viewer (Hanya Melihat)</option>
                </select>
              </div>

              {/* Modal Footer (Buttons) */}
              <div style={{ display: 'flex', gap: '12px', justifyContent: 'flex-end', borderTop: '1px solid #f1f3f5', paddingTop: '20px' }}>
                <button type="button" onClick={() => setIsModalOpen(false)} style={{ padding: '12px 24px', backgroundColor: '#f1f3f5', color: '#495057', border: 'none', borderRadius: '8px', cursor: 'pointer', fontWeight: '700', fontSize: '14px' }}>Batalkan</button>
                <button type="submit" style={{ padding: '12px 24px', backgroundColor: '#28a745', color: 'white', border: 'none', borderRadius: '8px', cursor: 'pointer', fontWeight: '700', fontSize: '14px', boxShadow: '0 4px 10px rgba(40, 167, 69, 0.2)' }}>Simpan Pengguna</button>
              </div>
            </form>

          </div>
        </div>
      )}

      {/* ========================================================= */}
      {/* 2. MODAL EDIT ROLE */}
      {/* ========================================================= */}
      {isEditModalOpen && selectedUser && (
        <div style={{ position: 'fixed', top: 0, left: 0, right: 0, bottom: 0, backgroundColor: 'rgba(0,0,0,0.6)', zIndex: 999, display: 'flex', justifyContent: 'center', alignItems: 'center', backdropFilter: 'blur(3px)' }}>
          <div style={{ backgroundColor: '#ffffff', padding: '30px', borderRadius: '12px', width: '100%', maxWidth: '400px', boxShadow: '0 10px 30px rgba(0,0,0,0.2)' }}>
            <h3 style={{ margin: '0 0 15px', color: '#333', fontSize: '18px' }}>⚙️ Ubah Hak Akses</h3>
            <p style={{ margin: '0 0 20px', fontSize: '14px', color: '#6c757d' }}>
              Pilih role baru untuk <strong>{selectedUser.nama_lengkap}</strong>:
            </p>
            
            <form onSubmit={handleEditRoleSubmit}>
              <select 
                value={newRole} 
                onChange={(e) => setNewRole(e.target.value)} 
                style={{ width: '100%', padding: '12px', border: '2px solid #cce5ff', borderRadius: '6px', outline: 'none', backgroundColor: '#e7f5ff', fontWeight: 'bold', color: '#0056b3', marginBottom: '20px' }}
              >
                <option value="Admin">Admin (Akses Penuh)</option>
                <option value="Investigator">Investigator (Bisa Edit Kasus)</option>
                <option value="Viewer">Viewer (Hanya Melihat)</option>
              </select>

              <div style={{ display: 'flex', gap: '10px', justifyContent: 'flex-end' }}>
                <button type="button" onClick={() => setIsEditModalOpen(false)} style={{ padding: '10px 15px', backgroundColor: '#e9ecef', color: '#495057', border: 'none', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold' }}>Batal</button>
                <button type="submit" style={{ padding: '10px 15px', backgroundColor: '#0056b3', color: 'white', border: 'none', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold' }}>Simpan Perubahan</button>
              </div>
            </form>
          </div>
        </div>
      )}

    </div>
  );
}

export default ManajemenPengguna;