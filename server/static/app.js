document.getElementById("scheduleForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const url = document.getElementById("url").value;
    const start_time = document.getElementById("startTime").value;
    const end_time = document.getElementById("endTime").value;
    const day_of_week = parseInt(document.getElementById("dayOfWeek").value);

    const res=await fetch("/api/schedules", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ url, start_time, end_time, day_of_week })
    });
    console.log("post reponse",res)
    
    fetchSchedules();
    e.target.reset();
});

async function fetchSchedules() {
    const res = await fetch("/api/schedules");
    console.log("response on fetch",res)
    const schedules = await res.json();
    console.log("schedules data",schedules);
    
    const schedulesList = document.getElementById("schedulesList");
    schedulesList.innerHTML = '';
    
    schedules.forEach(s => {
        const tr = document.createElement('tr');
        tr.className = 'hover:bg-blue-50';
        
        tr.innerHTML = `
            <td class="px-6 py-4 whitespace-nowrap">${s.URL}</td>
            <td class="px-6 py-4 whitespace-nowrap">${getDayName(s.day_of_week)}</td>
            <td class="px-6 py-4 whitespace-nowrap">${s.start_time} - ${s.end_time}</td>
            <td class="px-6 py-4 whitespace-nowrap">
                <button class="text-red-500 hover:text-red-700">
                    <i class="fas fa-trash"></i>
                </button>
            </td>
        `;
        
        const deleteButton = tr.querySelector('button');
        deleteButton.addEventListener('click', () => deleteSchedule(s.id));
        
        schedulesList.appendChild(tr);
    });
}
async function deleteSchedule(id) {
    const response = await fetch(`/api/schedules?id=${id}`, { method: "DELETE" });
    if (response.ok) {
        await fetchSchedules(); // Refetch schedules after successful deletion
    } else {
        console.error('Failed to delete schedule:', response.statusText);
    }
}

fetchSchedules();
