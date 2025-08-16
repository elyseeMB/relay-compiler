import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { ToastContextProvider } from "./hooks/useToast.tsx";
import { ButtonContextProvider } from "./hooks/useButton.tsx";
import { MessageContextProvider } from "./hooks/useTest.tsx";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <MessageContextProvider>
      <ButtonContextProvider>
        <ToastContextProvider>
          <App/>
        </ToastContextProvider>
      </ButtonContextProvider>
    </MessageContextProvider>
  </React.StrictMode>,
);
