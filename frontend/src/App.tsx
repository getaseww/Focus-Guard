import React, { useEffect, useState } from 'react';
import './App.css';
import Table from './components/Table';
import Modal from './components/Modal';
import { FaBeer,FaEdit,FaPen,FaRemoveFormat,FaTrash } from 'react-icons/fa';  // Font Awesome

function App() {

  const [isModalOpen,setIsModalOpen]=useState(false);
  const  [editingWebsite,setEditingWebsite]=useState(false)
  const [websites,setWebsites]=useState([])
  const [formData,setFormData]=useState<any>({})

  const handleEditWebsite=(row:any)=>{

  }

  const handleSubmit =()=>{

  }

  useEffect(()=>{
    const fetchWebsites = async () => {
      const res = await fetch("http://localhost:8081/api/schedules");
      const schedules = await res.json();
      console.log("response",schedules)
      setWebsites(schedules);
    };

    fetchWebsites();
  },[])
    
  return (
    <div className="App">
      <div className="container mx-auto p-4">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-2xl font-bold">Website Blocker</h1>
          <button
            onClick={() => setIsModalOpen(true)}
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          >
            Add Website
          </button>
        </div>

        <Table
          columns={[
            { 
              key: 'URL', 
              header: 'URL / IP',
              render: (value: string) => <div className="text-center">{value}</div>
            },
            { 
              key: 'StartTime', 
              header: 'Start Time',
              render: (value: string) => {
                const time = new Date(value);
                return (
                  <div className="text-center">
                    {time.toLocaleTimeString('en-US', {
                      hour: 'numeric',
                      minute: '2-digit', 
                      hour12: true
                    })}
                  </div>
                );
              }
            },
            { 
              key: 'EndTime', 
              header: 'End Time',
              render: (value: string) => {
                const time = new Date(value);
                return (
                  <div className="text-center">
                    {time.toLocaleTimeString('en-US', {
                      hour: 'numeric',
                      minute: '2-digit',
                      hour12: true
                    })}
                  </div>
                );
              }
            },
            { 
              key: 'DayOfWeek', 
              header: 'Day of Week',
              render: (value: number) => {
                const days = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
                return <div className="text-center">{days[value]}</div>;
              }
            },
            {
              key: 'actions',
              header: 'Actions',
              render: (value: any) => (
                <div className="flex gap-2 justify-center">
                  <button
                    onClick={(e) => {
                      e.stopPropagation();
                      handleEditWebsite(value);
                    }}
                    className="text-blue-500 hover:text-blue-700"
                  >
                    <FaPen/>
                  </button>
                  <button
                    onClick={async (e) => {
                      e.stopPropagation();
                      const response = await fetch(`http://localhost:8081/api/schedules?id=${value}`, {
                        method: 'DELETE'
                      });
                      if (response.ok) {
                        setWebsites(websites.filter((w: any) => w.ID !== value.ID));
                      }
                    }}
                    className="text-red-500 hover:text-red-700"
                  >
                    <FaTrash/>
                  </button>
                </div>
              )
            }
          ]}
          data={websites}
          onRowClick={(row) => handleEditWebsite(row)}
        />

        <Modal
          isOpen={isModalOpen}
          onClose={() => setIsModalOpen(false)}
          title={editingWebsite ? 'Edit Website' : 'Add Website'}
        >
          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">Website URL</label>
              <input
                type="text"
                value={formData.url}
                onChange={(e) => setFormData({...formData, url: e.target.value})}
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                required
              />
            </div>
            <div className="flex gap-4">
              <button
                type="submit"
                className="w-full bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
              >
                {editingWebsite ? 'Update' : 'Add'}
              </button>
            </div>
          </form>
        </Modal>
      </div>
    </div>
  );
}

export default App;
