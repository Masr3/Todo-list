import { useState, useEffect } from "react";
import "./navbar.css";
import { DataLoad } from "./DataLoad";
import { deleteTask, getTasks } from "./TaskApi";

export interface Task {
  id: number;
  title: string;
  description: string;
  created_at: string;
  deadline: string;
  active: boolean;
}

export const Navbar: React.FC = () => {
  const [data, setData] = useState<Array<Task>>([]);
  const [selectedTask, setSelectedTask] = useState<number>(1);

  const isActive = (active: boolean) => {
    return active ? "✅" : "❌"
  };


  const handleDelete=()=>{
    deleteTask(selectedTask)
  }


  const getTask = (id: number) => {
    setSelectedTask(id);
  };

  useEffect(() => {
    const fetchData = async () => {
      const tasks = await getTasks();
      setData(tasks);
    };
    fetchData();
  }, []);


  return (
    <>
      <div className="container">
        <nav className="navbar">
          <div>
            <h3 className="title">Tasks</h3>
            <ul>
              {data.map((todo: Task) => (
                <li key={todo.id}>
                  <a onClick={() => getTask(todo.id)} href="#">
                    {todo.title}
                  </a>
                  <p>Active: {isActive(todo.active)}</p>
                  <hr />
                 
                </li>
              ))}
            </ul>
            <button className="button button--add">Add Task</button>
          </div>
        </nav>
        <div className="content">
          <DataLoad selectedTask={selectedTask} />
        </div>
      </div>
    </>
  );
};
