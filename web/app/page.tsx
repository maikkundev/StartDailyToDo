"use client";

import React, { useEffect, useState } from "react";
import axios from "axios";
import { MantineProvider } from "@mantine/core";
import { Navbar } from "./components/navbar";
import TodoList from "./components/Todo";
import { Todo } from "./interfaces/Todo";

export default function Home() {
  const [todos, setTodos] = useState<Todo[]>([]);

  const toggleTodo = async (todoId: number, currentStatus: boolean) => {
    try {
      const response = await axios.put(
        `http://localhost:3000/todo/${todoId.toString()}`,
        {
          IsDone: !currentStatus,
        }
      );
      if (response.status === 200) {
        setTodos(
          todos.map((todo) =>
            todo.ID === Number(todoId)
              ? { ...todo, IsDone: !todo.IsDone }
              : todo
          )
        );
      }
    } catch (error) {
      console.error(error);
    }
  };

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
            <h1 className="py-2.5">Welcome to StartDailyTodo</h1>
            <TodoList todos={todos} toggleTodo={toggleTodo} />
          </div>
        </div>
      </main>
    </MantineProvider>
  );
}
