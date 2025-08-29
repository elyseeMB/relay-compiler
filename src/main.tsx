import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { ToastContextProvider } from "./hooks/useToast.tsx";
import { RelayProvider } from "./providers/RelayProviders.tsx";

class RootElement extends HTMLElement {
  // @ts-ignore
  name: string | void;

  static get observedAttributes() {
    return ["name"];
  }

  attributeChangedCallback(name: string, oldValue: string, newValue: string) {
    //   console.log(`Attribute ${name} has changed.
    //  old Value: ${oldValue}
    //   new value: ${newValue}
    //   `);
  }

  connectedCallback() {
    this.addEventListener("alert", (ev) => {
      this.name = this.setAttribute("name", ev.detail.name);
    });

    return ReactDOM.createRoot(this).render(
      <React.StrictMode>
        <RelayProvider>
          <ToastContextProvider>
            <App name={this.name!} />
          </ToastContextProvider>
        </RelayProvider>
      </React.StrictMode>
    );
  }
}

customElements.define("root-element", RootElement);

setTimeout(() => {
  document.querySelector("root-element")!.dispatchEvent(
    new CustomEvent("alert", {
      detail: { name: "rochi" },
      bubbles: true,
    })
  );
}, 1000);
