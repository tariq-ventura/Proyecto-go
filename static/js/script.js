const API_URL = 'http://localhost:3000/api/tasks';

let tasksToAdd = [];
let tasksToUpdate = [];
let tasksToDelete = [];

document.addEventListener('DOMContentLoaded', () => {
    fetchTasks();
    setupDragAndDrop();

    // Establecer la fecha mínima en el campo de fecha
    const today = new Date().toISOString().split('T')[0];
    document.getElementById('new-pending-date').setAttribute('min', today);
});

async function fetchTasks() {
    try {
        const response = await fetch(API_URL);
        const result = await response.json();
        const tasks = result.data || [];
        
        // Limpiar listas
        document.getElementById('pending-tasks').innerHTML = '';
        document.getElementById('in-progress-tasks').innerHTML = '';
        document.getElementById('done-tasks').innerHTML = '';

        tasks.forEach(task => {
            const listId = getListIdByStatus(task.status);
            if (listId) {
                const list = document.getElementById(listId);
                list.appendChild(createTaskElement(task));
            }
        });
    } catch (error) {
        console.error('Error fetching tasks:', error);
    }
}

async function fetchTasksByStatus() {
    const statusSelect = document.getElementById('status-select');
    const status = statusSelect.value;

    try {
        const response = await fetch(`${API_URL}/status/${status}`);
        if (!response.ok) {
            throw new Error(`Error fetching tasks for status "${status}": ${response.statusText}`);
        }

        const result = await response.json();
        const tasks = result.data || [];

        if (tasks.length === 0) {
            alert(`No tasks found for status "${status}".`);
        } else {
            const taskList = tasks.map(task => `- ${task.name} (Due: ${task.dueDate}, Priority: ${task.priority})`).join('\n');
            alert(`Tasks in status "${status}":\n${taskList}`);
        }
    } catch (error) {
        console.error('Error fetching tasks by status:', error);
        alert('An error occurred while fetching tasks. Please try again.');
    }
}

async function fetchTasksByDate() {
    const dateInput = document.getElementById('date-input');
    const date = dateInput.value;

    if (!date) {
        alert('Please select a date.');
        return;
    }

    try {
        const response = await fetch(`${API_URL}/dates/${date}`);
        if (!response.ok) {
            throw new Error(`Error fetching tasks for date "${date}": ${response.statusText}`);
        }

        const result = await response.json();
        const tasks = result.data || [];

        if (tasks.length === 0) {
            alert(`No tasks found for date "${date}".`);
        } else {
            const taskList = tasks.map(task => `- ${task.name} (Due: ${task.dueDate}, Priority: ${task.priority})`).join('\n');
            alert(`Tasks for date "${date}":\n${taskList}`);
        }
    } catch (error) {
        console.error('Error fetching tasks by date:', error);
        alert('An error occurred while fetching tasks. Please try again.');
    }
}

async function fetchTasksByPriority() {
    const prioritySelect = document.getElementById('priority-select');
    const priority = prioritySelect.value;

    try {
        const response = await fetch(`${API_URL}/priority/${priority}`);
        if (!response.ok) {
            throw new Error(`Error fetching tasks for priority "${priority}": ${response.statusText}`);
        }

        const result = await response.json();
        const tasks = result.data || [];

        if (tasks.length === 0) {
            alert(`No tasks found for priority "${priority}".`);
        } else {
            const taskList = tasks.map(task => `- ${task.name} (Due: ${task.dueDate}, Status: ${task.status})`).join('\n');
            alert(`Tasks with priority "${priority}":\n${taskList}`);
        }
    } catch (error) {
        console.error('Error fetching tasks by priority:', error);
        alert('An error occurred while fetching tasks. Please try again.');
    }
}

function getListIdByStatus(status) {
    switch (status) {
        case 'Pending':
            return 'pending-tasks';
        case 'In Progress':
            return 'in-progress-tasks';
        case 'Done':
            return 'done-tasks';
        default:
            return null;
    }
}

function createTaskElement(task) {
    const li = document.createElement('li');
    li.id = task.id;
    li.draggable = true;
    li.dataset.taskId = task.id;
    li.dataset.description = task.description;
    li.dataset.dueDate = task.dueDate;
    li.dataset.priority = task.priority;

    li.innerHTML = `
        <strong>${task.name}</strong><br>
        <small>Description: ${task.description}</small><br>
        <small>Priority: ${task.priority}</small><br>
        <small>Due Date: ${task.dueDate}</small>
    `;

    const deleteButton = document.createElement('button');
    deleteButton.textContent = 'Delete';
    deleteButton.onclick = (e) => {
        e.stopPropagation();
        deleteTask(task.id);
    };
    li.appendChild(deleteButton);

    li.addEventListener('dragstart', handleDragStart);
    return li;
}

function addTask(status) {
    const nameInput = document.getElementById(`new-pending-task`);
    const descriptionInput = document.getElementById(`new-pending-description`);
    const priorityInput = document.getElementById(`new-pending-priority`);
    const dateInput = document.getElementById(`new-pending-date`);

    const name = nameInput.value.trim();
    const description = descriptionInput.value.trim();
    const priority = priorityInput.value;
    const dueDate = dateInput.value;

    if (name && description && dueDate) {
        const newTask = {
            name: name,
            description: description,
            status: status,
            dueDate: dueDate,
            priority: priority
        };
        tasksToAdd.push(newTask);
        const listId = getListIdByStatus(status);
        const list = document.getElementById(listId);
        list.appendChild(createTaskElement(newTask));
        
        // Limpiar los campos después de agregar la tarea
        nameInput.value = '';
        descriptionInput.value = '';
        priorityInput.value = 'Medium';
        dateInput.value = '';
    } else {
        alert('Please fill in all fields before adding a task.');
    }
}

function deleteTask(taskId) {
    const taskElement = document.getElementById(taskId);
    if (taskElement) {
        const task = {
            id: taskId,
            name: taskElement.querySelector('strong').textContent.trim(),
            description: taskElement.dataset.description,
            status: taskElement.parentElement.dataset.status,
            dueDate: taskElement.dataset.dueDate,
            priority: taskElement.dataset.priority
        };
        
        // Si la tarea no es una recién agregada, añadir a la lista de eliminación
        if (!taskId.startsWith('new-')) {
            tasksToDelete.push(task);
        } else {
            // Si es nueva, simplemente la quitamos de la lista de agregar
            tasksToAdd = tasksToAdd.filter(t => t.id !== taskId);
        }
        
        taskElement.remove();
    }
}

async function acceptChanges() {
    try {
        const requests = [];

        if (tasksToAdd.length > 0) {
            // Quitamos el id temporal antes de enviar
            const newTasks = tasksToAdd.map(({ id, ...rest }) => rest);
            requests.push(fetch(API_URL, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(newTasks)
            }));
        }
        if (tasksToUpdate.length > 0) {
            requests.push(fetch(API_URL, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(tasksToUpdate)
            }));
        }
        if (tasksToDelete.length > 0) {
            requests.push(fetch(API_URL, {
                method: 'DELETE',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(tasksToDelete)
            }));
        }

        await Promise.all(requests);

        // Limpiar arrays y recargar tareas
        tasksToAdd = [];
        tasksToUpdate = [];
        tasksToDelete = [];
        fetchTasks();

        // Mostrar alerta de éxito
        alert('Changes have been successfully saved!');
    } catch (error) {
        console.error('Error accepting changes:', error);
        // Mostrar alerta de error
        alert('An error occurred while saving changes. Please try again.');
    }
}

function setupDragAndDrop() {
    const lists = document.querySelectorAll('.task-ul');
    lists.forEach(list => {
        list.addEventListener('dragover', handleDragOver);
        list.addEventListener('drop', handleDrop);
        list.addEventListener('dragenter', handleDragEnter);
        list.addEventListener('dragleave', handleDragLeave);
    });
}

function handleDragEnter(e) {
    e.preventDefault();
    e.target.classList.add('drag-over');
}

function handleDragLeave(e) {
    e.target.classList.remove('drag-over');
}

function handleDragOver(e) {
    e.preventDefault();
}

function handleDrop(e) {
    e.preventDefault();
    e.target.classList.remove('drag-over');
    const taskId = e.dataTransfer.getData('text/plain');
    const taskElement = document.getElementById(taskId);
    const targetList = e.target.closest('.task-ul');

    if (targetList && taskElement && targetList !== taskElement.parentElement) {
        targetList.appendChild(taskElement);
        const newStatus = targetList.dataset.status;
        
        const updatedTask = {
            id: taskId,
            name: taskElement.querySelector('strong').textContent.trim(),
            description: taskElement.dataset.description,
            status: newStatus,
            dueDate: taskElement.dataset.dueDate,
            priority: taskElement.dataset.priority
        };

        // Si la tarea ya estaba para actualizar, solo cambiamos el estado
        const existingUpdate = tasksToUpdate.find(t => t.id === taskId);
        if (existingUpdate) {
            existingUpdate.status = newStatus;
        } else {
            tasksToUpdate.push(updatedTask);
        }
    }
}

function handleDragStart(e) {
    e.dataTransfer.setData('text/plain', e.target.id);
}