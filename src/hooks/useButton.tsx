import {
  ComponentProps,
  createContext,
  PropsWithChildren, useCallback, useContext,
  useRef, useState,
} from "react";
import { Button } from "../components/Button.tsx";
import { constructor } from "@typescript-eslint/eslint-plugin";

type Params = ComponentProps<typeof Button>
type ButtonItems = ComponentProps<typeof Button> & {
  id: number
}
const defaultFunc = (button: Params) => {
};

const defaultValue = {
  pushButtonRef: {current: defaultFunc},
};

const ButtonContext = createContext(defaultValue);

export function ButtonContextProvider({children}: PropsWithChildren) {
  const pushButtonRef = useRef(defaultFunc);
  return <ButtonContext.Provider value={{pushButtonRef}}>
    {children}
    <Buttons/>
  </ButtonContext.Provider>;
}

export function useButton() {
  const {pushButtonRef} = useContext(ButtonContext);
  return {
    pushButton: useCallback((button: Params) => {
      pushButtonRef.current(button);
    }, [pushButtonRef]),
  };
}


function Buttons() {
  const [buttons, setButtons] = useState([] as ButtonItems[]);
  const {pushButtonRef} = useContext(ButtonContext);
  pushButtonRef.current = (button: Params) => {
    const id = Date.now();
    const buttonElement = {...button, id};
    setButtons((v: ButtonItems[]) => [...v, buttonElement]);
  };
  
  return <>
    {buttons.map((button) => (
      <div key={button.id} className="button__group">
        <Button {...button} />
      </div>
    ))}
  </>;
}