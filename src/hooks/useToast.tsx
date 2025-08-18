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

type Params = ComponentProps<typeof Toast> &
  Partial<{
    duration: number;
    id: number;
    timer: ReturnType<typeof setTimeout>;
  }>;

//@ts-ignore
const defaultFunc = (props: Params) => {};

const pushToastGlobal = {
  pushToastRef: { current: defaultFunc },
};

const toastContext = createContext(pushToastGlobal);

export function ToastContextProvider(props: PropsWithChildren) {
  const pushToastRef = useRef(defaultFunc);
  return (
    <toastContext.Provider value={{ pushToastRef }}>
      {props.children}
      <Toasts />
    </toastContext.Provider>
  );
}

export function useToast() {
  const { pushToastRef } = useContext(toastContext);
  return {
    pushToast: useCallback(
      (props: Params) => {
        pushToastRef.current(props);
      },
      [pushToastRef]
    ),
  };
}

export function Toasts() {
  const [toasts, setToasts] = useState<Params[]>([]);
  const { pushToastRef } = useContext(toastContext);

  pushToastRef.current = (props: Params) => {
    const id = Date.now();
    const timer = setTimeout(() => {
      setToasts((s) => s.filter((toast) => toast.id !== id));
    }, (props.duration ?? 5) * 1000);
    const toast = { ...props, timer, id };
    setToasts((v) => [...v, toast]);
  };

  const onRemove = (props: Params) => {
    clearTimeout(props.timer);
    setToasts((v) => v.filter((v) => v !== props));
  };

  return (
    <div>
      {toasts.map((props, index) => (
        <div key={index} onClick={() => onRemove(props)}>
          <Toast {...props} />
        </div>
      ))}
    </div>
  );
}
