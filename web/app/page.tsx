"use client";

import React, { useEffect, useState } from "react";
import axios from "axios";

interface Todo {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string;
  name: string;
  isDone: boolean;
}

export default function Home() {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    axios
      .get("http://localhost:3000/todos")
      .then((response) => {
        console.log(response.data);
        return setTodos(response.data);
      })
      .catch((error) => console.error(error));
  }, []);

  console.log(todos);

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <h1>Welcome to StartDailyTodo</h1>
      <div>
        <>
          {todos ? (
            todos.map((todo, id) => {
              return (
                <ol key={id}>
                  <li>{todo.name}</li>
                  <li>{todo.isDone}</li>
                </ol>
              );
            })
          ) : (
            <p>Loading...</p>
          )}
        </>
      </div>
    </main>
  );
}
