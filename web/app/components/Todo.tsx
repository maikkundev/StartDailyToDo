// components/Todo.tsx
import React from "react";
import { Todo } from "../interfaces/Todo";

interface TodoProps {
    todos: Todo[];
}

const TodoList: React.FC<TodoProps> = ({ todos }) => {
    return (
        <div>
            {todos.length > 0 ? (
                todos.map((todo) => (
                    <ol key={todo.ID}>
                        <li>
                            {todo.name} <button>x</button>
                        </li>
                        <li>{todo.isDone.toString()}</li>
                    </ol>
                ))
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};

export default TodoList;
