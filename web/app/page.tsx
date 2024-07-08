"use client";

import React, { useEffect, useState } from "react";
import axios from "axios";
import { MantineProvider } from "@mantine/core";
import { Navbar } from "./navbar";

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
    <MantineProvider>
      <main>
        <Navbar></Navbar>
        <h1>Welcome to StartDailyTodo</h1>
        <div>
          <>
            {todos ? (
              todos.map((todo, id) => {
                return (
                  <ol key={id}>
                    <li>
                      {todo.name} <button>x</button>
                    </li>
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
    </MantineProvider>
  );
}
