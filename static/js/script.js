const API_URL = 'http://localhost:3000/api/tasks';

let tasksToAdd = [];
let tasksToUpdate = [];
let tasksToDelete = [];

document.addEventListener('DOMContentLoaded', () => {
    fetchTasks();
    setupDragAndDrop();
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
    li.textContent = task.name;

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
    const input = document.getElementById(`new-pending-task`);
    const name = input.value.trim();
    if (name) {
        const newTask = {
            // El ID se asignará en el backend, pero usamos uno temporal para el frontend
            id: `new-${Date.now()}`, 
            name: name,
            description: "Default description",
            status: status,
            dueDate: new Date().toISOString().split('T')[0],
            priority: "Medium"
        };
        tasksToAdd.push(newTask);
        const listId = getListIdByStatus(status);
        const list = document.getElementById(listId);
        list.appendChild(createTaskElement(newTask));
        input.value = '';
    }
}

function deleteTask(taskId) {
    const taskElement = document.getElementById(taskId);
    if (taskElement) {
        const task = {
            id: taskId,
            name: taskElement.textContent.replace('Delete',''),
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
            name: taskElement.textContent.replace('Delete',''),
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