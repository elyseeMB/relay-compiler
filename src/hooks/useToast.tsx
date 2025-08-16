import {
  ComponentProps,
  PropsWithChildren,
  createContext,
  useCallback,
  useContext,
  useRef,
  useState,
} from "react";
import { Toast } from "../components/Toast.tsx";

type Params = ComponentProps<typeof Toast> & { duration?: number };
type ToastsItems = ComponentProps<typeof Toast> & {
  id: number;
  timer: ReturnType<typeof setTimeout>;
};

const defaultFunction = (toast: Params) => {};

const defaultValue = {
  pushToastRef: { current: defaultFunction },
};

const ToastContext = createContext(defaultValue);

export function ToastContextProvider({ children }: PropsWithChildren) {
  const pushToastRef = useRef(defaultFunction);
  return (
    <ToastContext.Provider value={{ pushToastRef }}>
      {children}
      <Toasts />
    </ToastContext.Provider>
  );
}

export function useToast() {
  const { pushToastRef } = useContext(ToastContext);
  return {
    pushToast: useCallback(
      (toast: Params) => {
        pushToastRef.current(toast);
      },
      [pushToastRef]
    ),
  };
}

function Toasts() {
  const [toasts, setToasts] = useState([] as ToastsItems[]);
  const { pushToastRef } = useContext(ToastContext);
  pushToastRef.current = ({ duration, ...props }: Params) => {
    const id = Date.now();
    const timer = setTimeout(() => {
      setToasts((v) => v.filter((t) => t.id !== id));
    }, duration ?? 5 * 1000);
    const toast = { ...props, id, timer };
    setToasts((v) => [...v, toast]);
  };

  function onRemove(toast: ToastsItems) {
    clearTimeout(toast.timer);
    setToasts((v) => v.filter((t) => t !== toast));
  }

  return (
    <div>
      {toasts.map((toast) => (
        <div
          key={toast.id}
          onClick={() => onRemove(toast)}
          className="container"
        >
          <Toast {...toast} />
        </div>
      ))}
    </div>
  );
}
