// page.tsx
"use client";

import React, { useEffect, useState } from "react";
import axios from "axios";
import { MantineProvider } from "@mantine/core";
import { Navbar } from "./navbar";
import TodoList from "./components/Todo";
import { Todo } from "./interfaces/Todo";

export default function Home() {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    axios
        .get("http://localhost:3000/todos")
        .then((response) => {
          console.log(response.data);
          setTodos(response.data);
        })
        .catch((error) => console.error(error));
  }, []);

  console.log(todos);

  return (
      <MantineProvider defaultColorScheme="light">
        <main className="flex h-screen">
          <Navbar />
          <div className="flex-grow flex justify-center items-start py-2.5">
            <div>
              <h1>Welcome to StartDailyTodo</h1>
              <TodoList todos={todos} />
            </div>
          </div>
        </main>
      </MantineProvider>
  );
}
