import { FC, useState, useEffect } from "react";
import { Task } from "./Navbar";
import "./task-content.css";
import { deleteTask, postTask, putTask } from "./TaskApi";
interface TodoProps {
  selectedTask: number;
}

export const DataLoad: FC<TodoProps> = ({ selectedTask }: TodoProps) => {
  const [taskData, setTaskData] = useState<Task>();

  useEffect(() => {
    console.log(selectedTask);

    const fetchTaskData = async () => {
      const res = await fetch(`http://localhost:80/tasks/${selectedTask}`);
      const data = await res.json();
      setTaskData(data);
    };
    fetchTaskData();
  }, [selectedTask]);

  if (!taskData) {
    return <div>Loading...</div>;
  }

  const isActive = (active: boolean) => {
    return active ? "✅" : "❌";
  };

  const handleDelete = () => {
    deleteTask(selectedTask);
  };

  const handleEdit = (task: Task) => {
    putTask(selectedTask, task);
  };

  const handleAdd = (task: Task) => {
    postTask(task);
  };

  return (
    <div className="content">
      <h1 className="task-title">{taskData.title}</h1>
      <hr />
      <p>{taskData.description}</p>
      <p>Started at : {taskData.created_at}</p>
      <p>Ending at: {taskData.deadline}</p>
      <p>Active: {isActive(taskData.active)}</p>
      <div className="button-wrapper">  
        <button className="button">
          Edit Task
          </button>
        <button onClick={handleDelete} className="button button--delete">
          Delete Task
        </button>
      </div>
    </div>
  );
};
