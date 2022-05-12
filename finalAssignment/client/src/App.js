import logo from './logo.svg';
import './App.css';
import { ToastContainer } from 'react-toastify';

import NewTaskForm from './features/tasks/NewTaskForm.js';
import TasksList from './features/tasks/TasksList.js';
import NewListForm from './features/lists/NewListForm.js';
import ListsList from './features/lists/ListsList';

function App() {
  return (
    <>
    <div className="container p-4">
      <nav className="navbar navbar-light bg-light mb-4">
        <div className="container-fluid">
          <span className="navbar-brand mb-0 h1">ScaleFocus Academy Todos</span>
        </div>
      </nav>
      <div className="row">
        <div className="col-3 h-100 px-4 border-end">
          <h5 className="mb-3">Lists</h5>
          <NewListForm />
          <ListsList />
        </div>
        <div className="col-5 px-4">
          <h5 className="mb-3">Tasks</h5>
          <NewTaskForm />
          <TasksList />
        </div>
      </div>
    </div>
    <ToastContainer position="bottom-right" />
  </>
  );
}


export default App;
