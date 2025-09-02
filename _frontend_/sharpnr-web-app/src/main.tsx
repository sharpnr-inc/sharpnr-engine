import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { ModalProvider } from "./context/ModalContext.tsx";
import { ToastProvider } from "./context/ToastContext.tsx";

import { LoaderProvider } from "./context/LoaderContext.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ToastProvider>
      <LoaderProvider>
        <ModalProvider>
          <App />
        </ModalProvider>
      </LoaderProvider>
    </ToastProvider>
  </StrictMode>
);
