import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

function Dashboard() {
  const navigate = useNavigate();
  const [nama, setNama] = useState('');
  const [daftarKasus, setDaftarKasus] = useState([]);

  // State Form
  const [formData, setFormData] = useState({
    nomor_lp: '',
    kategori_kasus_id: '',
    prioritas: 'NORMAL',
    status_id: '',
    lokasi_kejadian: ''
  });

  // State Penanda Edit
  const [isEditing, setIsEditing] = useState(false);
  const [editId, setEditId] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem('token');
    const namaUser = localStorage.getItem('namaLengkap');
    if (!token) navigate('/');
    else {
      setNama(namaUser);
      fetchDataKasus();
    }
  }, [navigate]);

  const fetchDataKasus = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await axios.get('http://localhost:8080/kasus', {
        headers: { Authorization: `Bearer ${token}` }
      });
      setDaftarKasus(response.data.data);
    } catch (error) {
      console.error("Gagal mengambil data:", error);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('namaLengkap');
    navigate('/');
  };

  // 1. FUNGSI SUBMIT (Bisa POST untuk Simpan, bisa PUT untuk Update)
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const token = localStorage.getItem('token');
      const config = { headers: { Authorization: `Bearer ${token}` } };

      if (isEditing) {
        // UPDATE DATA LAMA (PUT)
        await axios.put(`http://localhost:8080/kasus/${editId}`, formData, config);
        alert("Data berhasil diperbarui! 🔄");
      } else {
        // SIMPAN DATA BARU (POST)
        await axios.post('http://localhost:8080/kasus', formData, config);
        alert("Laporan Kasus Berhasil Ditambahkan! 🚀");
      }

      resetForm();
      fetchDataKasus(); // Refresh tabel

    } catch (error) {
      alert("Gagal memproses data.");
      console.error(error);
    }
  };

  // 2. FUNGSI SAAT TOMBOL EDIT DITEKAN PADA TABEL
  const klikEdit = (kasus) => {
    setIsEditing(true);
    setEditId(kasus.id);
    setFormData({
      nomor_lp: kasus.nomor_lp,
      kategori_kasus_id: kasus.kategori_kasus_id,
      prioritas: kasus.prioritas,
      status_id: kasus.status_id,
      lokasi_kejadian: kasus.lokasi_kejadian
    });
  };

  // 3. FUNGSI SAAT TOMBOL HAPUS DITEKAN
  const klikHapus = async (id) => {
    const konfirmasi = window.confirm("Apakah Anda yakin ingin menghapus laporan ini? ⚠️");
    if (!konfirmasi) return;

    try {
      const token = localStorage.getItem('token');
      await axios.delete(`http://localhost:8080/kasus/${id}`, {
        headers: { Authorization: `Bearer ${token}` }
      });
      alert("Laporan berhasil dihapus! 🗑️");
      fetchDataKasus(); // Refresh tabel
    } catch (error) {
      alert("Gagal menghapus laporan.");
    }
  };

  const resetForm = () => {
    setIsEditing(false);
    setEditId(null);
    setFormData({
      nomor_lp: '',
      kategori_kasus_id: '',
      prioritas: 'NORMAL',
      status_id: '',
      lokasi_kejadian: ''
    });
  };

  return (
    <div style={{ padding: '20px', fontFamily: 'sans-serif', backgroundColor: '#f9fbfd', minHeight: '100vh' }}>
      <header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', backgroundColor: '#ffffff', padding: '15px 25px', borderRadius: '8px', boxShadow: '0 2px 5px rgba(0,0,0,0.05)' }}>
        <h2 style={{ margin: 0, color: '#0056b3' }}>SMIPI Dashboard 🕵️‍♂️</h2>
        <div>
          <span style={{ marginRight: '15px', fontWeight: 'bold' }}>Halo, Penyidik {nama}</span>
          <button onClick={handleLogout} style={{ padding: '8px 15px', backgroundColor: '#dc3545', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer', fontWeight: 'bold' }}>Logout</button>
        </div>
      </header>
      
      <main style={{ marginTop: '30px', display: 'flex', gap: '20px', alignItems: 'flex-start' }}>
        
        {/* PANEL KIRI: FORMULIR */}
        <div style={{ flex: '1', backgroundColor: '#ffffff', padding: '20px', borderRadius: '8px', boxShadow: '0 2px 5px rgba(0,0,0,0.05)' }}>
          <h3 style={{ marginTop: 0, color: isEditing ? '#ff9900' : '#333' }}>
            {isEditing ? '✏️ Edit Laporan' : 'Buat Laporan Baru'}
          </h3>
          <hr style={{ border: '0', borderTop: '1px solid #eee', marginBottom: '15px' }} />
          
          <form onSubmit={handleSubmit}>
            <div style={{ marginBottom: '10px' }}>
              <label style={{ fontSize: '14px', fontWeight: 'bold' }}>Nomor LP:</label>
              <input type="text" required value={formData.nomor_lp} onChange={(e) => setFormData({...formData, nomor_lp: e.target.value})} style={{ width: '100%', padding: '8px', marginTop: '5px', boxSizing: 'border-box' }} />
            </div>
            
            <div style={{ marginBottom: '10px' }}>
              <label style={{ fontSize: '14px', fontWeight: 'bold' }}>Kategori Kasus (ID):</label>
              <input type="text" required value={formData.kategori_kasus_id} onChange={(e) => setFormData({...formData, kategori_kasus_id: e.target.value})} style={{ width: '100%', padding: '8px', marginTop: '5px', boxSizing: 'border-box' }} />
            </div>

            <div style={{ marginBottom: '10px' }}>
              <label style={{ fontSize: '14px', fontWeight: 'bold' }}>Prioritas:</label>
              <select value={formData.prioritas} onChange={(e) => setFormData({...formData, prioritas: e.target.value})} style={{ width: '100%', padding: '8px', marginTop: '5px', boxSizing: 'border-box' }}>
                <option value="LOW">LOW</option>
                <option value="NORMAL">NORMAL</option>
                <option value="URGENT">URGENT</option>
              </select>
            </div>

            <div style={{ marginBottom: '10px' }}>
              <label style={{ fontSize: '14px', fontWeight: 'bold' }}>Status (ID):</label>
              <input type="text" required value={formData.status_id} onChange={(e) => setFormData({...formData, status_id: e.target.value})} style={{ width: '100%', padding: '8px', marginTop: '5px', boxSizing: 'border-box' }} />
            </div>

            <div style={{ marginBottom: '15px' }}>
              <label style={{ fontSize: '14px', fontWeight: 'bold' }}>Lokasi Kejadian:</label>
              <textarea required value={formData.lokasi_kejadian} onChange={(e) => setFormData({...formData, lokasi_kejadian: e.target.value})} style={{ width: '100%', padding: '8px', marginTop: '5px', boxSizing: 'border-box', resize: 'vertical' }}></textarea>
            </div>

            <button type="submit" style={{ width: '100%', padding: '10px', backgroundColor: isEditing ? '#ff9900' : '#28a745', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer', fontWeight: 'bold', marginBottom: '10px' }}>
              {isEditing ? 'Update Laporan' : 'Simpan Laporan'}
            </button>
            
            {/* Tombol Batal Edit muncul jika sedang mode Edit */}
            {isEditing && (
              <button type="button" onClick={resetForm} style={{ width: '100%', padding: '10px', backgroundColor: '#6c757d', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer', fontWeight: 'bold' }}>
                Batal Edit
              </button>
            )}
          </form>
        </div>

        {/* PANEL KANAN: TABEL KASUS */}
        <div style={{ flex: '2', backgroundColor: '#ffffff', padding: '20px', borderRadius: '8px', boxShadow: '0 2px 5px rgba(0,0,0,0.05)' }}>
          <h3 style={{ marginTop: 0, color: '#333' }}>Daftar Kasus Investigasi</h3>
          <hr style={{ border: '0', borderTop: '1px solid #eee', marginBottom: '15px' }} />
          
          <table style={{ width: '100%', borderCollapse: 'collapse', fontSize: '14px' }}>
            <thead>
              <tr style={{ backgroundColor: '#f4f4f4', color: '#333', textAlign: 'left' }}>
                <th style={{ padding: '12px', borderBottom: '2px solid #ddd' }}>No. LP</th>
                <th style={{ padding: '12px', borderBottom: '2px solid #ddd' }}>Prioritas</th>
                <th style={{ padding: '12px', borderBottom: '2px solid #ddd' }}>Status</th>
                <th style={{ padding: '12px', borderBottom: '2px solid #ddd' }}>Lokasi</th>
                <th style={{ padding: '12px', borderBottom: '2px solid #ddd', textAlign: 'center' }}>Aksi</th>
              </tr>
            </thead>
            <tbody>
              {daftarKasus.length > 0 ? (
                daftarKasus.map((kasus) => (
                  <tr key={kasus.id} style={{ borderBottom: '1px solid #eee' }}>
                    <td style={{ padding: '12px' }}>{kasus.nomor_lp}</td>
                    <td style={{ padding: '12px' }}>
                      <span style={{ backgroundColor: kasus.prioritas === 'URGENT' ? '#ffcccc' : (kasus.prioritas === 'NORMAL' ? '#e6f7ff' : '#e6ffe6'), padding: '4px 8px', borderRadius: '12px', fontSize: '12px', fontWeight: 'bold', color: kasus.prioritas === 'URGENT' ? '#cc0000' : (kasus.prioritas === 'NORMAL' ? '#0066cc' : '#006600') }}>
                        {kasus.prioritas}
                      </span>
                    </td>
                    <td style={{ padding: '12px' }}>{kasus.status_id}</td>
                    <td style={{ padding: '12px' }}>{kasus.lokasi_kejadian}</td>
                    <td style={{ padding: '12px', textAlign: 'center' }}>
                      <button onClick={() => klikEdit(kasus)} style={{ padding: '5px 10px', backgroundColor: '#ffc107', border: 'none', borderRadius: '4px', cursor: 'pointer', marginRight: '5px', fontSize: '12px', fontWeight: 'bold' }}>Edit</button>
                      <button onClick={() => klikHapus(kasus.id)} style={{ padding: '5px 10px', backgroundColor: '#dc3545', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer', fontSize: '12px', fontWeight: 'bold' }}>Hapus</button>
                    </td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td colSpan="5" style={{ padding: '20px', textAlign: 'center', fontStyle: 'italic', color: '#888' }}>Belum ada laporan.</td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </main>
    </div>
  );
}

export default Dashboard;