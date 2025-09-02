import React from "react";
import App from "./App.tsx";
import "./index.css";
import { ToastContextProvider } from "./hooks/useToast.tsx";
import { RelayProvider } from "./providers/RelayProviders.tsx";
import { render } from "preact";
import "./index.css";

render(
  <RelayProvider>
    <ToastContextProvider>
      <App name="hane" />
    </ToastContextProvider>
  </RelayProvider>,
  document.getElementById("root")!
);
