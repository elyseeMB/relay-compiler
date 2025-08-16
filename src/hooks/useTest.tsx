import {
  ComponentProps,
  createContext,
  PropsWithChildren,
  useCallback,
  useContext,
  useRef,
  useState,
} from "react";
import { Message } from "../components/Message.tsx";

type Params = ComponentProps<typeof Message> & { id: number };

const defaultFunc = {
  messages: (message: Params) => {},
  toast: (toast: Params) => {},
};

const defaultValue = {
  pushMessageRef: { current: defaultFunc.messages },
  pushToastRef: { current: defaultFunc.toast },
};

const MessageContext = createContext(defaultValue);

export function MessageContextProvider({ children }: PropsWithChildren) {
  const refElement = {
    pushMessageRef: useRef(defaultFunc.messages),
    pushToastRef: useRef(defaultFunc.toast),
  };
  return (
    <MessageContext.Provider value={refElement}>
      {children}
      <Messages />
    </MessageContext.Provider>
  );
}

export function useMessage() {
  const { pushMessageRef } = useContext(MessageContext);
  return {
    pushMessage: useCallback(
      (message: Params) => {
        pushMessageRef.current(message);
      },
      [pushMessageRef]
    ),
  };
}

function Messages() {
  const [messages, setMessages] = useState([] as Params[]);
  const { pushMessageRef } = useContext(MessageContext);
  pushMessageRef.current = (message: Params) => {
    const id = Date.now();
    const messageElement = { ...message, id };
    setMessages((v) => [...v, messageElement]);
  };

  return (
    <div>
      {messages.map((message) => (
        <span key={message.id}>
          <Message {...message} />
        </span>
      ))}
    </div>
  );
}
