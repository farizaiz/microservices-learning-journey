import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

function Dashboard() {
  const navigate = useNavigate();
  const [nama, setNama] = useState('');
  const [daftarKasus, setDaftarKasus] = useState([]);

  // State Form & Modal Edit/Create
  const [formData, setFormData] = useState({
    nomor_lp: '', kategori_kasus_id: '', prioritas: 'NORMAL', status_id: 'STAT-001', lokasi_kejadian: ''
  });
  const [isEditing, setIsEditing] = useState(false);
  const [editId, setEditId] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  // State Modal View Detail
  const [viewDetail, setViewDetail] = useState(null); 

  // State Search, Filter, dan Pagination
  const [searchTerm, setSearchTerm] = useState('');
  const [filterStatus, setFilterStatus] = useState('');
  const [filterPrioritas, setFilterPrioritas] = useState('');
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 5;

  useEffect(() => {
    const token = localStorage.getItem('token');
    const namaUser = localStorage.getItem('namaLengkap');
    if (!token) navigate('/');
    else {
      setNama(namaUser);
      fetchDataKasus();
    }
  }, [navigate]);

  useEffect(() => {
    setCurrentPage(1);
  }, [searchTerm, filterStatus, filterPrioritas]);

  const fetchDataKasus = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await axios.get('http://localhost:8080/api/kasus', {
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

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const token = localStorage.getItem('token');
      const config = { headers: { Authorization: `Bearer ${token}` } };

      if (isEditing) {
        await axios.put(`http://localhost:8080/api/kasus/${editId}`, formData, config);
        alert("Status Laporan berhasil diperbarui! 🔄");
      } else {
        await axios.post('http://localhost:8080/api/kasus', formData, config);
        alert("Laporan investigasi resmi dibuat! 🚀");
      }
      closeModal();
      fetchDataKasus();
    } catch (error) {
      alert("Gagal memproses data. Pastikan semua kolom terisi.");
      console.error(error);
    }
  };

  const klikEdit = (kasus) => {
    setIsEditing(true);
    setEditId(kasus.id);
    setFormData({
      nomor_lp: kasus.nomor_lp, kategori_kasus_id: kasus.kategori_kasus_id,
      prioritas: kasus.prioritas, status_id: kasus.status_id, lokasi_kejadian: kasus.lokasi_kejadian
    });
    setIsModalOpen(true);
  };

  const klikHapus = async (id) => {
    const konfirmasi = window.confirm("PERINGATAN AUDIT: Menghapus laporan akan menghilangkan jejak investigasi. Lanjutkan? ⚠️");
    if (!konfirmasi) return;
    try {
      const token = localStorage.getItem('token');
      await axios.delete(`http://localhost:8080/api/kasus/${id}`, { headers: { Authorization: `Bearer ${token}` } });
      fetchDataKasus();
    } catch (error) {
      alert("Gagal menghapus laporan.");
    }
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setIsEditing(false);
    setEditId(null);
    setFormData({ nomor_lp: '', kategori_kasus_id: '', prioritas: 'NORMAL', status_id: 'STAT-001', lokasi_kejadian: '' });
  };

  const openNewModal = () => {
    closeModal();
    // FITUR BARU: AUTO-GENERATE NOMOR LP
    const year = new Date().getFullYear();
    const randomDigits = Math.floor(1000 + Math.random() * 9000);
    const autoLP = `LP/SMIPI/${year}/${randomDigits}`;
    
    setFormData(prev => ({ ...prev, nomor_lp: autoLP, status_id: 'STAT-001' }));
    setIsModalOpen(true);
  };

  // --- LOGIKA FILTER & PENCARIAN ---
  const filteredData = daftarKasus.filter(kasus => {
    const matchSearch = kasus.nomor_lp.toLowerCase().includes(searchTerm.toLowerCase()) || 
                        kasus.lokasi_kejadian.toLowerCase().includes(searchTerm.toLowerCase());
    const matchStatus = filterStatus === '' || kasus.status_id === filterStatus;
    const matchPrioritas = filterPrioritas === '' || kasus.prioritas === filterPrioritas;
    return matchSearch && matchStatus && matchPrioritas;
  });

  // --- LOGIKA PAGINASI ---
  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentItems = filteredData.slice(indexOfFirstItem, indexOfLastItem);
  const totalPages = Math.ceil(filteredData.length / itemsPerPage);

  // --- KALKULASI METRIK ---
  const totalKasus = daftarKasus.length;
  const kasusSelesai = daftarKasus.filter(k => k.status_id === 'STAT-SELESAI').length;
  const kasusAktif = totalKasus - kasusSelesai;
  const kasusUrgent = daftarKasus.filter(k => k.prioritas === 'URGENT').length;

  // --- KOMPONEN HELPER ---
  const PriorityBadge = ({ prioritas }) => {
    let bgColor, textColor;
    if (prioritas === 'URGENT') { bgColor = '#ffe3e3'; textColor = '#c92a2a'; }
    else if (prioritas === 'NORMAL') { bgColor = '#e7f5ff'; textColor = '#1864ab'; }
    else { bgColor = '#ebfbee'; textColor = '#2b8a3e'; }
    return <span style={{ backgroundColor: bgColor, color: textColor, padding: '4px 10px', borderRadius: '4px', fontSize: '11px', fontWeight: 'bold', border: `1px solid ${textColor}40` }}>{prioritas}</span>;
  };

  const getKategoriName = (id) => {
    const cats = { 'KAT-001': 'Pembalakan Liar', 'KAT-002': 'Perburuan Satwa', 'KAT-003': 'Pencemaran Limbah', 'KAT-004': 'Perambahan Hutan' };
    return cats[id] || id;
  };

  // FITUR BARU: TIMELINE VISUAL
  const StatusTimeline = ({ currentStatus }) => {
    const steps = [
      { id: 'STAT-001', label: 'DILAPORKAN' },
      { id: 'STAT-002', label: 'DIINVESTIGASI' },
      { id: 'STAT-003', label: 'VALIDASI' },
      { id: 'STAT-SELESAI', label: 'SELESAI' }
    ];
    const currentIndex = steps.findIndex(s => s.id === currentStatus);

    return (
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginTop: '15px', padding: '15px', backgroundColor: '#f8f9fa', borderRadius: '8px', border: '1px solid #e9ecef' }}>
        {steps.map((step, index) => {
          const isCompleted = index <= currentIndex;
          const isActive = index === currentIndex;
          return (
            <div key={step.id} style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', flex: 1, position: 'relative' }}>
              <div style={{ width: '24px', height: '24px', borderRadius: '50%', backgroundColor: isCompleted ? '#0056b3' : '#e9ecef', color: isCompleted ? '#fff' : '#adb5bd', display: 'flex', alignItems: 'center', justifyContent: 'center', fontSize: '12px', fontWeight: 'bold', zIndex: 2, border: isActive ? '3px solid #cce5ff' : 'none' }}>
                {isCompleted ? '✓' : index + 1}
              </div>
              <span style={{ fontSize: '10px', fontWeight: 'bold', marginTop: '8px', color: isActive ? '#0056b3' : (isCompleted ? '#495057' : '#adb5bd'), textAlign: 'center' }}>
                {step.label}
              </span>
              {index < steps.length - 1 && (
                <div style={{ position: 'absolute', top: '12px', left: '50%', width: '100%', height: '3px', backgroundColor: index < currentIndex ? '#0056b3' : '#e9ecef', zIndex: 1 }}></div>
              )}
            </div>
          );
        })}
      </div>
    );
  };

  // SUMMARY CARD INTERAKTIF
  const SummaryCard = ({ title, value, color, icon, onClickAction }) => (
    <div onClick={onClickAction} style={{ backgroundColor: '#fff', padding: '20px', borderRadius: '12px', boxShadow: '0 2px 10px rgba(0,0,0,0.03)', display: 'flex', alignItems: 'center', justifyContent: 'space-between', borderLeft: `5px solid ${color}`, cursor: 'pointer', transition: 'transform 0.2s' }} onMouseOver={(e) => e.currentTarget.style.transform = 'translateY(-3px)'} onMouseOut={(e) => e.currentTarget.style.transform = 'translateY(0)'}>
      <div>
        <p style={{ margin: '0 0 5px 0', fontSize: '12px', color: '#6c757d', fontWeight: 'bold', textTransform: 'uppercase' }}>{title}</p>
        <h2 style={{ margin: 0, fontSize: '28px', color: '#333', fontWeight: '900' }}>{value}</h2>
      </div>
      <div style={{ fontSize: '32px', opacity: '0.8' }}>{icon}</div>
    </div>
  );

  return (
    <div style={{ backgroundColor: '#f4f7f6', minHeight: '100vh', fontFamily: '"Segoe UI", Roboto, Helvetica, Arial, sans-serif' }}>
      
      <header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', backgroundColor: '#ffffff', padding: '15px 30px', boxShadow: '0 2px 10px rgba(0,0,0,0.05)', position: 'sticky', top: 0, zIndex: 100 }}>
        <div style={{ display: 'flex', alignItems: 'center' }}>
          <img src="/smipi_shield.png" alt="Logo" style={{ height: '40px', marginRight: '15px' }} />
          <div>
            <h2 style={{ margin: 0, color: '#0056b3', fontSize: '20px', fontWeight: '900', letterSpacing: '1px' }}>SMIPI DASHBOARD</h2>
            <p style={{ margin: 0, fontSize: '11px', color: '#6c757d', fontWeight: '600' }}>PORTAL INVESTIGASI TERPADU</p>
          </div>
        </div>
        <div style={{ display: 'flex', alignItems: 'center', gap: '20px' }}>
          <div style={{ textAlign: 'right' }}>
            <p style={{ margin: 0, fontSize: '12px', color: '#6c757d' }}>Unit Penindakan:</p>
            <p style={{ margin: 0, fontSize: '14px', fontWeight: 'bold', color: '#333' }}>Penyidik {nama}</p>
          </div>
          <button onClick={handleLogout} style={{ padding: '8px 16px', backgroundColor: '#f8d7da', color: '#dc3545', border: '1px solid #f5c6cb', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold' }}>Keluar</button>
        </div>
      </header>
      
      <main style={{ padding: '30px', maxWidth: '1400px', margin: '0 auto', display: 'flex', flexDirection: 'column', gap: '25px' }}>
        
        {/* SUMMARY CARDS KLIKABEL */}
        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(4, 1fr)', gap: '20px', width: '100%' }}>
          <SummaryCard title="Total Kasus" value={totalKasus} color="#0056b3" icon="📑" onClickAction={() => {setFilterStatus(''); setFilterPrioritas('');}} />
          <SummaryCard title="Kasus Aktif" value={kasusAktif} color="#fd7e14" icon="🔥" onClickAction={() => setFilterStatus('STAT-002')} />
          <SummaryCard title="Urgent (High)" value={kasusUrgent} color="#dc3545" icon="🚨" onClickAction={() => setFilterPrioritas('URGENT')} />
          <SummaryCard title="Telah Selesai" value={kasusSelesai} color="#28a745" icon="✅" onClickAction={() => setFilterStatus('STAT-SELESAI')} />
        </div>

        <div style={{ backgroundColor: '#ffffff', padding: '25px', borderRadius: '12px', boxShadow: '0 4px 15px rgba(0,0,0,0.05)' }}>
          
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '20px' }}>
            <h3 style={{ margin: 0, color: '#333', fontSize: '18px' }}>Papan Kendali Operasional</h3>
            <button onClick={openNewModal} style={{ padding: '10px 20px', backgroundColor: '#0056b3', color: 'white', border: 'none', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold', fontSize: '14px', boxShadow: '0 2px 6px rgba(0, 86, 179, 0.2)' }}>
              + Register Kasus Baru
            </button>
          </div>

          <div style={{ display: 'flex', gap: '15px', marginBottom: '20px', backgroundColor: '#f8f9fa', padding: '15px', borderRadius: '8px', border: '1px solid #e9ecef' }}>
            <div style={{ flex: '2', position: 'relative' }}>
              <span style={{ position: 'absolute', left: '12px', top: '10px', fontSize: '14px' }}>🔍</span>
              <input type="text" placeholder="Cari Nomor LP atau Lokasi..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} style={{ width: '100%', padding: '10px 10px 10px 35px', border: '1px solid #ced4da', borderRadius: '6px', outline: 'none', boxSizing: 'border-box' }} />
            </div>
            <select value={filterPrioritas} onChange={(e) => setFilterPrioritas(e.target.value)} style={{ flex: '1', padding: '10px', border: '1px solid #ced4da', borderRadius: '6px', outline: 'none', backgroundColor: '#fff', fontWeight: 'bold', color: filterPrioritas ? '#0056b3' : '#495057' }}>
              <option value="">Semua Prioritas</option>
              <option value="URGENT">🔴 URGENT</option>
              <option value="NORMAL">🔵 NORMAL</option>
              <option value="LOW">🟢 LOW</option>
            </select>
            <select value={filterStatus} onChange={(e) => setFilterStatus(e.target.value)} style={{ flex: '1', padding: '10px', border: '1px solid #ced4da', borderRadius: '6px', outline: 'none', backgroundColor: '#fff', fontWeight: 'bold', color: filterStatus ? '#0056b3' : '#495057' }}>
              <option value="">Semua Status</option>
              <option value="STAT-001">DILAPORKAN</option>
              <option value="STAT-002">DIINVESTIGASI</option>
              <option value="STAT-003">VALIDASI</option>
              <option value="STAT-SELESAI">SELESAI</option>
            </select>
          </div>
          
          <div style={{ overflowX: 'auto' }}>
            <table style={{ width: '100%', borderCollapse: 'collapse', fontSize: '13px', textAlign: 'left' }}>
              <thead>
                <tr style={{ backgroundColor: '#f8f9fa', borderBottom: '2px solid #dee2e6' }}>
                  <th style={{ padding: '12px 16px', color: '#495057', fontWeight: '700' }}>No. Register (LP)</th>
                  <th style={{ padding: '12px 16px', color: '#495057', fontWeight: '700' }}>Kategori</th>
                  <th style={{ padding: '12px 16px', color: '#495057', fontWeight: '700' }}>Unit Penanganan</th>
                  <th style={{ padding: '12px 16px', color: '#495057', fontWeight: '700' }}>Prioritas</th>
                  <th style={{ padding: '12px 16px', color: '#495057', fontWeight: '700', textAlign: 'center' }}>Tindakan</th>
                </tr>
              </thead>
              <tbody>
                {currentItems.length > 0 ? (
                  currentItems.map((kasus) => (
                    <tr key={kasus.id} style={{ borderBottom: '1px solid #e9ecef' }} onMouseOver={(e) => e.currentTarget.style.backgroundColor = '#f8f9fa'} onMouseOut={(e) => e.currentTarget.style.backgroundColor = 'transparent'}>
                      <td style={{ padding: '12px 16px' }}>
                        <div style={{ fontWeight: '700', color: '#0056b3', fontSize: '14px' }}>{kasus.nomor_lp}</div>
                        <div style={{ fontSize: '11px', color: '#6c757d', marginTop: '4px' }}>📍 {kasus.lokasi_kejadian.substring(0, 25)}...</div>
                      </td>
                      <td style={{ padding: '12px 16px', color: '#495057', fontWeight: '600' }}>{getKategoriName(kasus.kategori_kasus_id)}</td>
                      
                      {/* KOLOM MOCKUP: ASSIGNED TO */}
                      <td style={{ padding: '12px 16px' }}>
                        <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                          <div style={{ width: '24px', height: '24px', borderRadius: '50%', backgroundColor: '#e9ecef', display: 'flex', alignItems: 'center', justifyContent: 'center', fontSize: '10px' }}>👤</div>
                          <span style={{ fontWeight: '600', color: '#333' }}>Tim Alpha ({nama || 'Penyidik'})</span>
                        </div>
                      </td>

                      <td style={{ padding: '12px 16px' }}><PriorityBadge prioritas={kasus.prioritas} /></td>
                      
                      <td style={{ padding: '12px 16px', textAlign: 'center' }}>
                        {/* UPDATE INLINE: Ganti Edit menjadi Action Button lebih bermakna */}
                        <button onClick={() => setViewDetail(kasus)} style={{ padding: '6px 12px', backgroundColor: '#e7f5ff', color: '#0056b3', border: '1px solid #b8daff', borderRadius: '4px', cursor: 'pointer', marginRight: '5px', fontWeight: 'bold', fontSize: '12px' }}>
                          👁️ Pantau Kasus
                        </button>
                        <button onClick={() => klikEdit(kasus)} style={{ padding: '6px 12px', backgroundColor: '#fff3cd', color: '#856404', border: '1px solid #ffeeba', borderRadius: '4px', cursor: 'pointer', fontWeight: 'bold', fontSize: '12px' }}>
                          ⚙️ Update Status
                        </button>
                      </td>
                    </tr>
                  ))
                ) : (
                  <tr>
                    <td colSpan="5" style={{ padding: '50px', textAlign: 'center', color: '#868e96' }}>
                      <div style={{ fontSize: '40px', marginBottom: '10px' }}>📁</div>
                      <strong>Papan Kendali Kosong.</strong><br/>Tidak ada data yang sesuai dengan filter operasi saat ini.
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>

          {filteredData.length > 0 && (
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginTop: '20px', paddingTop: '15px', borderTop: '1px solid #e9ecef' }}>
              <span style={{ fontSize: '12px', color: '#6c757d', fontWeight: 'bold' }}>
                Menampilkan {indexOfFirstItem + 1} - {Math.min(indexOfLastItem, filteredData.length)} dari {filteredData.length} laporan aktif
              </span>
              <div style={{ display: 'flex', gap: '5px' }}>
                <button onClick={() => setCurrentPage(prev => Math.max(prev - 1, 1))} disabled={currentPage === 1} style={{ padding: '6px 12px', border: '1px solid #ced4da', backgroundColor: currentPage === 1 ? '#e9ecef' : '#fff', borderRadius: '4px', cursor: currentPage === 1 ? 'not-allowed' : 'pointer', fontSize: '12px', fontWeight: 'bold' }}>&laquo; Prev</button>
                <button onClick={() => setCurrentPage(prev => Math.min(prev + 1, totalPages))} disabled={currentPage === totalPages} style={{ padding: '6px 12px', border: '1px solid #ced4da', backgroundColor: currentPage === totalPages ? '#e9ecef' : '#fff', borderRadius: '4px', cursor: currentPage === totalPages ? 'not-allowed' : 'pointer', fontSize: '12px', fontWeight: 'bold' }}>Next &raquo;</button>
              </div>
            </div>
          )}
        </div>
      </main>

      {/* MODAL FORM CREATE/EDIT */}
      {isModalOpen && (
        <div style={{ position: 'fixed', top: 0, left: 0, right: 0, bottom: 0, backgroundColor: 'rgba(0,0,0,0.6)', zIndex: 999, display: 'flex', justifyContent: 'center', alignItems: 'center', backdropFilter: 'blur(3px)' }}>
          <div style={{ backgroundColor: '#ffffff', padding: '30px', borderRadius: '12px', width: '100%', maxWidth: '600px', boxShadow: '0 10px 30px rgba(0,0,0,0.2)' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '20px' }}>
              <h3 style={{ margin: 0, color: isEditing ? '#fd7e14' : '#0056b3', fontSize: '20px' }}>
                {isEditing ? '⚙️ Update Status Kasus' : '📝 Registrasi Kasus Baru'}
              </h3>
              <button onClick={closeModal} style={{ background: 'none', border: 'none', fontSize: '24px', cursor: 'pointer', color: '#6c757d' }}>&times;</button>
            </div>
            
            {/* Visual Reminder di Form Edit */}
            {isEditing && <StatusTimeline currentStatus={formData.status_id} />}
            <hr style={{ border: '0', borderTop: '2px solid #f1f3f5', margin: '20px 0' }} />

            <form onSubmit={handleSubmit}>
              <div style={{ display: 'flex', gap: '15px', marginBottom: '15px' }}>
                <div style={{ flex: '1' }}>
                  <label style={{ display: 'block', fontSize: '12px', fontWeight: 'bold', color: '#495057', marginBottom: '6px' }}>Nomor Register (Otomatis)</label>
                  <input type="text" readOnly={!isEditing} value={formData.nomor_lp} onChange={(e) => setFormData({...formData, nomor_lp: e.target.value})} style={{ width: '100%', padding: '10px 12px', border: '1px solid #ced4da', borderRadius: '6px', boxSizing: 'border-box', outline: 'none', backgroundColor: '#f8f9fa', color: '#495057', fontWeight: 'bold' }} />
                </div>
                <div style={{ flex: '1' }}>
                  <label style={{ display: 'block', fontSize: '12px', fontWeight: 'bold', color: '#495057', marginBottom: '6px' }}>Kategori Pelanggaran</label>
                  <select required value={formData.kategori_kasus_id} onChange={(e) => setFormData({...formData, kategori_kasus_id: e.target.value})} style={{ width: '100%', padding: '10px 12px', border: '1px solid #ced4da', borderRadius: '6px', boxSizing: 'border-box', outline: 'none', backgroundColor: '#fff' }}>
                    <option value="" disabled>-- Pilih Kategori --</option>
                    <option value="KAT-001">Pembalakan Liar (Illegal Logging)</option>
                    <option value="KAT-002">Perburuan Satwa Dilindungi</option>
                    <option value="KAT-003">Pencemaran Limbah B3</option>
                    <option value="KAT-004">Perambahan Kawasan Hutan</option>
                  </select>
                </div>
              </div>

              <div style={{ display: 'flex', gap: '15px', marginBottom: '15px' }}>
                <div style={{ flex: '1' }}>
                  <label style={{ display: 'block', fontSize: '12px', fontWeight: 'bold', color: '#495057', marginBottom: '6px' }}>Tingkat Prioritas</label>
                  <select value={formData.prioritas} onChange={(e) => setFormData({...formData, prioritas: e.target.value})} style={{ width: '100%', padding: '10px 12px', border: '1px solid #ced4da', borderRadius: '6px', boxSizing: 'border-box', outline: 'none', backgroundColor: '#fff' }}>
                    <option value="LOW">🟢 LOW (Biasa)</option>
                    <option value="NORMAL">🔵 NORMAL (Menengah)</option>
                    <option value="URGENT">🔴 URGENT (Tinggi/Atensi)</option>
                  </select>
                </div>
                <div style={{ flex: '1' }}>
                  <label style={{ display: 'block', fontSize: '12px', fontWeight: 'bold', color: '#dc3545', marginBottom: '6px' }}>Tindakan Operasional</label>
                  <select required value={formData.status_id} onChange={(e) => setFormData({...formData, status_id: e.target.value})} style={{ width: '100%', padding: '10px 12px', border: '2px solid #cce5ff', borderRadius: '6px', boxSizing: 'border-box', outline: 'none', backgroundColor: '#e7f5ff', fontWeight: 'bold', color: '#0056b3' }}>
                    <option value="STAT-001">DILAPORKAN (Tahap 1)</option>
                    <option value="STAT-002">DIINVESTIGASI (Tahap 2)</option>
                    <option value="STAT-003">VALIDASI BUKTI (Tahap 3)</option>
                    <option value="STAT-SELESAI">KASUS SELESAI (Final)</option>
                  </select>
                </div>
              </div>
              <div style={{ marginBottom: '25px' }}>
                <label style={{ display: 'block', fontSize: '12px', fontWeight: 'bold', color: '#495057', marginBottom: '6px' }}>Lokasi Kejadian Perkara (TKP)</label>
                <textarea required value={formData.lokasi_kejadian} onChange={(e) => setFormData({...formData, lokasi_kejadian: e.target.value})} style={{ width: '100%', padding: '10px 12px', border: '1px solid #ced4da', borderRadius: '6px', boxSizing: 'border-box', outline: 'none', minHeight: '60px', resize: 'vertical' }}></textarea>
              </div>
              <div style={{ display: 'flex', gap: '10px', justifyContent: 'flex-end' }}>
                <button type="button" onClick={closeModal} style={{ padding: '12px 20px', backgroundColor: '#e9ecef', color: '#495057', border: 'none', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold' }}>Batalkan</button>
                <button type="submit" style={{ padding: '12px 20px', backgroundColor: isEditing ? '#28a745' : '#0056b3', color: 'white', border: 'none', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold' }}>{isEditing ? 'Simpan Progress' : 'Register Laporan'}</button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* MODAL VIEW DETAIL INVESTIGASI */}
      {viewDetail && (
        <div style={{ position: 'fixed', top: 0, left: 0, right: 0, bottom: 0, backgroundColor: 'rgba(0,0,0,0.7)', zIndex: 999, display: 'flex', justifyContent: 'center', alignItems: 'center', backdropFilter: 'blur(4px)' }}>
          <div style={{ backgroundColor: '#ffffff', padding: '0', borderRadius: '12px', width: '100%', maxWidth: '600px', boxShadow: '0 10px 40px rgba(0,0,0,0.3)', overflow: 'hidden' }}>
            
            {/* Modal Header */}
            <div style={{ backgroundColor: '#0056b3', padding: '20px 30px', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
              <div>
                <span style={{ backgroundColor: '#fff', color: '#0056b3', padding: '4px 8px', borderRadius: '4px', fontSize: '10px', fontWeight: '900', letterSpacing: '1px', marginBottom: '8px', display: 'inline-block' }}>DOKUMEN INVESTIGASI</span>
                <h3 style={{ margin: 0, color: '#fff', fontSize: '22px' }}>{viewDetail.nomor_lp}</h3>
              </div>
              <button onClick={() => setViewDetail(null)} style={{ background: 'none', border: 'none', fontSize: '28px', cursor: 'pointer', color: '#fff', opacity: '0.7' }}>&times;</button>
            </div>
            
            <div style={{ padding: '30px' }}>
              {/* TIMELINE PROGRESS */}
              <div style={{ marginBottom: '25px' }}>
                <p style={{ margin: '0 0 5px 0', fontSize: '12px', color: '#6c757d', fontWeight: 'bold' }}>STATUS INVESTIGASI TERKINI:</p>
                <StatusTimeline currentStatus={viewDetail.status_id} />
              </div>

              <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '20px', marginBottom: '20px' }}>
                <div><span style={{ fontSize: '11px', color: '#6c757d', fontWeight: 'bold', display: 'block', marginBottom: '4px' }}>KATEGORI PELANGGARAN</span><div style={{ fontSize: '14px', fontWeight: 'bold', color: '#333' }}>{getKategoriName(viewDetail.kategori_kasus_id)}</div></div>
                <div><span style={{ fontSize: '11px', color: '#6c757d', fontWeight: 'bold', display: 'block', marginBottom: '4px' }}>TINGKAT PRIORITAS</span><PriorityBadge prioritas={viewDetail.prioritas} /></div>
              </div>

              <div style={{ marginBottom: '20px' }}>
                <span style={{ fontSize: '11px', color: '#6c757d', fontWeight: 'bold', display: 'block', marginBottom: '4px' }}>LOKASI KEJADIAN PERKARA (TKP)</span>
                <div style={{ fontSize: '14px', backgroundColor: '#f8f9fa', padding: '15px', borderRadius: '8px', border: '1px solid #e9ecef', color: '#495057', lineHeight: '1.5' }}>
                  {viewDetail.lokasi_kejadian}
                </div>
              </div>

              {/* MOCKUP ACTIVITY LOG (Persiapan Database) */}
              <div>
                <span style={{ fontSize: '11px', color: '#6c757d', fontWeight: 'bold', display: 'block', marginBottom: '10px' }}>CATATAN AKTIVITAS (AUDIT TRAIL)</span>
                <div style={{ borderLeft: '2px solid #e9ecef', paddingLeft: '15px', marginLeft: '5px' }}>
                  <div style={{ position: 'relative', marginBottom: '15px' }}>
                    <div style={{ position: 'absolute', left: '-21px', top: '2px', width: '10px', height: '10px', borderRadius: '50%', backgroundColor: '#ced4da' }}></div>
                    <span style={{ fontSize: '10px', color: '#adb5bd', fontWeight: 'bold' }}>HARI INI</span>
                    <p style={{ margin: '2px 0', fontSize: '13px', color: '#495057' }}>Laporan didaftarkan ke dalam sistem oleh <b>Tim Alpha</b>.</p>
                  </div>
                  <div style={{ position: 'relative' }}>
                    <div style={{ position: 'absolute', left: '-21px', top: '2px', width: '10px', height: '10px', borderRadius: '50%', backgroundColor: '#fff', border: '2px dashed #ced4da' }}></div>
                    <p style={{ margin: '2px 0', fontSize: '13px', color: '#adb5bd', fontStyle: 'italic' }}>Menunggu penambahan bukti lapangan...</p>
                  </div>
                </div>
              </div>

            </div>
            
            {/* Modal Footer dengan Aksi Cepat */}
            <div style={{ backgroundColor: '#f8f9fa', padding: '15px 30px', borderTop: '1px solid #e9ecef', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
              <button style={{ background: 'none', border: 'none', color: '#0056b3', fontWeight: 'bold', fontSize: '13px', cursor: 'pointer' }}>📎 Lampirkan Bukti (Segera Hadir)</button>
              <div style={{ display: 'flex', gap: '10px' }}>
                <button onClick={() => setViewDetail(null)} style={{ padding: '10px 20px', backgroundColor: '#e9ecef', color: '#495057', border: 'none', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold' }}>Tutup Layar</button>
                <button onClick={() => { setViewDetail(null); klikEdit(viewDetail); }} style={{ padding: '10px 20px', backgroundColor: '#0056b3', color: '#fff', border: 'none', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold' }}>⚙️ Update Kasus Ini</button>
              </div>
            </div>

          </div>
        </div>
      )}

    </div>
  );
}

export default Dashboard;