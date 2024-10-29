document.getElementById("scheduleForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const url = document.getElementById("url").value;
    const startTime = document.getElementById("startTime").value;
    const endTime = document.getElementById("endTime").value;
    const dayOfWeek = parseInt(document.getElementById("dayOfWeek").value);

    await fetch("/api/schedules", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ url, startTime, endTime, dayOfWeek })
    });

    fetchSchedules();
    e.target.reset();
});

async function fetchSchedules() {
    const res = await fetch("/api/schedules");
    const schedules = await res.json();
    const schedulesList = document.getElementById("schedulesList");
    schedulesList.innerHTML = schedules.map(s => `
        <li class="p-2 bg-white shadow rounded flex justify-between">
            <span>${s.url} from ${s.start_time} to ${s.end_time} on day ${s.day_of_week}</span>
            <button onclick="deleteSchedule(${s.id})" class="text-red-500">Delete</button>
        </li>
    `).join("");
}

async function deleteSchedule(id) {
    await fetch(`/api/schedules/${id}`, { method: "DELETE" });
    fetchSchedules();
}

fetchSchedules();
