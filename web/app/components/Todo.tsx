import React from "react";
import { Todo } from "../interfaces/Todo";
import { Checkbox, UnstyledButton, Text } from "@mantine/core";
import classes from "./Navbar.module.css";

interface TodoProps {
  todos: Todo[];
  toggleTodo: (todoId: number, currentStatus: boolean) => Promise<void>;
}

const TodoList: React.FC<TodoProps> = ({ todos, toggleTodo }) => {
  return (
    <div>
      {todos.length > 0 ? (
        todos.map((todo) => (
          <ol key={todo.ID}>
            <li>
              <UnstyledButton
                onClick={() => toggleTodo(todo.ID, todo.IsDone)}
                className={classes.button}
              >
                <Checkbox
                  checked={todo.IsDone || false}
                  tabIndex={-1}
                  size="md"
                  mr="xl"
                  styles={{ input: { cursor: "pointer" } }}
                  aria-hidden
                ></Checkbox>
                <Text>{todo.Name}</Text>
              </UnstyledButton>
            </li>
          </ol>
        ))
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};

export default TodoList;
