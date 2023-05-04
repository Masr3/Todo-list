
const url = 'http://localhost/tasks'

 interface Task {
    id: number;
    title: string;
    description: string;
    created_at: string;
    deadline: string;
    active: boolean;
  }

export const getTasks = async()=>{

    const res = await fetch(url)
    const tasks = await res.json()
    return tasks
}

console.log(getTasks())

export const postTask = async (task:Task) => {
    
    fetch(url,{
        method:"POST",
        body:JSON.stringify({
            
                "title": task.title,
                "description": task.description,
                "created_at": task.created_at,
                "deadline": task.deadline,
                "active": task.active
              
        }), headers:{
            "Content-type": "application/json; charset=UTF-8"
        }
    }
    )
}

export const deleteTask = (id:number)=>{

    const r = confirm("Are you sure you want to delete the task?")
    if (r==true){
        fetch(url+"/"+id, {
            method: "DELETE"
        })
    }    
}

export const putTask = (id:number, task:Task)=>{
    
    fetch(url+"/"+id, {
        method: "PUT",
        body:JSON.stringify({

            "title": task.title,
            "description": task.description,
            "created_at": task.created_at,
            "deadline": task.deadline,
            "active": task.active
          
            
        }), headers:{
            "Content-type": "application/json; charset=UTF-8"
        }
    })

}
